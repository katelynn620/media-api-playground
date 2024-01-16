package youtube

import (
	"media-api-playground/media"

	yt "google.golang.org/api/youtube/v3"
)

type YoutubeServiceInterface interface {
	ChannelsList(part []string) *yt.ChannelsListCall
	VideosList(part []string) *yt.VideosListCall
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
