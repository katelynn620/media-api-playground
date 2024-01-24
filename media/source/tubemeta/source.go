package tubemeta

import (
	"media-api-playground/media"
	"strings"

	"github.com/katelynn620/tubemeta"
)

type TubemetaInterface interface {
	GetVideo(videoId string) (*media.Video, error)
	GetChannel(channelId string) (*media.MediaUser, error)
}

type TubemetaService struct {
	TubemetaInterface
}

type SourceTubemeta struct {
	media.Source
	ts *TubemetaService
}

func (ts *TubemetaService) GetVideo(videoId string) (*media.Video, error) {
	video, err := tubemeta.GetVideo(videoId)
	if err != nil {
		return nil, err
	}

	return &media.Video{
		Id:          video.Id,
		Title:       video.Title,
		Description: video.Description,
		Thumbnail:   video.Thumbnails[0],
		URL:         video.URL,
		IsLive:      video.LiveContent,
	}, nil

}

func (ts *TubemetaService) GetChannel(channelId string) (*media.MediaUser, error) {
	channel, err := tubemeta.GetChannel(channelId)
	if err != nil {
		return nil, err
	}
	nameUrl := strings.Split(channel.CustomUrl, "/")
	name := nameUrl[len(nameUrl)-1]

	return &media.MediaUser{
		Id:          channel.Id,
		Name:        name,
		Title:       channel.Name,
		Description: channel.Description,
		Avatar:      channel.Avatar,
		URL:         channel.Url,
		Platform:    "youtube",
		IsLive:      channel.Live,
	}, nil
}

func (st *SourceTubemeta) GetVideo(id string) (*media.Video, error) {
	return st.ts.GetVideo(id)
}

func (st *SourceTubemeta) GetMediaUser(id string) (*media.MediaUser, error) {
	return st.ts.GetChannel(id)
}

func NewSourceTubemeta() *SourceTubemeta {
	return &SourceTubemeta{
		ts: &TubemetaService{},
	}
}
