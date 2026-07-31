# `./note`
A CLI/terminal application that acts as a fancy notetaking/timestamping tool with various ways of managing personal notes.

## Usage
Just build the application in your terminal and use `./note randomstuff i want to note down` to save a timestamp of "randomstuff i want to note down" to a default file in the given notes directory (defaults to user's home directory).

## Planned Features (Not Yet Implemented)
- Use `note -sh` to run the application in "shell mode" (a looped input stream), where every newline enters a new timestamp for each line. This allows you to use special characters and formatting that your shell would interpret. Exit shell mode with an empty newline.
- `note -serve` starts a background process that runs a simple API webserver for interacting across networks (defaults to 8080). Append the port to you want to use with `-serve:PortNumber` to specify a different port.
- `note -api https://www.mywebsite.me` runs interactive web requests to manage and add notes from other devices and interfaces. You use the same endpoints given by `note -serve`.
