package main

import (
	"fmt"
	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegramBot/auth"
	"telegramBot/handlers"
)

func main() {
	botAPI, err := tb.NewBotAPI(auth.Token)
	if err != nil {
		panic(err)
	}

	botAPI.Debug = true
	fmt.Printf("Authorized on account %s\n", botAPI.Self.UserName)

	u := tb.NewUpdate(0)
	u.Timeout = 60

	updates := botAPI.GetUpdatesChan(u)

	var products []handlers.Product

	for update := range updates {
		if update.Message == nil { // игнорируем пустые обновления
			continue
		}
		switch update.Message.Text {
		case "start":
			fmt.Println("ENTER")
			//botAPI.SendMessage(update.Message.Chat.ID, "Привет! Я продуктовый бот.")

			handlers.Start(botAPI, update.Message, &products)
		case "/add_product":
			botAPI.Send(tb.NewMessage(update.Message.Chat.ID, "ничего удивительного, пока нихрена не делает"))
			//addProduct(botAPI, &products, update)
		case "/stop_add_p":

			//stopAddingProducts(botAPI, &products, update)
		case "/list_p":
			//listProducts(botAPI, products, update)
		}
	}
}
