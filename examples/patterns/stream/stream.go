package stream

import (
	"errors"
	"fmt"
)

type Stream struct{}

func (s *Stream) Read(bb ...byte) error {
	if len(bb) == 0 {
		return errors.New("specified empty bytes to read")
	}
	for _, b := range bb {
		fmt.Print(string(b))
	}
	return nil
}
