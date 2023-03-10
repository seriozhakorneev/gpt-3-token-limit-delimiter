package tokenizer

import (
	"fmt"

	gpt3encoder "github.com/samber/go-gpt-3-encoder"
)

type Encoder interface {
	Encode(string) ([]int, error)
	Decode([]int) string
}

// GPT3 encoder wrapper, current realisation: github.com/samber/go-gpt-3-encoder v0.3.1
type GPT3 struct {
	e *gpt3encoder.Encoder
}

func New() (*GPT3, error) {
	encoder, err := gpt3encoder.NewEncoder()
	if err != nil {
		return nil, fmt.Errorf("encoder - New-gpt3encoder.NewEncoder: %w", err)
	}

	return &GPT3{e: encoder}, nil
}

func (g *GPT3) Encode(s string) ([]int, error) {
	t, err := g.e.Encode(s)
	if err != nil {
		return nil, fmt.Errorf("len - encoder.Encode: %w", err)
	}

	return t, nil
}

func (g *GPT3) Decode(a []int) string {
	return g.e.Decode(a)
}
