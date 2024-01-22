package twitch

import (
	"media-api-playground/media"

	"github.com/nicklaw5/helix/v2"
)

func (ts *TwitchSource) GetVideo(vid string) (*media.Video, error) {
	resp, err := ts.s.GetVideos(&helix.VideosParams{
		IDs: []string{vid},
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Data.Videos) == 0 {
		return nil, media.ErrVideoNotFound{}
	}

	return &media.Video{
		Id:          resp.Data.Videos[0].ID,
		Title:       resp.Data.Videos[0].Title,
		Description: resp.Data.Videos[0].Description,
		Thumbnail:   resp.Data.Videos[0].ThumbnailURL,
		URL:         resp.Data.Videos[0].URL,
		IsLive:      resp.Data.Videos[0].Type == "archive",
		Platform:    "twitch",
	}, nil
}
