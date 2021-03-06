package loaders

// A Loader loads spec content from the provided location, sending each line
// to the returned channel. If any error is received on the provided channel,
// the Loader will halt and close the output channel.
type Loader interface {
	Load(string) ([]byte, error)
	LoadAll(chan string, chan error) chan File
}

// A File is a wrapper for a file's name and contents.
type File struct {
	Name string
	Body []string
}
