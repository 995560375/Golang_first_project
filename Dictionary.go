package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc         int        `json:"rc"`
	Wiki       struct{}   `json:"wiki"`
	Dictionary Dictionary `json:"dictionary"`
}

type Dictionary struct {
	Prons        Pron       `json:"prons"`
	Explanations []string   `json:"explanations"`
	Synonym      []string   `json:"synonym"`
	Antonym      []string   `json:"antonym"`
	WqxExample   [][]string `json:"wqx_example"`
	Entry        string     `json:"entry"`
	Type         string     `json:"type"`
	Related      []string   `json:"related"`
	Source       string     `json:"source"`
}

type Pron struct {
	EnUs string `json:"en-us"`
	En   string `json:"en"`
}

func main() {
	fmt.Println("Please input English word, input wk to exit")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("gg", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		if input == "wk" {
			break
		} else {
			search(input)
		}
	}
}

func search(word string) {
	client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewBuffer(buf)
	req, err := http.NewRequest("POST", "https://lingocloud.caiyunapp.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "_gcl_au=1.1.648653805.1711174569; _ga=GA1.2.869905283.1711174569; _gid=GA1.2.1991306244.1711174570; _gat_gtag_UA_185151443_2=1; _ga_65TZCJSDBD=GS1.1.1711174569.1.1.1711174626.0.0.0; _ga_R9YPR75N68=GS1.1.1711174569.1.1.1711174626.3.0.0")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("device-id", "24f57ffece36b3b5b2920f0930e699fc")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad status code:", resp.StatusCode, bodyText)
	}
	var dicResponse DictResponse
	err = json.Unmarshal(bodyText, &dicResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", dicResponse.Dictionary.Prons.En, "US:", dicResponse.Dictionary.Prons.EnUs)
	for _, item := range dicResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}
