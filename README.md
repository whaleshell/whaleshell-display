<h1 align="center">osg-display</h1>

<p align="center">
  <strong>GUI / noVNC helpers</strong><br>
  Passworded noVNC surfaces for Debian GUI sandboxes.
</p>
<p align="center">
  <a href="https://github.com/zorneth/osg-display/actions/workflows/ci.yml"><img src="https://github.com/zorneth/osg-display/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://pkg.go.dev/github.com/zorneth/osg-display"><img src="https://pkg.go.dev/badge/github.com/zorneth/osg-display.svg" alt="Go Reference"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  <a href="https://github.com/zorneth/osg-display"><img src="https://img.shields.io/badge/Go-1.27+-00ADD8?logo=go" alt="Go Version"></a>
</p>
<p align="center">
  <sub>Part of the <a href="https://github.com/zorneth">zorneth / osg</a> ecosystem</sub>
</p>

---

## Overview

**osg-display** helpers wire noVNC publish mode, password generation, and browser open helpers used by `osg sandbox create --display novnc`.

### Key Features

| Category | Capabilities |
|----------|--------------|
| **Mode** | none / novnc parse + defaults |
| **Auth** | Password file helpers for x11vnc / noVNC |
| **UX** | Open published `vnc.html` URL on the host |

---

## Installation

```bash
go get github.com/zorneth/osg-display@latest
```

**Requirements:** Go 1.27+. GUI image: `ghcr.io/zorneth/osg/sandboxes/gui`.

---

## Quick Start

```bash
osg sandbox create --name gui --from gui --display novnc \
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
| Organization | [https://github.com/zorneth](https://github.com/zorneth) |
| Organization overview | [github.com/zorneth](https://github.com/zorneth) |
| pkg.go.dev | [`github.com/zorneth/osg-display`](https://pkg.go.dev/github.com/zorneth/osg-display) |

## License

[MIT](./LICENSE) © zorneth
