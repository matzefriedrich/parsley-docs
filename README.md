# Parsley Documentation

This repository contains the source files for the [Parsley](https://github.com/matzefriedrich/parsley) documentation website. 

Parsley is an easy-to-use, reflection-based dependency injection package that fits seamlessly into any Go application.

## Live documentation

The documentation is hosted at: [https://matzefriedrich.github.io/parsley-docs](https://matzefriedrich.github.io/parsley-docs)

## Toolset

- [Retype](https://retype.com/) - A static site generator for building documentation websites from Markdown files.

- [gomarkdoc](https://github.com/princjef/gomarkdoc) - Tool used to generate documentation for Go packages from source code.

### Updating package documentation

To update the Go package documentation in `library/pkg/`, you can run the provided script. This requires `gomarkdoc` to be installed and the `parsley` source repository to be present in the parent directory (`../parsley`).

```bash
./update-package-docs.sh
```

### Building locally

To build and preview the documentation locally using Retype:

1. Install Retype (see [Retype installation guide](https://retype.com/guides/installation/)).
2. Run the following command in the project root:

```bash
retype start
```

---

Copyright 2024 - 2026 by Matthias Friedrich