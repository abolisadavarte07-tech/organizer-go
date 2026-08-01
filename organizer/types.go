package organizer

type Category struct {
	Name       string
	Folder     string
	Extensions []string
}

type FileInfo struct {
	Name      string
	Extension string
}