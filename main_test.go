package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"protobufPOC/schema"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func makeGetBenchmark(acceptHeader string) func(b *testing.B) {
	return func(b *testing.B) {
		r, _ := http.NewRequest("GET", "/articles", nil)
		r.Header.Add("Accept", acceptHeader)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(articleHandler)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			handler.ServeHTTP(w, r)
		}
	}
}

type marshaller func(v any) ([]byte, error)

func proto_marshal(v any) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func makePostBenchmark(contentType string, fn marshaller) func(b *testing.B) {
	article := &schema.Article{
		Id:    200,
		Title: "Lorem Ipsum",
		Authors: []*schema.Author{
			{
				FirstName: "Dennis",
				LastName:  "Ritchie",
			},
			{
				FirstName: "Brian",
				LastName:  "Kernighan",
			}},
		Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
		Categories: []uint32{
			uint32(3),
		},
		State:        &Draft,
		DateCreated:  timestamppb.New(time.Now()),
		DateModefied: timestamppb.Now(),
	}
	data, err := fn(article)
	if err != nil {
		return nil
	}
	reader := bytes.NewReader(data)

	return func(b *testing.B) {
		r, _ := http.NewRequest("POST", "/articles", reader)
		r.Header.Add("Content-Type", contentType)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(articleHandler)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			handler.ServeHTTP(w, r)
		}
	}
}

func BenchmarkHandleGetRequest(b *testing.B) {
	b.Run("Endpoint with JSON: GET /articles", makeGetBenchmark("application/json"))
	b.Run("Endpoint with XML: GET /articles", makeGetBenchmark("application/xml"))
	b.Run("Endpoint with Protobuf: GET /articles", makeGetBenchmark("application/x-protobuf"))
}

func BenchmarkHandlePostRequest(b *testing.B) {
	b.Run("Endpoint with JSON: POST /articles", makePostBenchmark("application/json", json.Marshal))
	b.Run("Endpoint with XML: POST /articles", makePostBenchmark("application/xml", xml.Marshal))
	b.Run("Endpoint with Protobuf: POST /articles", makePostBenchmark("application/x-protobuf", proto_marshal))
}
