package handler

import (
	"fmt"
	"net/http"

	"github.com/arthureichelberger/trailrcore/user/service"
	"github.com/arthureichelberger/trailrcore/user/transport/request"
	"github.com/gin-gonic/gin"
)

func LoginHandler(signinService service.SigninService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		auth, err := signinService.Login(ctx.Request.Context(), req.Email, req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.Writer.Header().Add("Authorization", fmt.Sprintf("Bearer %s", auth))
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
