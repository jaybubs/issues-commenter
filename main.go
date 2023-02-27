package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// TODO have a json implementation i guess, env will do for now
	// "issues-commenter/config"
	. "issues-commenter/utilities"
)

func main() {

	type Message struct {
		Body string
	}

	// set vars
	token := os.Getenv("TOKEN")
	// report_path := os.Getenv("REPORT_PATH")
	comment_body := os.Getenv("COMMENT_BODY")
	url := os.Getenv("ISSUE_URL")

	// load token
	// err := config.Read_config()
	// CE(err)
	// token := config.Token

	// move the entire file content into the body
	// f, _ := os.ReadFile(config.Report_path)
	// f, _ := os.ReadFile(report_path)
	// m := Message{string(f)}
	m := Message{comment_body}
	j, _ := json.Marshal(m)

	// url of the git issue api
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	CE(err)

	// token sent as auth, set correct content-type, otherwise won't work
	req.Header.Add("authorization", fmt.Sprintf("token %s", token))
	req.Header.Add("content-type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	
	resp, err := client.Do(req)
	CE(err)

	log.Print(resp)
	
}
