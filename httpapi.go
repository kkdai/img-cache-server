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
	"strings"
	"time"
)

var tempImg map[int64][]byte

func imgDownload(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.RequestURI, "/imgs")
	id = strings.Trim(id, "?")
	io.WriteString(w, "")
}

func urlGet(w http.ResponseWriter, r *http.Request) {
	escapeUrl := strings.Trim(r.RequestURI, "/url?")
	rawUrl, err := url.QueryUnescape(escapeUrl)
	if err != nil {
		log.Println("url err:", err)
	}
	response, err := http.Get(rawUrl)
	if err != nil {
		log.Println("Error while downloading", rawUrl, "-", err)
		return
	}
	defer response.Body.Close()

	if strings.EqualFold(response.Header.Get("Content-Type"), "image/jpeg") {
		log.Println("Not image URL:", response.Header.Get("Content-Type"))
		io.WriteString(w, "Not image URL")
		return
	}

	totalBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while downloading", rawUrl, "-", err)
		return
	}

	checkInt64 := time.Now().Unix()
	if _, ok := tempImg[checkInt64]; !ok {
		checkInt64 = time.Now().Unix()
		log.Println("Coflict, do again...")
	}

	tempImg[checkInt64] = totalBody

	//QueryEscape(s string) string
	io.WriteString(w, "ret")
}

func serveHttpAPI(port string, existC chan bool) {
	go func() {
		if err, ok := <-existC; ok {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/imgs", imgDownload)
	mux.HandleFunc("/url", urlGet)
	http.ListenAndServe(":"+port, mux)
}
