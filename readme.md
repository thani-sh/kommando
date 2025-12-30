# Kommando

Kommando is a tool to build command line tools with many nested sub commands. Instead of writing code to stitch sub commands together, use the directory structure itself to organize sub commands.

```shell
# Create a shell alias for your tool
alias calc=/path/to/repo/example/kommando

calc add 1 3          # 4
calc sub 5 3          # 2
calc sub -i 5 3       # -2
calc say hello Earth  # Hello Earth!
```

## Getting Started

First, copy the `kommando` executable file to the root of your project. It can be compiled using `go build -o kommando ./cmd/kommando/...` or you can find pre-built binaries under [releases](https://github.com/thani-sh/kommando/releases). The next step is to add an alias in your `.bashrc` or `.zshrc` file pointing to the kommando binary you copied.

It will search for commands from the directory where it's located. Each sub command can be written in any language/framework you wish as long as they are executable. Check the `example` directory on this repo for an example project with Nodejs and Python commands.

**Note for Windows Users:** This tool is designed for Unix-like environments. If you are developing on Windows, we strongly recommend using [WSL (Windows Subsystem for Linux)](https://learn.microsoft.com/en-us/windows/wsl/install).
