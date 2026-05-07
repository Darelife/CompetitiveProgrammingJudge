package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"judge/internal/db"
	"judge/internal/models"
	"judge/internal/queue"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/submit", SubmitHandler)
	router.GET("/result/:id", ResultHandler)
}

func SubmitHandler(c *gin.Context) {
	var submission models.Submission

	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := db.CreateSubmission(
		submission.Code,
		submission.Language,
		submission.QuestionID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	submission.ID = int(id)

	queue.JobQueue <- submission

	c.JSON(http.StatusOK, gin.H{
		"submission_id": id,
	})
}

func ResultHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	result, err := db.GetSubmission(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Submission not found",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
