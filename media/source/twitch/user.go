package twitch

import (
	"media-api-playground/media"

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

	liveStatus := false
	if len(channelResp.Data.Streams) > 0 {
		liveStatus = true
	}

	return &media.MediaUser{
		Id:          userResp.Data.Users[0].ID,
		Title:       userResp.Data.Users[0].DisplayName,
		Description: userResp.Data.Users[0].Description,
		Avatar:      userResp.Data.Users[0].ProfileImageURL,
		URL:         baseURL + "/" + userResp.Data.Users[0].Login,
		Platform:    platform,
		IsLive:      liveStatus,
	}, nil
}
