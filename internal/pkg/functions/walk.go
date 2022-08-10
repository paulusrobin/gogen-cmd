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
