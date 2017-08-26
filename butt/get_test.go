package butt_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/gapi/butt"
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ButtDLFake struct {
}

func FakeButtDataContextMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		buttDl := &ButtDLFake{}
		c.Set("Db", buttDl)

		c.Next()
	}
}

func (b *ButtDLFake) RetrieveAllButts() ([]Butt, error) {
	butts := []Butt{{Id: 1, Name: "(("}, {Id: 2, Name: "))"}, {Id: 3, Name: "(( ))"}}
	return butts, nil
}

func (b *ButtDLFake) RetrieveSingleButt(buttId int) (Butt, error) {
	butt := Butt{Id: 1, Name: "(("}
	return butt, nil
}

var _ = Describe("Butt", func() {
	Context("When RetrieveAllButts is called  ", func() {
		It("farts hehe", func() {
			gin.SetMode(gin.TestMode)

			r := gin.Default()
			r.Use(FakeButtDataContextMW())
			r.GET("/butt", RetrieveAllButts)

			req, _ := http.NewRequest("GET", "/butt", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			statusOK := w.Code == http.StatusOK
			fmt.Println(w)

			expected := `{"data":[{"Id":1,"name":"(("},{"Id":2,"name":"))"},{"Id":3,"name":"(( ))"}],"status":200}`
			Expect(statusOK).To(Equal(true))
			Expect(w.Body.String()).To(Equal(expected))
		})
	})
})
