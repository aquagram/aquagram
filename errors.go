package aquagram

import (
	"errors"
	"fmt"
)

var (
	ErrStopPropagation = errors.New("stop event propagation")

	// user errors
	ErrUserError         = errors.New("user error")
	ErrEmptyToken        = fmt.Errorf("%w: empty bot token", ErrUserError)
	ErrUnknownFileSource = fmt.Errorf("%w: unknown file source", ErrUserError)
	ErrUnknownMarkup     = fmt.Errorf("%w: unknown reply markup", ErrUserError)

	// telegram errors
	ErrTelegramError = errors.New("telegram error")
	ErrTgBadRequest  = fmt.Errorf("%w: bad request", ErrTelegramError)
	ErrExpectedTrue  = fmt.Errorf("%w: the result is not true", ErrTelegramError)

	ErrUpdaterError = errors.New("updater error")
)
