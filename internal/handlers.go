package internal

import (
	"holistic-herbal-encyclopedia/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    //entries := model.New().Select() 
    c.HTML(http.StatusOK, "index.html", gin.H{"Create": "/create"})
}

type NewHerb struct {
	Name string `form:"name" binding:"required"`
	Dosage string `form:"dosage" binding:"required"`
	Uses string `form:"uses" binding:"required"`
	Precautions string `form:"precautions" binding:"required"`
	Preparations string `form:"preparations" binding:"required"`
}

func Create(c *gin.Context) {
	var form NewHerb
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	_ = model.New().Insert(form.Name, form.Dosage, form.Uses, form.Precautions, form.Preparations)
	c.Redirect(http.StatusSeeOther, "/")
}