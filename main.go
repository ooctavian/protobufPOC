package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"protobufPOC/schema"

	"google.golang.org/protobuf/proto"
)

func getHandleProtobuf() []byte {
	data, err := proto.Marshal(articles)
	if err != nil {
		return nil
	}
	return data
}

func getHandleXML() []byte {
	data, err := xml.Marshal(articles)
	if err != nil {
		fmt.Printf("Failed to marshal XML")
	}

	return data
}

func getHandlerJson() []byte {
	data, err := json.Marshal(articles)
	if err != nil {
		return nil
	}

	return data
}

func postHandleProtobuf(body []byte) {
	article := &schema.Article{}
	err := proto.Unmarshal(body, article)
	articles.Data = append(articles.Data, article)
	if err != nil {
		return
	}
}

func postHandleJSON(body []byte) {
	article := &schema.Article{}
	err := json.Unmarshal(body, article)
	articles.Data = append(articles.Data, article)
	if err != nil {
		return
	}
}

func postHandleXML(body []byte) {
	article := &schema.Article{}
	err := xml.Unmarshal(body, article)
	articles.Data = append(articles.Data, article)
	if err != nil {
		return
	}
}

func articleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		switch req.Header.Get("Accept") {
		case "application/x-protobuf":
			w.Write(getHandleProtobuf())
		case "application/xml":
			w.Write(getHandleXML())
		default:
			w.Write(getHandlerJson())
		}
		return
	}

	if req.Method == http.MethodPost {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		switch req.Header.Get("Content-Type") {
		case "application/x-protobuf":
			postHandleProtobuf(body)
		case "application/xml":
			postHandleXML(body)
		default:
			postHandleJSON(body)
		}
	}
}

func main() {
	http.HandleFunc("/articles", articleHandler)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Failed to start server")
	}
}
