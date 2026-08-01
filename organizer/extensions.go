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
	{
	    Name:   "Word",
	    Folder: "Word",
	    Extensions: []string{
		    ".doc", ".docx",
	    },
    },
    {
	    Name:   "PowerPoints",
	    Folder: "PowerPoints",
	    Extensions: []string{
		    ".ppt", ".pptx",
	    },
    },
    {
	    Name:   "Excels",
	    Folder: "Excels",
	    Extensions: []string{
		    ".xls", ".xlsx",
	    },
    },
    {
	    Name:   "Publisher",
	    Folder: "Publisher",
	    Extensions: []string{
	    	".pub",
	    },
    },
    {
    	Name:   "Access",
	    Folder: "Access",
	    Extensions: []string{
		    ".accdb",
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