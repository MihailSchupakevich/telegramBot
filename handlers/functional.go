package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func AddProduct(nameProduct string, count string, productMap map[string]*Product) error {
	pr := Product{nameProduct, count}
	productMap[nameProduct] = &pr
	return nil
}

func Start(botApi *tgbotapi.BotAPI, m *tgbotapi.Message) {
	_, err := botApi.Send(tgbotapi.NewMessage(m.Chat.ID, "Добро пожаловать в продуктового тг бота!"))
	if err != nil {
		log.Fatal(err)
	}
}

//func
