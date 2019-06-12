package main

import (
	"delycarapi/middleware"
	_ "delycarapi/middleware"
	"delycarapi/repository"
	_ "delycarapi/repository"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	_ "time"

	"github.com/gin-gonic/gin"
)

func main() {

	n, err := strconv.Atoi(os.Getenv("PERIOD"))
	if err != nil {
		n = 5
	}
	fmt.Println(n)

	go func() {
		for {
			repository.PostData()
			time.Sleep(time.Second * time.Duration(n))
			break
		}
	}()

	r := gin.Default()
	r.GET("/near", middleware.GetCarsNearByG)
	if err := r.Run(":8081"); err != nil {
		log.Fatalln(err)
	}
}
