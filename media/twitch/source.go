package twitch

import (
	"media-api-playground/media"

	"github.com/nicklaw5/helix/v2"
)

const (
	platform = "twitch"
	baseURL  = "https://www.twitch.tv"
)

type TwitchServiceInterface interface {
	GetVideos(params *helix.VideosParams) (*helix.VideosResponse, error)
	GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error)
}

type TwitchService struct {
	TwitchServiceInterface
	client *helix.Client
}

type TwitchSource struct {
	media.Source
	s TwitchServiceInterface
}

func (ts *TwitchService) GetVideos(params *helix.VideosParams) (*helix.VideosResponse, error) {
	return ts.client.GetVideos(params)
}

func (ts *TwitchService) GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error) {
	return ts.client.GetUsers(params)
}

func NewTwitchService(client *helix.Client) *TwitchService {
	return &TwitchService{client: client}
}

func NewTwitchSource(s TwitchServiceInterface) *TwitchSource {
	return &TwitchSource{s: s}
}
