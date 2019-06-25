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

var ImgMap map[int64][]byte

var spportType = []struct {
	ImgType     string
	ContentType string
}{
	{"JPG", " image/jpeg "},
	{"PNG", " image/png "},
}

//GetImgCache : Get image cache image content map ID
func GetImgCache(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		response.Body.Close()
		return "", err
	}

	defer response.Body.Close()

	// support := false
	// for _, v := range spportType {

	// }
	log.Println("url:", url, " contains:", response.Header.Get("Content-Type"))
	if strings.EqualFold(response.Header.Get("Content-Type"), " image/jpeg ") {
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

	ImgMap[checkInt64] = totalBody
	return strconv.FormatInt(checkInt64, 10), nil
}
