package youtube

import (
	"fmt"

	media "github.com/katelynn620/mediameta"

	"strings"

	"google.golang.org/api/googleapi"
)

type channelResponse struct {
	Id struct {
		ChannelId string `json:"channelId"`
	} `json:"id"`
	Snippet struct {
		Title       string `json:"title"`
		CustomUrl   string `json:"customUrl"`
		Description string `json:"description"`
		Thumbnails  struct {
			Default struct {
				Url string `json:"url"`
			} `json:"default"`
		} `json:"thumbnails"`
	} `json:"snippet"`
}

type videoResponse struct {
	Id struct {
		VideoId string `json:"videoId"`
	} `json:"id"`
	Snippet struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Thumbnails  struct {
			Default struct {
				Url string `json:"url"`
			} `json:"default"`
		} `json:"thumbnails"`
		ChannelId string `json:"channelId"`
	} `json:"snippet"`
}

func (sy *SourceYoutube) getChannelByName(name string) (*channelResponse, error) {
	q := sy.yts.SearchList([]string{"snippet", "id"})

	xs, err := q.Do(
		googleapi.QueryParameter("q", name[1:]),
		googleapi.QueryParameter("type", "channel"),
		googleapi.QueryParameter("maxResults", "1"),
	)
	if err != nil {
		return nil, err
	}

	if len(xs.Items) == 0 {
		return nil, fmt.Errorf("no channel found")
	}

	channel := &channelResponse{}
	channel.Id.ChannelId = xs.Items[0].Snippet.ChannelId
	channel.Snippet.CustomUrl = name
	channel.Snippet.Title = xs.Items[0].Snippet.Title
	channel.Snippet.Description = xs.Items[0].Snippet.Description
	channel.Snippet.Thumbnails.Default.Url = xs.Items[0].Snippet.Thumbnails.Default.Url

	return channel, nil
}

func (sy *SourceYoutube) getChannelById(uid string) (*channelResponse, error) {
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

	channel := &channelResponse{}
	channel.Id.ChannelId = xs.Items[0].Id
	channel.Snippet.Title = xs.Items[0].Snippet.Title
	channel.Snippet.Description = xs.Items[0].Snippet.Description
	channel.Snippet.Thumbnails.Default.Url = xs.Items[0].Snippet.Thumbnails.Default.Url

	return channel, nil
}

// getVideoByChannelId get videos by channel id
// eventType: none, live, upcoming, completed
// https://developers.google.com/youtube/v3/docs/search/list
// maxResults: 0-50, default 5
func (sy *SourceYoutube) getVideoByChannelId(channelId string, eventType string) ([]*videoResponse, error) {
	// default event type is none
	if eventType == "" {
		eventType = "none"
	} else if eventType != "live" && eventType != "upcoming" && eventType != "completed" {
		return nil, fmt.Errorf("invalid event type")
	}

	q := sy.yts.SearchList([]string{"snippet", "id"})

	callOption := []googleapi.CallOption{
		googleapi.QueryParameter("channelId", channelId),
		googleapi.QueryParameter("type", "video"),
		googleapi.QueryParameter("eventType", eventType),
		// default maxResults is 5
		googleapi.QueryParameter("maxResults", "10"),
	}

	xs, err := q.Do(callOption...)
	if err != nil {
		return nil, err
	}

	var videos []*videoResponse
	for _, item := range xs.Items {
		video := &videoResponse{}
		video.Id.VideoId = item.Id.VideoId
		video.Snippet.Title = item.Snippet.Title
		video.Snippet.Description = item.Snippet.Description
		video.Snippet.Thumbnails.Default.Url = item.Snippet.Thumbnails.Default.Url
		video.Snippet.ChannelId = item.Snippet.ChannelId

		videos = append(videos, video)
	}

	return videos, nil
}

func (sy *SourceYoutube) GetMediaUser(uid string) (*media.MediaUser, error) {
	var (
		channel *channelResponse
		err     error
		user    *media.MediaUser
	)
	// use search api if uid starts with @
	if strings.HasPrefix(uid, "@") {
		channel, err = sy.getChannelByName(uid)
		if err != nil {
			return nil, err
		}
	} else {
		channel, err = sy.getChannelById(uid)
		if err != nil {
			return nil, err
		}
	}

	currentStream := ""
	liveVideo, err := sy.getVideoByChannelId(channel.Id.ChannelId, "live")
	if err != nil {
		return nil, err
	}
	if len(liveVideo) > 0 {
		currentStream = liveVideo[0].Id.VideoId
	}

	ongoingStreamIds := []string{}
	ongoingStreams, err := sy.getVideoByChannelId(channel.Id.ChannelId, "upcoming")
	if err != nil {
		return nil, err
	}

	for _, stream := range ongoingStreams {
		if stream.Id.VideoId == currentStream {
			continue
		}
		ongoingStreamIds = append(ongoingStreamIds, stream.Id.VideoId)
	}

	user = &media.MediaUser{
		Id:             channel.Id.ChannelId,
		Name:           channel.Snippet.CustomUrl,
		Title:          channel.Snippet.Title,
		Description:    channel.Snippet.Description,
		Avatar:         channel.Snippet.Thumbnails.Default.Url,
		URL:            fmt.Sprintf("https://www.youtube.com/channel/%s", channel.Id.ChannelId),
		Platform:       "youtube",
		IsLive:         len(liveVideo) > 0,
		OngoingStreams: ongoingStreamIds,
		CurrentStream:  currentStream,
	}

	return user, nil
}
