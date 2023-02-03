# instacheck
Automated verification of available Instagram usernames
Simple program to check if any usernames in a list of handles are taken.
Using this program is as simple as making a text file containing usernames that you want to check, then running the tool and proving the filepath to the usernames.

Ex: `./instacheck -f usernames.txt`

## Compiling
#
To compile, you just need the golang language installed. You can build the binary for any operating system using `go build` inside of the project directory.
### Dependencies
#
This project relies on having the cURL commandline tool installed.

To install cURL on linux:
```
sudo apt install curl
```

To install cURL on mac:

```
brew install curl
```
