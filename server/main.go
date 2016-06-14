package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/dgraph-io/dgraph/x"
	"github.com/dgraph-io/gru/server/interact"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
)

const (
	DEMO     = "demo"
	TEST     = "test"
	END      = "END"
	CORRECT  = 1
	WRONG    = 2
	DURATION = 60 * time.Minute
)

var (
	quizFile = flag.String("quiz", "test.yml", "Input question file")
	port     = flag.String("port", ":8888", "Port on which server listens")
	candFile = flag.String("cand", "candidates.txt", "Candidate inforamation")
	quizInfo map[string][]Question
	cmap     map[string]Candidate
	glog     = x.Log("Gru Server")
	// List of question ids.
	qnList     []string
	demoQnList []string
)

type server struct{}

func checkToken(token string) (Candidate, bool) {
	var c Candidate
	var ok bool

	if c, ok = cmap[token]; ok && time.Now().Before(c.validity) {
		// Initially testStart is zero, but after candidate has taken the
		// test once, it shouldn't be zero.
		if !c.testStart.IsZero() && time.Now().After(c.testStart.Add(DURATION)) {
			return c, false
		}
		return c, true
	}
	return c, false
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func authenticate(t *interact.Token) (*interact.Session, error) {
	var c Candidate
	var valid bool

	if c, valid = checkToken(t.Id); !valid {
		return nil, errors.New("Invalid token")
	}

	session := interact.Session{Id: RandStringBytes(36)}
	c.qnList = qnList[:]
	c.demoQnList = demoQnList[:]
	f, err := os.OpenFile(fmt.Sprintf("logs/%s.log", t.Id),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	c.logFile = f
	log.SetOutput(c.logFile)
	log.Println("auth_token", t.Id)
	log.Println("session_token", session.Id)
	cmap[t.Id] = c
	return &session, nil
}

func (s *server) Authenticate(ctx context.Context,
	t *interact.Token) (*interact.Session, error) {
	return authenticate(t)
}

type Option struct {
	Uid string
	Str string
}

type Question struct {
	Id       string
	Str      string
	Correct  []string
	Opt      []Option
	Positive float32
	Negative float32
	Tag      string
}

func nextQuestion(c Candidate, testType string) (int, *interact.Question) {
	var list []string
	if testType == DEMO {
		list = c.demoQnList
	} else {
		list = c.qnList
	}

	idx := rand.Intn(len(list))
	q := quizInfo[testType][idx]

	var opts []*interact.Answer
	for _, o := range q.Opt {
		a := &interact.Answer{Id: o.Uid, Str: o.Str}
		opts = append(opts, a)
	}

	var isM bool
	if len(q.Correct) > 1 {
		isM = true
	}
	que := &interact.Question{Id: q.Id, Str: q.Str, Options: opts,
		IsMultiple: isM, Positive: q.Positive, Negative: q.Negative,
		Totscore: c.score,
	}
	return idx, que
}

func getQuestion(req *interact.Req) (*interact.Question, error) {
	var c Candidate
	var ok bool
	if c, ok = cmap[req.Token]; !ok {
		return nil, errors.New("Invalid token.")
	}

	testType := req.TestType
	if testType == DEMO {
		if len(c.demoQnList) == 0 {
			// If demo qns are over, indicate end of demo to client.
			q := &interact.Question{Id: END, Totscore: 0}
			c.score = 0
			cmap[req.Token] = c
			return q, nil
		}
		idx, q := nextQuestion(c, testType)
		c.demoQnList = append(c.demoQnList[:idx], c.demoQnList[idx+1:]...)
		cmap[req.Token] = c
		return q, nil
	}
	// TOOD(pawan) - Check if time is up
	if len(c.qnList) == 0 {
		q := &interact.Question{Id: END, Totscore: c.score}
		c.logFile.Close()
		return q, nil
	}
	// This means its the first test question.
	if len(c.qnList) == len(qnList) {
		log.SetOutput(c.logFile)
		log.Println("test_start")
		c.testStart = time.Now()
	}
	idx, q := nextQuestion(c, testType)
	c.qnList = append(c.qnList[:idx], c.qnList[idx+1:]...)
	cmap[req.Token] = c
	return q, nil
}

func (s *server) GetQuestion(ctx context.Context,
	req *interact.Req) (*interact.Question, error) {
	return getQuestion(req)
}

func isCorrectAnswer(resp *interact.Response) (int, int64, error) {
	for i, que := range quizInfo[resp.TestType] {
		if que.Id == resp.Qid {
			if reflect.DeepEqual(resp.Aid, que.Correct) {
				return i, CORRECT, nil
			} else {
				return i, WRONG, nil
			}
		}
	}
	return -1, -1, errors.New("No matching question")
}

func (s *server) SendAnswer(ctx context.Context,
	resp *interact.Response) (*interact.Status, error) {
	var c Candidate
	var ok bool
	if c, ok = cmap[resp.Token]; !ok {
		return &interact.Status{Status: 0}, errors.New("Invalid token.")
	}

	var status interact.Status
	var err error
	var idx int

	idx, status.Status, err = isCorrectAnswer(resp)

	if len(resp.Aid) > 0 && resp.Aid[0] != "skip" {
		if status.Status == 1 {
			c.score += quizInfo[resp.TestType][idx].Positive
		} else {
			c.score -= quizInfo[resp.TestType][idx].Negative
		}
	} else {
		if len(resp.Aid) > 1 {
			glog.Error("Got extra optoins with SKIP")
		}
	}

	cmap[resp.Token] = c
	// We log only if its a actual test question.
	if resp.TestType == TEST {
		log.SetOutput(c.logFile)
		log.Println("response", resp.Qid, resp.Aid, c.score)
	}

	return &status, err
}

func (s *server) StreamChan(stream interact.GruQuiz_StreamChanServer) error {

	var stat interact.ServerStatus
	var wg sync.WaitGroup

	msg, err := stream.Recv()
	if err != nil {
		glog.Error(err)
	}
	token := msg.Token

	wg.Add(1)
	go func() {
		for {
			stat.TimeLeft = time.Now().Sub(cmap[token].testStart).String()
			if err := stream.Send(&stat); err != nil {
				glog.Error(err)
			}
			time.Sleep(5 * time.Second)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					glog.Error(err)
				} else {
					break
				}
			}
			log.SetOutput(cmap[token].logFile)
			log.Println("ping", msg.CurrQuestion)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

/*
func (s *server) StreamChanClient(stream interact.GruQuiz_StreamChanClientClient) error {


		go func() {
			for {
				msg, err := stream.Recv()
				if err != nil {
					if err != io.EOF {
						glog.Error(err)
					} else {
						break
					}
				}
				log.Println("ping", msg.CurrQuestion)
			}
		}()

	return nil
}
*/
func runGrpcServer(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		glog.WithField("err", err).Fatalf("Error running quiz server")
		return
	}
	glog.WithField("address", ln.Addr()).Info("Server listening")

	s := grpc.NewServer()
	interact.RegisterGruQuizServer(s, &server{})
	if err = s.Serve(ln); err != nil {
		glog.Fatalf("While serving gRpc requests", err)
	}
	return
}

type Candidate struct {
	name     string
	email    string
	validity time.Time
	score    float32
	// List of test questions that have not been asked to the candidate yet.
	qnList     []string
	demoQnList []string
	logFile    *os.File
	testStart  time.Time
}

func parseCandidateInfo(file string) error {
	cmap = make(map[string]Candidate)
	format := "2006/01/02 (MST)"
	f, err := os.Open(*candFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line[0] == ';' {
			continue
		}

		var c Candidate
		splits := strings.Split(line, " ")

		if len(splits) < 6 {
			continue
		}

		c.name = strings.Join(splits[:2], " ")
		c.email = splits[2]
		c.validity, err = time.Parse(format,
			fmt.Sprintf("%s (%s)", splits[3], splits[4]))
		if err != nil {
			log.Fatal(err)
		}

		token := splits[5]
		cmap[token] = c
	}
	return nil
}

func extractQids(testType string) []string {
	var list []string
	for _, q := range quizInfo[testType] {
		list = append(list, q.Id)
	}
	return list
}

func extractQuizInfo(file string) map[string][]Question {
	var info map[string][]Question
	b, err := ioutil.ReadFile(file)
	if err != nil {
		glog.WithField("err", err).Fatal("Error while reading quiz info file")
	}
	err = yaml.Unmarshal(b, &info)
	if err != nil {
		glog.WithField("err", err).Fatal("Error while unmarshalling into yaml")
	}
	return info
}

func main() {
	flag.Parse()
	quizInfo = extractQuizInfo(*quizFile)
	qnList = extractQids(TEST)
	demoQnList = extractQids(DEMO)
	parseCandidateInfo(*candFile)
	// TODO(pawan) - Read testStart timings for candidates.
	runGrpcServer(*port)
}
