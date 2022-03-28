package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	const BotToken = "5271893873:AAExaXdEsAwbLO6zddr7Pw-LhxE2z8RmUXM"
	baseUrl := fmt.Sprintf("https://api.telegram.org/bot%s/", BotToken)
	GetMeReq(baseUrl, "getMe")
}

func GetMeReq(baseUrl string, endPoint string) error {
	url := baseUrl + endPoint
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("http Error -> %s", err.Error())
	}
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	var nGetme GetMe
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&nGetme)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", nGetme)
	return nil
}

type GetMe struct {
	Ok   bool `json:"ok"`
	User User `json:"result"`
}

type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	Username                string `json:"username"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}
