package main

// echobot
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}
type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}
type Chat struct {
	ChatId int `json:"id"`
}
type RestResponse struct {
	Result []Update `json:"result"`
}
type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

// dogs
type Jpg struct {
	Message string `json:"message"`
}

// forecast
type Forecast struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	CurrentUnits struct {
		Temperature string `json:"temperature_2m"`
	} `json:"current_units"`
	Current struct {
		Temperature float64 `json:"temperature_2m"`
	} `json:"current"`
}
