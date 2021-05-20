package server

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/PalindromeCorp/TODOBOT/external/dgrouter/exrouter"
)

type orderCommand struct {
	route *exrouter.Route

	logger *log.Entry
}

func (c *orderCommand) order(ctx *exrouter.Context) {
	ctx.Reply("sub1 called with arguments:\n", strings.Join(ctx.Args, ";"))
}

func (c *orderCommand) subOrder(ctx *exrouter.Context) {
	ctx.Reply("sub2 called with arguments:\n", strings.Join(ctx.Args, ";"))
}
