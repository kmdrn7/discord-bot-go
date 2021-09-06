package main

import (
	"github.com/diamondburned/arikawa/v3/bot"
	"github.com/diamondburned/arikawa/v3/gateway"
	"os"
)

func main() {
	bot.Run(os.Getenv("DISCORD_TOKEN"), &Bot{},
		func(ctx *bot.Context) error {
			ctx.HasPrefix = bot.NewPrefix("!")
			return nil
		},
	)
}

type Bot struct {
	Ctx *bot.Context
}

func (b *Bot) Ping(*gateway.MessageCreateEvent) (string, error) {
	return "Pong!", nil
}