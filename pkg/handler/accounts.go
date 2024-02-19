package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) filter(c *gin.Context) {
	fmt.Print("Yeeees\n")
}
