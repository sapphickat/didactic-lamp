// package didacticlamp
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	GuildID 	= ""
	BotToken	= ""
)

var s *discordgo.Session

func init() {
	fmt.Println("registering slash commands")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}

	BotToken = os.Getenv("BOT_TOKEN")

	s, err := discordgo.New("Bot " + BotToken)

	if err != nil {
		fmt.Println("error creating discord session", err)
		return 
	}

	s.AddHandler(interactionHandler)

	err = s.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return 
	}

	command := &discordgo.ApplicationCommand{
		Name:		"hello",
		Description: "Replies with helloo",
	}

	_, err = s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		fmt.Println("cannot create slash command", err)
	}

	fmt.Println("bot is running. press ctrl-c t exit")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc

	s.Close()

}


func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "hello": 
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hellowww katt",
				},
			})
		}
	}
}