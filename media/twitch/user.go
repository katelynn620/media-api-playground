package twitch

import (
	"media-api-playground/media"

	"github.com/nicklaw5/helix/v2"
)

func (ts *TwitchSource) GetMediaUser(uid string) (*media.MediaUser, error) {
	resp, err := ts.s.GetUsers(&helix.UsersParams{
		Logins: []string{uid},
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Data.Users) == 0 {
		return nil, media.ErrMediaUserNotFound{}
	}

	return &media.MediaUser{
		Id:          resp.Data.Users[0].ID,
		Title:       resp.Data.Users[0].DisplayName,
		Description: resp.Data.Users[0].Description,
		Avatar:      resp.Data.Users[0].ProfileImageURL,
		URL:         baseURL + "/" + resp.Data.Users[0].Login,
		Platform:    platform,
	}, nil
}
