package organizer

// Category represents a group of file extensions
// and the destination folder where they should be moved.
type Category struct {
	Name       string
	Folder     string
	Extensions []string
}