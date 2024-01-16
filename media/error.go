package media

type ErrVideoNotFound struct{}
type ErrMediaUserNotFound struct{}

func (e ErrVideoNotFound) Error() string {
	return "video not found"
}

func (e ErrMediaUserNotFound) Error() string {
	return "media user not found"
}
