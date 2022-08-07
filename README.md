# gogo-gigashare
GoLang based CLI file sharing software
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
- Socks5 proxy (by default will use tor or mullvad [undecided])
- Embeded TOR upload option using Bine wrapper (this would not be doable on windows so debating doing this, it may just be easier to run torify then add this... (un decided)
- possibly built in support for other upload services 
- stdin functionality (shuold be dummy easy to add just pump the "files" channel. 
- download functionality, personally i dont see the need for this but at the same time its like 10 lines of code to add so why not :/ )

### Known issues
- Encrypted files when decrypted are fucked (basically encryption is jacked up right now, dont mess with it unless you wanna help fix it, pretty sure its due to the tempfile creation and how im choosing to read the file data (os.open vs ioutil.read etc etc) lol) 
