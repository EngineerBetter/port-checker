package main_test

import (
	. "github.com/EngineerBetter/port-checker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Port Checker Server", func() {
	Describe("/", func() {
		It("returns 200 OK", func() {
			server := httptest.NewServer(http.HandlerFunc(ServeIndex))
			defer server.Close()

			resp, err := http.Get(server.URL)
			Ω(err).ShouldNot(HaveOccurred())
			body, err := ioutil.ReadAll(resp.Body)
			Ω(string(body)).Should(Equal("Up"))
		})
	})
})
