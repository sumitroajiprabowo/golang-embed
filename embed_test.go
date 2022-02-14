package main_test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed files/version.txt
var version string

//go:embed files/version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)

	fmt.Println(version2)
}

//go:embed files/logo.jpg
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo-new.jpg", logo, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

}

//go:embed files/first.txt
//go:embed files/second.txt
//go:embed files/third.txt
var files embed.FS

func TestMultipleFile(t *testing.T) {

	first, err := files.ReadFile("files/first.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(first))

	second, err := files.ReadFile("files/second.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(second))

	third, err := files.ReadFile("files/third.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(third))

}

//go:embed files/*.txt
var path embed.FS

func TestPath(t *testing.T) {
	dirEntries, err := path.ReadDir("files")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
