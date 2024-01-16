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
	Id          string
	Title       string
	Description string
	Avatar      string
	URL         string
	Platform    string
}

type Source interface {
	GetVideo(vid string) (*Video, error)
	GetMediaUser(uid string) (*MediaUser, error)
}

type MediaService struct {
	source Source
}

func NewMediaService(source Source) *MediaService {
	return &MediaService{source: source}
}

func (ms *MediaService) GetVideo(id string) (*Video, error) {
	return ms.source.GetVideo(id)
}

func (ms *MediaService) GetMediaUser(id string) (*MediaUser, error) {
	return ms.source.GetMediaUser(id)
}
