package core

import "github.com/pkg/errors"

func New(msg string) error {
	return errors.New(msg)
}
