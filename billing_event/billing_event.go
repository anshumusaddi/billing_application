package billing_event

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func GetBillingEvent(c *gin.Context) {
	Articles := []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	}
	c.JSON(http.StatusOK, Articles)
}

func postBillingEvent() {

}
