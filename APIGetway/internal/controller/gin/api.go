package gin

import (
	"context"

	"github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/jwt"
	userdomain "github.com/DencCPU/gRPCServices/APIGetway/internal/domain/user"
	"github.com/gin-gonic/gin"
)

type Service interface {
	RegistrationUser(ctx context.Context, newUser userdomain.User) (jwt.PairToken, error)
}
type GinAPI struct {
	r       *gin.Engine
	service Service
}

func NewGinAPI(service Service) GinAPI {
	api := GinAPI{}
	api.r = gin.Default()
	api.service = service
	api.endpoints()
	return api
}

func (api *GinAPI) endpoints() {
	api.r.POST("/order", api.CreateOrderHandler)
	api.r.GET("/status/{id}")
	api.r.GET("/realtime_status/{id}")
	api.r.POST("/user/reg", api.RegistrationUser)
}

func (api *GinAPI) Router() *gin.Engine {
	return api.r
}
