package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("Hello, main!")

	// etapa1 OMIT

	var p *plugin.Plugin
	var err error
	p, err = plugin.Open("./plugin/plugin.so")
	if err != nil {
		fmt.Println(err)
		return
	}

	// etapa2 OMIT

	var s plugin.Symbol
	s, err = p.Lookup("Hello")
	if err != nil {
		fmt.Println(err)
		return
	}

	Hello := s.(func() error)

	// etapa3 OMIT

	err = Hello()
	if err != nil {
		fmt.Println(err)
		return
	}

}
