package cmd

import (
	"bytes"

	"github.com/Cloud-Fortress/syborg/pkg/framework"
)

func DebugCommand(ctx framework.Context) {
	if ctx.Conf.OwnerId != ctx.User.ID {
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("no current session")
		return
	}
	queue := sess.Queue
	q := queue.Get()
	buff := bytes.Buffer{}
	for _, song := range q {
		buff.WriteString(song.Id + " ")
	}
	ctx.Reply(buff.String())
}
