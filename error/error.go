package error

type IoError struct {
	Msg string
}

func (e IoError) Error() string {
	return e.Msg
}

type SamePathError struct {
	Msg string
}

func (e SamePathError) Error() string {
	return e.Msg
}
