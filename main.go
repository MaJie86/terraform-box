package main

import (
	"fmt"
	"github.com/majie86/terraform-box/api"
	"github.com/majie86/terraform-box/taskpool"
)

func main() {
	fmt.Println("terraform box is starting")

	taskpool.Init()
	api.Init(":8888")
}
