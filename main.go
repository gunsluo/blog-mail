package main

import (
	"flag"
	"os"

	"github.com/gunsluo/blog-mail/common"
	"github.com/gunsluo/blog-mail/http"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	common.ParseConfig(*cfg)
	go http.Start()

	select {}
}
