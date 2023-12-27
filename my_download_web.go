package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Web struct {
	Url  string
	Body string
	Len  int
}

func downloadWeb(i int, url string, webChannel chan Web) {
	response, err := http.Get(url)

	// step-1，先檢查Get()本身有沒有問題
	if err != nil {
		log.Fatal(fmt.Sprintf("Goroutine[%d] getting url: \"%s\" encountered exception!\n", i, url))
	}

	defer response.Body.Close()

	// step-2，再檢查response本身有沒有問題
	if response.StatusCode != http.StatusOK {
		log.Fatal(fmt.Sprintf("Goroutine[%d] http response url: \"%s\" encountered exception!\n", i, url))
	}

	// step-3，嘗試剖析body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(fmt.Sprintf("Goroutine[%d] parsing response body: \"%s\" encountered exception!\n", i, url))
	}

	// step-4，收納剖析結果
	web := Web{Url: url, Body: string(body), Len: len(string(body))}
	webChannel <- web
}

func main() {
	urls := []string{
		"https://example.com/", "https://go.dev/", "https://go.dev/doc/",
	}

	webChannel := make(chan Web)

	// step-1，開始平行下載
	for i, url := range urls {
		go downloadWeb(i, url, webChannel)
	}

	// step-2，依序印出結果
	for i := 0; i < len(urls); i++ {
		web := <-webChannel
		fmt.Printf("[%s][%d]\n", web.Url, web.Len)
	}
}
