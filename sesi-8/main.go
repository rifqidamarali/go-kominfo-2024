package main

import (
	"encoding/json"
	"fmt"
)

type User struct{

}

func main() {
	fmt.Print("Hi")
}

func jsonEncodingDecoding(){
	
	jsonBody := `{
		"id":1,
		"name": "rifqi",
		"email": "rifqi@gmail.com",
	}`

	user := User{}
	err := json.Unmarshal([]byte(jsonBody), &user)
	if err != nil{
		panic(err)
	}
	fmt.Print(user)
}