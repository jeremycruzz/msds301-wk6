package regression

var (
	// Removed rooms and age
	OmitA = map[string]bool{"rooms": true, "age": true}
	// Removed tax and crim
	OmitB = map[string]bool{"tax": true, "crim": true}
)
