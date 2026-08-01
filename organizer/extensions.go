package organizer

var Categories = []Category{
	{
		Name:   "Images",
		Folder: "Images",
		Extensions: []string{
			".jpg", ".jpeg", ".png", ".bmp", ".gif", ".webp",
		},
	},
	{
		Name:   "Audio",
		Folder: "Audio",
		Extensions: []string{
			".mp3", ".wav", ".ogg",
		},
	},
	{
		Name:   "PDFs",
		Folder: "PDFs",
		Extensions: []string{
			".pdf",
		},
	},
	{
		Name:   "Python",
		Folder: "Python",
		Extensions: []string{
			".py",
		},
	},
}

var ExtensionMap = make(map[string]string)

func init() {
	for _, category := range Categories {
		for _, ext := range category.Extensions {
			ExtensionMap[ext] = category.Folder
		}
	}
}