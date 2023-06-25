package controller

import (
	"AllenSo/logic"
	"AllenSo/model"
	"AllenSo/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/6/24
  @desc:
  @modified by:
**/

type pictureController struct {
}

var Picture = new(pictureController)

// GetPictureList 获取图片列表
func (pictureController) GetPictureList(c *gin.Context) {
	// 初始化结构体指定初始参数
	p := new(model.SearchDTO)
	// 获取参数，绑定query参数到结构体上
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("[controller post] get user list with invalid query params", zap.Error(err))
	}
	// 限制爬虫
	if p.Size > 20 {
		zap.L().Error("pictureController GetPictureList to much ", zap.Error(util.ErrorSpiderRequest))
		ResponseError(c, ErrorInvalidParams)
	}
	// 业务处理
	data, err := logic.GetPictureList(p)
	if err != nil {
		zap.L().Error("controller GetPictureList error ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	ResponseOK(c, data)
}
