# ASCII Art FS

## Description

ASCII Art FS is a Go program that generates ASCII art from a given string using predefined banner files.

The program extends the basic ASCII-Art functionality by introducing file system handling, allowing banner templates to be loaded dynamically from external files.

---

## Author

- Ooja Omale

---

## Usage

### Display ASCII Art in the Terminal

```bash
go run . "hello" standard
```
Using a Different Banner
go run . "Hello There!" shadow
View Output in Terminal (with piping)
go run . "hello" standard | cat -e

**Implementation Details**

1. Argument Validation

The program validates command-line arguments to ensure correct usage.

Supported formats:

go run . "STRING"
go run . "STRING" BANNER

Any invalid format prints a usage message:

Usage: go run . [STRING] [BANNER]
EX: go run . something standard
2. Banner Loading

Banner files are loaded from the banners/ directory using Go’s file system API (os.ReadFile).

Each printable ASCII character (ASCII 32–126) is mapped to its 8-line ASCII representation and stored in:

map[rune][]string

This allows efficient lookup during rendering.

3. Input Validation

All characters in the input string are validated to ensure they fall within the printable ASCII range (32–126).

If an invalid character is found, the program stops execution and returns an error message.

4. ASCII Art Rendering

The input string is processed character by character.

For each character:

The corresponding 8-line ASCII pattern is retrieved from the banner map
Each line is concatenated horizontally with other characters
This produces the final ASCII-art output
5. File System Handling

The program reads banner templates directly from external files located in the banners/ directory.

This demonstrates the use of Go’s file system API for:

Reading files
Parsing structured text data
Mapping file content into usable program structures
Allowed Packages

Only standard Go packages are used.

Example
go run . "hello" standard

Output:

 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
Author
Ooja Omale

---