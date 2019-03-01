<h1 align="center">Git Info</h1>
<p align="center">
    <a href="https://travis-ci.org/docwhat/go-git-info">
        <img src="https://travis-ci.org/docwhat/go-git-info.svg?branch=master" alt="Build Status" />
    </a>
    <a href="https://goreportcard.com/report/github.com/docwhat/go-git-info">
        <img src="https://goreportcard.com/badge/github.com/docwhat/go-git-info" alt="Go Report Card" />
    </a>
    <a href="http://godoc.org/github.com/docwhat/go-git-info">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square" alt="Go Doc">
    </a>
    <a href="https://github.com/golang-standards/docwhat/go-git-info">
        <img src="https://img.shields.io/github/release/docwhat/go-git-info.svg" alt="Release">
    </a>
    <img alt="License" src="https://img.shields.io/github/license/docwhat/go-git-info.svg">
    <br />
    <a href="https://github.com/docwhat/go-git-info/graphs/contributors" alt="Contributors">
        <img src="https://img.shields.io/github/contributors/docwhat/go-git-info.svg" />
    </a>
    <a href="https://github.com/docwhat/go-git-info/issues" alt="GitHub issues">
        <img alt="GitHub issues" src="https://img.shields.io/github/issues/docwhat/go-git-info.svg">
    </a>
    <a href="https://github.com/docwhat/go-git-info/pulls" alt="GitHub pull requests">
        <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/docwhat/go-git-info.svg">
    </a>
</p>

> A fast, pure go tool to get information from git; useful (for example)
> for creating prompts.

## Why?

You need your shell prompt to be quick. Gathering information about the
current git repository can be slow using the normal `git` command.

`git-info` replaces the `git` command for gathering prompt information.

`git-info` is also cross platform so that you can use it on all
platforms you use your shell prompt on.

## Similar projects

-   [romkatv](https://github.com/romkatv)'s
    [`gitstatus`](https://github.com/romkatv/gitstatus) utility inspired
    this project.
-   ZSH's
    [`vcs_info`](http://zsh.sourceforge.net/Doc/Release/User-Contributions.html#Version-Control-Information)
