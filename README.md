# 📂 Organizer CLI (Go)

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-Completed-brightgreen)
![Platform](https://img.shields.io/badge/Platform-Windows-blue)

A fast, lightweight, and modular **Command Line File Organizer** built in **Go** that automatically organizes files into categorized folders based on their extensions.

This project is a Go reimplementation of the original Organizer CLI, redesigned using Go's modular architecture while preserving the original command-line experience.

---

# 📌 Table of Contents

- Overview
- Features
- Technologies Used
- Project Architecture
- Project Structure
- Supported Commands
- Supported File Types
- Installation
- Usage
- Example Output
- Logging
- Testing
- Benchmark
- Future Enhancements
- License
- Author

---

# 📖 Overview

Managing files in folders such as **Downloads**, **Desktop**, or **Documents** often becomes difficult as files accumulate over time.

Organizer CLI automatically detects supported file types and organizes them into dedicated folders such as Images, Audio, Videos, Office Documents, Programming Files, Executables, Fonts, Vector Graphics, and many more.

The application provides an interactive command-line interface with confirmation prompts, duplicate filename handling, and detailed movement logging.

---

# ✨ Features

- Automatically organizes files into categorized folders
- Interactive confirmation prompt before moving files
- Supports 26+ file types
- Multiple organization commands
- Duplicate filename handling
- Automatic movement logging
- Clean modular Go architecture
- Lightweight and fast execution
- Easy to extend with new categories

---

# 🛠 Technologies Used

- Go
- Go Standard Library
- File System APIs
- Git
- GitHub

---

# 🏗 Project Architecture

```
                 +----------------+
                 |    main.go     |
                 +-------+--------+
                         |
                         v
                 +----------------+
                 | Organize()     |
                 +-------+--------+
                         |
        +----------------+----------------+
        |                                 |
        v                                 v
 ScanDirectory()                  CountMovableFiles()
        |
        v
 GetCategory()
        |
        v
 MoveFile()
        |
        v
 LogMove()
```

---

# 📁 Project Structure

```
organizer-go/
│
├── organizer/
│   ├── commands.go
│   ├── extensions.go
│   ├── logger.go
│   ├── organize.go
│   ├── types.go
│   └── utils.go
│
├── tests/
│
├── README.md
├── BENCHMARK.md
├── TESTING.md
├── LICENSE
├── go.mod
├── main.go
└── .gitignore
```

---

# 🚀 Supported Commands
__________________________________________________________________
|         Command         |              Description              |
|-------------------------|---------------------------------------|
| `all`                   | Organize all supported files          |
| `safe`                  | Organize common file types            |
| `image`                 | Organize image files                  |
| `audio`                 | Organize audio files                  |
| `video`                 | Organize video files                  |
| `text`                  | Organize text files                   |
| `pdf`                   | Organize PDF files                    |
| `font`                  | Organize font files                   |
| `vector`                | Organize vector graphics              |
| `gif`                   | Organize GIF files                    |
| `photoshop`             | Organize Photoshop files              |
| `office`                | Organize Microsoft Office files       |
| `code`                  | Organize programming files            |
| `python`                | Organize Python files                 |
| `program`               | Organize executable and APK files     |
-------------------------------------------------------------------
---

# 📄 Supported File Types

### Images

`.jpg` `.jpeg` `.png` `.bmp` `.gif` `.webp`

### Audio

`.mp3` `.wav` `.ogg`

### Videos

`.mp4` `.avi` `.mov`

### Documents

`.txt` `.pdf`

### Fonts

`.ttf` `.otf`

### Vector Graphics

`.svg` `.ai`

### Photoshop

`.psd`

### Microsoft Office

`.doc` `.docx` `.ppt` `.pptx` `.xls` `.xlsx` `.pub` `.accdb`

### Programming Files

`.html` `.css` `.js` `.java` `.php` `.c` `.cpp` `.swift` `.vb` `.py`

### Programs

`.exe` `.msi` `.apk`

---

# ⚙️ Installation

Clone the repository:

```bash
git clone https://github.com/abolisadavarte07-tech/organizer-go.git
```

Navigate to the project:

```bash
cd organizer-go
```

Run the application:

```bash
go run . all
```

Or build an executable:

```bash
go build -o organizer.exe
```

Run the executable:

```bash
organizer.exe all
```

---

# 🚀 Usage

```bash
go run . image
```

```bash
go run . office
```

```bash
go run . code
```

```bash
go run . program
```

```bash
go run . safe
```

```bash
go run . all
```

---

# 💻 Example Output

```text
ATTENTION: 26 file(s) will be moved.
Continue? (y/n): y

Organizing files...
----------------------------------------
✓ photo.jpg            -> Images
✓ song.mp3             -> Audio
✓ report.docx          -> Word
✓ installer.exe        -> Executables
...
----------------------------------------
Organization completed successfully!
26 file(s) moved.
Log saved to: TestFiles/Moved-Files-Log.txt
----------------------------------------
```

---

# 📝 Logging

Each successful execution generates a log file named:

```
Moved-Files-Log.txt
```

Example:

```text
photo.jpg -> Images
resume.pdf -> PDFs
song.mp3 -> Audio
report.docx -> Word
installer.exe -> Executables
```

---

# 🧪 Testing

The project has been tested with:

- ✅ Image files
- ✅ Audio files
- ✅ Video files
- ✅ Text files
- ✅ PDF files
- ✅ Font files
- ✅ Vector graphics
- ✅ Photoshop files
- ✅ Microsoft Office documents
- ✅ Programming files
- ✅ Executables
- ✅ APK files
- ✅ Duplicate filename handling
- ✅ Confirmation prompt
- ✅ Logging functionality
- ✅ Mixed file organization using the `all` command

---

# 📊 Benchmark

Compared with the original Python implementation:
___________________________________________________
|     Python Version      |      Go Version       |
|-------------------------|-----------------------|
| Interpreted             |  Compiled             |
| Requires Python runtime | Standalone executable |
| Dynamic typing          | Static typing         |
| Multiple dictionaries   | Modular architecture  |
| Runtime dependency      | Minimal dependencies  |
| Slower startup          | Faster startup        |
---------------------------------------------------
---

# 🔮 Future Enhancements

- Recursive directory organization
- User-defined categories
- Configuration file support
- Undo last organization
- Drag-and-drop support
- GUI version
- Concurrent file processing
- Cross-platform installers
- Automatic duplicate conflict resolution

---

# 📜 License

This project is licensed under the **MIT License**.

---

# 👩‍💻 Author

**Aboli Sadavarte**

GitHub: https://github.com/abolisadavarte07-tech

---

⭐ If you found this project useful, consider giving it a star on GitHub.