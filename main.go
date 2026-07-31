package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	input := ""
	args := os.Args[1:]
	for i:=0;i<len(args);i++ {
		input += args[i]+" "
	}
	OpenDateAsFile(input)
}

func GetDateAsIntString() (string, string) {
	loc, _ := time.LoadLocation("Local") // TODO if on termux plz enter your location here... for now, until I can set it up better with like a config file
	t := time.Now().In(loc)
	// Equivalent to "YYYYMMDD and "HHmmSS", but [20]06 is the year, 01 is the month, 02 is the day, 03 is hour, 04 minute... 07 time zone
	// This way of formatting is absolute hell lol
	formatted_date := t.Format("20060102")
	formatted_time := t.Format("150405")
	return formatted_date, formatted_time
}

func OpenDateAsFile(newtxt string) {
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
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
