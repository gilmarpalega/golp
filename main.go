package main

import (
	"os"
	"fmt"
	"time"
	"strings"
)

// This will rename file to dist/app_20170801-084212.min.js
const file = "dist/app.min.js"

// This file will contain app_20170801-084212.min.js, so the app can read it and set in the proper main templates
const filedesc = "dist/filename.txt"


func datetimeHash() (ret string) {

	t := time.Now()

	ret = fmt.Sprintf("%04d/%02d/%02d-%02d%02d%02d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())

	return
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	parts := strings.SplitN(file, ".", 2)
	hash := datetimeHash()
	finalName := fmt.Sprintf("%v_%v%v", parts[0], hash, parts[1])

	f, err := os.Create(filedesc)
	check(err)
	defer f.Close()
	_, err = f.WriteString("writes\n")
	check(err)
	f.Sync()

	err =  os.Rename(file, finalName)

	check(err)

}