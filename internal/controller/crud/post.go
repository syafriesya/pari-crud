package crud

import (
	"net/http"
	"pari/internal/domain"

	"github.com/gin-gonic/gin"
)

func (cc CrudController) Post(c *gin.Context) {
	ctx := c.Request.Context()
	var newItem domain.Items

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	err := cc.crudUsecase.PostItem(ctx, newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Failed to create item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Item created successfully",
	})

}
