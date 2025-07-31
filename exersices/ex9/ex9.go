package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const RawResp = `{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
} `

type (
	Response struct {
		Head HeadResp     `json:"header"`
		Data ResponseData `json:"data"`
	}

	HeadResp struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty" `
	}

	ResponseData []DataItem

	DataItem struct {
		Type       string         `json:"type"`
		Id         int            `json:"id"`
		Attributes AttributesData `json:"attributes"`
	}

	AttributesData struct {
		Email       string `json:"email"`
		Article_ids []int  `json:"article_ids"`
	}
)

func ReadResponse(rawResp string) (Response, error) {
	var response Response
	err := json.Unmarshal([]byte(rawResp), &response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	man := Person{
		Name:        "Alex",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}
	manToJs, err := json.Marshal(man)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Printf("Man %v", string(manToJs))

	resp, err := ReadResponse(RawResp)
	if err != nil {
		panic("error")
	}
	fmt.Printf("%+v\n", resp)

}
