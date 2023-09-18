package helper

type showOutput func(string) string

var ShowName showOutput = func(name string) string {
	return name
}
