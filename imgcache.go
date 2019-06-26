package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ImgContent struct {
	ImgType string
	Content []byte
}

var ImgMap map[int64]ImgContent

var spportType = []struct {
	ImgType     string
	ContentType string
}{
	{"JPG", "image/jpeg"},
	{"PNG", "image/png"},
}

//GetImgCache : Get image cache image content map ID ("1234-JPG" or "4567-PNG")
func GetImgCache(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		response.Body.Close()
		return "", err
	}

	defer response.Body.Close()

	supportType := ""

	for _, v := range spportType {
		if strings.EqualFold(response.Header.Get("Content-Type"), v.ContentType) {
			log.Println("IMG is", v.ImgType)
			supportType = v.ImgType
			break
		}
	}
	log.Println("url:", url, " contains:", response.Header.Get("Content-Type"))
	if len(spportType) == 0 {
		log.Println("Not image URL:", url)
		return "", fmt.Errorf("Not image URL:%s", url)
	}

	totalBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return "", err
	}

	checkInt64 := time.Now().UnixNano()
	if _, ok := ImgMap[checkInt64]; ok {
		checkInt64 = time.Now().UnixNano()
		log.Println("Coflict, do replace...")
	}

	ImgMap[checkInt64] = ImgContent{ImgType: supportType, Content: totalBody}
	return strconv.FormatInt(checkInt64, 10), nil
}

//GetImageContent :
func GetImageContent(idURL string) (ImgContent, error) {
	id := strings.TrimRight(idURL, ".jpg")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("input not a number for id:", id)
		return ImgContent{}, fmt.Errorf("input not a number for id:%s", id)
	}

	if _, ok := ImgMap[i]; !ok {
		return ImgContent{}, fmt.Errorf("ImgMap doesn't have this data in :%d", i)
	}

	return ImgMap[i], nil
}
