package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// base()
	// rel()
	walk()
}

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}

func walk() {
	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	subDirToSkip := "skip"

	fmt.Println("On Unix:")
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
}

func rel() {
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"

	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}
}

func base() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Base("/foo/bar/baz.js"))
	fmt.Println(filepath.Base("/foo/bar/baz"))
	fmt.Println(filepath.Base("/foo/bar/baz/"))
	fmt.Println(filepath.Base("dev.txt"))
	fmt.Println(filepath.Base("../todo.txt"))
	fmt.Println(filepath.Base(".."))
	fmt.Println(filepath.Base("."))
	fmt.Println(filepath.Base("/"))
	fmt.Println(filepath.Base(""))
}
