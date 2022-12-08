package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func handleProtobuf(w http.ResponseWriter) {
	data, err := proto.Marshal(articles)
	if err != nil {
		return
	}
	_, err = w.Write(data)
	if err != nil {
		return
	}
}

func handleXML(w http.ResponseWriter) {
	data, err := xml.Marshal(articles)
	if err != nil {
		fmt.Printf("Failed to marshal XML")
	}
	_, err = w.Write(data)
	if err != nil {
		return
	}
}

func handleJson(w http.ResponseWriter) {
	data, err := json.Marshal(articles)
	if err != nil {
		_, err2 := fmt.Fprintf(w, "Failed to marshal JSON %v", err)
		if err2 != nil {
			return
		}

	}
	_, err = w.Write(data)
	if err != nil {
		return
	}
}

func articleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		switch req.Header.Get("Accept") {
		case "application/x-protobuf":
			handleProtobuf(w)
		case "application/xml":
			handleXML(w)
		default:
			handleJson(w)
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
