package youtube

import (
	"fmt"
	"media-api-playground/media"

	"google.golang.org/api/googleapi"
)

func (sy *SourceYoutube) GetVideo(vid string) (*media.Video, error) {
	q := sy.yts.VideosList([]string{"snippet"})

	xs, err := q.Do(
		googleapi.QueryParameter("id", vid),
		googleapi.QueryParameter("maxResults", "1"),
	)
	if err != nil {
		return nil, err
	}

	if len(xs.Items) == 0 {
		return nil, fmt.Errorf("no video found")
	}

	return &media.Video{
		Id:          xs.Items[0].Id,
		Title:       xs.Items[0].Snippet.Title,
		Description: xs.Items[0].Snippet.Description,
		Thumbnail:   xs.Items[0].Snippet.Thumbnails.Default.Url,
		URL:         fmt.Sprintf("https://www.youtube.com/watch?v=%s", xs.Items[0].Id),
		IsLive:      xs.Items[0].Snippet.LiveBroadcastContent == "live",
		Platform:    "youtube",
	}, nil
}
