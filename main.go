package main

import (
	"fmt"
	"github.com/MaJie86/terraform-box/api"
	"github.com/MaJie86/terraform-box/cmd"
)

func main() {
	fmt.Println("abc")

	var params = []string{}
	_ = cmd.Exec("test.log", "d://", "ipconfig", params)

	api.Create()
}
