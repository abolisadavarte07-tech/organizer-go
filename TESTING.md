# 🧪 Organizer CLI Testing Report

## Project

**Organizer CLI (Go)**

---

# Test Environment
_________________________________________
|    Parameter     |       Value        |
|------------------|--------------------|
| Operating System | Windows 10/11      |
| Language         | Go                 |
| Go Version       | Go 1.24+           |
| IDE              | Visual Studio Code |
| Test Folder      | TestFiles          |
-----------------------------------------
---

# Objective

The objective of testing was to verify that Organizer CLI correctly:

- Detects supported file types.
- Organizes files into appropriate folders.
- Handles duplicate filenames.
- Displays confirmation prompts.
- Generates movement logs.
- Executes all supported commands successfully.

---

# Test Cases

## 1. Image Command

**Command**

```bash
go run . image
```

**Expected Result**

- All supported image files should be moved to:

```
Organized/Images
```

**Result**

✅ Passed

---

## 2. Audio Command

**Command**

```bash
go run . audio
```

**Expected Result**

- Audio files moved to:

```
Organized/Audio
```

**Result**

✅ Passed

---

## 3. Video Command

**Command**

```bash
go run . video
```

**Expected Result**

- Video files moved to:

```
Organized/Videos
```

**Result**

✅ Passed

---

## 4. Text Command

**Command**

```bash
go run . text
```

**Expected Result**

- Text files moved to:

```
Organized/Texts
```

**Result**

✅ Passed

---

## 5. PDF Command

**Command**

```bash
go run . pdf
```

**Expected Result**

- PDF files moved to:

```
Organized/PDFs
```

**Result**

✅ Passed

---

## 6. Font Command

**Command**

```bash
go run . font
```

**Expected Result**

- Font files moved to:

```
Organized/Fonts
```

**Result**

✅ Passed

---

## 7. Vector Command

**Command**

```bash
go run . vector
```

**Expected Result**

- Vector graphics moved to:

```
Organized/Vectors
```

**Result**

✅ Passed

---

## 8. Photoshop Command

**Command**

```bash
go run . photoshop
```

**Expected Result**

- Photoshop files moved to:

```
Organized/Photoshop
```

**Result**

✅ Passed

---

## 9. Office Command

**Command**

```bash
go run . office
```

**Expected Result**

- Word documents
- Excel spreadsheets
- PowerPoint presentations
- Publisher files
- Access databases

should be organized into their respective folders.

**Result**

✅ Passed

---

## 10. Code Command

**Command**

```bash
go run . code
```

**Expected Result**

Programming files should be organized into:

- HTML
- CSS
- JavaScript
- Java
- PHP
- C
- C++
- Swift
- Visual Basic
- Python

**Result**

✅ Passed

---

## 11. Program Command

**Command**

```bash
go run . program
```

**Expected Result**

Executable files should be organized into:

- Executables
- APKs

**Result**

✅ Passed

---

## 12. Safe Command

**Command**

```bash
go run . safe
```

**Expected Result**

Only predefined safe file categories should be organized.

**Result**

✅ Passed

---

## 13. All Command

**Command**

```bash
go run . all
```

**Expected Result**

All supported files should be organized into their corresponding folders.

**Result**

✅ Passed

---

# Functional Testing
___________________________________________
|           Feature           | Status     |
|-----------------------------|------------|
| Directory scanning          | ✅ Passed |
| Extension detection         | ✅ Passed |
| Category identification     | ✅ Passed |
| Folder creation             | ✅ Passed |
| File movement               | ✅ Passed |
| Duplicate filename handling | ✅ Passed |
| Confirmation prompt         | ✅ Passed |
| Logging                     | ✅ Passed |
| Unsupported file handling   | ✅ Passed |
| CLI output                  | ✅ Passed |
-------------------------------------------
---

# Edge Cases Tested

### Empty Directory

Expected

```
No matching files found.
```

Result

✅ Passed

---

### Unsupported Files

Expected

Unsupported files remain in the original directory.

Result

✅ Passed

---

### Existing Destination Folder

Expected

Application should reuse the existing folder.

Result

✅ Passed

---

### Duplicate Files

Expected

Duplicate filenames should be handled safely without overwriting existing files.

Result

✅ Passed

---

### Log File

Expected

`Moved-Files-Log.txt` should remain in the root directory and should not be reorganized.

Result

✅ Passed

---

# Test Summary
__________________________________
|      Category      |   Result   |
|--------------------|------------|
| Command Testing    | ✅ Passed |
| Functional Testing | ✅ Passed |
| Edge Case Testing  | ✅ Passed |
| Logging            | ✅ Passed |
| Duplicate Handling | ✅ Passed |
----------------------------------
---

# Conclusion

All implemented commands and features were tested successfully. The application correctly organized supported file types into their designated folders, handled duplicate filenames safely, generated movement logs, and provided an interactive command-line interface. The testing confirms that the Organizer CLI is stable and ready for deployment and demonstration.