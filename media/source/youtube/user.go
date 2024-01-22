package youtube

import (
	"fmt"
	"media-api-playground/media"
	"strings"

	"github.com/katelynn620/tubemeta"
	"google.golang.org/api/googleapi"
)

func (sy *SourceYoutube) GetMediaUser(uid string) (*media.MediaUser, error) {
	c, err := tubemeta.GetChannel(uid)
	if err != nil {
		return nil, err
	}

	// change uid if prefixed with @
	if strings.HasPrefix(uid, "@") {
		uid = c.Id
	}

	q := sy.yts.ChannelsList([]string{"snippet"})

	xs, err := q.Do(
		googleapi.QueryParameter("id", uid),
		googleapi.QueryParameter("maxResults", "1"),
	)
	if err != nil {
		return nil, err
	}

	if len(xs.Items) == 0 {
		return nil, fmt.Errorf("no user found")
	}

	return &media.MediaUser{
		Id:          xs.Items[0].Id,
		Title:       xs.Items[0].Snippet.Title,
		Description: xs.Items[0].Snippet.Description,
		Avatar:      xs.Items[0].Snippet.Thumbnails.Default.Url,
		URL:         fmt.Sprintf("https://www.youtube.com/channel/%s", xs.Items[0].Id),
		Platform:    "youtube",
		IsLive:      c.Live,
	}, nil
}
