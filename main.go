package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"protobufPOC/definition"
	"protobufPOC/schema"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	a = []definition.Articles{{
		Title: "The C Programming Language",
		Authors: []definition.Author{
			{
				FirstName: "Dennis",
				LastName:  "Ritchie",
			},
			{
				FirstName: "Brian",
				LastName:  "Kernighan",
			},
		},
		Content:      "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
		Draft:        true,
		Category:     3,
		DateCreated:  time.Now(),
		DateModefied: time.Now(),
	}}
)

func handleProtobuf(w http.ResponseWriter) {
	art := a[0]
	articles := &schema.Articles{
		Articles: []*schema.Article{{
			Title: art.Title,
			Authors: []*schema.Author{{
				FirstName: a[0].Authors[0].FirstName,
				LastName:  a[0].Authors[0].LastName,
			},
				{
					FirstName: a[0].Authors[1].FirstName,
					LastName:  a[0].Authors[1].LastName,
				},
			},
			Content:      art.Content,
			Draft:        &art.Draft,
			DateCreated:  timestamppb.New(art.DateCreated),
			DateModefied: timestamppb.New(art.DateModefied),
		}},
	}
	data, err := proto.Marshal(articles)
	if err != nil {
		return
	}
	w.Write(data)
}

func handleXML(w http.ResponseWriter) {
	data, err := xml.Marshal(a)
	if err != nil {
		fmt.Fprintf(w, "Failed to marshal XML %v", err)
	}
	w.Write(data)
}

func handleJson(w http.ResponseWriter) {
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Fprintf(w, "Failed to marshal JSON %v", err)

	}
	w.Write(data)
}

func articleHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Header.Get("Content-Type") {
	case "application/x-protobuf":
		handleProtobuf(w)
	case "application/xml":
		handleXML(w)
	default:
		handleJson(w)
	}
}

func main() {
	http.HandleFunc("/articles", articleHandler)
	http.ListenAndServe(":8090", nil)
}
