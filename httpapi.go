// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var tempImg map[int64][]byte

func imgOnfly(w http.ResponseWriter, r *http.Request) {
	escapeUrl := strings.Trim(r.RequestURI, "/go?")
	w.Header().Set("Content-Type", "image/jpeg")

	rawUrl, err := url.QueryUnescape(escapeUrl)
	log.Println("Get url:", rawUrl)

	if err != nil {
		log.Println("url err:", err)
	}

	response, err := http.Get(rawUrl)
	if err != nil {
		log.Println("Error while downloading", rawUrl, "-", err)
		return
	}
	defer response.Body.Close()

	totalBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while downloading", rawUrl, "-", err)
		return
	}

	io.WriteString(w, string(totalBody))
}

func imgDownload(w http.ResponseWriter, r *http.Request) {
	idStr := strings.Trim(r.RequestURI, "/imgs?")
	id := strings.TrimRight(idStr, ".jpg")
	w.Header().Set("Content-Type", "image/jpeg")

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("input not a number for id:", id)
		return
	}
	if v, ok := tempImg[i]; ok {
		io.WriteString(w, string(v))
	}
	io.WriteString(w, "No data")
}

func urlGet(w http.ResponseWriter, r *http.Request) {
	escapeUrl := strings.Trim(r.RequestURI, "/url?")
	rawUrl, err := url.QueryUnescape(escapeUrl)
	log.Println("Get url:", rawUrl)

	if err != nil {
		log.Println("url err:", err)
	}

	response, err := http.Get(rawUrl)
	defer response.Body.Close()

	if err != nil {
		log.Println("Error while downloading", rawUrl, "-", err)
		return
	}

	if strings.EqualFold(response.Header.Get("Content-Type"), " image/jpeg ") {
		log.Println("Not image URL:", response.Header.Get("Content-Type"))
		io.WriteString(w, "Not image URL")
		return
	}

	totalBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while downloading", rawUrl, "-", err)
		return
	}

	checkInt64 := time.Now().UnixNano()
	if _, ok := tempImg[checkInt64]; ok {
		checkInt64 = time.Now().UnixNano()
		log.Println("Coflict, do again...")
	}

	tempImg[checkInt64] = totalBody

	str := strconv.FormatInt(checkInt64, 10)
	io.WriteString(w, str)
}

func serveHttpAPI(port string, existC chan bool) {
	go func() {
		if err, ok := <-existC; ok {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/go", imgOnfly)
	mux.HandleFunc("/imgs", imgDownload)
	mux.HandleFunc("/url", urlGet)
	http.ListenAndServe(":"+port, mux)
}
