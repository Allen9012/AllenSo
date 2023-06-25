package util

import "errors"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/6/24
  @desc:
  @modified by:
**/

var (
	ErrorOutRange      = errors.New("图片不存在")
	ErrorIdNotExist    = errors.New("用户不可用")
	ErrorTokenNotExist = errors.New("Token 不存在")
	ErrorPostExist     = errors.New("文章已存在")
	ErrorUserExist     = errors.New("用户已存在")
	ErrorSpiderRequest = errors.New("爬虫请求失败")
)
