package main

import (
	"fmt"
	"os"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Too few argments")
		return
	}

	url := os.Args[1]               // https://hogehoge.com/hoge/fuga
	urls := strings.Split(url, "/") // ["https:", "", "hogehoge.com", "hoge", fuga]
	if len(urls) != 5 {
		fmt.Println("invalid url")
		return
	}

	var path string
	if gitHome := os.Getenv("GITGET_HOME"); gitHome != "" {
		path = fmt.Sprintf("%s/%s/%s/%s", gitHome, urls[2], urls[3], urls[4])
	} else {
		path = fmt.Sprintf("%s/src/%s/%s/%s", os.Getenv("HOME"), urls[2], urls[3], urls[4])
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		fmt.Println(err)
	}
}
