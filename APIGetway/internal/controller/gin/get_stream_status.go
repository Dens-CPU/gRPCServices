package gin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	orderdto "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/order_service"
	"github.com/gin-gonic/gin"
)

func (api *GinAPI) GetStreamStatus(c *gin.Context) {
	var orderInfo orderdto.GetInput

	//Get order info
	//Get orderID from query row
	orderInfo.OrderId = c.Query("id")
	if orderInfo.OrderId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unfinde orderID in query row",
		})
		return
	}

	//Get userId from context
	uid, exist := c.Get("x-user-id")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "user_id missing",
		})
		return
	}
	user_id, ok := uid.(string)
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "error casting user_id to string type",
		})
		return
	}
	orderInfo.UserId = user_id

	msgChan := make(chan orderdto.StreamOutput, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(msgChan)
		err := api.service.GetStreamStatus(c.Request.Context(), orderInfo, msgChan)
		select {
		case <-c.Request.Context().Done():
			return
		default:
			if err != nil {
				errChan <- err
				return
			}
		}

	}()

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	var update struct {
		Status     string `json:"status"`
		UpdateTime string `json:"update_time"`
	}

	c.Stream(func(w io.Writer) bool {
		select {
		case err := <-errChan:
			fmt.Fprintf(w, "event: error\ndata: %s\n\n", err.Error())
			return false
		case <-c.Request.Context().Done():
			return false
		case msg, ok := <-msgChan:
			if !ok {
				fmt.Fprintf(w, "event: message\ndata: transfer completed\n\n")
				return false
			}
			update.UpdateTime = msg.UpdateTime.Format("2006-01-02 15:04:05")
			update.Status = msg.OrderStatus
			data, _ := json.Marshal(update)
			fmt.Fprintf(w, "event: message\ndata: %s\n\n", data)
			return true
		}
	})
}
