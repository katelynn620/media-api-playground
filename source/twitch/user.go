package twitch

import (
	media "github.com/katelynn620/mediameta"

	"github.com/nicklaw5/helix/v2"
)

func (ts *TwitchSource) GetMediaUser(uid string) (*media.MediaUser, error) {
	userResp, err := ts.s.GetUsers(&helix.UsersParams{
		Logins: []string{uid},
	})
	if err != nil {
		return nil, err
	}

	if len(userResp.Data.Users) == 0 {
		return nil, media.ErrMediaUserNotFound{}
	}

	// use webpage to check if user is live instead of API?
	channelResp, err := ts.s.GetStreams(&helix.StreamsParams{
		UserIDs: []string{userResp.Data.Users[0].ID},
		First:   1,
	})
	if err != nil {
		return nil, err
	}

	currentStream := ""
	if len(channelResp.Data.Streams) > 0 {
		currentStream = channelResp.Data.Streams[0].ID
	}

	return &media.MediaUser{
		Id:            userResp.Data.Users[0].ID,
		Name:          userResp.Data.Users[0].Login,
		Title:         userResp.Data.Users[0].DisplayName,
		Description:   userResp.Data.Users[0].Description,
		Avatar:        userResp.Data.Users[0].ProfileImageURL,
		URL:           baseURL + "/" + userResp.Data.Users[0].Login,
		Platform:      platform,
		IsLive:        len(channelResp.Data.Streams) > 0,
		CurrentStream: currentStream,
	}, nil
}
