package repositories

import (
	"go-rest-postgres/domain/auth/requests"
	"go-rest-postgres/domain/users/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"
	"go-rest-postgres/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) (user models.User, err error) {
	var payload *requests.SignUpRequest
	badRequestErr := errorhandlers.BadRequestError{}

	if err1 := ctx.ShouldBindJSON(&payload); err1 != nil {
		badRequestErr.SetError(ctx, err.Error())
		err = err1
		return
	}

	if payload.Password != payload.PasswordConfirm {
		badRequestErr.SetError(ctx, "Passwords do not match")
	}

	hashedPassword, err2 := utils.HashPassword(payload.Password)
	if err2 != nil {
		badRequestErr.SetError(ctx, err.Error())
		err = err2
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  true,
		Photo:     payload.Photo,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		conflictErr := errorhandlers.ConflictError{}
		conflictErr.SetError(ctx, "User with that email already exists")
		err = result.Error
		return
	} else if result.Error != nil {
		badGatewayErr := errorhandlers.ConflictError{}
		badGatewayErr.SetError(ctx, "Something bad happened, please try again later")
		err = result.Error
		return
	}

	return newUser, err
}
