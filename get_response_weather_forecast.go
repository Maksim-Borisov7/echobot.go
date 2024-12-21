package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func get_response_weather_forecast(str string, botMessage BotMessage) {
	botMessage.Text = str
	buf, err := json.Marshal(botMessage)
	if err != nil {
		fmt.Println("Error Marshal", err)
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		fmt.Println("Error POST", err)
	}
	fmt.Println("Weather shown, good job")
}
