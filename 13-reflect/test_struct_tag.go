package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type resume struct{
	Name string `info:"name" doc:"我的名字"`
	Sex string  `info:"sex" `

}

func findTag(str interface{}){
	t := reflect.TypeOf(str).Elem()

	for i:= 0;i<t.NumField();i++{
		tagstring := t.Field(i).Tag.Get("info")
		fmt.Println("info: ",tagstring)
	}
}



type Movie struct{
	Title string 	`json:"title"`
	Year int 	`json:"year"`
	Price int 	`json:"rmb"`
	Actors  []string   `json:"actors"`
}


func main(){
	var re resume
	findTag(&re)


	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye","zhangbozhi"}}

	// 编码的过程  结构体->json
	jsonStr, err := json.Marshal(movie)

	if err != nil{
		fmt.Println("json marshal error,", err)
		return
	}
	fmt.Printf("jsonStr = %s\n",jsonStr)

	// 解码的过程 jsonstr -> 结构体
	// jsonStr = {"title":"喜剧之王","year":2000,"rmb":10,"actors":["xingye","zhangbozhi"]}

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr,&myMovie)
	if err != nil{
		fmt.Println("json unmarshal error,", err)
		return
	}
	fmt.Printf("myMovie = %v\n",myMovie)
}