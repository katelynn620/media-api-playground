package media

type MediaService struct {
	source Source
}

func (ms *MediaService) GetVideo(id string) (*Video, error) {
	return ms.source.GetVideo(id)
}

func (ms *MediaService) GetMediaUser(id string) (*MediaUser, error) {
	return ms.source.GetMediaUser(id)
}

func NewMediaService(source Source) (*MediaService, error) {
	return &MediaService{
		source: source,
	}, nil
}
