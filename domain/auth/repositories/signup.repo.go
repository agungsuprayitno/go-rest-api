package repositories

import (
	"go-rest-postgres/domain/auth/requests"
	"go-rest-postgres/domain/users/models"
	errorhandlers "go-rest-postgres/error-handlers"
	"go-rest-postgres/initializers"
	"go-rest-postgres/utils"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SignUp(ctx *gin.Context) (user models.User, err error) {
	var payload *requests.SignUpRequest
	badRequestErr := errorhandlers.BadRequestError{}

	if err1 := ctx.Bind(&payload); err1 != nil {
		badRequestErr.SetError(ctx, err1.Error())
		err = err1
		return
	}

	if payload.Password != payload.PasswordConfirm {
		badRequestErr.SetError(ctx, "Passwords do not match")
	}

	hashedPassword, err2 := utils.HashPassword(payload.Password)
	if err2 != nil {
		badRequestErr.SetError(ctx, err2.Error())
		err = err2
		return
	}
	
	photo, _ := ctx.FormFile("photo")
	// Retrieve file information
	extension := filepath.Ext(photo.Filename)
	 // Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension
	photoPath := "images/profile/" + newFileName

	// ctx.SaveUploadedFile(photo, photoPath)
	if err3 := ctx.SaveUploadedFile(photo, photoPath); err != nil {
		internalServerErr := errorhandlers.InternalServerError{}

        internalServerErr.SetError(ctx, err3.Error())
		err = err3
		return
    }

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  true,
		Photo:     photoPath,
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
