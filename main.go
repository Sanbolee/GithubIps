package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
	"time"
)

func fetch(key string, url string, ch chan bool) {
	soup.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36",
	}

	source, err := soup.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(source)
	ip := doc.Find("ul", "class", "comma-separated").Find("li").Text()
	if ip != "" {
		fmt.Println(ip + "   " + key)
	}
	time.Sleep(2 * time.Second)
	ch <- true
}

func main() {
	urls := map[string]string{
		"github.com":                     "https://github.com.ipaddress.com",
		"gist.github.com":                "https://github.com.ipaddress.com/gist.github.com",
		"assets-cdn.github.com":          "https://github.com.ipaddress.com/assets-cdn.github.com",
		"spotlights-feed.github.com":     "https://github.com.ipaddress.com/spotlights-feed.github.com",
		"raw.githubusercontent.com":      "https://githubusercontent.com.ipaddress.com/raw.githubusercontent.com",
		"gist.githubusercontent.com":     "https://githubusercontent.com.ipaddress.com/gist.githubusercontent.com",
		"cloud.githubusercontent.com":    "https://githubusercontent.com.ipaddress.com/cloud.githubusercontent.com",
		"camo.githubusercontent.com":     "https://githubusercontent.com.ipaddress.com/camo.githubusercontent.com",
		"avatars0.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars0.githubusercontent.com",
		"avatars1.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars1.githubusercontent.com",
		"avatars2.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars2.githubusercontent.com",
		"avatars3.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars3.githubusercontent.com",
		"avatars4.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars4.githubusercontent.com",
		"avatars5.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars5.githubusercontent.com",
		"avatars6.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars6.githubusercontent.com",
		"avatars7.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars7.githubusercontent.com",
		"avatars8.githubusercontent.com": "https://githubusercontent.com.ipaddress.com/avatars8.githubusercontent.com",
	}

	fmt.Println("正在输出结果，请稍等：")

	ch := make(chan bool)
	for key, url := range urls {
		go fetch(key, url, ch)
	}

	for range urls {
		<-ch
	}

	fmt.Println("输出结果完毕，请复制粘贴到hosts文件。")
	fmt.Println("Mac系统和Linux系统hosts文件位置：/etc/hosts")
	fmt.Println("Win系统hosts文件位置：c:/windows/system32/drivers/etc/hosts")
	// 防止windows下立刻退出
	time.Sleep(60 * time.Second)
}
