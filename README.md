# mes3hacklab.ssh

Same as [mes3hacklab.org](https://mes3hacklab.org) but SSH served using
[Wish](https://github.com/charmbracelet/wish) and built with the [BubbleTea](https://github.com/charmbracelet/bubbletea)
TUI stack.

## Development

Needed programs and tools for development:

- [Go](https://go.dev/doc/install) v1.26.2
- [Just](https://github.com/casey/just)
- [Air](https://github.com/air-verse/air)

### Settings

Since `.env` files are ignored by `.gitignore`, if one want to contribute to the
project at least one `.env` file has to be created.

Here are the mandatory environment variables and their default values:

| Nome | Descrizione | Default |
|--|--|--|
| SRV_HOST | Server IP address | localhost |
| SRV_PORT | Server port | 46593 |
| SSH_KEY_PATH | SSH key absolute path used by the server | ~/.ssh/id_ed25519 |
| GITHUB_REPO_ID | Github repository ID using the format `<user>/<repo>` | mes3hacklab/mes3hacklab.github.io |

