package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"shootyhoops/handlers"
	"syscall"
)

func main() {
	// get bot token
	token := os.Getenv("DISCORD_TOKEN")

	// initialize new discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating discord session: " + err.Error())
		return
	}

	// open a websocket connection to discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection: ", err)
		return
	}

	// initialize bot handlers
	handler := handlers.NewHandler(dg.State.User)

	// mount the baseHandler
	dg.AddHandler(handler.BaseHandler)

	// make channel to catch CTRL-C or other term signals
	fmt.Println("shootyhoops is now raining buckets.  Press CTRL-C (or send a term signal of your choice) to make it do the epic Rockets-Warriors meltdown and go ice cold from downtown.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// close down the discord session after channel receives term signal
	err = dg.Close()
	if err != nil {
		panic("failed to close discord session: " + err.Error())
	}
}
