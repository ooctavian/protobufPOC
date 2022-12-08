package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func makeBenchmark(acceptHeader string) func(b *testing.B) {
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

func BenchmarkHandleJson(b *testing.B) {
	b.Run("Endpoint with JSON: GET /articles", makeBenchmark("application/json"))
	b.Run("Endpoint with XML: GET /articles", makeBenchmark("application/xml"))
	b.Run("Endpoint with Protobuf: GET /articles", makeBenchmark("application/x-protobuf"))
}
