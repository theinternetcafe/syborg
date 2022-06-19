package cmd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Cloud-Fortress/syborg/pkg/framework"
	"github.com/Cloud-Fortress/youtube"
	"github.com/bwmarrin/discordgo"
)

const result_format = "\n`%d` %s - %s (%s)"

var ytSessions ytSearchSessions = make(ytSearchSessions)

type (
	ytSearchSessions map[string]ytSearchSession

	ytSearchSession struct {
		results youtube.SearchResult
	}
)

func ytSessionIdentifier(user *discordgo.User, channel *discordgo.Channel) string {
	return user.ID + channel.ID
}

func formatDuration(input string) string {
	return parseISO8601(input).String()
}

func YoutubeCommand(ctx framework.Context) {
	if len(ctx.Args) == 0 {
		ctx.Reply("Usage: `music youtube <search query>`")
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
		return
	}
	query := strings.Join(ctx.Args, " ")
	results, err := youtube.Search(query, 0)
	if err != nil {
		ctx.Reply("An error occured!")
		fmt.Println("Error searching youtube,", err)
		return
	}
	if len(results.Items) == 0 {
		ctx.Reply("No results found for your query `" + query + "`.")
		return
	}
	buffer := bytes.NewBufferString("__Search results__ for `" + query + "`:\n")
	for index, result := range results.Items {
		fmt.Printf("Title: %s\nVideo Id: %s\n\n", result.Title, result.ID)
		buffer.WriteString(fmt.Sprintf(result_format, index+1, result.Title, result.Author,
			formatDuration(result.Duration)))
	}
	buffer.WriteString("\n\nTo pick a song, use `music pick <number>`.")
	ytSessions[ytSessionIdentifier(ctx.User, ctx.TextChannel)] = ytSearchSession{results}
	ctx.Reply(buffer.String())
}
