# WC-Tool

This is my own version of the Unix command line tool `wc`.

## How to build

Run `make build` to build the executable file for the cli

## How to run

Execute the binary with the given command and specify the file
```bash
# Outputs the number of bytes

> ./main bytes test.text
342190

# Outputs the number of lines

> ./main lc test.text
7145

# Outputs the number of words

> ./main wc test.text
58164

# Outputs the number of characters

> ./main cc test.txt
339292

# Output all values
> ./main all test.text
342190 7145 58164 339292
```