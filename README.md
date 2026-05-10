# mes3hacklab.ssh

Stessa cosa di mes3hacklab.org, ma servito in SSH con [Wish](https://github.com/charmbracelet/wish) e
costruito con lo stack TUI [BubbleTea](https://github.com/charmbracelet/bubbletea).

## Come accedere

Essendo una TUI app servita con SSH, basta eseguire il seguente comando per accederci.

```bash
ssh mes3hacklab.org
```

## Settings

Dal momento che i file .env vengono ignorati dal .gitignore, durante lo sviluppo
ed il deploy bisogna crearne uno per far funzionare il server.

Di seguito Vengono elencate le variabili d'ambiente necessarie ed i rispettivi
valori di default:

| Nome | Descrizione | Default |
|--|--|--|
| SRV_HOST | Indirizzo IP del server | localhost |
| SRV_PORT | Porta del server | 46593 |
| SSH_KEY_PATH | Path assoluto della chiave SSH privata utilizzata dal server | ~/.ssh/id_ed25519 |
| GITHUB_REPO_ID | ID della repository GitHub in formato `<user>/<repo>` | mes3hacklab/mes3hacklab.github.io |

## Sviluppo

- [Go](https://go.dev/doc/install) v1.26.2
- [Just](https://github.com/casey/just)
- [Air](https://github.com/air-verse/air)

