package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/chromedp/chromedp/kb"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), // debug使用
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	var b2 []byte
	if err := chromedp.Run(ctx,

		// reset
		chromedp.Emulate(device.Reset),
		chromedp.Navigate("https://blog.csdn.net/stubbornness1219/article/details/53446121"),
		chromedp.KeyEvent(kb.F12),

		chromedp.Sleep(time.Second*2),
		chromedp.FullScreenshot(&b2, 100),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("baidu_PC.png", b2, 0777); err != nil {
		log.Fatal(err)
	}
}
