package tcp_test

import (
	"github.com/EngineerBetter/port-checker/tcp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net"
	"strconv"
)

var _ = Describe("TCP Connection Checking", func() {
	Context("when there is no process to listen", func() {
		It("returns a connection refused error", func() {
			port := getUnusedPort()
			err := tcp.Check(net.ParseIP("127.0.0.1"), port)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("getsockopt: connection refused"))
		})
	})
})

func getUnusedPort() int {
	listener, err := net.Listen("tcp", "")
	Ω(err).ShouldNot(HaveOccurred())

	addr := listener.Addr().String()
	_, port, err := net.SplitHostPort(addr)
	Ω(err).ShouldNot(HaveOccurred())
	portNum, err := strconv.Atoi(port)
	Ω(err).ShouldNot(HaveOccurred())
	Ω(portNum).Should(BeNumerically(">", 0))
	err = listener.Close()
	Ω(err).ShouldNot(HaveOccurred())
	return portNum
}
