package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func main() {
	browser := rod.New().Connect()

	defer browser.Close()

	page := browser.Timeout(time.Minute).Page("https://github.com")

	page.Window(0, 0, 1200, 600)
	time.Sleep(time.Millisecond * 10)

	page.Element("input").Input("git").Press(input.Enter)

	text := page.Element(".codesearch-results p").Text()

	fmt.Println(text)

	fmt.Println(len(page.Elements("input")))

	page.Eval(`console.log("hello world")`)

	fmt.Println(page.Eval(`(a, b) => a + b`, 1, 2).Int())

	fmt.Println(page.Element("title").Eval(`this.innerText`).String())

	err := rod.Try(func() {
		page.Timeout(time.Second * 10).Element("body")
	})
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("after 10 seconds, the element is still not rendered")
	}
}
