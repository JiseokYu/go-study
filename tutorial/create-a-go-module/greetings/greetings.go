package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// hello returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	// var message2 string
	// message2 = fmt.Sprintf("Hi, %v. Welcome too!", name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
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

func init() {
	rand.Seed((time.Now().UnixNano()))
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Greate to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
