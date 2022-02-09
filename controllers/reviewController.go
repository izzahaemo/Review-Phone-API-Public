package controllers

import (
	"net/http"

	"dataphone/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewInput struct {
	Isi     string `json:"isi"`
	PhoneID uint   `json:"phoneID"`
	UserID  uint   `json:"userID"`
}

// GetReviewById godoc
// @Summary Get Review.
// @Description Get a Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} models.Review
// @Router /review/{id} [get]
func GetReviewById(c *gin.Context) { // Get model if exist
	var review models.Review

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetReviewByPhone godoc
// @Summary Get Review by Phone id.
// @Description Get a Review by Phone id.
// @Tags Review
// @Produce json
// @Param id path string true "Phone id"
// @Success 200 {object} models.Review
// @Router /review/phone/{id} [get]
func GetReviewByPhone(c *gin.Context) { // Get model if exist
	var review []models.Review

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("phone_id = ?", c.Param("id")).Find(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// CreateReview godoc
// @Summary Create New Review.
// @Description Creating a new Review.
// @Tags Review
// @Param Body body reviewInput true "the body to create a new Review"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Produce json
// @Success 200 {object} models.Review
// @Failure 400 {object} FailureUser{} "If Username not Found"
// @Router /review/ [post]
func CreateReview(c *gin.Context) {
	// Validate input
	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var reviewc models.Review
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("phone_id = ?", input.PhoneID).First(&reviewc).Error; err == nil {
		if err := db.Where("user_id = ?", input.UserID).First(&reviewc).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already write review"})
			return
		}
	}

	// Create Rating
	review := models.Review{Isi: input.Isi, PhoneID: input.PhoneID, UserID: input.UserID}
	db.Create(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReview godoc
// @Summary Update Review.
// @Description Update Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Param Body body reviewInput true "the body to update age rating category"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} models.Review
// @Failure 400 {object} FailureRecord{} "If the review not found"
// @Router /review/{id} [patch]
func UpdateReview(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Review
	updatedInput.Isi = input.Isi
	updatedInput.PhoneID = input.PhoneID
	updatedInput.UserID = input.UserID

	db.Model(&review).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview godoc
// @Summary Delete one Review.
// @Description Delete a Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} Success{}
// @Failure 400 {object} FailureRecord{} "If the review not found"
// @Router /review/{id} [delete]
func DeleteReview(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
