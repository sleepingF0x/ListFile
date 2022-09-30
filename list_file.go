package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	FileExt   string
	WithDir   bool
	TimeStart string
	TimeStop  string
	WalkPath  string
)

func walkDIR() {
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", TimeStart, time.Local)
	if err != nil {
		fmt.Println(err)
		return
	}
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", TimeStop, time.Local)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = filepath.Walk(WalkPath, func(path string, info os.FileInfo, err error) error {
		matched, err := filepath.Match(FileExt, filepath.Base(path))
		if err != nil {
			fmt.Println(err)
			return err
		}
		if matched && (!info.IsDir() || WithDir) {
			if info.ModTime().Sub(t1).Seconds() > 0 && t2.Sub(info.ModTime()).Seconds() > 0 {
				fmt.Printf("%s\n%s\n\n", path, info.ModTime().Format("2006-01-02 15:04:05"))
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	flag.StringVar(&FileExt, "ext", "*", "only show specify file extension")
	flag.BoolVar(&WithDir, "with-dir", false, "show dir?")
	flag.StringVar(&TimeStart, "start", "", "mod time start")
	flag.StringVar(&TimeStop, "stop", "", "mod time stop")
	flag.StringVar(&WalkPath, "path", "", "search path")
	flag.Parse()

	fmt.Printf("searching dir %s from %s to %s with extension %s\n\n", WalkPath, TimeStart, TimeStop, FileExt)
	walkDIR()
}
