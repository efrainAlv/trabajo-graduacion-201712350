package main

import(
	"fmt"
	"webscraper/collector"
)



func main(){

	coins := map[string]int{"btc": 1, "eth": 2, "bnb": 3, "ada": 4, "sol": 5}
	response := collector.Scrap("https://www.coingecko.com/", "table.table tbody tr", "data-coin-symbol", coins)

	fmt.Println(response)
}


