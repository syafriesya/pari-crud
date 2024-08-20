package crud

import (
	"net/http"
	"pari/internal/domain"

	"github.com/gin-gonic/gin"
)

func (cc CrudController) Put(c *gin.Context) {

	id := c.Query("id")

	var updatedItem domain.Items

	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	err := cc.crudUsecase.PutItem(ctx, id, updatedItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Failed to update item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Item updated successfully",
	})
}
