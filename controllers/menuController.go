package controllers

import (
	"database/sql"
	"net/http"

	"dataphone/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type dataMenu struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Icon  string `json:"icon"`
}

// GetMenuByRoleId godoc
// @Summary Get Menu.
// @Description Get an Menu by Role id.
// @Tags Menu
// @Produce json
// @Param id path string true "Role id"
// @Produce json
// @Success 200 {object} dataMenu
// @Router /menu/role/{id} [get]
func GetMenuByRoleId(c *gin.Context) { // Get model if exist
	var dataa []dataMenu

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Raw("SELECT    * FROM `menus` JOIN `accessmenus` JOIN `submenus` ON `menus`.`id` = `accessmenus`.`menu_id` WHERE `accessmenus`.`role_id` = @roleidd AND `accessmenus`.`menu_id`=`submenus`.`menu_id`", sql.Named("roleidd", c.Param("id"))).Find(&dataa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// if err := db.Raw("SELECT    * FROM menus JOIN accessmenus JOIN submenus ON menus.id = accessmenus.menu_id WHERE accessmenus.role_id = @roleidd AND accessmenus.menu_id=submenus.menu_id", sql.Named("roleidd", c.Param("id"))).Find(&dataa).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": dataa})
}

type menuInput struct {
	Name string `json:"name"`
}

// CreateMenu godoc
// @Summary Create New Menu`.
// @Description Creating a new Menu.
// @Tags Menu
// @Param Body body menuInput true "the body to create new Menu"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Menu
// @Router /menu/ [post]
func CreateMenu(c *gin.Context) {
	// Validate input
	var input menuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	menu := models.Menu{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&menu)

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

// UpdateMenu godoc
// @Summary Update Menu.
// @Description Update Menu by id.
// @Tags Menu
// @Produce json
// @Param id path string true "Menu id"
// @Param Body body menuInput true "the body to update age rating category"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} models.Menu
// @Router /menu/{id} [patch]
func UpdateMenu(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var menu models.Menu
	if err := db.Where("id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input menuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Menu
	updatedInput.Name = input.Name

	db.Model(&menu).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

// DeleteMenu godoc
// @Summary Delete one Menu.
// @Description Delete a Menu by id.
// @Tags Menu
// @Produce json
// @Param id path string true "Menu id"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} map[string]boolean
// @Router /menu/{id} [delete]
func DeleteMenu(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var menu models.Menu
	if err := db.Where("id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&menu)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

type submenuInput struct {
	MenuID uint   `json:"menuID"`
	Title  string `json:"title"`
	Url    string `json:"url"`
	Icon   string `json:"icon"`
}

// GetAllSubmenu godoc
// @Summary Get all Submenu.
// @Description Get a list of Submenu.
// @Tags Submenu
// @Produce json
// @Success 200 {object} []models.Submenu
// @Router /submenu [get]
func GetAllSubmenu(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var submenu []models.Submenu
	db.Find(&submenu)

	c.JSON(http.StatusOK, gin.H{"data": submenu})
}

// GetSubmenuById godoc
// @Summary Get Submenu.
// @Description Get an Submenu by id.
// @Tags Submenu
// @Produce json
// @Param id path string true "Submenu id"
// @Success 200 {object} models.Submenu
// @Router /submenu/{id} [get]
func GetSubmenuById(c *gin.Context) { // Get model if exist
	var submenu models.Submenu

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&submenu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": submenu})
}

// CreateSubmenu godoc
// @Summary Create New Submenu`.
// @Description Creating a new Submenu.
// @Tags Submenu
// @Param Body body submenuInput true "the body to create a new Submenu"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Produce json
// @Success 200 {object} models.Submenu
// @Router /submenu/ [post]
func CreateSubmenu(c *gin.Context) {
	// Validate input
	var input submenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	submenu := models.Submenu{MenuID: input.MenuID, Title: input.Title, Url: input.Url, Icon: input.Icon}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&submenu)

	c.JSON(http.StatusOK, gin.H{"data": submenu})
}

// UpdateSubmenu godoc
// @Summary Update Submenu.
// @Description Update Submenu by id.
// @Tags Submenu
// @Produce json
// @Param id path string true "Submenu id"
// @Param Body body submenuInput true "the body to update age rating category"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} models.Submenu
// @Router /submenu/{id} [patch]
func UpdateSubmenu(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var submenu models.Submenu
	if err := db.Where("id = ?", c.Param("id")).First(&submenu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input submenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Submenu
	updatedInput.MenuID = input.MenuID
	updatedInput.Title = input.Title
	updatedInput.Url = input.Url
	updatedInput.Icon = input.Icon

	db.Model(&submenu).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": submenu})
}

// DeleteSubmenu godoc
// @Summary Delete one Submenu.
// @Description Delete a Submenu by id.
// @Tags Submenu
// @Produce json
// @Param id path string true "Submenu id"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} map[string]boolean
// @Router /submenu/{id} [delete]
func DeleteSubmenu(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var submenu models.Submenu
	if err := db.Where("id = ?", c.Param("id")).First(&submenu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&submenu)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

type accessmenuInput struct {
	RoleID uint `json:"roleID"`
	MenuID uint `json:"menuID"`
}

// CreateAccessmenu godoc
// @Summary Create New Accessmenu`.
// @Description Creating a new Accessmenu.
// @Tags Accessmenu
// @Param Body body accessmenuInput true "the body to create a new Accessmenu"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Produce json
// @Success 200 {object} models.Accessmenu
// @Router /accessmenu/ [post]
func CreateAccessmenu(c *gin.Context) {
	// Validate input
	var input accessmenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	accessmenu := models.Accessmenu{RoleID: input.RoleID, MenuID: input.MenuID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&accessmenu)

	c.JSON(http.StatusOK, gin.H{"data": accessmenu})
}

// UpdateAccessmenu godoc
// @Summary Update Accessmenu.
// @Description Update Accessmenu by id.
// @Tags Accessmenu
// @Produce json
// @Param id path string true "Accessmenu id"
// @Param Body body accessmenuInput true "the body to update age rating category"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} models.Accessmenu
// @Router /accessmenu/{id} [patch]
func UpdateAccessmenu(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var accessmenu models.Accessmenu
	if err := db.Where("id = ?", c.Param("id")).First(&accessmenu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input accessmenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Accessmenu
	updatedInput.RoleID = input.RoleID
	updatedInput.MenuID = input.MenuID

	db.Model(&accessmenu).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": accessmenu})
}

// DeleteAccessmenu godoc
// @Summary Delete one Accessmenu.
// @Description Delete a Accessmenu by id.
// @Tags Accessmenu
// @Produce json
// @Param id path string true "Accessmenu id"
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
// @Success 200 {object} map[string]boolean
// @Router /accessmenu/{id} [delete]
func DeleteAccessmenu(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var accessmenu models.Accessmenu
	if err := db.Where("id = ?", c.Param("id")).First(&accessmenu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&accessmenu)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
