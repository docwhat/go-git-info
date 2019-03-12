# Contributing

I love pull requests from everyone!

## Getting Started

### Install Go

First you'll need to make sure you have go version 1.12 or later.
golang.org has some
[good instructions on installing Go](https://golang.org/doc/install).

### Getting the source

If you will be contributing, then you'll want to
[fork the repository](https://help.github.com/articles/fork-a-repo/).

Once you've forked it, then you can clone the source:

    git clone git@github.com:<your-username>/<repository-name>.git

Fetch the required dependencies:

```sh
$ script/bootstrap
```

After bootstrapping, you can start to build the project:
```
GO111MODULE=on
CGO_ENABLED=0
go build ./...
```
After this, you can run the main file:
```
go run cmd/git-info/git-info.go
```

Before you do any changes, make sure the tests pass:

```sh
$ bin/golangci-lint run ./...
$ go test ./..
```

Make your change. Add tests for your change. Make the tests pass:

```sh
$ bin/golangci-lint run ./...
$ go test ./..
```

Push to your fork and
[submit a pull request](https://help.github.com/articles/creating-a-pull-request/).

At this point you're waiting on me. I try to be responsive to pull
requests, but you know life can get in the way. I may suggest some
changes or improvements or alternatives.

Some things that'll increase the chance that I accept your pull request:

-   Write tests.
-   Write a
    [good commit message](https://seesparkbox.com/foundry/semantic_commit_messages).
