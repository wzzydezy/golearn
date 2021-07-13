package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct{
	Title string      `json:"title"`
	Year int          `json:"year"`
	Price int         `json:"price"`
	Actors []string   `json:"actor"` 
}

func main(){
	movie := Movie{"a",2000,10,[]string{"a","b"}}

	//编码过程 结构体 to json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	fmt.Printf("%s\n",jsonStr)

	//解码过程 
	//{"title":"a","year":2000,"price":10,"actor":["a","b"]}
	mymovie := Movie{}
	err = json.Unmarshal(jsonStr, &mymovie)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	fmt.Println("my",mymovie)

}