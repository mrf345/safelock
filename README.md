<p align='center'>
  <img width='10%' src='docs/logo.webp' />
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

