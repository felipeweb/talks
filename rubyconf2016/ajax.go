package main

import (
	"net/http"
	"io/ioutil"
	"github.com/felipeweb/gopher-utils"
	"honnef.co/go/js/dom"
)

func main() {
	res, err := http.Get("https://cursos.alura.com.br/api/cursos")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	d := dom.GetWindow().Document()
	p := d.CreateElement("p").(*dom.HTMLParagraphElement)
	p.SetTextContent(gopher_utils.ToStr(bytes))
	content := d.GetElementByID("content").(*dom.HTMLDivElement)
	content.AppendChild(p)
}
