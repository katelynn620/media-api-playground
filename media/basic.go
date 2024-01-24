package media

type Video struct {
	Id          string
	Title       string
	Description string
	Thumbnail   string
	URL         string
	IsLive      bool
	Platform    string
}

type MediaUser struct {
	Id             string
	Name           string
	Title          string
	Description    string
	Avatar         string
	URL            string
	Platform       string
	IsLive         bool
	OngoingStreams []string
	CurrentStream  string
}

type Source interface {
	GetVideo(vid string) (*Video, error)
	GetMediaUser(uid string) (*MediaUser, error)
}
