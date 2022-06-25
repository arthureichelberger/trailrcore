package transport

import (
	"github.com/arthureichelberger/trailrcore/user/service"
	"github.com/arthureichelberger/trailrcore/user/transport/handler"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	SignupService service.SignupService
	SigninService service.SigninService
}

func NewUserRouter(signupService service.SignupService, signinService service.SigninService) UserRouter {
	return UserRouter{
		SignupService: signupService,
		SigninService: signinService,
	}
}

func (ur UserRouter) InitRouter(g *gin.RouterGroup) {
	userGroup := g.Group("/user")
	userGroup.POST("/", handler.CreateUserHandler(ur.SignupService))
	userGroup.POST("/auth", handler.LoginHandler(ur.SigninService))
}
