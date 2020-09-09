package main

import (
	"fmt"

	"github.com/GreatGodApollo/gospacebin"
)

func main() {
	doc, err := gospacebin.NewClient("https://api.spaceb.in").GetDocument("hxUzfoVy")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(doc, err)
}
