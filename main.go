package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func usage() {
	log.Fatalln("usage: prefix-modtime /path/to/files")
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		usage()
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		info, err := e.Info()
		if err != nil {
			log.Fatalln(err)
		}
		time := info.ModTime()

		o := filepath.Join(dir, e.Name())
		n := filepath.Join(dir, time.Format("20060102")+" - "+e.Name())
		err = os.Rename(o, n)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
