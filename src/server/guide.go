package server

import (
	log "github.com/sirupsen/logrus"

	"github.com/PalindromeCorp/TODOBOT/external/dgrouter/exrouter"
)

type helpCommand struct {
	route *exrouter.Route

	logger *log.Entry
}

func (c *helpCommand) router(name string) *exrouter.Route {
	c.route.Default = c.route.On(name, c.help).Desc("Use the emojis to navigate this help guide:")
	return c.route
}

func (c *helpCommand) help(ctx *exrouter.Context) {
	var f func(depth int, r *exrouter.Route) string
	f = func(depth int, r *exrouter.Route) string {
		text := ""
		for _, v := range r.Routes {
			text += v.Description
			text += f(depth+1, &exrouter.Route{Route: v})
		}
		return text
	}
	_, err := ctx.ReplyEmbed(f(0, c.route))
	if err != nil {
		c.logger.Error(err)
	}
}
