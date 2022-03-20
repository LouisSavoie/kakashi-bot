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

func main() {
	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("The .env scroll is sealed:", err)
	}
	token := os.Getenv("KAKASHIBOT_TOKEN")

	// Start Discord session
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Not enough chakra to use Discordjutsu:", err)
		return
	} else {
		fmt.Println("Sharingan!")
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Pigeon sent to Discord did not return:,", err)
		return
	}

	// Wait here until CTRL-C or other termination signal is received.
	fmt.Println("Kakashi is now copying signs. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore Kakachi's own messages
	if message.Author.ID == session.State.User.ID {
		return
	}
}
