package scraper

import(
	"strings"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)


type Coin struct {
    Symbol string
    Price  string
}


func Scrap(url, query, target string, coins map[string]int) []Coin {

	c := colly.NewCollector()

	foundCount := 0
	collected := []Coin{}

	c.OnHTML(query, func(e *colly.HTMLElement) {
		
		if foundCount<len(coins) {
			for _, node := range e.DOM.Nodes {
				if coin := searchElement(target, Coin{Symbol:"", Price:""}, node); coins[coin.Symbol]>0 {
					collected = append(collected, coin)
					foundCount++;
				}
			}
		}else{
			e.Request.Abort()
		}
	})

	c.Visit(url)

	return collected
}


func searchElement(target string, found Coin, node *html.Node) Coin {

	if node.Type == html.ElementNode {

		if node.FirstChild!=nil {

			var s *html.Node

			if node.FirstChild.Type == html.ElementNode {
				s=node.FirstChild
				
			}else {

				for _, attr:= range node.Attr {
					if attr.Key == target {
						
						found.Symbol = strings.ToLower(attr.Val)
						found.Price = strings.Replace(node.FirstChild.Data, ",", "", -1)
						break
					}
				}

				s=node.FirstChild.NextSibling
			}

			for s!=nil{
				found = searchElement(target, found, s)
				s = s.NextSibling
			}

		}
		
		return found
	
	}
	
	return found
}