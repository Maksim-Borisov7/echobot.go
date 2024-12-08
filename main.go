package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	botToken := "8001922451:AAGfXYpu1tUhGXNz_jMmchfLs7pDTge5R-Q"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println(err)
		}
		for _, update := range updates {
			fmt.Println(update)
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
		}
	}
}

// функция которая запрашивает обновления
func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	fmt.Println(restResponse.Result)
	return restResponse.Result, nil
}

// отвечает на обновления
func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
