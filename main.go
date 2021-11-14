package main

import (
	"fmt"
	"github.com/majie86/terraform-box/api"
	"github.com/majie86/terraform-box/cmd"
)

func main() {
	fmt.Println("abc")

	var params = []string{}
	_ = cmd.Exec("test.log", "d://", "ipconfig", params)

	lines, total := cmd.ReadLog(6)
	for i := lines.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println(total)

	api.Create()
}
