# Memepulation

**Memepulation** is a Command-Line Interface (CLI) tool for performing various text file manipulations such as viewing file content, detecting repeated lines, displaying statistics, and more.

## Features

- Show the content of a file.
- Display all repeated lines or the number of repeated lines.
- Display various file statistics (total lines, unique lines, empty lines, duplicate lines, size, and bytes).
- Autocompletion support for Bash, Zsh, Fish, and PowerShell.

## Installation via Script (Linux/macOS)

Run the following command in your terminal to install Memepulation:

```bash
curl -fsSL https://raw.githubusercontent.com/sharon-xa/memepulation/main/install.sh | sudo bash
```

## Usage

### Command Line Usage

```bash
memepulation [flags] <command>
```

### Global Flags

| Flag         | Description                     |
| ------------ | ------------------------------- |
| `-f, --file` | Specify the file to manipulate. |

### Commands and Their Flags

#### `show`
Display the file content or repeated lines.

**Flags:**
| Flag                    | Description                          |
| ----------------------- | ------------------------------------ |
| `-c, --content`         | Show file content.                   |
| `-r, --repeated`        | Show all repeated lines in the file. |
| `-n, --repeated-number` | Show the number of repeated lines.   |

**Example:**
```bash
memepulation -f text-files/names.txt show --content
memepulation -f text-files/names.txt show --repeated
```

#### `stats`
Display file statistics, including total, unique, empty, duplicate lines, size, and bytes.

**Flags:**
| Flag              | Description                           |
| ----------------- | ------------------------------------- |
| `-t, --total`     | Show the total number of lines.       |
| `-u, --unique`    | Show the number of unique lines.      |
| `-d, --duplicate` | Show the number of duplicate lines.   |
| `-e, --empty`     | Show the number of empty lines.       |
| `-s, --size`      | Show the size of the file.            |
| `-b, --bytes`     | Show the number of bytes in the file. |
| `-a, --all`       | Display all available statistics.     |

**Example:**
```bash
memepulation -f text-files/names.txt stats --all
```

#### `completion`
Generate shell completion scripts for Bash, Zsh, Fish, or PowerShell.

**Usage:**
```bash
memepulation completion <shell>
```

**Example:**
```bash
memepulation completion bash > /etc/bash_completion.d/memepulation
```

### Example Scenarios

#### Show File Content
```bash
memepulation -f example.txt show --content
```

Output:
```
File content:
line1
line2
line3
```

#### Display Repeated Lines
```bash
memepulation -f example.txt show --repeated
```

Output:
```
Repeated lines:
line1
line3
```

#### Display Statistics
```bash
memepulation -f example.txt stats --all
```

Output:
```
📊 File statistics:
- Total lines: 6
- Unique lines: 3
- Duplicate lines: 2
- Empty lines: 0
- Size: 24 bytes
```

## Autocompletion

Memepulation supports shell autocompletion. To enable it, use the `completion` command as shown below:

- **Bash:**
  ```bash
  memepulation completion bash > /etc/bash_completion.d/memepulation
  ```
- **Zsh:**
  ```bash
  memepulation completion zsh > "${fpath[1]}/_memepulation"
  ```
- **Fish:**
  ```bash
  memepulation completion fish > ~/.config/fish/completions/memepulation.fish
  ```
- **PowerShell:**
  ```powershell
  memepulation completion powershell > memepulation.ps1
  ```

## Contributing

Feel free to submit issues or pull requests on the [GitHub repository](https://github.com/sharon-xa/memepulation).
