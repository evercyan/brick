package xdraw

import (
	"fmt"
	"testing"
	"time"

	"github.com/evercyan/brick/ximg"
)

func TestCompress(t *testing.T) {
	//Compress(ximg.Read("/Users/Cyan/Downloads/img/001.jpg"), 0.5)
	//Compress(ximg.Read("/Users/Cyan/Downloads/img/005.png"), 0.5)

	//fmt.Println(ximg.Type("/Users/Cyan/Downloads/img/005.png"))
	//fmt.Println(ximg.Type("/Users/Cyan/Downloads/img/001.jpg"))
	//fmt.Println(ximg.Type("/Users/Cyan/Downloads/img/004.png"))

	//Compress("/Users/Cyan/Downloads/img/001.jpg", 0.5)

	defer func(b time.Time) {
		fmt.Println(time.Since(b).String())
	}(time.Now())

	name := "004"

	target, err := Compress("/Users/Cyan/Downloads/img/"+name+".png", 1)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	tpath := "/Users/Cyan/Downloads/img/" + name + "_c.png"
	ximg.Write(tpath, target)
	fmt.Println(tpath)
}
