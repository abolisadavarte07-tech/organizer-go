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
		"Python",
		"Fonts",
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
}

