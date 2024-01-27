package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FilterEmail struct {
	ClassifiedWords []string `json:"classified_words"`
	EmailBody       string   `json:"body"`
}

func emailFilter(classifiedWords []string, emailText string) (bool, string) {
	isClassidied := false
	filteredEmailText := emailText

	for _, list := range classifiedWords {
		if strings.Contains(emailText, list) {
			isClassidied = true
			filteredEmailText = strings.ReplaceAll(filteredEmailText, list, "*****")
		}
	}

	return isClassidied, filteredEmailText
}

func filterEmailHandler(c *gin.Context) {
	var filterEmail FilterEmail

	if err := c.BindJSON(&filterEmail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasClassifiedWords, emailText := emailFilter(filterEmail.ClassifiedWords, filterEmail.EmailBody)

	c.JSON(http.StatusOK, gin.H{
		"classified": hasClassifiedWords,
		"message":    emailText,
	})
}

func main() {
	r := gin.Default()

	r.POST("/filter-email", filterEmailHandler)

	r.Run(":8080")
}
