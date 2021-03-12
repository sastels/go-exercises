package erratum

// Use a resource
func Use(o ResourceOpener, input string) (result error) {
	var resource Resource
	var err error

	defer func() {
		if r := recover(); r != nil {
			frobError, isFrobError := r.(FrobError)
			if isFrobError {
				resource.Defrob(frobError.defrobTag)
			}
			resource.Close()
			result = r.(error)
		}
	}()

	for {
		resource, err = o()
		if err != nil {
			_, isTransientError := err.(TransientError)
			if !isTransientError {
				return err
			}
		} else {
			break
		}
	}
	resource.Frob(input)
	resource.Close()
	return nil
}
