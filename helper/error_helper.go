package helper

func ResultHandler(result any, err error) (any, error) {
	if err != nil {
		return nil, err
	}

	return result, nil
}
