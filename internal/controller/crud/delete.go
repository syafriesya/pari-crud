package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cc CrudController) Delete(c *gin.Context) {
	id := c.Query("id")

	ctx := c.Request.Context()
	if err := cc.crudUsecase.DeleteItem(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Failed to delete item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Item deleted successfully",
	})
}
