package errors

type Kind string

type Error struct {
	Kind Kind
	Err error
}

func (e Error) Error() string {
	return string(e.Kind)
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Kind:
			e.Kind = arg
		case error:
			e.Err = arg
		default:
			panic("unknown behaviour while constructing Error struct")
		}
	}

	return e
}

func GetKind(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return UnexpectedErr
	}

	return e.Kind
}

