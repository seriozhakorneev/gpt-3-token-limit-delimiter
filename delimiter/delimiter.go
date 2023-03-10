package delimiter

import (
	"fmt"

	"github.com/seriozhakorneev/gpt-3-token-limit-divider/pkg/tokenizer"
)

type Delimiter struct {
	limit int
	e     tokenizer.Encoder
}

func New(limit int) (*Delimiter, error) {
	if limit < 1 {
		return &Delimiter{}, fmt.Errorf("token limit should be more than 0")
	}

	encoder, err := tokenizer.New()
	if err != nil {
		return &Delimiter{}, fmt.Errorf("New Delimiter - tokenizer.New: %w", err)
	}

	return &Delimiter{
		limit: limit,
		e:     encoder,
	}, nil
}

func (d *Delimiter) Split(s string) ([]string, error) {
	tokens, err := d.e.Encode(s)
	if err != nil {
		return nil, fmt.Errorf("Delimiter - Split - d.e.Encode: %w", err)
	}

	splitS := make([]string, (len(tokens)/d.limit)+1)

	for i, t := range d.limitSpl(tokens) {
		splitS[i] = d.e.Decode(t)
	}

	return splitS, nil
}

// limitSpl returns slice of tokens split by limit
func (d *Delimiter) limitSpl(blob []int) (split [][]int) {
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
