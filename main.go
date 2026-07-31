package main

import (
	"fmt"
	"os"
	"time"
)

// My ugly enum attempt for making it easier to figure out wat this abonination is trying to do...
type Action int
const (
	Default Action = iota
	Client
	Flag
	Search
	Shell
	WebServer
)
var actionName = map[Action]string{
	Default: "default",
	Client: "client",
	Flag: "flag",
	Search: "search",
	Shell: "shell",
	WebServer: "webserver",
}
func (a Action) String() string {
	return actionName[a]
}


// Main programming logic and funky-funcs //


func main() {
	input := ""
	args := os.Args[1:]
	action := Default
	// Check for any flags via 1st argument's leading char
	if len(args) >= 1 {
		switch args[0][0] {
		case '/':
			action = Search
		case '-':
			action = Flag
			flagOptTextIn := args[0][1:]
			switch flagOptTextIn {
			case "api", "client":
				action = Client
			case "srv", "serve", "server", "webserver":
				action = WebServer
			case "sh", "shell":
				action = Shell
			}
		default:
			break
		}
	}

	// Get input for the majority of functions (may call out as needed to save performace but imma lazy :P
	for i:=0;i<len(args);i++ {
		input += args[i]+" "
	}

	// Perform Actions based on validated user-given args
	switch action {
	case Client:
		fmt.Println("TODO: CLIENT")
		fallthrough
	case Default:
		OpenDateAsFile(input)
	case Search:
		fmt.Println("TODO: SEARCH")
	case Shell:
		fmt.Println("TODO: SHELL")
	case WebServer:
		fmt.Println("TODO: SERVER")
	default:
		PrintHelp()
	}
}


func GetDateAsIntString() (string, string) {
	// Returns the current timestamp as two seperate strings:
	// The first is the date in YYYYMMDD format.
	// The second is the time in the HHmmSS format.
	loc, _ := time.LoadLocation("Local") // TODO if on termux plz enter your location here... for now, until I can set it up better with like a config file
	t := time.Now().In(loc)
	// Equivalent to "YYYYMMDD and "HHmmSS", but [20]06 is the year, 01 is the month, 02 is the day, 03 is hour, 04 minute... 07 time zone
	// This way of formatting is absolute hell lol
	formatted_date := t.Format("20060102")
	formatted_time := t.Format("150405")
	return formatted_date, formatted_time
}


func OpenDateAsFile(newtxt string) {
	// TODO - Find a better name reflecting writing the note text to the date-named file
	// Takes the user's arg-provided text and writes it to a file named after the current date.
	// Prefixes the text with the current time (HHmmSS) and enters it as a new line.
	date, time := GetDateAsIntString()
	filepath, _ := os.UserHomeDir()
	filepath += "/"+date
	dateheader := date+":\n\n"
	logline := " - "+time+":  "+newtxt+"\n"
	// Check for file at ~/YYYYMMDD
	if !CheckFileExists(filepath+"/"+date) {
		// Create and append to file
		file, _ := os.Create(filepath)
		file.WriteString(dateheader+logline)
		return
	}
	// Open and append file
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(logline)
	if err != nil {
		fmt.Println(err)
	}
}


func CheckFileExists(filepath string) bool {
	// A simple abstraction to make checkig if files exist an easier 1-liner instead of 2+.
	// Returns a boolean of if it exists (true) or not (false).
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}


func PrintHelp() {
	// Prints out a simple usage and help message.
	fmt.Println(`USAGE: ./note [/search|-opt] contents
	Search:
	Just use the '/' like you would in ed/vim/etc: /searchtarget
	Having extra arguments after will also loop and search for those, then display the file and line where the terms are located.

	Options:
	-api www.website.me, -client www.website.me: Act as a client to make api callouts to the given webserver.
	Everything after the first 2 args is the content that you'll sync with the given website.

	-h, -help: Print this message. Amazing! /s

	-srv, -serve, -server, -webserver: Starts a webserver (defaults to port 8080). If you want to specify a
	different port (like HTTP port 80), simply supply -P:80 as the second argument.

	-sh, -shell: Run a looped input "shell", which is like inputting text in ed.
	Each newline is a new entry, like running ./note args without having to type in ./note
	This mode allows special characters that your shell would normally interpret and cause problems
	saving to the file. Just like ed, have a newline with a period to exit. You can also escape with Ctrl+d
	`)
}
