package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	userToken = os.Getenv("DISCORD_USER_TOKEN")
	channels = []string{"595224241742413844", "692712497844584448"}
	legit=0
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(userToken)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()
}

func ready(_ *discordgo.Session, _ *discordgo.Ready) {
	fmt.Println("Automation ready!  Press CTRL-C to exit.")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, item := range channels {
		if item == m.ChannelID {
			if m.Author.ID == "755580145078632508" && m.Embeds[0].Title == "A trick-or-treater has stopped by!" {
				c, _ := s.Channel(m.ChannelID)
				fmt.Println("Trick Or Treat in Channel #" + c.Name)
				if legit == 0 {
					_, _= s.ChannelMessageSend(m.ChannelID, "h!treat")
					fmt.Println("Got Trick Or Treat")

					legit++
				} else {
					fmt.Println("LEGIT! Lost Trick Or Treat")
					legit++
					if legit == 2 {
						legit = 0
					}
				}
			}
		}
	}
}