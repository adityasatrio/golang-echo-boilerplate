package helper

//goland:noinspection GoUnusedExportedFunction
func ResultHandler(result any, err error) (any, error) {
	if err != nil {
		return nil, err
	}

	return result, nil
}
