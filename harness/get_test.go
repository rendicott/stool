package harness_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/chrisevett/stool/harness"
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RunAllTests", func() {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(FakeMiddleware())
	r.GET("/test", RunAllTests)
	Context("When we send get with no arguments", func() {
		It("returns a 200", func() {
			req, _ := http.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusOK := w.Code == http.StatusOK

			Expect(statusOK).To(Equal(true))
		})
	})

})
