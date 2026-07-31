# `./note`
A CLI/terminal application that acts as a fancy notetaking/timestamping tool with various ways of managing personal notes.

## Usage
Just build the application in your terminal and use `./note randomstuff i want to note down` to save a timestamp of "randomstuff i want to note down" to a default file in the given notes directory (defaults to user's home directory).

- **Shell Mode**: Use `note -sh` to run the application in "shell mode" (a looped input stream), where every newline enters a new timestamp for each line. This allows you to use special characters and formatting that your shell would be wonky with. Exit shell mode with an empty newline, or a single period, as a shoutout to `ed`, the true champion's text editor!

## Planned Features (Not Yet Implemented)
- `note -serve` starts a background process that runs a simple API webserver for interacting across networks (defaults to 8080). Append the port to you want to use with `-serve:PortNumber` to specify a different port.
- `note -api https://www.mywebsite.me` runs interactive web requests to manage and add notes from other devices and interfaces. You use the same endpoints given by `note -serve`.
- The `.note.config` configuration file. This holds things like overrides for timezones, user-defined layouts, and persisting some options like defualt ports and peers to sync api notes to.
