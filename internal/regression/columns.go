package regression

var (
	// Removed rooms and age
	ColumnsA = []string{"neighborhood", "crim", "zn", "indus", "chas", "nox", "dis", "rad", "tax", "ptratio", "lstat", "mv"}
	// Removed tax and crim
	ColumnsB = []string{"neighborhood", "zn", "indus", "chas", "nox", "rooms", "age", "dis", "rad", "ptratio", "lstat", "mv"}
)
