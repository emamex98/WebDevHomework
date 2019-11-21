package main 

import (
	//"io/ioutil" 
	"log" 
	//"net/http" 
	"encoding/json"
)


type person struct {
	Name string
}

func addCookie(w http.ResponseWriter, name string, value string) {

    expire := time.Now().AddDate(0, 0, 1)

    cookie := http.Cookie{
        Name:    name,
        Value:   value,
        Expires: expire,
    }

    http.SetCookie(w, &cookie)
}

func main() { 

	resp, err := http.Get("https://swapi.co/api/people/1/") 

	if err != nil { 
		log.Fatalln(err) 
	} 

	defer resp.Body.Close() 

	body, err := ioutil.ReadAll(resp.Body) 
	if err != nil { 
		log.Fatalln(err) 
	}

	var bodyJsonString string = string(body)
	json.Unmarshal([]byte(bodyJsonString), &resultStruct)
	
	log.Println(resultStruct)
	
	addCookie(w, "Person", resultStruct.Name)

}

