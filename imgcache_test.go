package main

import (
	"strconv"
	"testing"
)

func init() {
	ImgMap = make(map[int64]ImgContent)
}
func TestGetImgCacheJPG(t *testing.T) {

	var testingUrls = []string{
		"https://www.coa.gov.tw/files/masthead/24/A01_1.jpg",
	}

	for _, v := range testingUrls {
		r, err := GetImgCache(v)
		if err != nil {
			t.Errorf("GetImgCache failed with %x", err)
		}
		retMapID, err := strconv.ParseInt(r, 10, 64)
		if err != nil {
			t.Errorf("GetImgCache strconv failed with %x", err)
		}
		if _, ok := ImgMap[retMapID]; !ok {
			t.Errorf("GetImgCache map is not exist")
		}
	}
}

func TestGetImgCachePNG(t *testing.T) {

	var testingUrls = []string{
		"http://lineat.blogimg.jp/tw/imgs/f/3/f36f199a.png",
	}

	for _, v := range testingUrls {
		r, err := GetImgCache(v)
		if err != nil {
			t.Errorf("GetImgCache failed with %x", err)
		}
		retMapID, err := strconv.ParseInt(r, 10, 64)
		if err != nil {
			t.Errorf("GetImgCache strconv failed with %x", err)
		}
		if _, ok := ImgMap[retMapID]; !ok {
			t.Errorf("GetImgCache map is not exist: %d", retMapID)
		}
	}
}
