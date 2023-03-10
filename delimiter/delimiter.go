package delimiter

import (
	"fmt"

	"github.com/seriozhakorneev/gpt-text-delimiter/pkg/tokenizer"
)

type Delimiter struct {
	limit int
	e     tokenizer.Encoder
}

func New(limit int) (*Delimiter, error) {
	if limit < 1 {
		return &Delimiter{}, fmt.Errorf("token limit should be more than 0")
	}

	encoder, err := tokenizer.NewGPT3()
	if err != nil {
		return &Delimiter{}, fmt.Errorf("New Delimiter - tokenizer.New: %w", err)
	}

	return &Delimiter{
		limit: limit,
		e:     encoder,
	}, nil
}

// Split returns slice of strings split by token limit.
func (d *Delimiter) Split(s string) ([]string, error) {
	tokens, err := d.e.Encode(s)
	if err != nil {
		return nil, fmt.Errorf("Delimiter - Split - d.e.Encode: %w", err)
	}

	splitS := make([]string, (len(tokens)/d.limit)+1)

	for i, t := range d.TokenSplit(tokens) {
		splitS[i] = d.e.Decode(t)
	}

	return splitS, nil
}

// TokenSplit returns slice of tokens split by token limit.
func (d *Delimiter) TokenSplit(blob []int) (split [][]int) {
	if len(blob) < 1 {
		return nil
	}

	split = make([][]int, (len(blob)/d.limit)+1)

	for i, f, s := 0, 0, d.limit; f < len(blob); i, f, s = i+1, f+d.limit, s+d.limit {
		if s > len(blob) {
			split[i] = blob[f:]
			continue
		}

		split[i] = blob[f:s]
	}

	return split
}
