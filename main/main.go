package main

import (
	"goworkshop/web"
	"goworkshop/persistence"
)

func main() {
	if err := persistence.InitDB(); err != nil{
		panic(err)
	}
	web.StartServer()
}

