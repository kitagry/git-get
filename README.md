## git-get

This is wrapper for `git clone` command.

## Background

I manage some repositories in $HOME/src/github.com/USER_NAME/REPO_NAME.
So, when I download an repository, I always run below command.

```
$ cd $HOME/src/github.com
$ mkdir USER_NAME
$ cd USER_NAME
$ git clone $HOME/src/github.com/USER_NAME/REPO_NAME
```

This is shortcut for above commands.

## Installation

```
$ go get github.com/kitagry/git-get
```

## Usage

```
$ git get https://github.com/hoge/fuga
```

So, git clone the above repository in `$GITGET_HOME/github.com/hoge/fuga`
