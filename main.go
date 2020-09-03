package main

import "gin-user/router"

func main() {
	r := router.InitRouter()
	panic(r.Run())
}
