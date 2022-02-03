package controllers

import (
	"dataphone/models"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type brandInput struct {
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	Username string `json:"username"`
}

type brandUser struct {
	Username string `json:"username"`
}

// GetAllBrand godoc
// @Summary Get all Brand.
// @Description Get a list of Brand.
// @Tags Brand
// @Produce json
// @Success 200 {object} []models.Brand
// @Router /brand [get]
func GetAllBrand(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var brand []models.Brand
	db.Find(&brand)

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// GetBrandById godoc
// @Summary Get Brand.
// @Description Get a Brand by id.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} models.Brand
// @Router /brand/{id} [get]
func GetBrandById(c *gin.Context) { // Get model if exist
	var brand models.Brand

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// CreateBrand godoc
// @Summary Create New Brand.
// @Description Creating a new Brand (only username with role id 1 can do this).
// @Tags Brand
// @Param Body body brandInput true "the body to create a new Brand"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Produce json
// @Success 200 {object} models.Brand
// @Router /brand/ [post]
func CreateBrand(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input brandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tru, err := models.UserBrand(input.Username, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot found"})
		return
	}

	if !tru {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot access this menu"})
		return
	}

	// Create Brand
	brand := models.Brand{Name: input.Name, Logo: input.Logo}

	db.Create(&brand)

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// UpdateBrand godoc
// @Summary Update Brand.
// @Description Update Brand by id (only username with role id 1 can do this).
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Param Body body brandInput true "the body to update the Brand"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} models.Brand
// @Router /brand/{id} [patch]
func UpdateBrand(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input brandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tru, err := models.UserBrand(input.Username, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot found"})
		return
	}

	if !tru {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot access this menu"})
		return
	}

	var updatedInput models.Brand
	updatedInput.Name = input.Name
	updatedInput.Logo = input.Logo

	db.Model(&brand).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// DeleteBrand godoc
// @Summary Delete one Brand.
// @Description Delete a Brand by id (only username with role id 1 can do this).
// @Tags Brand
// @Produce json
// @Param Body body brandUser true "the User"
// @Param id path string true "Brand id"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} map[string]boolean
// @Router /brand/{id} [delete]
func DeleteBrand(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var input brandUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tru, err := models.UserBrand(input.Username, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot found"})
		return
	}

	if !tru {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot access this menu"})
		return
	}

	var brand models.Brand

	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&brand)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// UploadBrandPicture godoc
// @Summary Upload Brand Picture.
// @Description Upload Brand Picture (.png only!).
// @Tags Brand
// @Accept multipart/form-data
// @Param id path string true "Brand id"
// @Param file  formData  file  true  "image (.png only)"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Produce json
// @Success 200
// @Router /brand/upload/{id} [post]
func UploadBrand(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// var input brandUser
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// tru, err := models.UserBrand(input.Username, db)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot found"})
	// 	return
	// }

	// if !tru {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot access this menu"})
	// 	return
	// }

	namebrand, err := models.NameBrand(c.Param("id"), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brand cannot found"})
		return
	}
	namep := namebrand + ".png"

	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	// Set Folder untuk menyimpan filenya
	path := "assets/brands/" + namep
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput models.Brand
	updatedInput.Name = namebrand
	updatedInput.Logo = namep
	db.Model(&brand).Updates(updatedInput)

	// Response
	// c.String(http.StatusOK, fmt.Sprintf("File : %s, namep : %s", file.Filename, namep))
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetBrandPicture godoc
// @Summary Get Brand Picture.
// @Description Get a Brand Picture.
// @Tags Brand
// @Produce octet-stream
// @Param id path string true "Brand id"
// @Success 200
// @Router /brand/picture/{id} [get]
func GetBrandPicture(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	namebrand, err := models.NameBrand(c.Param("id"), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brand cannot found"})
		return
	}
	namep := namebrand + ".png"

	onfile := "assets/brands/" + namep
	file, err := os.Open(onfile) //Create a file
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "test": onfile})
		return
	}
	defer file.Close()
	c.Writer.Header().Add("Content-type", "	image/png")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
}
