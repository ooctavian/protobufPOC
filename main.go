package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func handleProtobuf() []byte {
	data, err := proto.Marshal(articles)
	if err != nil {
		return nil
	}
	return data
}

func handleXML() []byte {
	data, err := xml.Marshal(articles)
	if err != nil {
		fmt.Printf("Failed to marshal XML")
	}

	return data
}

func handleJson() []byte {
	data, err := json.Marshal(articles)
	if err != nil {
		return nil
	}

	return data
}

func articleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		switch req.Header.Get("Accept") {
		case "application/x-protobuf":
			w.Write(handleProtobuf())
		case "application/xml":
			w.Write(handleXML())
		default:
			w.Write(handleJson())
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
