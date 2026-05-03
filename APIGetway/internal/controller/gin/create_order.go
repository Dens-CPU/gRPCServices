package gin

import (
	"fmt"
	"net/http"

	orderdomain "github.com/DencCPU/gRPCServices/APIGetway/internal/domain/order"
	"github.com/gin-gonic/gin"
)

func (api *GinAPI) CreateOrderHandler(c *gin.Context) {
	var order orderdomain.OrderInfo

	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, exist := c.Get("x-user-role")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "user role missing",
		})
		c.Abort()
		return
	}

	role, ok := r.(string)
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "error casting user_role to string type",
		})
		return
	}

	switch role {
	case "basic":
		order.UserRole = orderdomain.USER_ROLE_BASIC_USER
	case "premium":
		order.UserRole = orderdomain.USER_ROLE_PREMIUM_USER
	}

	uid, exist := c.Get("x-user-id")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "user id missing",
		})
		c.Abort()
		return
	}

	order.UserId, ok = uid.(string)
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "error casting user_id to string type",
		})
		return
	}
	fmt.Println(order)
	output, err := api.service.CreateOrder(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order_id":     output.OrderId,
		"order_status": output.OrderStatus,
	})
}
