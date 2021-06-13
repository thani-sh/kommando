# Kommando

Kommando is a tool to build command line tools with many nested sub commands. Instead of writing code to stitch sub commands together, use the directory structure itself to organize sub commands.

```shell
# Create a shell alias for your tool
alias calc=`pwd`/example/kommando.linux

calc add 1 3          # 4
calc sub 5 3          # 2
calc sub -i 5 3       # -2
calc say hello Earth  # Hello Earth!
```

## Getting Started

There are 2 steps to get started:
 - Copy the appropriate binary (ex: `build/kommando.linux`) to the root of your project.
 - Add an alias on your `.bashrc` file or `.zshrc` file to the kommando binary you copied.

It will search for commands from the directory where it's located. Each sub command can be written in any language/framework you wish as long as they are executable. Check the examples directory for an example with Nodejs and Python.
