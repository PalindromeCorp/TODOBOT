package server

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type helpCommand struct {
	logger *log.Entry
}

func (c *helpCommand) command(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(s)
	fmt.Println(m.ID, m.Content)
	c.logger.Info("Log Run")
}
