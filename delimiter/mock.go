package delimiter

type MockEncoder struct {
	EncodeFunc func(string) ([]int, error)
	DecodeFunc func([]int) string
}

func (m MockEncoder) Encode(s string) ([]int, error) {
	if m.EncodeFunc != nil {
		return m.EncodeFunc(s)
	}
	return nil, nil
}

func (m MockEncoder) Decode(a []int) string {
	if m.DecodeFunc != nil {
		return m.DecodeFunc(a)
	}
	return ""
}
