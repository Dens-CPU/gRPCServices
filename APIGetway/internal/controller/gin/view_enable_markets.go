package gin

import (
	"net/http"

	spotservicedto "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/spot_service"
	"github.com/gin-gonic/gin"
)

func (api *GinAPI) ViewEnableMarkets(c *gin.Context) {
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

	uid, exist := c.Get("x-user-id")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "user id missing",
		})
		c.Abort()
		return
	}

	userId, ok := uid.(string)
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "error casting user_id to string type",
		})
		return
	}

	input := spotservicedto.Input{
		UserID:   userId,
		UserRole: role,
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	markets, err := api.service.ViewEnableMarkets(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "error getting available markets",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"available markets": markets,
	})
}
