package handler

import (
	"congo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) filter(c *gin.Context) {
// 	fmt.Print("Yeeees\n")
// }

type getAllListsResponse struct {
	Data []congo.Account `json:"data"`
}

func (h *Handler) getAll(c *gin.Context) {

	lists, err := h.services.AccountsList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) filterSex(c *gin.Context) {
	sex, isNot := c.GetQuery("sex")
	if !isNot {
		fmt.Println("not found param Sex")
	}

	lists, err := h.services.AccountsList.FilterSex(sex)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}
