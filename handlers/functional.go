package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func AddProduct(nameProduct string, count string, productMap map[string]*Product) error {
	pr := Product{nameProduct, count}
	productMap[nameProduct] = &pr
	return nil
}

func Start(botApi *tgbotapi.BotAPI, m *tgbotapi.Message) {
	_, err := botApi.Send(tgbotapi.NewMessage(m.Chat.ID, "Добро пожаловать в продуктового тг бота!\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func List(botApi *tgbotapi.BotAPI, m *tgbotapi.CallbackQuery, productMap map[string]*Product) {
	if len(productMap) == 0 {
		botApi.Send(tgbotapi.NewMessage(m.Message.Chat.ID, "Еще ничего не добавлено!"))
		return
	}
	result := ""
	arguments := 1
	for nameProduct, value := range productMap {
		result += fmt.Sprintf("[%v] %s в количестве: %s\n", arguments, nameProduct, value.count)
		arguments += 1
	}
	botApi.Send(tgbotapi.NewMessage(m.Message.Chat.ID, result))
}
