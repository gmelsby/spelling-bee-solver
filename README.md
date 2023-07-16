# Spelling Bee Puzzle Solver

### Description
CLI tool to generate solutions for a New York Times Spelling Bee-style puzzle. \
Prompts user for letters and prints a list of possible puzzle solutions. \
Uses [words_alpha.txt](https://github.com/dwyl/english-words/blob/master/words_alpha.txt) from [dwly's english-words repo](https://github.com/dwyl/english-words/tree/master) as a source of valid words. \
This list contains many words that the official NYT Spelling Bee game does not recognize as valid words. \
However, other word lists omit certain valid answers, and it is preferable to generate too many words than too few.

### How to run
- Make sure you have Go installed
- Clone this repo
- In the project root, run the command `go run . -i`
- Follow the instructions printed to stdout

### Output to file
If you wish to output to file and are on a Unix (e.g. Linux, MacOS) system, you can run the command \
`echo [letters] | go run . > output.txt` \
with `[letters]` replaced with a string of the letters in the puzzle with the required letter first.