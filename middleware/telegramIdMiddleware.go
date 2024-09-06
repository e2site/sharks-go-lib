package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const CONTEXT_TG_NAME = "telegram-id"

type TGMiddleware struct {
	HttpHeaderName string
}

func (tgMiddleware *TGMiddleware) TGMiddleware(c *gin.Context) {
	telegramId := c.GetHeader(tgMiddleware.HttpHeaderName)
	var tgID int64
	if telegramId == "" {
		c.Status(404)
		return
	} else {
		tg, err := strconv.ParseInt(telegramId, 10, 64)
		if err != nil {
			c.Status(404)
			return
		} else {
			tgID = tg
		}
	}
	c.Set(CONTEXT_TG_NAME, tgID)

	c.Next()
}
