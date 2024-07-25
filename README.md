### Using this tool, the selected text is translated into the desired language using the xsel tool and displayed using kdialog.
## Essentials package
```bash
apt install xsel kdialog
```
## Run:
```bash
./translator_linux_amd64 -lang en -dir ltr
```
OR
```bash
./translator_linux_amd64 -lang fa -dir rtl
```
>  Use it as a shortcut for easier use

## Options
### -dir
```
The dir attribute can have the following values:

    ltr - means left-to-right text direction
    rtl - means right to left text direction
```

### -lang
```
The lang attribute specifies the language of the element's content.

Common examples are "en" for English, "es" for Spanish, "fr" for French and so on.
```
## Screenshot
![alt text](https://github.com/nopeEOF/translator/blob/main/image/example1.png?raw=true)
## build
> Please make sure Go version >= 1.21.4
```bash
env GOOS=linux GOARCH=amd64 go build -o translator_linux_amd64 main.go
```