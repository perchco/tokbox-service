package main

import (
	"github.com/perchco/tokbox-service/tokbox"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
	"os"
	"fmt"
)

func main() {
	//setup the api to use your credentials
	tb := tokbox.New(os.Getenv("TOKBOX_API"),os.Getenv("TOKBOX_KEY"))

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
	    Origins:        "*",
	    Methods:        "GET, PUT, POST, DELETE",
	    RequestHeaders: "Origin, Authorization, Content-Type",
	    ExposedHeaders: "",
	    MaxAge: 50 * time.Second,
	    Credentials: true,
	    ValidateHeaders: false,
	}))


	router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "all": "ok",
        })
    })

    router.POST("/", func (c *gin.Context) {
		//create a session
		session, err := tb.NewSession("", true) //no location, peer enabled

		//create a token
		token, err := session.Token("publisher", "", 86400) //type publisher, no connection data, expire in 24 hours

		if err != nil{
			fmt.Println(err)
			c.JSON(200, gin.H{"not": "ok","token":""})
		} else {
			c.JSON(200, gin.H{"all": "ok","token":token})
		}

		
	})

	

	router.Run()//8080 http port, localhost
}

