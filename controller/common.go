package controller

import (
	"AllenSo/util"
	"github.com/gin-gonic/gin"
)

func getUserId(c *gin.Context) (int64, error) {
	value, exist := c.Get(util.KeyUserId)
	if !exist {
		return -1, util.ErrorIdNotExist
	}
	userId, ok := value.(int64)
	if !ok {
		return -1, util.ErrorIdNotExist
	}
	return userId, nil
}
