package handler

import (
	"errors"
	"net/http"

	"github.com/arthureichelberger/trailrcore/user/exception"
	"github.com/arthureichelberger/trailrcore/user/service"
	"github.com/arthureichelberger/trailrcore/user/transport/request"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(signInService service.SignInService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.CreateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		user, err := signInService.CreateUser(ctx.Request.Context(), req.Email, req.Password, req.PasswordConfirmation)
		if err != nil {
			switch {
			case errors.As(err, new(exception.CouldNotValidateEmailError)), errors.As(err, new(exception.CouldNotValidatePasswordError)):
				ctx.JSON(http.StatusBadRequest, gin.H{})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"id": user.ID,
		})
	}
}
