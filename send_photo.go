package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

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
