# BookGo

> A CLI tool to write books in Markdown and compile them into PDF.

BookGo is a personal project built to practice Go while solving a real problem: I've always wanted a simple way to write books in Markdown and export them as proper PDFs. This tool lets you structure your content into chapters and sections, manage everything through a CLI, and eventually compile it all into a final document.

---

## Tech stack

| Tool | Purpose |
|---|---|
| [Go 1.25](https://go.dev/) | Main language |
| [Cobra](https://github.com/spf13/cobra) | CLI framework |
| [gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) | Book metadata serialization |
| [golang.org/x/text](https://pkg.go.dev/golang.org/x/text) | String utilities |
| Makefile | Build tooling |

---

## Architecture

```
bookgo/
├── cmd/bookgo/
│   └── main.go                 # Entrypoint
├── internal/
│   ├── command/                # Cobra command definitions
│   │   ├── command.go          # Root command
│   │   ├── setup.go            # bookgo new
│   │   ├── addchapter.go       # bookgo add-chapter
│   │   └── build.go            # bookgo build
│   ├── service/                # Business logic
│   │   ├── setup.go            # Book initialization
│   │   ├── addchapter.go       # Chapter creation
│   │   ├── build.go            # Build pipeline
│   │   └── templates/          # Embedded Go templates
│   │       ├── README.md       # Generated in each new book
│   │       └── new-section.md  # Generated for each new chapter
│   ├── book/
│   │   └── model.go            # Book, Chapter, Section structs + YAML marshaling
│   ├── filesystem/
│   │   └── book.go             # Filesystem helpers (book root detection)
│   └── util/
│       ├── constant.go         # Shared path constants
│       └── sanitize.go         # Name sanitization and capitalization
└── go.mod
```

The code is organized around a clean separation between commands (CLI layer) and services (logic layer). Commands parse user input and delegate to the corresponding service. The `book` package owns the data model and its persistence to `book.yaml` via YAML marshaling.

Templates are embedded directly into the binary using Go's `embed` package, so the CLI is fully self-contained with no external files needed at runtime.

---

## How it works

### 1. Creating a new book

```bash
bookgo new <name> --author <author> --description <description>
```

This creates a new directory with the following structure:

```
your-book/
├── book.yaml    # Auto-generated metadata file, do not edit
├── content/     # Your chapters go here
├── images/      # Assets referenced in your markdown
└── README.md    # Quick reference guide
```

The `book.yaml` file stores the book's metadata (name, author, description, creation date) and the full chapter/section tree. It is managed automatically by BookGo.

### 2. Adding chapters

```bash
bookgo add-chapter <name>
```

Creates a new folder inside `content/` named `{number}-chapter-{name}`, with a starter section file. The numbering is handled automatically based on the existing chapters.

Chapters must always be created with this command to ensure the folder naming convention is respected and the build order stays consistent.

### 3. Writing sections

Sections are plain `.md` files inside a chapter folder. There is no command for this, just create the files directly. A few conventions:

- **Order** - sections are compiled in filesystem order. Prefix filenames with numbers to control it: `01-intro.md`, `02-deep-dive.md`.
- **Title** - the section title is derived from the filename: hyphens become spaces and words are capitalized. `my-first-section.md` becomes *My First Section*.

### 4. Building the book

```bash
bookgo build
```

> **Work in progress.** The build command currently scans the `content/` directory and updates `book.yaml`. PDF generation is the next step.

---

## Status

BookGo is a work-in-progress learning project. Core features (project setup, chapter management, content scanning) are functional. PDF generation is still being implemented.

---

## Motivation

I started this project to get hands-on experience with Go, working with the standard library, structuring a real CLI application, and handling file I/O. Writing a book tool felt like the right scope: concrete enough to be useful, complex enough to be a good exercise.
