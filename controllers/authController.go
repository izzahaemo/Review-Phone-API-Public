package controllers

import (
	"dataphone/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	RoleID   int    `json:"roleID" binding:"required"`
}

type LoginResult struct {
	Message  string `json:"message"`
	Token    string `json:"token"`
	UserID   string `json:"userID"`
	RoleID   string `json:"RoleID"`
	Username string `json:"Username"`
}
type LoginFailed struct {
	Error string `json:"error" example:"username or password is incorrect."`
}

type Success struct {
	Message string `json:"message" example:"done"`
}

type FailureUser struct {
	Error string `json:"error" example:"username cannot found"`
}

type FailureForbiden struct {
	Error string `json:"error" example:"username cannot access this menu"`
}

type FailureRecord struct {
	Error string `json:"error" example:"Record not found!"`
}

// LoginUser godoc
// @Summary Login User.
// @Description Logging in to get jwt token to access admin or Member api by roles.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} LoginResult{} "Result Login"
// @Failure 400 {object} LoginFailed{} "If Login Failed"
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	u.Pict = "https://www.tenforums.com/geek/gars/images/2/types/thumb_15951118880user.png"

	iduser, role, token, err := models.LoginCheck(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	user := map[string]string{
		"userID":   strconv.Itoa(iduser),
		"roleID":   strconv.Itoa(role),
		"username": u.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})

}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access.
// @Tags Auth
// @Param Body body RegisterInput true "the body to register a user (Role ID 1 = Admin, Role ID 2 = Member)"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password
	u.RoleID = uint(input.RoleID)

	_, err := u.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": input.Username,
		"email":    input.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}
