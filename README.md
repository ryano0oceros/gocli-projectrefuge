# csv2json

Go CLI app for porting CSV file to JSON

`csv2json` is a command-line utility written in Go that converts a CSV file to a JSON lines (.jl) file.

## Installation
First, clone the repository to your local machine:

`git clone https://github.com/ryano0oceros/gocli-projectrefuge.git`

Then, navigate to the csv2json directory:


`cd gocli-projectrefuge`

Build the Go program to create an executable:

`go build csv2json.go`

This will create an executable file named csv2json in the same directory.

If you want to be able to run your program from any directory, move the `csv2json` executable to a directory that's in your system's PATH. (i.e. for Mac users, you could move it to `/usr/local/bin`):

`mv csv2json /usr/local/bin`

## Usage
After installation, you can use `csv2json` to convert a CSV file to a JSON lines (.jl) file. The command takes two arguments: the input CSV file and the output JSON file.

Here's the general syntax (replace with `./csv2json` if you didn't move the executable to a directory in your PATH):

`csv2json housesInput.csv housesOutput.jl`

Replace housesInput.csv with the path to your input CSV file and housesOutput.jl with the path to your output JSON lines file.

To convert a CSV file named housesInput.csv to a JSON lines file named json.jl, you would run the above snippet.
