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

	Context("When we send get with no arguments and a runner and path are defined", func() {
		It("returns a 200", func() {
			gin.SetMode(gin.TestMode)

			r := gin.Default()
			r.Use(FakeMiddlewareGood())
			r.GET("/test", RunAllTests)
			req, _ := http.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusOK := w.Code == http.StatusOK

			Expect(statusOK).To(Equal(true))
		})
	})
	Context("When we send get with no arguments and a runner is not defined", func() {
		It("returns internal server error", func() {
			gin.SetMode(gin.TestMode)

			r := gin.Default()
			r.Use(FakeMiddlewareRunnerNotDefined())
			r.GET("/test", RunAllTests)
			req, _ := http.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusError := w.Code == http.StatusInternalServerError

			Expect(statusError).To(Equal(true))
		})

	})
	Context("When we send get with no arguments and a test path is not defined", func() {
		It("returns a server error", func() {
			gin.SetMode(gin.TestMode)

			r := gin.Default()
			r.Use(FakeMiddlewarePathNotDefined())
			r.GET("/test", RunAllTests)
			req, _ := http.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusError := w.Code == http.StatusInternalServerError

			Expect(statusError).To(Equal(true))
		})
	})
})
