package twitch

import (
	media "github.com/katelynn620/mediameta"

	"github.com/nicklaw5/helix/v2"
)

const (
	platform = "twitch"
	baseURL  = "https://www.twitch.tv"
)

type TwitchServiceInterface interface {
	GetVideos(params *helix.VideosParams) (*helix.VideosResponse, error)
	GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error)
	GetStreams(params *helix.StreamsParams) (*helix.StreamsResponse, error)
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

func (ts *TwitchService) GetStreams(params *helix.StreamsParams) (*helix.StreamsResponse, error) {
	return ts.client.GetStreams(params)
}

func NewTwitchService(client *helix.Client) *TwitchService {
	return &TwitchService{client: client}
}

func NewTwitchSource(s TwitchServiceInterface) *TwitchSource {
	return &TwitchSource{s: s}
}
