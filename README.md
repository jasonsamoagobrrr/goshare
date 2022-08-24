[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/gojp/goreportcard) [![License](https://img.shields.io/github/license/imayberoot/goshare)](https://github.com/imayberoot/goshare/blob/main/LICENSE)
# gogo-gigashare
GoLang based CLI file sharing script
#### Install:
```
go install github.com/imayberoot/ggshare@latest
```

## Features
- Concurrent file upload to services that use RFC 2388 multipart/form-data. By default the script uses https://pomf.cat Alternatives could be https://x0.at and https://0x0.st (just know they are known to block upload from all tor exit nodes) Also note the field name and response data will probably be diffrent so slight tweaks there.
- Concurrent AES file encrytion with automatic 16 bit hash generation from passprase
- File decryption using passprase
- Directory (automatic detection) and Recursive Directory upload 

## How to use
```
ggshare {flags} {file or directory}
```
Examples:
```
ggshare UPLOADME.file

ggshare -e -k=bilbobaggins UPLOADME.file

ggshare -r ~/documents/notporn
```

### Coming soon:
- Socks5 proxy (by default will use tor or mullvad [undecided, if mullvad will make sure to only select from wireguard])
- Embeded TOR upload option using Bine wrapper (this would not be doable on windows so debating doing this, it may just be easier to run torify then add this... (un decided)
- possibly built in support for other upload services (this would mean changing the form title and response data based of selected provider, probably best to add provider classes if we do this)
- stdin functionality (shuold be dummy easy to add just pump the "files" channel, although I am very tempeted to add an init func based off configurations to then we can better injest/map the files before just running) 
- download functionality, personally i dont see the need for this but at the same time its like 10 lines of code to add so why not :/ )
- I may add in a configuration class to then have a default config as well as a few pre-defined configs and better error logging as we need to have services running before the crawling etc begin. (im lazy and probably wont ever do this)

### Known issues
- Encrypted files when decrypted are fucked (basically encryption is jacked up right now, dont mess with it unless you wanna help fix it, pretty sure its due to the tempfile creation and how im choosing to read the file data (os.open vs ioutil.read etc etc) lol) 

### why did i make this?
Mainly to get better at working with files and coding in golang (this is my like 3rd time doing el go). Also have a binary aval that can run on most systems thats ready to dump files fast onto a secure 3rd part. 
