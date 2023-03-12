package delimiter

type Encoder interface {
	Encode(string) ([]int, error)
	Decode([]int) string
}
