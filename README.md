# ASCII Art FS (Go)

## Description

This project is a simple ASCII Art generator written in Go.  
It reads a text input and converts it into ASCII art using predefined banner templates stored in files.

The program supports multiple banner styles such as:

- `standard`
- `shadow`
- `thinkertoy`

---

## Features

- Converts any printable ASCII text into ASCII art
- Supports multiple banner templates
- Handles multi-line input using `\n`
- Validates input for non-printable characters
- Reads banner files using Go's filesystem API
- Clean modular structure (validation, rendering, loading, generation)

---

## Usage

### Basic format

```bash
go run . [STRING] [BANNER]