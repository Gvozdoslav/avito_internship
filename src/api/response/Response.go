package response

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (e errorResponse) Error() string {
	return e.Message
}

type statusResponse struct {
	Message string `json:"status"`
}

func NewErrorResponse(context *gin.Context, statusCode int, message string) {
	log.Println(message)
	if err := context.AbortWithError(statusCode, errorResponse{message}); err != nil {
		log.Fatalf("error occured: %s", err)
	}
}
