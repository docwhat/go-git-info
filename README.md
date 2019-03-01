[![Build Status](https://travis-ci.org/docwhat/go-git-info.svg?branch=master)](https://travis-ci.org/docwhat/go-git-info)

# git-info

A fast, pure go tool to get information from git; useful (for example)
for creating prompts.

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
