package main 

import (
	"io/ioutil" 
	"log" 
	"net/http" 
	"encoding/json"
	"strings"
)

type result struct {
	Name string
	Description string
}

type itemListElement struct {
	ResultScore float32
	Result result
}

type gkb struct {
	ItemListElement []itemListElement
}

func main() { 

	resp, err := http.Get("https://kgsearch.googleapis.com/v1/entities:search?query=tecnologico+de+monterrey&key=[ApiKey]&indent=True") 

	if err != nil { 
		log.Fatalln(err) 
	} 

	defer resp.Body.Close() 

	body, err := ioutil.ReadAll(resp.Body) 
	if err != nil { 
		log.Fatalln(err) 
	}

	var bodyJsonString string = string(body)
	var resultStruct gkb

	json.Unmarshal([]byte(bodyJsonString), &resultStruct)

	log.Println("Google Knowledge Graph knows that", resultStruct.ItemListElement[0].Result.Name, "is a", resultStruct.ItemListElement[0].Result.Description, "with an accuracy score of", resultStruct.ItemListElement[0].ResultScore)


}
