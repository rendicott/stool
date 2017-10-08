package player

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(FakePlayerDataContextMW())
	r.GET("/player/:Id", RetrieveSinglePlayer)
	r.GET("/players", RetrieveAllPlayers)

	Context("When RetrieveAllPlayers is called  ", func() {
		It("returns a 200", func() {

			req, _ := http.NewRequest("GET", "/players", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			statusOK := w.Code == http.StatusOK
			fmt.Println(w)

			expected := `{"data":[{"Id":1,"name":"Vlaada Chvatil"},{"Id":2,"name":"Uwe Rosenberg"},{"Id":3,"name":"Stefan Feld"}],"status":200}`
			Expect(statusOK).To(Equal(true))
			Expect(w.Body.String()).To(Equal(expected))
		})
	})

	Context("When RetrieveSinglePlayer is called ", func() {

		It("Returns a 200 when the player exists", func() {
			req, _ := http.NewRequest("GET", "/player/1", nil)
			fmt.Println(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusOK := w.Code == http.StatusOK
			fmt.Println(w)

			expected := `{"data":{"Id":1,"name":"Klaus Teuber"},"status":200}`
			Expect(statusOK).To(Equal(true))
			Expect(w.Body.String()).To(Equal(expected))

		})
		It("Returns a 404 when the player does not exist", func() {
			req, _ := http.NewRequest("GET", "/player/2", nil)
			fmt.Println(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusMatch := w.Code == http.StatusNotFound
			fmt.Println(w)

			Expect(statusMatch).To(Equal(true))
		})
		It("Returns a 500 when input is invalid", func() {
			req, _ := http.NewRequest("GET", "/player/a", nil)
			fmt.Println(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusMatch := w.Code == http.StatusInternalServerError
			fmt.Println(w)

			Expect(statusMatch).To(Equal(true))
		})
	})

})
