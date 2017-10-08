package game

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(FakeGameDataContextMW())
	r.DELETE("/game/:Id", DeleteGame)

	Context("When DeleteGameis called  ", func() {
		It("returns a 200", func() {

			req, _ := http.NewRequest("DELETE", "/game/1", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			statusOK := w.Code == http.StatusOK
			fmt.Println(w)

			Expect(statusOK).To(Equal(true))
		})
		It("Returns a 404 when the game does not exist", func() {
			req, _ := http.NewRequest("DELETE", "/game/2", nil)
			fmt.Println(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusMatch := w.Code == http.StatusNotFound
			fmt.Println(w)

			Expect(statusMatch).To(Equal(true))
		})
		It("Returns a 500 when input is invalid", func() {
			req, _ := http.NewRequest("DELETE", "/game/a", nil)
			fmt.Println(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusMatch := w.Code == http.StatusInternalServerError
			fmt.Println(w)

			Expect(statusMatch).To(Equal(true))
		})
	})

})
