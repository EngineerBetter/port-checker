package main_test

import (
	. "github.com/EngineerBetter/port-checker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("Port Checker Server", func() {
	Describe("/", func() {
		It("returns 200 OK", func() {
			handler := new(IndexHandler)
			server := httptest.NewServer(handler)
			defer server.Close()

			_, err := http.Get(server.URL)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
