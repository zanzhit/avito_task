package errs

type ErrBannerNotFound struct{}

func (e *ErrBannerNotFound) Error() string {
	return "banner not found"
}

type ErrBannerNotUnique struct{}

func (e *ErrBannerNotUnique) Error() string {
	return "banner not unique"
}
