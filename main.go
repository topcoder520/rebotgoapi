package main

import (
	"github.com/topcoder520/rebotgoapi/router"
)

func main() {
	r := router.GetRouter()
	r.Run()
}
