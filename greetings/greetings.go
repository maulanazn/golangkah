package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(RandomGreetings(), name)

	return message, nil
}

func HelloToPeople(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func RandomGreetings() string {
	formats := []string{
		"Hi %v. Welcome",
		"Good morning, %v",
		"Good night, %v",
	}

	return formats[rand.Intn(len(formats))]
}
