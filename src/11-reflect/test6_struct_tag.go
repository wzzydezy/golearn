package main

import "fmt"
import "reflect"

type resume struct {
	Name string `info:"name" doc:"我的mingzi"`
	Sex string `info:"sex"`
}

func findTag(str interface{}){
	t := reflect.TypeOf(str).Elem()
    for i:=0;i<t.NumField();i++{
		tagstring := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println(tagstring, tagdoc)
	}
}

func main(){
	var re resume
	findTag(&re)
}