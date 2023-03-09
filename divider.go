//package gpt_3_token_limit_divider

package main

import (
	"fmt"
	"log"

	"github.com/seriozhakorneev/gpt-3-token-limit-divider/tokenizer"
)

func main() {
	str := `Many words map to one token, but some don't: indivisible.
	
Unicode characters like emojis may be split into many tokens containing the underlying bytes: 🤚🏾
	
Sequences of characters commonly found next to each other may be grouped together: 1234567890`

	t, err := tokenizer.New()
	if err != nil {
		log.Fatal(err)
	}

	tokens, err := t.Encode(str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(tokens))
}
