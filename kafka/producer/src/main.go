package main

import(
	"fmt"
	"net/http"
	// "log"
	// "os"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	k "kafka/kafka"
	ts "kafka/types"
	sc "kafka/scraper"
)

var coins = map[string]int{"btc": 1, "eth": 2, "bnb": 3, "ada": 4, "sol": 5}


func defaultRoute(c *gin.Context){

	c.JSON(http.StatusOK, "Hola desde Kafka Publisher")
}


func publish(c *gin.Context){

	var newMsg ts.Message

    // Call BindJSON to bind the received JSON to
    if err := c.BindJSON(&newMsg); err != nil {
        c.JSON(http.StatusInternalServerError, "Bad request body")
    }else{
		err, res := k.SendMessage(newMsg, "topic")
		if err !=nil{
			c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
		}else{
			c.JSON(http.StatusOK, res)
		}
	}
	
}


func lastPriceRoute(c *gin.Context){

	c.JSON(http.StatusOK, sc.Scrap("https://www.coingecko.com/", "table.table tbody tr", "data-coin-symbol", coins))
}


func collect() {

	for {
		
		response := sc.Scrap("https://www.coingecko.com/", "table.table tbody tr", "data-coin-symbol", coins)
	
		for _, r := range response {
			newMsg := ts.Message{Time:0, Symbol:r.Symbol, Price:r.Price}
			err, res := k.SendMessage(newMsg, strings.ToLower(r.Symbol)+"-topic")
			if err !=nil{
				fmt.Println("====> ", fmt.Sprint(err))
				// return fmt.Sprint(err)
			}else{
				fmt.Println("==>", res)
				// return res
			}
		}

		time.Sleep(15*time.Second)
	}
	
}


// func initTopic() {

// 	err:=k.CreteTopic(os.Getenv("TOPIC"))
// 	if err !=nil{
// 		log.Print(err)
// 		fmt.Println(" ******* RETRYING IN 3s ******* ")
// 		time.Sleep(3*time.Second)
// 		initTopic()
// 	}else{
// 		k.ListTopics()
// 	}
// }


func main(){

	err := godotenv.Load("e.env")
	if err!=nil{
		fmt.Println("Error loading enviroment variables")
	}

	go collect()

	// fmt.Println("============================== CREATING TOPIC ==============================")
	// 	initTopic()
	// fmt.Println("============================================================================")

	router := gin.Default()
    router.Use(gin.Recovery()) // Para recuperarse de Errores y enviar un 500

    router.GET("/kafka", defaultRoute)
	router.GET("/getPrice", lastPriceRoute)
    router.POST("/send", publish)

    router.Run("0.0.0.0:8080")
}