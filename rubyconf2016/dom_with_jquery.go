package main

import "github.com/gopherjs/jquery"

func main() {
	jquery.NewJQuery().Ready(func() {
		jquery.NewJQuery("body").SetCss("background-color", "red")
		jquery.NewJQuery("body").Append("<img src=\"gopher.png\"/>")
	})
}
