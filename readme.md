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

First, copy the `kommando` executable file to the root of your project. It can be compiled using `make` or you can find pre-built under [releases](https://github.com/thani-sh/kommando/releases). The next step is to add an alias on your `.bashrc` file or `.zshrc` file to the kommando binary you copied.

It will search for commands from the directory where it's located. Each sub command can be written in any language/framework you wish as long as they are executable. Check the `example` directory on this repo for an example project with Nodejs and Python commands.

## Auto Completion

Kommando supports auto completion for Bash and Zsh. To enable it, you need to add the initialization script to your shell configuration file.

### Bash

Add the following to your `.bashrc` file (replace `calc` with your alias):

```bash
source <( /path/to/repo/example/kommando --kommando-init-bash calc )
```

### Zsh

Add the following to your `.zshrc` file (replace `calc` with your alias):

```zsh
source <( /path/to/repo/example/kommando --kommando-init-zsh calc )
```
