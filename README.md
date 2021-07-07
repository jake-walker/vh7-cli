# Quark CLI

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/quark-links/quark-cli/CI/master) [![DeepSource](https://deepsource.io/gh/quark-links/quark-cli.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/quark-links/quark-cli/?ref=repository-badge) ![License](https://img.shields.io/github/license/quark-links/quark-cli)

A work-in-progress command line interface app for Quark.

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

A lovely command-line interface made for Quark. Quark command-line allows you to quickly create short links, share your code and share files without leaving your terminal.

- **Free.** Quark is not only free to use on the [official instance](https://vh7.uk) but is also free to download and run for yourself.
- **Open Source.** All of Quark's source code is available here for the community to take a look under the hood. _We also accept community contributions, just open a pull request!_
- **Multi-purpose.** Unlike other mainstream URL shorteners, Quark also provides file sharing and a pastebin also with short links.

## Installation

There are many different ways that Quark is distributed, Chocolatey is the preferred ways of installation as it provides easy updating, whereas if you use one of the binaries you will need to update manually.

### Windows (using Chocolatey, recommended)

The best way for insalling Quark CLI on Windows is by using [Chocolatey](https://chocolatey.org/).

_Please note Chocolatey packages are only published on minor versions, not patches._

Firstly, install Chocolatey, if you haven't already, by reading the guide [here](https://chocolatey.org/install). Now, you can install the CLI by running:

```powershell
choco install quark-cli
```

And upgrade using:

```powershell
choco upgrade quark-cli
```

### Linux (using install script, recommended)

_Install script coming soon..._

### Pre-built binaries

If you'd rather not use Chocolatey, pre-built binaries are available on [GitHub Releases](https://github.com/quark-links/quark-cli/releases). These are portable and can be run through your command line without installation or administrator privileges. They are built for most popular operating systems and architectures.

### Using Go

If you have Go installed, you can clone this repository and run the following to install Quark CLI on to your system.

```bash
go install
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please run tests and style checking before committing.

## License

This project is licensed under the MIT License - see the [license file](LICENSE) for details.
