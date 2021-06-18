package video

type Video struct {
	ID    int32
	Title string
	data  []byte
}

func (v *Video) Write(p []byte) (int, error) {
	v.data = p
	return len(v.data), nil
}

func (v *Video) ReadAll() ([]byte, error) {
	return v.data, nil
}
