package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

const (
	botToken = "8001922451:AAGfXYpu1tUhGXNz_jMmchfLs7pDTge5R-Q"
	botApi   = "https://api.telegram.org/bot"
	botUrl   = botApi + botToken
)

func main() {
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println(err)
		}
		for _, update := range updates {
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
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
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
	if botMessage.Text == "/get_photo_dog" {
		get_photo_dogs()
		sendPhoto(int64(update.Message.Chat.ChatId), "C:\\Users\\maxva\\GolandProjects\\echobot\\img.jpg")
		return nil
	} else if botMessage.Text == "/get_weather_forecast" {
		get_weather_forecast(int64(botMessage.ChatId))
		return nil
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

type Jpg struct {
	Message string `json:"message"`
}

func get_photo_dogs() {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Не удалось прочитать файл", err)
	}
	str := ""
	for _, v := range body {
		if string(v) == "\\" {
			continue
		} else {
			str += string(v)
		}
	}
	var jpg Jpg
	err = json.Unmarshal([]byte(str), &jpg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	response2, err := http.Get(jpg.Message)
	if err != nil {
		fmt.Println("Не удалость сделать запрос", err)
	}
	outFile, err := os.Create("img.jpg")
	if err != nil {
		fmt.Println("Не удалось создать файл", err)
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, response2.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func sendPhoto(chatID int64, photoPath string) {
	file, err := os.Open(photoPath)
	if err != nil {
		fmt.Println("Error opening photo file:", err)
		return
	}
	defer file.Close()
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("photo", file.Name()) // Здесь мы указываем имя файла
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file contents:", err)
		return
	}
	writer.Close()
	resp, err := http.Post(botUrl+"/sendPhoto?chat_id="+fmt.Sprint(chatID), writer.FormDataContentType(), &b)
	if err != nil {
		fmt.Println("Error sending photo:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Photo sent, response:", string(body))
}

func get_weather_forecast(chatid int64) {
	API := "https://api.open-meteo.com/v1/forecast?latitude=54.7261409&longitude=55.947499&current=temperature_2m&timezone=auto&forecast_days=1"
	var botMessage BotMessage
	botMessage.ChatId = int(chatid)
	response, err := http.Get(API)
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
