package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func get_weather_forecast(chatid int64) {
	var botMessage BotMessage
	botMessage.ChatId = int(chatid)
	response, err := http.Get(APIweather)
	if err != nil {
		fmt.Println("Ошибка GET-запроса", err)
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	var forecast Forecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		fmt.Println("Error Unmarshal", err)
	}
	str := fmt.Sprintf("Текущая погода %.1f%s", forecast.Current.Temperature, forecast.CurrentUnits.Temperature)
	get_response_weather_forecast(str, botMessage)
}
