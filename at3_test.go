package main

import (
	"os"
	"testing"
	"bufio"
	"io/ioutil"
	"strings"
)

func TestAt3_transform(t *testing.T) {
	outpath := "testdata/IO/out.txt"

	in, err := os.Open("testdata/access_log")
	if err != nil {
		t.Fatal("in file error.")
	}
	defer in.Close()

	out, err := os.OpenFile(outpath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		t.Fatal("out file error")
	}
	defer out.Close()

	at3 := NewAt3()

	// output formatted csv to file
	writer := bufio.NewWriter(out)
	err = at3.Transform(in, writer)
	if err != nil {
		t.Fatal("transform error")
	}

	// assert output file
	contents, _ := ioutil.ReadFile(outpath)
	text := string(contents)

	// assert host
	if count := strings.Count(text, "127.0.0.1"); count != 3 {
		t.Fatalf("parsed error. actual %d", count)
	}
}
