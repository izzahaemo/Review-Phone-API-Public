package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"dataphone/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type phoneInput struct {
	Name        string `json:"name"`
	Display     string `json:"display"`
	Camerafront string `json:"Camerafront"`
	Cameraback  string `json:"Cameraback"`
	Chipset     string `json:"chipset"`
	Ram         string `json:"ram"`
	Battery     string `json:"battery"`
	Pict        string `json:"pict"`
	BrandID     uint   `json:"brandID"`
	Username    string `json:"username"`
}

type phoneUser struct {
	Username string `json:"username"`
}

// GetAllPhone godoc
// @Summary Get all Phone.
// @Description Get a list of Phone.
// @Tags Phone
// @Produce json
// @Success 200 {object} []models.Phone
// @Router /phone [get]
func GetAllPhone(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var phone []models.Phone
	db.Find(&phone)

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// GetPhoneById godoc
// @Summary Get Phone.
// @Description Get a Phone by id.
// @Tags Phone
// @Produce json
// @Param id path string true "Phone id"
// @Success 200 {object} models.Phone
// @Router /phone/{id} [get]
func GetPhoneById(c *gin.Context) { // Get model if exist
	var phone models.Phone

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// GetPhoneByBrand godoc
// @Summary Get Phone By Brand.
// @Description Get a Phone by Brand id.
// @Tags Phone
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} models.Phone
// @Router /phone/brand/{id} [get]
func GetPhoneByBrand(c *gin.Context) { // Get model if exist
	var phone []models.Phone

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("brand_id = ?", c.Param("id")).Find(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// CreatePhone godoc
// @Summary Create New Phone.
// @Description Creating a new Phone (only username with role id 1 can do this).
// @Tags Phone
// @Param Body body phoneInput true "the body to create a new Phone"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Produce json
// @Success 200 {object} models.Phone
// @Failure 400 {object} FailureUser{} "If Username not Found"
// @Router /phone/ [post]
func CreatePhone(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input phoneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tru, err := models.UserPhone(input.Username, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot found"})
		return
	}

	if !tru {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot access this menu"})
		return
	}

	// Create Phone
	phone := models.Phone{Name: input.Name, Display: input.Display, Camerafront: input.Camerafront, Cameraback: input.Cameraback, Chipset: input.Chipset, Ram: input.Ram, Battery: input.Display, Pict: input.Pict, BrandID: input.BrandID}

	db.Create(&phone)

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// UpdatePhone godoc
// @Summary Update Phone.
// @Description Update a Phone by id (only username with role id 1 can do this).
// @Tags Phone
// @Produce json
// @Param id path string true "Phone id"
// @Param Body body phoneInput true "the body to update Phone"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} models.Phone
// @Failure 400 {object} FailureForbiden{} "If Username cannot using this"
// @Router /phone/{id} [patch]
func UpdatePhone(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var phone models.Phone
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input phoneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tru, err := models.UserPhone(input.Username, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot found"})
		return
	}

	if !tru {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot access this menu"})
		return
	}

	var updatedInput models.Phone
	updatedInput.Name = input.Name
	updatedInput.Display = input.Display
	updatedInput.Camerafront = input.Camerafront
	updatedInput.Cameraback = input.Cameraback
	updatedInput.Chipset = input.Chipset
	updatedInput.Ram = input.Ram
	updatedInput.Battery = input.Battery
	updatedInput.Pict = input.Pict
	updatedInput.BrandID = input.BrandID

	db.Model(&phone).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// DeletePhone godoc
// @Summary Delete one Phone.
// @Description Delete a Phone by id (only username with role id 1 can do this).
// @Tags Phone
// @Produce json
// @Param Body body phoneUser true "the User"
// @Param id path string true "Phone id"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} Success{}
// @Failure 400 {object} FailureRecord{} "If the phone not found"
// @Router /phone/{id} [delete]
func DeletePhone(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)

	var input phoneUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tru, err := models.UserPhone(input.Username, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot found"})
		return
	}

	if !tru {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot access this menu"})
		return
	}

	var phone models.Phone
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&phone)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// UploadPicture godoc
// @Summary Upload Picture Phone.
// @Description Upload Phone Picture (.png only!).
// @Tags Phone
// @Accept       multipart/form-data
// @Param id path string true "Phone id"
// @Param file  formData  file  true  "image (.png only)"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Produce json
// @Success 200
// @Failure 400 {object} FailureRecord{} "If the phone not found"
// @Router /phone/upload/{id} [post]
func UploadPhone(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// var input phoneUser
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// tru, err := models.UserPhone(input.Username, db)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot found"})
	// 	return
	// }

	// if !tru {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot access this menu"})
	// 	return
	// }

	namephone, err := models.NamePhone(c.Param("id"), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone cannot found"})
		return
	}
	namep := namephone + ".png"

	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	// Set Folder untuk menyimpan filenya
	path := "assets/phones/" + namep
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}
	var phone models.Phone
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput models.Phone
	updatedInput.Pict = namep

	db.Model(&phone).Updates(updatedInput)

	// Response
	// c.String(http.StatusOK, fmt.Sprintf("File : %s, namep : %s", file.Filename, namep))
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetPhonePicture godoc
// @Summary Get Phone Picture.
// @Description Get a Phone Picture.
// @Tags Phone
// @Produce octet-stream
// @Param id path string true "Phone id"
// @Success 200
// @Router /phone/picture/{id} [get]
func GetPhonePicture(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	namephone, err := models.NamePhone(c.Param("id"), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone cannot found"})
		return
	}
	namep := namephone + ".png"

	onfile := "assets/phones/" + namep
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
