package server

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"os"
)

func New(logger *log.Entry) *Server {
	s := &Server{
		Logger: logger,
	}

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	dg.AddHandler(startup)
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.MessageCreate) {
		h := helpCommand{logger: logger}
		h.command(s, i)
	})

	s.session = dg

	return s
}

type Server struct {
	session *discordgo.Session

	Logger *log.Entry
}

func (s *Server) Serve() *discordgo.Session {
	if err := s.session.Open(); err != nil {
		panic(err)
	}

	return s.session
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func startup(ds *discordgo.Session, event *discordgo.Ready) {
	err := ds.UpdateGameStatus(0, "!help")
	if err != nil {
		panic(err)
	}
	fmt.Println(event)
}
