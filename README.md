# VH7 CLI

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/jake-walker/vh7-cli/CI/master) [![DeepSource](https://deepsource.io/gh/jake-walker/vh7-cli.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/jake-walker/vh7-cli/?ref=repository-badge) ![License](https://img.shields.io/github/license/jake-walker/vh7-cli)

A work-in-progress command line interface app for VH7.

<!-- TOC -->

- [Overview](#overview)
- [Installation](#installation)
    - [Windows (using Chocolatey, recommended)](#windows-using-chocolatey-recommended)
    - [Linux (using install script, recommended)](#linux-using-install-script-recommended)
    - [Pre-built binaries](#pre-built-binaries)
    - [Using Go](#using-go)
- [Contributing](#contributing)
- [License](#license)

<!-- /TOC -->

## Overview

A lovely command-line interface made for VH7. VH7 command-line allows you to quickly create short links, share your code and share files without leaving your terminal.

- **Free.** VH7 is not only free to use on the [official instance](https://vh7.uk) but is also free to download and run for yourself.
- **Open Source.** All of VH7's source code is available here for the community to take a look under the hood. _We also accept community contributions, just open a pull request!_
- **Multi-purpose.** Unlike other mainstream URL shorteners, VH7 also provides file sharing and a pastebin also with short links.

## Installation

There are many different ways that VH7 is distributed, Chocolatey is the preferred ways of installation as it provides easy updating, whereas if you use one of the binaries you will need to update manually.

### Windows (using Chocolatey, recommended)

The best way for insalling VH7 CLI on Windows is by using [Chocolatey](https://chocolatey.org/).

_Please note Chocolatey packages are only published on minor versions, not patches._

Firstly, install Chocolatey, if you haven't already, by reading the guide [here](https://chocolatey.org/install). Now, you can install the CLI by running:

```powershell
choco install vh7-cli
```

And upgrade using:

```powershell
choco upgrade vh7-cli
```

### Linux (using install script, recommended)

_Install script coming soon..._

### Pre-built binaries

If you'd rather not use Chocolatey, pre-built binaries are available on [GitHub Releases](https://github.com/jake-walker/vh7-cli/releases). These are portable and can be run through your command line without installation or administrator privileges. They are built for most popular operating systems and architectures.

### Using Go

If you have Go installed, you can clone this repository and run the following to install VH7 CLI on to your system.

```bash
go install
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please run tests and style checking before committing.

## License

This project is licensed under the MIT License - see the [license file](LICENSE) for details.
