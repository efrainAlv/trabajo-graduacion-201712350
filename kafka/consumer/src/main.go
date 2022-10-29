package main

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/joho/godotenv"
	k "kafka/kafka"
	an "kafka/analyzer"
	ts "kafka/types"
)


func defaultRoute(c *gin.Context){

	c.JSON(http.StatusOK, "Hola desde Kafka Consumer")
}


func storeHistoryPrice(c *gin.Context){

	var newHistoryPrice ts.HistoryPrice

	if err := c.BindJSON(&newHistoryPrice); err != nil {
        c.JSON(http.StatusInternalServerError, "Bad request body")
    }else{
		
		res, err := an.OpenFile(newHistoryPrice)

		if err != nil{
			c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
		}else{
			c.JSON(http.StatusOK, res)
		}
	}
}


func StartConsumer(){

	fmt.Println("=========== CONSUMER LISTENING ===========")

	finish:=k.Consume()
	for finish=="done"{
		fmt.Println(" ***** RETRYING IN 3s ***** ")
		time.Sleep(5*time.Second)
		finish=k.Consume()
	}

	fmt.Println("")
}


func main(){
	
	err := godotenv.Load("e.env")
	if err!=nil{
		fmt.Println("Error loading enviroment variables")
	}
	
	// fmt.Println("=========================== INITIALIZING COSUMER ===========================")
	// go StartConsumer()
	// fmt.Println("============================================================================")

	router := gin.Default()
    router.Use(gin.Recovery()) // Para recuperarse de Errores y enviar un 500

    router.GET("/kafka", defaultRoute)
	router.POST("/getPrice", storeHistoryPrice)
    // router.POST("/send", publish)

    router.Run("0.0.0.0:8080")
}