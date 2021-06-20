package video

type Video struct {
	ID       int32
	Title    string
	Filename string
	data     []byte
}

func (v *Video) Write(p []byte) (int, error) {
	if v.data == nil {
		v.data = []byte{}
	}

	v.data = append(v.data, p...)
	return len(p), nil
}

func (v *Video) ReadAll() ([]byte, error) {
	return v.data, nil
}
