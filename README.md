# Parsley Documentation

This repository contains the source files for the [Parsley](https://github.com/matzefriedrich/parsley) documentation website. 

Parsley is an easy-to-use, reflection-based dependency injection package that fits seamlessly into any Go application.

## Live Documentation

The documentation is hosted at: [https://matzefriedrich.github.io/parsley-docs](https://matzefriedrich.github.io/parsley-docs)

## Prerequisites

* Retype; see [retype.com](https://retype.com/)
* [gomarkdoc](https://github.com/princjef/gomarkdoc)

## Usage

A `Makefile` is provided to simplify common tasks.

### Initialization

Install the required tools:

```bash
make init
```

Note: This will install `gomarkdoc` to a local `bin/` folder. `retype` should be installed globally as a dotnet tool.

### Retype Commands

Retype comes in different flavors; see [https://retype.com/guides/installation/](https://retype.com/guides/installation/) for further information on available installation options.

```bash
dotnet tool install retypeapp --global
```

Run the following command to build and preview the documentation locally using Retype:

```bash
make start
```

To build the documentation:

```bash
make build
```

### Updating Package Documentation

To update the Go package documentation in `library/pkg/`, run:

```bash
make update-docs
```

This requires the `parsley` source repository to be present in the parent directory (`../parsley`).

### Docker

You can also build a Docker image of the documentation site to view it locally:

```bash
make docker-build
```

Run the documentation using the following command:

```bash
make docker-run
```

---

Copyright 2024 - 2026 by Matthias Friedrich
