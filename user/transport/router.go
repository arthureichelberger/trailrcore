package transport

import (
	"github.com/arthureichelberger/trailrcore/user/service"
	"github.com/arthureichelberger/trailrcore/user/transport/handler"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	SignInService service.SignInService
}

func NewUserRouter(signInService service.SignInService) UserRouter {
	return UserRouter{
		SignInService: signInService,
	}
}

func (ur UserRouter) InitRouter(g *gin.RouterGroup) {
	userGroup := g.Group("/user")
	userGroup.POST("/", handler.CreateUserHandler(ur.SignInService))
}
