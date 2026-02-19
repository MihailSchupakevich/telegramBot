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
var chatId int64

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

	keyboard := [][]tb.InlineKeyboardButton{
		{tb.NewInlineKeyboardButtonData("Добавить продукт", "add_product")},
		{tb.NewInlineKeyboardButtonData("Возможности", "possibilities")},
		{tb.NewInlineKeyboardButtonData("Список", "list")},
	}

	replyMarkup := tb.NewInlineKeyboardMarkup(keyboard...)

	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}
		if update.CallbackQuery != nil {
			chatId = update.CallbackQuery.Message.Chat.ID
		} else {
			chatId = update.Message.Chat.ID
		}
		currentState := states[chatId]
		if currentState == 0 || currentState == StateIdle {
			message := tb.NewMessage(chatId, "Выберите действие")
			message.ReplyMarkup = &replyMarkup
			botAPI.Send(message)
			states[chatId] = StateIdle
		}
		switch currentState {
		case StateIdle:
			if update.CallbackQuery != nil {
				cbq := update.CallbackQuery
				switch cbq.Data {
				case "add_product":
					botAPI.Send(tb.NewMessage(chatId, "Добавление продукта. Введите название продукта"))
					states[chatId] = StateWaitingForProductName
				case "possibilities":
					botAPI.Send(tb.NewMessage(chatId, "ничего удивительного, пока нихрена не делает"))
				case "list":
					botAPI.Send(tb.NewMessage(chatId, "Выводим список продуктов"))
					handlers.List(botAPI, update.CallbackQuery, productMap)
				}
			}
			//else if update.Message != nil {
			//	switch update.Message.Text {
			//	case "/start":
			//		handlers.Start(botAPI, update.Message)
			//	case "/add_product":
			//		botAPI.Send(tb.NewMessage(chatId, "Добавление продукта. Введите название продукта"))
			//		states[chatId] = StateWaitingForProductName
			//	case "/possibilities":
			//		botAPI.Send(tb.NewMessage(chatId, "ничего удивительного, пока нихрена не делает"))
			//	case "/list":
			//		botAPI.Send(tb.NewMessage(chatId, "Выводим список продуктов"))
			//		handlers.List(botAPI, update.Message, productMap)
			//	}
			//}
		case StateWaitingForProductName:
			fmt.Println("ENTER")
			productName := update.Message.Text
			globalStorageForProduct[chatId] = productName
			botAPI.Send(tb.NewMessage(chatId, "Отлично! Укажите количество"))
			states[chatId] = StateWaitingForQuantity
		case StateWaitingForQuantity:
			count := update.Message.Text
			errAddProduct := handlers.AddProduct(globalStorageForProduct[chatId], count, productMap)
			if errAddProduct != nil {
				botAPI.Send(tb.NewMessage(chatId, errAddProduct.Error()))
			} else {
				botAPI.Send(tb.NewMessage(chatId, "Добавление продукта прошло успешно!"))
			}
			states[chatId] = StateIdle
			delete(globalStorageForProduct, chatId)
		}
	}
}
