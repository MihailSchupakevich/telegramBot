package main

import (
	"fmt"
	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegramBot/auth"
	"telegramBot/handlers"
)

const (
	StateIdle = iota // начальное состояние
	StateWaitingForProductName
	StateWaitingForQuantity
)

var states = make(map[int64]int)
var globalStorageForProduct = make(map[int64]string)

func main() {
	botAPI, err := tb.NewBotAPI(auth.Token)
	if err != nil {
		panic(err)
	}
	productMap := map[string]*handlers.Product{}
	botAPI.Debug = true
	fmt.Printf("Authorized on account %s\n", botAPI.Self.UserName)

	u := tb.NewUpdate(0)
	u.Timeout = 60

	updates := botAPI.GetUpdatesChan(u)

	//var products []handlers.Product

	for update := range updates {
		if update.Message == nil { // игнорируем пустые обновления
			continue
		}

		chatID := update.Message.Chat.ID
		currentState := states[chatID]

		switch currentState {
		case StateIdle:
			switch update.Message.Text {
			case "/start":
				handlers.Start(botAPI, update.Message)
			case "/add_product":
				botAPI.Send(tb.NewMessage(update.Message.Chat.ID, "Добавление продукта.Введите сначала наименование продукта"))
				states[chatID] = StateWaitingForProductName
			case "/possibilities":
				botAPI.Send(tb.NewMessage(update.Message.Chat.ID, "ничего удивительного, пока нихрена не делает"))

			case "/list":
				botAPI.Send(tb.NewMessage(update.Message.Chat.ID, "Выводим список продуктов"))

				//botAPI.Send(tb.NewMessage(update.Message.Chat.ID, productMap))
			}
		case StateWaitingForProductName:
			productName := update.Message.Text
			globalStorageForProduct[chatID] = productName
			botAPI.Send(tb.NewMessage(update.Message.Chat.ID, "Отлично!Укажите количество"))
			states[chatID] = StateWaitingForQuantity
		case StateWaitingForQuantity:
			count := update.Message.Text
			errAddProduct := handlers.AddProduct(globalStorageForProduct[chatID], count, productMap)
			if errAddProduct != nil {
				botAPI.Send(tb.NewMessage(chatID, errAddProduct.Error()))
			} else {
				botAPI.Send(tb.NewMessage(chatID, "Добавление продукта прошло успешно!"))
			}
			fmt.Println(productMap)
			states[chatID] = StateIdle
			delete(globalStorageForProduct, chatID)
		}
	}
}
