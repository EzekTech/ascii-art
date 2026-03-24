# ASCII Art Generator in Go

## Overview

This Go program takes text input and converts it into **ASCII art** using a banner file (`standard.txt`). ASCII art is a way of creating pictures or stylized text using normal characters.

The program can handle:

* Regular text (e.g., "Hello")
* Uppercase, lowercase letters
* Numbers and symbols (if present in your banner file)
* Newline characters (`\n`) to separate lines

---

## Files

* **main.go** – The main program file.
* **standard.txt** – The banner file that contains ASCII art for all characters.
* **README.doc** – This explanation file.

---

## Usage

1. Make sure you have Go installed.
2. Place `standard.txt` in the same folder as `main.go`.
3. Open a terminal and run the program with:

```
go run . "Hello"
```

* To include a newline in your text, use `\n`:

```
go run . "Hello\nWorld"
```

---

## How It Works

1. **Read input from the command line**

   * `os.Args[1]` gets the text the user typed.

2. **Read the banner file**

   * The program reads `standard.txt` using `os.ReadFile`.
   * Each line in the file represents part of a character in ASCII art.

3. **Split the banner into lines**

   * `strings.Split` converts the file into a slice of strings (lines).

4. **Convert each character to ASCII art**

   * The program loops through each character in the input.
   * It calculates which lines in the banner file correspond to that character.
   * Then it builds the ASCII version using `strings.Builder`.

5. **Handle newlines**

   * If the input contains `\n`, the program starts a new block of ASCII art.

6. **Print the result**

   * Finally, the program prints the ASCII art to the terminal.

---

## Examples

```
go run . "Hello"
```

Output (example, depends on your banner file):

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/  
                                 
                                 
```

```
go run . "Hello\nWorld"
```

Output (two blocks separated by a blank line):

```
[ASCII for Hello]
[blank line]
[ASCII for World]
```

---

## Notes

* Only characters present in the banner file will display correctly.
* Spaces are handled automatically.
* For multiple lines, always use `\n` in the input string.

---

## License

This project is free to use and modify.
