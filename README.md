<h1 align="center">whaleshell-display</h1>

<p align="center">
  <strong>GUI / noVNC helpers</strong><br>
  Passworded noVNC surfaces for Debian GUI sandboxes.
</p>
<p align="center">
  <a href="https://github.com/whaleshell/whaleshell-display/actions/workflows/ci.yml"><img src="https://github.com/whaleshell/whaleshell-display/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://pkg.go.dev/github.com/whaleshell/whaleshell-display"><img src="https://pkg.go.dev/badge/github.com/whaleshell/whaleshell-display.svg" alt="Go Reference"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  <a href="https://github.com/whaleshell/whaleshell-display"><img src="https://img.shields.io/badge/Go-1.27+-00ADD8?logo=go" alt="Go Version"></a>
</p>
<p align="center">
  <sub>Part of the <a href="https://github.com/whaleshell">whaleshell / whaleshell</a> ecosystem</sub>
</p>

---

## Overview

**whaleshell-display** helpers wire noVNC publish mode, password generation, and browser open helpers used by `whaleshell sandbox create --display novnc`.

### Key Features

| Category | Capabilities |
|----------|--------------|
| **Mode** | none / novnc parse + defaults |
| **Auth** | Password file helpers for x11vnc / noVNC |
| **UX** | Open published `vnc.html` URL on the host |

---

## Installation

```bash
go get github.com/whaleshell/whaleshell-display@latest
```

**Requirements:** Go 1.27+. GUI image: `ghcr.io/whaleshell/whaleshell/sandboxes/gui`.

---

## Quick Start

```bash
whaleshell sandbox create --name gui --from gui --display novnc \
  --workspace . --policy ./policies/default.yaml
# open printed http://127.0.0.1:6080/vnc.html?password=…
```

---

## Package Structure

| Path | Purpose |
|------|---------|
| `novnc/` | Mode parsing, password, open helpers |


---

## Related

| Resource | Link |
|----------|------|
| Roadmap | [ROADMAP.md](./ROADMAP.md) |
| Organization | [https://github.com/whaleshell](https://github.com/whaleshell) |
| Organization overview | [github.com/whaleshell](https://github.com/whaleshell) |
| pkg.go.dev | [`github.com/whaleshell/whaleshell-display`](https://pkg.go.dev/github.com/whaleshell/whaleshell-display) |

## License

[MIT](./LICENSE) © whaleshell
