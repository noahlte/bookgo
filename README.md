# BookGo - A simple book creator application in Go

BookGo is an CLI application that help you create books in Markdown format. Once your book is finished, you can simply convert it into a full pdf file by saying `bookgo build`.

## 1. Technology used

- Golang 1.25.4
- Makefile

## 2. How to use it

### a. Create a new book

At first, you will need a new project, for that open your terminal and write this command :

```bash
bookgo new <book-name> --author <author-name> --description <description>
```

Once your project created you should have a file architecture like this :

```
book-name
  └── content
  └── images
  book.yaml
```

**`content`** is where you will add the chapter of your book with the **`add-chapter`** command.

**`images`** is where you will stored the image you used for your book

**`book.yaml`** **DO NOT TOUCH**, it's an auto-generated file that will help the program for the build.

### b. Add a new chapter

To add a new chapter simply write :

```bash
bookgo add-chapter <name> --description <description>
```

This will create a new folder called `x-chapter-name` inside the `content` folder. The x is for the number of the chapter.

Inside this folder you will find an auto-generated section. The section are the core of your books. You can have multiple section file by chapters.

### c. Finaly build your books

Work in progress...
