package main

import (
	"strconv"
	"testing"
)

func init() {
	ImgMap = make(map[int64][]byte)
}
func TestGetImgCache(t *testing.T) {

	var testingUrls = []string{
		"http://asms.coa.gov.tw/amlapp/upload/pic/c74571c1-0088-4d9f-a853-1177516d627b_org.jpg",
		"http://asms.coa.gov.tw/amlapp/upload/pic/acd5caf0-1c33-43e5-a97e-b670e15adc6f_org.JPG",
		"http://asms.coa.gov.tw/amlapp/upload/pic/44632381-d1a0-4aa9-af2e-334f157f9847_org.JPG",
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
