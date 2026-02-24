# BookGo - Write your book, one markdown file at a time

BookGo is a CLI tool that lets you write a book using simple markdown files and compile it into a PDF. This file is a quick reference — keep it around or delete it once you're comfortable.

---

## Project structure - {{ .BookName }}

```
{{ .BookPath }}/
├── book.yaml          # Book metadata (managed by BookGo - do not edit manually)
├── images/            # Put your images here
├── content/
│   ├── 1-chapter-introduction/
│   │   ├── my-first-section.md
│   │   └── another-section.md
│   └── 2-chapter-getting-started/
│       └── new-section.md
└── README.md          # This file
```

---

## Commands

### Create a new chapter

```bash
bookgo add-chapter <name>
```

**Aliases:** `ac`, `add-c`, `a-chapter`

Always use this command to create chapters - never create chapter folders by hand. BookGo names them with the correct prefix (`{number}-chapter-{name}`) so the build order stays correct.

> Example: `bookgo add-chapter The Beginning` creates `content/2-chapter-the-beginning/`

---

### Write sections

There is no command to create sections. Simply add `.md` files inside a chapter folder.

A few things to know:

- **Ordering** - Sections are compiled in the order the filesystem lists them. Prefix your filenames with numbers to control order: `01-intro.md`, `02-concepts.md`.
- **Naming** - The section title is derived from the filename: hyphens become spaces and each word is capitalized. `my-first-section.md` becomes _My First Section_.
- **Content** - Write standard markdown. Each chapter folder must contain at least one `.md` file.

---

### Build the book

```bash
bookgo build
```

**Alias:** `b`

Compiles your entire book into a PDF file, ready to read or share. It scans the `content/` directory, assembles all chapters and sections in order, and generates the final document.

---

## The `book.yaml` file

This file stores the metadata and structure of your book (name, author, description, chapters, sections). It is generated and updated automatically by BookGo.

**Do not edit it manually** - run `bookgo build` instead to keep it in sync.

---

## Tips

- Work inside the book folder when running commands (BookGo looks for `book.yaml` in the current directory).
- Rename section files carefully - the filename is the source of the section title.
- The `images/` folder is available for assets you want to reference in your markdown.

---

**Thank you for using BookGo, do not hesitate to send me some feedback on [github](https://github.com/noahlte/bookgo) !**
