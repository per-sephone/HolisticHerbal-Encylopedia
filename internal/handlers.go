package internal

import (
	"holistic-herbal-encyclopedia/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    entries := model.New().Select() 
    c.HTML(http.StatusOK, "index.html", gin.H{
		"Edit": "/edit",
		"Create": "/create",
		"Search": "/search",
		"Entries": entries,
	})
}

func PostCreate(c *gin.Context) {
	var form model.Herb
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	_ = model.New().Insert(form.Name, form.Dosage, form.Uses, form.Precautions, form.Preparations)
	c.Redirect(http.StatusSeeOther, "/")
}

func GetCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"Create": "/create",
	})
}

func Search(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", gin.H{
		"NameSearch": "/nameSearch",
		"UseSearch": "/useSearch",
	})
}

func SearchByName(c *gin.Context) {

}

func SearchByUse(c *gin.Context) {

}