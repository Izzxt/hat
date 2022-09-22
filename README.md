# Hat
[![Habbo Resource](https://github.com/Izzxt/hat/actions/workflows/main.yml/badge.svg?event=workflow_dispatch)](https://github.com/Izzxt/hat/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Izzxt/hat)](https://goreportcard.com/report/github.com/Izzxt/hat)

Download various files directly from Habbo.

```sh
➜ hat
Habbo Downloader Tools

Usage:
  hat [command]

Available Commands:
  articles    Download habbo web articles
  badgeparts  Download habbo group badgeparts
  badges      Download habbo badges
  clothes     Download habbo clothes
  completion  Generate the autocompletion script for the specified shell
  effects     Download habbo effects
  furni       Download habbo furni or icons
  gamedata    Download habbo gamedata
  gordon      Download habbo gordon assets
  habboswf    Download Habbo.swf
  help        Help about any command
  hotelview   Download habbo hotel view
  icons       Download habbo catalogue icons
  mp3         Download habbo mp3 songs
  pets        Download habbo pets
  promo       Download habbo web promo

Flags:
  -c, --config string       Config file
  -d, --domain string       com.br, com.tr, com, de, es, fi, fr, it, nl (default "com")
  -h, --help                help for hat
  -o, --output string       Folder output
  -p, --production string   Habbo gordon production

Use "hat [command] --help" for more information about a command.
```

## Install

## MacOS & Linux Homebrew
```sh
brew install Izzat/tap/hat
```

## Linux

#### Manual
```sh
# Choose desired version, Architecture & target OS
export HAT_VERSION="1.0"
export ARCH="x86_64"
export OS="linux"
wget -q https://github.com/Izzxt/hat/releases/download/v${VERSION}/hat_${VERSION}_${OS}_${ARCH}.tar.gz && \
tar -xf hat_${VERSION}_${OS}_${ARCH}.tar.gz && \
chmod +x hat && \
sudo mv hat /usr/local/bin/hat
```

## Windows
```sh
scoop bucket add hat https://github.com/Izzxt/scoop-bucket.git
scoop install hat
```

Alternatively you can download `deb`, `rpm` or `windows` from [release pages](https://github.com/Izzxt/hat/releases)  
you need to add the binaries to System [PATH](https://docs.oracle.com/en/database/oracle/machine-learning/oml4r/1.5.1/oread/creating-and-modifying-environment-variables-on-windows.html#GUID-DD6F9982-60D5-48F6-8270-A27EC53807D0) if you want to use it as command line.

## Usage Example
*Reminder:* if you download single file, be aware of file does not output as what u expected. It will download but if you open the file it will show you 404 HTML code.
```sh
# Download all files
hat gamedata

# If you provide optional flags, it will download single file
hat badges --code <badge code>
```

## Contributing
Contributions, issues and feature requests are welcome!

## Similar projects
> Disclaimer: This project only for educational purpose, I create this project just to gain more knowledge in programming.

  - [higoka/habbo-downloader](https://github.com/higoka/habbo-downloader)

## Support
Give a ⭐ if you like this project!
