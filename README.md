# Parsley Documentation

This repository contains the source files for the [Parsley](https://github.com/matzefriedrich/parsley) documentation website. 

Parsley is an easy-to-use, reflection-based dependency injection package that fits seamlessly into any Go application.

## Live Documentation

The documentation is hosted at: [https://matzefriedrich.github.io/parsley-docs](https://matzefriedrich.github.io/parsley-docs)

## Prerequisites

* Retype; see [retype.com](https://retype.com/)
* [gomarkdoc](https://github.com/princjef/gomarkdoc)

## Usage

Retype comes in different flavors; see [https://retype.com/guides/installation/](https://retype.com/guides/installation/) for further information on available installation options.

```bash
dotnet tool install retypeapp --global
```

Install `gomarkdoc`:

```bash
go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
```

### Updating Package Documentation

To update the Go package documentation in `library/pkg/`, run the provided script. This requires `gomarkdoc` to be installed and the `parsley` source repository to be present in the parent directory (`../parsley`).

```bash
./update-package-docs.sh
```

### Building Locally

Run the following command to build and preview the documentation locally using Retype:

```bash
retype start
```

You can also build a Docker image of the documentation site to view it locally:

```bash
docker build --rm -t parsley-docs -f Dockerfile .
```

Run the documentation using the following command:

```bash
docker run --rm -p 27821:80 parsley-docs
```

---

Copyright 2024 - 2026 by Matthias Friedrich
