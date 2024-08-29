<p align='center'>
  <img width='24%' src='docs/logo.webp' />
</p>

<p align='center'>
  <a href='https://github.com/mrf345/safelock/actions/workflows/ci.yml'>
    <img src='https://github.com/mrf345/safelock/workflows/Build/badge.svg'>
  </a>
  <img alt="Static Badge" src="https://img.shields.io/badge/OS-_Linux_%7C_Windows_%7C_MacOS-blue">
  <img alt="Static Badge" src="https://img.shields.io/badge/Arch-_amd64_%7C_arm64-black">
</p>

<p align='center'>
  Fast drag & drop cross-platform files encryption tool, based on <a href="https://github.com/mrf345/safelock-cli" target="_blank">safelock-cli</a> and built with
  <a href="https://github.com/wailsapp/wails" target="_blank">Wails</a> and <a href="https://github.com/angular/angular" target="_blank">Angular</a>.
</p>

<hr />

### Install

With the [Go](https://go.dev/) package manager

```shell
go install https://github.com/mrf345/safelock
```

Or using one of the latest compiled binaries [here](https://github.com/mrf345/safelock/releases)


### Changelog

##### v1.0.0

Should expect great improvement in performance (~6x) when compared to the last release 0.0.5, better overall encryption and cross-platform support.

Unfortunately this version breaks backward compatibility. Any files encrypted with a prior version can't be decrypted with this version, and vice versa.

### Development

- *Run tests*: `make test`
- *Run style check*: `make lint`
- *Compile binary*: `make pkg`


### Demo

![Demo](docs/demo.gif)

