package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quandoan/shorten_url/db"
	"github.com/quandoan/shorten_url/modules/shorten/transport"
)



func main() {
	log.Println("Shortening URL app!!")
	runService()
}

func runService() error {
	r := gin.Default()

	r.POST("/create", transport.CreateVirtualLink(db.NewMemDb(), r))

	return r.Run()
}
