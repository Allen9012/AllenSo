package logic

import (
	"AllenSo/job/once"
	"AllenSo/model"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/6/23
  @desc:
  @modified by:
**/

func GetPictureList(p *model.SearchDTO) ([]*model.PictureVO, error) {
	return once.GetPicture(p)
}
