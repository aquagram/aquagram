package aquagram

type FilterFunc = func(bot *Bot, event Event) (bool, error)

func BuildMiddleware(filter FilterFunc) Middleware {
	return func(next MiddlewareFunc) MiddlewareFunc {
		return func(bot *Bot, event Event) error {
			ok, err := filter(bot, event)
			if err != nil {
				return err
			}

			if !ok {
				return nil
			}

			return next(bot, event)
		}
	}
}

func And(a FilterFunc, b FilterFunc) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		aOut, err := a(bot, event)
		if err != nil {
			return false, err
		}

		bOut, err := b(bot, event)
		if err != nil {
			return false, err
		}

		return aOut && bOut, nil
	}
}

func Or(a FilterFunc, b FilterFunc) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		aOut, err := a(bot, event)
		if err != nil {
			return false, err
		}

		bOut, err := b(bot, event)
		if err != nil {
			return false, err
		}

		return aOut || bOut, nil
	}
}

func Not(a FilterFunc) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		aOut, err := a(bot, event)
		if err != nil {
			return false, err
		}

		return !aOut, nil
	}
}

func Nand(a FilterFunc, b FilterFunc) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		aOut, err := a(bot, event)
		if err != nil {
			return false, err
		}

		bOut, err := b(bot, event)
		if err != nil {
			return false, err
		}

		return !(aOut && bOut), nil
	}
}

func Xor(a FilterFunc, b FilterFunc) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		aOut, err := a(bot, event)
		if err != nil {
			return false, err
		}

		bOut, err := b(bot, event)
		if err != nil {
			return false, err
		}

		return (aOut && !bOut) || (!aOut && bOut), nil
	}
}

func AllOf(filters ...FilterFunc) FilterFunc {
	var fn FilterFunc

	for _, filter := range filters {
		if fn == nil {
			fn = filter
		} else {
			fn = And(fn, filter)
		}
	}

	return fn
}

func AnyOf(filters ...FilterFunc) FilterFunc {
	var fn FilterFunc

	for _, filter := range filters {
		if fn == nil {
			fn = filter
		} else {
			fn = Or(fn, filter)
		}
	}

	return fn
}

func NoneOf(filters ...FilterFunc) FilterFunc {
	var fn FilterFunc

	for _, filter := range filters {
		if fn == nil {
			fn = Not(filter)
		} else {
			fn = And(fn, Not(filter))
		}
	}

	return fn
}
