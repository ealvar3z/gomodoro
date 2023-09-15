# gomodoro
minimal pomo timer in Go  

## Overview

This is a simple command-line interface (CLI) tool implemented in Golang to help you manage your time using the Pomodoro Technique. The tool alternates between work and break intervals, with a longer break after every four work intervals.

## Features

- 25-minute work intervals
- 5-minute short breaks
- 15-minute long breaks after every four work intervals
- Graceful exit on interrupt (Ctrl+C)

## Prerequisites

- Golang 1.16 or higher

#Installation

To install the latest version of the program, run:

```bash
go install github.com/ealvar3z/gomodoro@latest
```

This will download the source code, compile it, and place the executable into $GOPATH/bin. Make sure this directory is in your system's PATH to run the program from any location.# Installation

## Usage

After running the program, it will automatically start the first 25-minute work interval. The CLI will update every second to show the current time.

To exit the program, press `Ctrl+C`.

## Contributing

Feel free to fork the project and submit a pull request with your changes!

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

This README provides a comprehensive guide for users and potential contributors, covering the essential aspects of the project. Feel free to modify it according to your project's specific needs.
