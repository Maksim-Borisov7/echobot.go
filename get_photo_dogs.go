package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func get_photo_dogs() {
	response, err := http.Get(APIdogs)
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
