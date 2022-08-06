# gogo-gigashare
GoLang based CLI file sharing software
#### Install (temporarily broken):
```
go install github.com/imayberoot/ggshare@latest
```

## Features
- Concurrent file upload to services that use RFC 2388 multipart/form-data. By default the script uses https://pomf.cat Alternatives could be https://x0.at and https://0x0.st (just know they are known to block upload from all tor exit nodes)
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
- Embeded TOR upload option using Bine wrapper
- Compression of files before encryption or upload 

### Known issues
- Encrypted files when decrypted are fucked (basically encryption is jacked up right now, dont mess with it lol) 
