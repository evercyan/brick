package main

import (
	"fmt"
	"time"

	"github.com/evercyan/brick/xcli/xloading"
)

func main() {
	fmt.Println("-------------------------------- basic")

	l1 := xloading.New("message success").Start()
	time.Sleep(2 * time.Second)
	l1.Success()

	l2 := xloading.New("message fail").Start()
	time.Sleep(2 * time.Second)
	l2.Fail()

	fmt.Println("\n-------------------------------- style")

	for style := xloading.Style1; style <= xloading.Style5; style++ {
		l := xloading.New(style.Elements()...).Style(style).Start()
		time.Sleep(2 * time.Second)
		l.Success()
	}

	fmt.Println("")

	fmt.Println("\n-------------------------------- symbol")

	for symbol := xloading.Symbol1; symbol <= xloading.Symbol5; symbol++ {
		l := xloading.New(symbol.Elements()...).Symbol(symbol).Start()
		time.Sleep(2 * time.Second)
		l.Success()
	}
}
