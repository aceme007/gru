package tag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgraph-io/gru/gruadmin/server"
)

type Tag struct {
	Uid       string `json:"_uid_"`
	Name      string
	Is_delete bool
}

// fetch all the tags
func Index(w http.ResponseWriter, r *http.Request) {
	server.AddCorsHeaders(w)
	tag_mutation := "{debug(_xid_: rootQuestion) { question { question.tag { name _uid_} }}}"
	tag_response, err := http.Post("http://localhost:8080/query", "application/x-www-form-urlencoded", strings.NewReader(tag_mutation))
	if err != nil {
		panic(err)
	}
	defer tag_response.Body.Close()
	tag_body, err := ioutil.ReadAll(tag_response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tag_body))

	jsonResp, err := json.Marshal(string(tag_body))
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
