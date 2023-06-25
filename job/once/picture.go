package once

import (
	"AllenSo/model"
	"AllenSo/util"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
)

func GetPicture(p *model.SearchDTO) ([]*model.PictureVO, error) {
	var escape, urlStr string
	// 使用url包自动url转换
	if p.Text == "" {
		escape = url.QueryEscape("壁纸")
	} else {
		escape = url.QueryEscape(p.Text)
	}
	// 拼接url
	urlStr = "https://cn.bing.com/images/search?q=" + escape + "&first=1"
	c := colly.NewCollector()

	// 绑定 字段 m 中的 json 数据
	var picMap map[string]string

	var bingPicture []*model.PictureVO
	// 对于是一个HTML的文件第一个参数填写对应的css选择器
	c.OnHTML(".iuscp.isv", func(e *colly.HTMLElement) {
		// 拿到图片属性
		pic := e.ChildAttr(".iusc", "m")
		// 获取标题
		title := e.ChildAttr(".inflnk", "aria-label")

		err := json.Unmarshal([]byte(pic), &picMap)
		if err != nil {
			fmt.Println("Unmarshal error ", err)
			return
		}

		bingPicture = append(bingPicture, &model.PictureVO{
			Title:   title,
			Picture: picMap["murl"],
			Purl:    picMap["purl"],
		})
	})
	// c.Visit() 正式启动网页访问。
	err := c.Visit(urlStr)
	if err != nil {
		fmt.Println("c.Visit error ", err)
		return nil, err
	}
	// 获取某一页的图片
	start := (p.Page - 1) * p.Size
	end := start + p.Size
	// 超过所有图片的数据
	if end >= len(bingPicture) {
		return nil, util.ErrorOutRange
	}

	return bingPicture[start:end], nil
}
