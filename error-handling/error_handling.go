package erratum

// Use a resource
func Use(o ResourceOpener, input string) (result error) {

	var resource Resource

	defer func() {
		if r := recover(); r != nil {
			err, frobError := r.(FrobError)
			if frobError {
				resource.Defrob(err.defrobTag)
				resource.Close()
			}
			result = err
		}
	}()

	resource, err := o()

	if err != nil {
		return err
	}
	resource.Frob(input)
	resource.Close()

	return nil
}
