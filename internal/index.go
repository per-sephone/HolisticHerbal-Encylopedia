package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    //entries := model.New().Select() 
    c.HTML(http.StatusOK, "index.html", gin.H{})
}