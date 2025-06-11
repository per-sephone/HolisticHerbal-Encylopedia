package internal

import (
	"holistic-herbal-encyclopedia/model"
	"net/http"
    "strconv"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Index(c *gin.Context) {
	logger := c.MustGet("logger").(*log.Entry)

	logger.Info("Received request to render index page") 

    entries := model.New().Select() 

	logger.WithFields(log.Fields{
		"entry_count": len(entries),
	}).Info("Fetched entries for index page")

    c.HTML(http.StatusOK, "index.html", gin.H{
		"Edit": "/edit",
		"Create": "/create",
		"Search": "/search",
		"Entries": entries,
	})
}

func PostCreate(c *gin.Context) {
	logger := c.MustGet("logger").(*log.Entry)

	var form model.Herb
	if err := c.ShouldBind(&form); err != nil {

		logger.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Failed to bind form data for herb creation")

		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	logger.WithFields(log.Fields{
		"name":         form.Name,
		"dosage":       form.Dosage,
		"uses":         form.Uses,
		"precautions":  form.Precautions,
		"preparations": form.Preparations,
	}).Info("Creating new herb entry")

	err := model.New().Insert(form.Name, form.Dosage, form.Uses, form.Precautions, form.Preparations)
	if err != 0 {
		log.Error("Failed to insert herb into the database")
	}
	c.Redirect(http.StatusSeeOther, "/")
}

func GetCreate(c *gin.Context) {
	logger := c.MustGet("logger").(*log.Entry)

	logger.Info("Rendering herb creation form")

	c.HTML(http.StatusOK, "create.html", gin.H{
		"Create": "/create",
	})
}

func Search(c *gin.Context) {
	logger := c.MustGet("logger").(*log.Entry)

	logger.Info("Rendering search page")

	c.HTML(http.StatusOK, "search.html", gin.H{
		"NameSearch": "/nameSearch",
		"UseSearch": "/useSearch",
	})
}

func SearchByName(c *gin.Context) {
	logger := c.MustGet("logger").(*log.Entry)

	name := c.Query("name")
	entries := model.New().SelectByName(name) 

	logger.WithFields(log.Fields{
		"name": name,
	}).Info("Searching for herbs by name")

    c.HTML(http.StatusOK, "nameSearch.html", gin.H{
		"Index": "/",
		"Entries": entries,
	})
}

func SearchByUse(c *gin.Context) {
	logger := c.MustGet("logger").(*log.Entry)

	use := c.Query("use")
	entries := model.New().SelectByUse(use) 

	logger.WithFields(log.Fields{
		"use": use,
	}).Info("Searching for herbs by use")

    c.HTML(http.StatusOK, "useSearch.html", gin.H{
		"Index": "/",
		"Entries": entries,
	})
}

func GetEdit(c *gin.Context) {
	logger := c.MustGet("logger").(*log.Entry)

	num := c.Query("id")
	logger.WithFields(log.Fields{
		"id": num,
	}).Info("Received request to display herb with ID")

	id, err := strconv.ParseInt(num, 10, 64)
    if err != nil {

		logger.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Error parsing herb ID")

		c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid value provided",
        })
        return
    }

	entry := model.New().SelectById(id)

	logger.WithFields(log.Fields{
		"id": entry.Id,
		"name": entry.Name,
		"dosage": entry.Dosage,
		"uses": entry.Uses,
		"precautions": entry.Precautions,
		"preparations": entry.Preparations,
	}).Info("Old herb entry")


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
	logger := c.MustGet("logger").(*log.Entry)

	num := c.Query("id")
	logger.WithFields(log.Fields{
		"id": num,
	}).Info("Received request to edit herb with ID")

	id, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		logger.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Error parsing herb ID")

		c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid value provided",
        })
        return
    }
	var form model.Herb
	if err := c.ShouldBind(&form); err != nil {

		logger.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Failed to bind form data for edit")

		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	model.New().Update(id, form.Name, form.Dosage, form.Uses, form.Precautions, form.Preparations)

	logger.WithFields(log.Fields{
		"id": id,
		"name": form.Name,
		"dosage": form.Dosage,
		"uses": form.Uses,
		"precautions": form.Precautions,
		"preparations": form.Preparations,
	}).Info("Updated herb entry")
	c.Redirect(http.StatusSeeOther, "/")
}