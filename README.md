# Turn Off Monitor tool

This tool is a simple application developed in Go using Fyne for User Interface, it actually put the monitor in a sleep mode. It provides a minimalistic interface offering a solution for anyone wanting to simply turn off their monitor without using hardware buttons or changing the computer sleep time.


![](https://i.imgur.com/vOYbVmS.png)
## Features

- **Turn Off Monitor**: Instantly puts your monitor into sleep mode with the click of a button.
- **System Tray**: Minimizes to the system tray and still running in background.
- **Cross-Platform**: Works on Windows, Linux, and macOS*(supported, but I don't have a build for it yet).

## Installation
### You have two Options to use this tool:
1. Download the [latest release from the releases page](https://github.com/the-eduardo/TurnOffMonitor/releases/latest) 
2. Build the application from source:

To build from source, you'll need to have Go installed on your machine and also have the Fyne library installed. You can download Go from: 
[The official website](https://go.dev/doc/install). 

Once Go is set up, you must install [Fyne Prerequisites](https://docs.fyne.io/started/), after that can install Fyne by running the following command:
```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
```
Clone the repository to your machine:
```
git clone https://github.com/the-eduardo/TurnOffMonitor.git
```
To build the application, run:
```
fyne package -os <linux OR windows OR darwin> -icon icon.png
```

This command compiles the program and creates an executable in the current directory.

## Usage
To start the application, simply double-click the executable file.
If you want to your current application system-wide you can execute the following:
```
fyne install -icon icon.png
```

## Supported Platforms

- **Windows**: Utilizes PowerShell commands to turn off the monitor
- **Linux**: Uses `xset`
- **macOS**: Uses `pmset`

## Contributing

Contributions are welcome! If you have a bug fix, feature request, or improvement, feel free to fork the repository and submit a pull request. <3

TODO:
- ~~License
- ~~Startup with machine option
- ~~Remember the last selected key after exiting the app
