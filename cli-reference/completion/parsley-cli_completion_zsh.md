---
label: completion zsh
icon: '<svg viewBox="0 0 24 24"><path d="M15.4,14.33c.3-1.4-.45-2.81-2.39-3.13l.61-2.88c.87.21,1.17.7,1.19.73,0,0,0-.01,0-.01l1.62-1.04c-.04-.07-.25-.44-.76-.81-.46-.33-1.03-.54-1.67-.65l.08-.36c.1-.48-.21-.87-.69-.87s-.96.39-1.06.87l-.08.36c-2.08.32-3.42,1.73-3.72,3.13s.45,2.81,2.39,3.13l-.61,2.85c-.81-.25-1.25-.78-1.26-.79h0s-1.5,1.24-1.5,1.24c.04.05.84,1.05,2.37,1.34l-.08.37c-.1.48.21.87.69.87s.96-.39,1.06-.87l.08-.36c2.08-.32,3.42-1.73,3.72-3.13ZM10.28,9.67c.12-.55.65-1.12,1.6-1.35l-.58,2.71c-.84-.23-1.14-.81-1.02-1.35ZM12.06,15.68l.58-2.71c.32.09.58.23.76.42.24.25.33.59.26.94-.12.55-.65,1.12-1.6,1.35Z" /></svg>'
---
## parsley-cli completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(parsley-cli completion zsh)

To load completions for every new session, execute once:

#### Linux:

	parsley-cli completion zsh > "${fpath[1]}/_parsley-cli"

#### macOS:

	parsley-cli completion zsh > $(brew --prefix)/share/zsh/site-functions/_parsley-cli

You will need to start a new shell for this setup to take effect.


```
parsley-cli completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [parsley-cli completion](./index.md)	 - Generate the autocompletion script for the specified shell
