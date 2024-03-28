package internal

import (
	"holistic-herbal-encyclopedia/model"
	"net/http"
    "strconv"
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
	name := c.Query("name")
	entries := model.New().SelectByName(name) 
    c.HTML(http.StatusOK, "nameSearch.html", gin.H{
		"Index": "/",
		"Entries": entries,
	})
}

func SearchByUse(c *gin.Context) {
	use := c.Query("use")
	entries := model.New().SelectByUse(use) 
    c.HTML(http.StatusOK, "nameSearch.html", gin.H{
		"Index": "/",
		"Entries": entries,
	})
}

func GetEdit(c *gin.Context) {
	num := c.Query("id")
	id, err := strconv.ParseInt(num, 10, 64)
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid value provided",
        })
        return
    }
	entry := model.New().SelectById(id)
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"Id": entry.Id,
		"Name": entry.Name,
		"Dosage": entry.Dosage,
		"Uses": entry.Uses,
		"Precautions": entry.Precautions,
		"Preparations": entry.Preparations,
	})
}

func PostEdit(c *gin.Context) {
	num := c.Query("id")
	id, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid value provided",
        })
        return
    }
	var form model.Herb
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	model.New().Update(id, form.Name, form.Dosage, form.Uses, form.Precautions, form.Preparations)
	c.Redirect(http.StatusSeeOther, "/")
}