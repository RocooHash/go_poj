package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go_poj/model"
	"go_poj/router"
	"log"
)

func main() {
	if err := model.Init(); err != nil {
		panic(err)
	}
	//g := gin.Default()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	router.InitRouter(g)
	log.Printf("Start to listening the incoming requests on http address: 8080")
	//log.Println(http.ListenAndServe(viper.GetString("addr"), g).Error())
	if err := g.Run(); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
