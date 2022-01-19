package utils

func DealErrors(errors ...error) {
	for _, err := range errors {
		if err != nil {
			panic(err)
		}
	}
}
