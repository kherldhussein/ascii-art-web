# ASCII Art Output

This project involves creating a Go program that generates ASCII art based on a given string and banner style, and writes the result into a file when output flag is specified.Additionally, it runs with a single [STRING] argument.

## Features

- Converts strings into ASCII art
- Supports numbers, letters, spaces, special characters, and newline characters ('\n')
- Utilizes specific graphical templates for ASCII representation
- Writes the ascii art into a file if output flag is specified

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/doonyango/ascii-art-output.git
    ```

2. Navigate to the project directory:

    ```bash
    cd ascii-art-output/
    ```


## Usage

To generate ASCII art for a string or write the ascii art into a file, run the following command:

```bash
go run . "string"
```

Example:

```bash
go run . "Hello\n" | cat -e
```

Output:

```ruby
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$                                                  
```
```
go run . --output=<fileName.txt> something standard
```
Output

cat -e fileName.txt
```
                                  _     _       _                  
                                 | |   | |     (_)                 
 ___    ___    _ __ ___     ___  | |_  | |__    _   _ __     __ _  
/ __|  / _ \  | '_ ` _ \   / _ \ | __| |  _ \  | | | '_ \   / _` | 
\__ \ | (_) | | | | | | | |  __/ \ |_  | | | | | | | | | | | (_| | 
|___/  \___/  |_| |_| |_|  \___|  \__| |_| |_| |_| |_| |_|  \__, | 
                                                             __/ | 
                                                            |___/  
```

### Testing
```
cd tests/
```
```
go test -v
```



## File Formats

- `standard.txt`: Standard ASCII character set
- `shadow.txt`: Shadowed ASCII character set
- `thinkertoy.txt`: ASCII character set with thinkertoy style

## File Integrity Verification

This program ensures file integrity using SHA-256 checksums. When downloading or verifying files (standard.txt, shadow.txt, thinkertoy.txt), it calculates the checksum of the downloaded file and compares it with a pre-defined expected checksum (expectedChecksum map). If the checksums do not match, it indicates that the file has been tampered with or corrupted.

## Contributing

If you have suggestions for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## Authors

This project was build and maintained by:

 * [Doreen Onyango](https://github.com/Doreen-Onyango)
 * [Kherld Hussein](https://github.com/kherldhussein)
 * [Stephen Omotto](https://github.com/somotto)