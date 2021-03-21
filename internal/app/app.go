package app

import (
	"bot/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	botApi = "https://api.telegram.org/bot1734095584:AAE9Pqw1XGCVtVCfmsbpPWToIj6LAiz-Y04/"
)

func InitApp() {
	fmt.Println("Starting bot...")
	offset := 0
	for {
		updates, err := getUpdates(botApi, offset)
		if err != nil {
			log.Fatal(err)
		}
		for _, update := range updates {
			err = respond(botApi, update)
			if err != nil {
				log.Fatal(err)
			}
			offset = update.ID + 1
		}
		fmt.Println(updates)
	}
}

func getUpdates(botApi string, offset int) ([]models.Update, error) {
	res, err := http.Get(botApi + "getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var data []byte
	var result models.GetUpdatesResult
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result.Result, nil
}
func respond(botApi string, update models.Update) error {
	var botResponse models.BotMessage
	botResponse.ChatID = update.Message.Chat.ID
	botResponse.Text = update.Message.Text
	data, err := json.Marshal(botResponse)
	if err != nil {
		return err
	}
	_, err = http.Post(botApi+"sendMessage", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	return nil
}
