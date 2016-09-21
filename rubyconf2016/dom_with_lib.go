package main

import "honnef.co/go/js/dom"

func main() {
	dom.GetWindow().AddEventListener("load", false, func(e dom.Event) {
		go func() {
			body := dom.GetWindow().Document().GetElementByID("body")
			body.SetAttribute("style", "background-color: red")
			img := dom.GetWindow().Document().CreateElement("img")
			img.SetAttribute("src", "gopher.png")
			body.AppendChild(img)
		}()
	})
}