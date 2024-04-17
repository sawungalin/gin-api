package main

import "gin-api/routers"

func main() {
	routersInit := routers.InitRouter()
	routersInit.Run(":8080")
}
