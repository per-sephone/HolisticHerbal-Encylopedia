package holistic

import (
	"holistic-herbal-encyclopedia/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    entries := model.New().Select() 
    c.HTML(http.StatusOK, "index.html", gin.H{
        "Entries": entries,
    })
}