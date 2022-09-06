package utils

import (
	"context"
	cdp "github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"time"
)

func AutoScreenshot(urlstr string, selector string, height int, width int) {
	// 创建新的cdp上下文
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()
	cdp.WindowSize(width, height)
	var buf []byte
	// 需要截图的元素，支持CSS selector以及XPath query
	if err := cdp.Run(ctx, elementScreenshot(urlstr, selector, &buf)); err != nil {
		log.Fatal(err)
	}
	// 写入文件
	if err := ioutil.WriteFile("360_so.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// 截图方法
func elementScreenshot(urlstr, sel string, res *[]byte) cdp.Tasks {
	return cdp.Tasks{
		// 打开url指向的页面
		cdp.Navigate(urlstr),

		// 等待待截图的元素渲染完成
		cdp.WaitVisible(sel, cdp.ByID),
		// 也可以等待一定的时间
		cdp.Sleep(time.Duration(10) * time.Second),

		// 执行截图
		cdp.Screenshot(sel, res, cdp.NodeVisible, cdp.ByID),
	}
}
