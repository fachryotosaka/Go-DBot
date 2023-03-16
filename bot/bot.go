package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/fachryotosoka/Go-DBot/config"
)

var BotID string
var bot *discordgo.Session

func StartBot() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot started >//<")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "Ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pings")
	}
	if m.Content == "Lexa" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Lexa Is a God Of Programmer")
	}
	if m.Content == "Hallo" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi , ><")
	}
	if m.Content == "Who" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi ! , I'm Axel , Trying to make a bot with the go programming language 😎😁")
	}
}
