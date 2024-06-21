package controller

import (
	"chatim/dao/sqlites"
	models "chatim/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpText(c *gin.Context) {
	p := new(models.UpText)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("uptext json err", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	if p.Text == "" {
		zap.L().Error("uptext json err", zap.String("error", "Missing text field"))
		c.JSON(400, gin.H{"error": "Missing text field"})
		return
	}
	textValue := p.Text
	sqlites.Text(textValue)
	c.JSON(200, gin.H{"Success": "200"})
}
