package organizer

var Commands = map[string][]string{

	// Organize every supported category
	"all": {
        "Images",
        "Audio",
        "Videos",
        "Vectors",
        "GIFs",
        "Photoshop",
        "Texts",
        "PDFs",
        "Fonts",

        "Word",
        "PowerPoints",
        "Excels",
        "Publisher",
        "Access",

        "HTML",
        "CSS",
        "JavaScript",
        "Java",
        "PHP",
        "C",
        "C++",
        "Swift",
        "Visual Basic",
        "Python",

        "Executables",
        "APKs",
    },

	// Similar to the original Python "safe" command
	"safe": {
		"Images",
		"Audio",
		"Videos",
		"Texts",
		"PDFs",
		"Fonts",
	},

	"image": {
		"Images",
	},

	"audio": {
		"Audio",
	},

	"video": {
		"Videos",
	},

	"text": {
		"Texts",
	},

	"vector": {
		"Vectors",
	},

	"gif": {
		"GIFs",
	},

	"photoshop": {
		"Photoshop",
	},

	"pdf": {
		"PDFs",
	},

	"python": {
		"Python",
	},

	"font": {
		"Fonts",
	},

	"office": {
	"Word",
	"PowerPoints",
	"Excels",
	"Publisher",
	"Access",
    },

	"code": {
    "HTML",
    "CSS",
    "JavaScript",
    "Java",
    "PHP",
    "C",
    "C++",
    "Swift",
    "Visual Basic",
    "Python",
    },

	"program": {
	"Executables",
	"APKs",
    },
}

