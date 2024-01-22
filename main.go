package main

import (
	"fmt"
	"media-api-playground/media"
	"media-api-playground/media/source"

	"github.com/katelynn620/tubemeta"
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
		UserId:  "fps_shaka",
	}

	all = append(all, data)

	for _, data := range all {
		source, err := source.NewSource(data.Source)
		if err != nil {
			panic(err)
		}
		svc, err := media.NewMediaService(source)
		if err != nil {
			panic(err)
		}
		if data.VideoId != "" {
			video, err = svc.GetVideo(data.VideoId)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%+v\n", video)
		}
		if data.UserId != "" {
			user, err = svc.GetMediaUser(data.UserId)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%+v\n", user)
		}
	}

	v, err := tubemeta.GetVideo("https://www.youtube.com/watch?v=hTVrE4BYkwA")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", v)
}
