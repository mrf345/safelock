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

<details>
<summary><h5>Linux</h5></summary>

  - With binaries

  ```bash
  wget -qO- https://github.com/mrf345/safelock/releases/latest/download/safelock-linux-amd64.tar.gz | tar xvz -C ~ && ~/safelock
  ```

  - Or from the source code

    Make sure you have [go](https://go.dev/doc/install), [npm](https://nodejs.org/en/download/package-manager) and [git](https://git-scm.com/downloads) are installed, then run:

    ```bash
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    git clone https://github.com/mrf345/safelock.git
    cd safelock
    wails dev
    ```
</details>
<details>
<summary><h5>Windows</h5></summary>

  Download, extract and install [this](https://github.com/mrf345/safelock/releases/latest/download/safelock-windows-amd64.zip) or [this](https://github.com/mrf345/safelock/releases/latest/download/safelock-windows-arm64.zip) for `arm64` processors. If you want a portable version download [this](https://github.com/mrf345/safelock/releases/latest/download/safelock-windows-portable-amd64.zip) or [this](https://github.com/mrf345/safelock/releases/latest/download/safelock-windows-portable-arm64.zip) for `arm64`.

</details>
<details>
<summary><h5>MacOS</h5></summary>

  Download and run [this universal .app](https://github.com/mrf345/safelock/releases/latest/download/safelock-darwin-universal.zip), Note that you'll need to enable running apps from unknown developers follow [this guide](https://www.wikihow.com/Install-Software-from-Unsigned-Developers-on-a-Mac).

</details>


### Performance

> [!NOTE]
> Check [safelock-cli/performance](https://github.com/mrf345/safelock-cli?tab=readme-ov-file#performance) for more detailed benchmarks.

<p align="center">
  <a href="https://raw.githubusercontent.com/mrf345/safelock-cli/master/benchmark/encryption-time.webp" target="_blank">
    <img src="https://raw.githubusercontent.com/mrf345/safelock-cli/master/benchmark/encryption-time.webp" alt="encryption time" />
  </a>
</p>

<p align="center">
  <a href="https://raw.githubusercontent.com/mrf345/safelock-cli/master/benchmark/decryption-time.webp" target="_blank">
    <img src="https://raw.githubusercontent.com/mrf345/safelock-cli/master/benchmark/decryption-time.webp" alt="decryption time" />
  </a>
</p>


### Breaking changes

##### v1.0.0

Should expect great improvement in performance (about **23.2** times faster) compared to the last release `0.5`, with better overall encryption and cross-platform support.

However, this version **breaks backward compatibility**. Any files encrypted with a prior versions can't be decrypted with this version, and vice versa.

<p style="margin-top: 35px;">
  <a href="https://raw.githubusercontent.com/mrf345/safelock/master/docs/demo.gif" target="_blank">
    <img src="docs/demo.gif" alt="demo" />
  </a>
</p>
