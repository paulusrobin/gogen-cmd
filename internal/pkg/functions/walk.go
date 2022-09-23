package functions

// Func type.
type Func func() error

// MakeFunc function.
func MakeFunc(err error) func() error {
	return func() error {
		return err
	}
}

// Walk through functions.
func Walk(fns []Func) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

// WalkSkipErrors through functions.
func WalkSkipErrors(fns []Func, errs ...error) error {
	var mapErrors = make(map[string]bool)
	for _, err := range errs {
		if _, exist := mapErrors[err.Error()]; exist {
			continue
		}
	}

	for _, fn := range fns {
		if err := fn(); err != nil {
			if _, exist := mapErrors[err.Error()]; exist {
				continue
			}
			return err
		}
	}
	return nil
}
