package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getInt64Param(c *gin.Context, key string) (int64, bool) {
	param, err := strconv.ParseInt(c.Param(key), 10, 64)
	if err != nil {
		return 0, false
	}

	return param, true
}
