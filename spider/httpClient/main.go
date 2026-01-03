package main

// goquery 使用
import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var client = http.Client{}

func main() {

	file, _ := os.Create("某人文章.txt")
	defer file.Close()
	for i := 1; i <= 60; i++ {
		doRequest(i, file)

		time.Sleep(time.Second * 2)
	}

}

func doRequest(i int, file *os.File) {

	page := strconv.Itoa(i)
	request, _ := http.NewRequest("GET", "url?page="+page, nil)
	request.Header = map[string][]string{"User-Agent": {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36"}}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	baseUri := "baseurl"

	doc.Find(".clearfix").Each(func(i int, doc *goquery.Selection) {

		articlDoc := doc.Find(".wz02").Find("h3 a")
		url, _ := articlDoc.Attr("href")
		url = baseUri + url
		title := articlDoc.Text()

		tmpDoc := doc.Find(".wz02a")
		time := tmpDoc.Text()
		author := tmpDoc.Find("a").Text()

		time = strings.TrimSpace(strings.ReplaceAll(time, author, ""))

		if author == "作者" {
			file.WriteString(author + "  " + title + "  " + url + "  " + time + "\n")

		}

	})

}
