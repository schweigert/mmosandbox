package dont

// Panic errors
func Panic(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}
