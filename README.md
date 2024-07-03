# ASCII Art Web

This project involves creating a Go program that generates ASCII art based on a given string and banner style, using a web interface and displays the results on the same interface.

## Features

- Converts strings into ASCII art
- Supports numbers, letters, spaces, special characters, and newline characters ('\n')
- Utilizes specific graphical templates for ASCII representation
- A running a server and a web GUI (graphical user interface)

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/khahussein/ascii-art-web.git
    ```

2. Navigate to the project directory:

    ```bash
    cd ascii-art-web/
    ```


## Usage

Run the server to start the [localhost](http://localhost:8080), then use the web browser to interact with the system.

```bash
go run .
```

<!-- Example: -->
<!-- TODO: Add screenshot of our web interface here to show example of the UI -->

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
 * [Tomlee Abila](https://github.com/Tomlee-abila)