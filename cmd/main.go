package main

import (
	"fmt"

	media "github.com/katelynn620/mediameta"
	"github.com/katelynn620/mediameta/source"
)

type Data struct {
	Source  string
	VideoId string
	UserId  string
}

func main() {
	all := []Data{}

	var (
		user  *media.MediaUser
		video *media.Video
	)
	data := Data{
		Source:  "youtube",
		VideoId: "P3buv6P_u7c",
		UserId:  "UC_x5XG1OV2P6uZZ5FSM9Ttw",
	}

	all = append(all, data)

	data = Data{
		Source:  "twitch",
		VideoId: "2031892840",
		UserId:  "never_loses",
	}

	all = append(all, data)

	for _, data := range all {
		fmt.Printf("Source: %s\n", data.Source)
		s, err := source.NewSource(data.Source)
		if err != nil {
			panic(err)
		}

		if data.VideoId != "" {
			video, err = s.GetVideo(data.VideoId)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%+v\n", video)
		}
		if data.UserId != "" {
			fmt.Printf("UserId: %s\n", data.UserId)
			user, err = s.GetMediaUser(data.UserId)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%+v\n", user)
		}
	}
}
