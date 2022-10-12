package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func lint(file string) error {
	content, err := readFile(file)

	if err != nil {
		return err
	}

	dec := json.NewDecoder(strings.NewReader(content))
	var lastToken json.Token
	var currentToken json.Token

	for {
		lastToken = currentToken
		currentToken, err = dec.Token()

		if err != nil {
			if err == io.EOF {
				break
			}

			if lastToken == nil {
				return err
			} else {
				return fmt.Errorf("invalid token near \"%v\", %w", lastToken, err)
			}
		}
	}

	return nil
}
