# Memepulation

**Memepulation** is a Command-Line Interface (CLI) tool for performing various text file manipulations such as detecting duplicates, removing duplicates, and more.

## Features

- View the content of a file.
- Count and display repeated lines.
- Create a new file with no duplicates.
- View file statistics such as total lines, empty lines, unique lines, and duplicate lines.


## Installation via Script (Linux/macOS)

Run the following command in your terminal to install Memepulation:

```bash
curl -fsSL https://raw.githubusercontent.com/sharon-xa/memepulation/main/install.sh | bash
```

## Usage

### Command Line Usage

```bash
memepulation [flags] [file]
```

### Flags

| Flag        | Description                                           |
|-------------|-------------------------------------------------------|
| `-s`        | Show file content.                                    |
| `-r`        | Show the number of repeated lines in the file.        |
| `-sr`       | Show all repeated lines in the file.                  |
| `-n`        | Create a new file without duplicates.                 |
| `-h`        | Show the help message with available operations.      |
| `-stats`    | Display file statistics (total, unique, empty, and duplicate lines). |
| `-q`        | Exit the program.                                     |

### Examples

To view the file content:
```bash
memepulation -s example.txt
```

To show the number of repeated lines:
```bash
memepulation -r example.txt
```

To create a new file without duplicates:
```bash
memepulation -n example.txt
```

To show statistics about the file:
```bash
memepulation -stats example.txt
```

To see the help message:
```bash
memepulation -h
```

### Example File Manipulation

If `example.txt` contains:
```
apple
banana
apple
orange
banana
apple
```

Running:
```bash
memepulation -r example.txt
```

Output:
```
⚠️  There are 3 repeated lines in the file.
```

Running:
```bash
memepulation -stats example.txt
```

Output:
```
📊 File statistics:
- Total lines: 6
- Empty lines: 0
- Unique lines: 3
- Duplicate lines: 3
```

