//go:build bdd
// +build bdd

package tests

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health check endpoint", func() {
	Context("do check", func() {
		It("and return 200", func() {
			c := http.Client{}
			r, err := c.Get(fmt.Sprintf("http://localhost:%d/healthz", server.Port()))
			Expect(err).NotTo(HaveOccurred())
			Expect(r.StatusCode).To(Equal(http.StatusOK))
		})

		It("and return 500", func() {
			c := http.Client{}
			r, err := c.Get(fmt.Sprintf("http://localhost:%d/healthz", server.Port()))
			Expect(err).NotTo(HaveOccurred())
			Expect(r.StatusCode).To(Equal(http.StatusInternalServerError))
		})
	})
})