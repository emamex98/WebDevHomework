package main 

import (
	"io/ioutil" 
	"log" 
	"net/http" 
	"encoding/json"
)

type image struct {
	Url string
	ContentUrl string
}

type detailedDescription struct {
	Url string
	ArticleBody string
	License string
}

type result struct {
	Url string
	Id string
	Image image
	Type []string
	Name string
	Description string
	DetailedDescription detailedDescription
}

type ItemListElement struct {
	resultScore float32
	result Result
}

type gkb struct {
	itemListElement ItemListElement
}

func main() { 

	resp, err := http.Get("https://kgsearch.googleapis.com/v1/entities:search?query=tecnologico+de+monterrey&key=AIzaSyDsjVxULYfGOjQ0dpZiJVqUFPi489izEEY&limit=1&indent=True") 

	if err != nil { 
		log.Fatalln(err) 
	} 

	defer resp.Body.Close() 

	body, err := ioutil.ReadAll(resp.Body) 
	if err != nil { 
		log.Fatalln(err) 
	}

	log.Println(string(body)) 

	var bodyJsonString string = `{"url": "http://www.itesm.edu/","id": "kg:/m/01vr72","image": {"url": "https://en.wikipedia.org/wiki/Monterrey_Institute_of_Technology_and_Higher_Education","contentUrl": "http://t1.gstatic.com/images?q=tbn:ANd9GcRsoi9MgqTnycYnZFcY3YxCgSK18b9-QmJyZnWHUkt8IvZD5fan"},"name": "Monterrey Institute of Technology and Higher Education","type": ["Corporation","EducationalOrganization","Thing","Organization","CollegeOrUniversity","Place"],"description": "Private university in Monterrey, Mexico","detailedDescription": {"url": "https://en.wikipedia.org/wiki/Monterrey_Institute_of_Technology_and_Higher_Education","articleBody": "Instituto Tecnológico y de Estudios Superiores de Monterrey, also known as Tecnológico de Monterrey, is a secular and coeducational multi-campus private university based in Monterrey, Mexico. ","license": "https://en.wikipedia.org/wiki/Wikipedia:Text_of_Creative_Commons_Attribution-ShareAlike_3.0_Unported_License"}`
	var resultStruct result

	json.Unmarshal([]byte(bodyJsonString), &resultStruct)

	log.Println(resultStruct)

}
