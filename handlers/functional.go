package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (p *Product) addProduct(nameProduct string, count int, productMap map[string]*Product) {

}

func Start(botApi *tgbotapi.BotAPI, m *tgbotapi.Message, productSlice *[]Product) {
	message, err := botApi.Send(tgbotapi.NewMessage(m.Chat.ID, "Добро пожаловать в продуктового тг бота!"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message.Text)
}
