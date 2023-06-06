package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func requireParamInt(c *gin.Context, param string) int {
	value, err := strconv.Atoi(c.Query(param))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Missing query param `%s`", param))
	}
	return value
}
