package main

import (
	"context"
	"fmt"
	"media-api-playground/media"
	"media-api-playground/media/twitch"
	"media-api-playground/media/youtube"
	"os"

	"github.com/nicklaw5/helix/v2"
	yt "google.golang.org/api/youtube/v3"
)

func runYoutube() {
	var (
		mediaSource = "youtube"
		videoId     = "P3buv6P_u7c"
		userId      = "UC_x5XG1OV2P6uZZ5FSM9Ttw"
		video       *media.Video
		mediaUser   *media.MediaUser
		source      media.Source
		err         error
	)

	if mediaSource == "youtube" {
		ytService, err := yt.NewService(context.Background())
		if err != nil {
			panic(err)
		}
		ytServiceWrapper := youtube.NewYoutubeService(ytService)
		source = youtube.NewSourceYoutube(ytServiceWrapper)
	} else {
		panic("not implemented")
	}

	video, err = source.GetVideo(videoId)
	if err != nil {
		panic(err)
	}
	println(video.Title)

	mediaUser, err = source.GetMediaUser(userId)
	if err != nil {
		panic(err)
	}
	println(mediaUser.Title)
}

func runTwitch() {
	var (
		mediaSource = "twitch"
		videoId     = "2031892840"
		userId      = "fps_shaka"
		video       *media.Video
		mediaUser   *media.MediaUser
		source      media.Source
		err         error
	)

	if mediaSource == "twitch" {
		ClientId := os.Getenv("TWITCH_CLIENT_ID")
		ClientSecret := os.Getenv("TWITCH_TOKEN")
		client, err := helix.NewClient(&helix.Options{
			ClientID:        ClientId,
			UserAccessToken: ClientSecret,
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		twitchService := twitch.NewTwitchService(client)
		source = twitch.NewTwitchSource(twitchService)
	} else {
		panic("not implemented")
	}
	video, err = source.GetVideo(videoId)
	if err != nil {
		panic(err)
	}
	println(video.Title)
	mediaUser, err = source.GetMediaUser(userId)
	if err != nil {
		panic(err)
	}
	println(mediaUser.Title)
}

func main() {
	runYoutube()
	runTwitch()
}
