package game_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/gapi/game"
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GET /game", func() {
	Context("When the request body is empty", func() {
		It("returns a 200", func() {
			gin.SetMode(gin.TestMode)

			r := gin.Default()
			// todo: read this route from the real place
			r.GET("/game", GameIndex)

			req, _ := http.NewRequest("GET", "/game", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusOK := w.Code == http.StatusOK
			Expect(statusOK).To(Equal(true))

			//p, err := ioutil.ReadAll(w.Body)
			// pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

			//return statusOK && pageOK
		})
	})
})

/*

			handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
*/
