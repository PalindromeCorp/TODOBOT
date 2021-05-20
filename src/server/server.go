package server

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/PalindromeCorp/TODOBOT/external/dgrouter/exrouter"
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

	router := exrouter.New()

	help := helpCommand{route: router, logger: logger}
	router = help.router("help")

	router.On("new", func(ctx *exrouter.Context) {
		order := orderCommand{route: router, logger: logger}
		order.order(ctx)
	}).On("new2", func(ctx *exrouter.Context) {
		order := orderCommand{route: router, logger: logger}
		order.subOrder(ctx)
	})

	// Add message handler
	dg.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		if err := router.FindAndExecute(dg, "todo!", dg.State.User.ID, m.Message); err != nil {
			logger.Error(err)
		}
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
func startup(ds *discordgo.Session, _ *discordgo.Ready) {
	err := ds.UpdateGameStatus(0, "Helo there !help")
	if err != nil {
		panic(err)
	}
}
