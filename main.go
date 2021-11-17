package main

import (
	"fmt"
	"github.com/majie86/terraform-box/api"
)

func main() {
	fmt.Println("terraform box is starting")

	api.Init()
}
