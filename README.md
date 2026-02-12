# parsley-docs

## Prerequisites

* Retype; see [retype.com](https://retype.com/)
* [gomarkdoc](https://github.com/princjef/gomarkdoc)

### Usage

Retype comes in different flavors; see [https://retype.com/guides/installation/](https://retype.com/guides/installation/) for further information on available installation options.

```sh
$ dotnet tool install retypeapp --global
```

Install `gomarkdoc`:

```sh
$ go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
```

## View locally

You can build a Docker image of the documentation site to view it locally:

```sh
$ docker build --rm -t parsley-docs -f Dockerfile .
```

Run the documentation using the following command:

```sh
$ docker run --rm -p 27821:80 parsley-docs
```

