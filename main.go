package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/joho/godotenv"
	"github.com/nsigel/ucblogmonitor/blog"
)

var (
	hook webhook.Client
)

func init() {
	godotenv.Load()

	var err error
	hook, err = webhook.NewWithURL(os.Getenv("DISCORD_WEBHOOK_URL"))
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	watch()
}

func watch() {
	c := blog.NewClient()

	log.Println("Searching for new blogs...")
	last, err := c.Search()
	if err != nil {
		log.Panic(err)
	}

	interval := time.Minute * 3

	for {
		log.Printf("Sleeping for %.0f minutes...", interval.Minutes())
		time.Sleep(interval)

		log.Printf("Searching for new blogs...")
		post, err := c.Search()
		if err != nil {
			log.Println(err)
			continue
		}

		if post.Url != last.Url {
			hook.CreateMessage(discord.WebhookMessageCreate{Content: "<@755796408791203850>", Embeds: []discord.Embed{embed(*post)}})
		}

		last = post
	}
}

func embed(post blog.Post) discord.Embed {
	t := time.Now()

	postUrl := "https://collegeadmissions.uchicago.edu" + post.Url

	return discord.Embed{
		Title: post.Title,
		URL:   postUrl,
		Author: &discord.EmbedAuthor{
			Name: "University of Chicago Uncommon Blog",
			URL:  "https://collegeadmissions.uchicago.edu/uncommon-blog",
		},

		Footer: &discord.EmbedFooter{
			Text:    "UChicago Blog Monitor",
			IconURL: "https://biocars.uchicago.edu/wp-content/uploads/2019/05/cropped-logo.png",
		},
		Description: strings.Replace(post.Content, "Read more...", fmt.Sprintf("[Read more...](%s)", postUrl), 1),
		Timestamp:   &t,
		Color:       0x800000,
	}
}
