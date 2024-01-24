package source

import (
	"context"
	"fmt"

	media "github.com/katelynn620/mediameta"

	"github.com/katelynn620/mediameta/source/tubemeta"
	"github.com/katelynn620/mediameta/source/twitch"
	"github.com/katelynn620/mediameta/source/youtube"

	"os"

	"github.com/nicklaw5/helix/v2"
	yt "google.golang.org/api/youtube/v3"
)

func NewSource(mediaType string) (media.Source, error) {
	var source media.Source

	if mediaType == "twitch" {
		ClientId := os.Getenv("TWITCH_CLIENT_ID")
		ClientSecret := os.Getenv("TWITCH_TOKEN")
		client, err := helix.NewClient(&helix.Options{
			ClientID:        ClientId,
			UserAccessToken: ClientSecret,
		})
		if err != nil {
			return nil, err
		}
		twitchService := twitch.NewTwitchService(client)
		source = twitch.NewTwitchSource(twitchService)
	} else if mediaType == "youtube" {
		if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
			source = tubemeta.NewSourceTubemeta()
		} else {
			ytService, err := yt.NewService(context.Background())
			if err != nil {
				return nil, err
			}
			ytServiceWrapper := youtube.NewYoutubeService(ytService)
			source = youtube.NewSourceYoutube(ytServiceWrapper)
		}
	} else {
		return nil, fmt.Errorf("invalid media type")
	}

	return source, nil
}
