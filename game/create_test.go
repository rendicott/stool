package game_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	. "github.com/gapi/game"
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	Context("When CreateGAme is called  ", func() {
		gin.SetMode(gin.TestMode)

		r := gin.Default()
		r.Use(FakeGameDataContextMW())
		r.POST("/game", CreateGame)

		data := url.Values{}
		data.Set("Name", "Busen memo")

		req, _ := http.NewRequest("POST", "/game", bytes.NewBufferString(data.Encode()))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		fmt.Println(w.Body.String())
		statusOK := w.Code == http.StatusOK
		fmt.Println(w)

		expected := `{"Id":1,"Name":"Busen memo"}`
		It("returns a 200", func() {
			Expect(statusOK).To(Equal(true))
		})
		It("returns a game with the name we passed in and an id", func() {
			Expect(w.Body.String()).To(Equal(expected))
		})
	})
})
