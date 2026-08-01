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
	{
	    Name:   "HTML",
	    Folder: "HTML",
	    Extensions: []string{
	    	".html", ".htm",
	    },
    },
	{
	    Name:   "CSS",
	    Folder: "CSS",
	    Extensions: []string{
	    	".css",
	    },
    },
	{
	    Name:   "JavaScript",
	    Folder: "JavaScript",
	    Extensions: []string{
		    ".js",
	    },
    },
	{
	    Name:   "Java",
	    Folder: "Java",
	    Extensions: []string{
		    ".java", ".class",
	    },
    },
	{
	    Name:   "PHP",
	    Folder: "PHP",
	    Extensions: []string{
	    	".php",
    	},
    },
	{
    	Name:   "C",
    	Folder: "C",
    	Extensions: []string{
	    	".c",
	    },
    },
	{
	    Name:   "C++",
    	Folder: "C++",
    	Extensions: []string{
	    	".cpp", ".cc", ".cxx",
	    },
    },
	{
	    Name:   "Swift",
	    Folder: "Swift",
	    Extensions: []string{
	    	".swift",
	    },
    },
	{
	    Name:   "Visual Basic",
	    Folder: "Visual Basic",
	    Extensions: []string{
		    ".vb",
	    },
    },
	{
	    Name:   "Executables",
    	Folder: "Executables",
    	Extensions: []string{
	    	".exe", ".msi",
    	},
    },
    {
    	Name:   "APKs",
    	Folder: "APKs",
    	Extensions: []string{
  	    	".apk",
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