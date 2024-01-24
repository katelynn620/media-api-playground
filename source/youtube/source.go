package youtube

import (
	media "github.com/katelynn620/mediameta"

	yt "google.golang.org/api/youtube/v3"
)

type YoutubeServiceInterface interface {
	ChannelsList(part []string) *yt.ChannelsListCall
	VideosList(part []string) *yt.VideosListCall
	SearchList(part []string) *yt.SearchListCall
}

type YoutubeService struct {
	YoutubeServiceInterface
	service *yt.Service
}

type SourceYoutube struct {
	media.Source
	yts YoutubeServiceInterface
}

func (ys *YoutubeService) ChannelsList(part []string) *yt.ChannelsListCall {
	return ys.service.Channels.List(part)
}

func (ys *YoutubeService) VideosList(part []string) *yt.VideosListCall {
	return ys.service.Videos.List(part)
}

func (ys *YoutubeService) SearchList(part []string) *yt.SearchListCall {
	return ys.service.Search.List(part)
}

func NewYoutubeService(service *yt.Service) *YoutubeService {
	return &YoutubeService{
		service: service,
	}
}

func NewSourceYoutube(service YoutubeServiceInterface) *SourceYoutube {
	return &SourceYoutube{
		yts: service,
	}
}
