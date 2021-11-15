package main

import (
	"fmt"
	"github.com/majie86/terraform-box/api"
	"github.com/majie86/terraform-box/cmd"
	"github.com/majie86/terraform-box/pool"
)

func main() {
	fmt.Println("terraform box is starting")

	var params = []string{}
	_ = cmd.Exec("test.log", "d://", "ipconfig", params)

	pool.Init()
	api.Init()
}
