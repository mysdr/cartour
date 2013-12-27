package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
)

const (
	autoHome = "autohome"
	bitAuto  = "bitauto"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	app := cli.NewApp()
	app.Name = "cartour"
	app.Usage = "fetch, update, query, delete threads and more"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		{
			Name:        "fetch",
			ShortName:   "f",
			Usage:       "fetch threads from bbs",
			Description: "Note: this command will overwrite data, if you want to update it, please use update command",
			Flags: []cli.Flag{
				cli.StringFlag{"sources, s", "", "the source we fetch from, may autohome or bitauto, separated by comma"},
				//cli.StringSliceFlag{"sources, s", &cli.StringSlice{}, "the source we fetch from, may autohome or bitauto"},
				cli.StringFlag{"db, d", "cartour", "mongodb database to use"},
				cli.StringFlag{"collection, c", "threads", "mongodb collection to use"},
				cli.IntFlag{"threads, t", 0, "threads to fetch"},
				cli.IntFlag{"pages, p", 1, "pages to fetch"},
			},
			Action: func(c *cli.Context) {
				Fetch(c.String("sources"), c.Int("pages"), c.Int("threads"))
			},
		},
	}

	app.Run(os.Args)
}

func Fetch(sources string, maxPages, maxThreads int) {
	for _, name := range strings.Split(sources, ",") {
		if len(name) == 0 {
			continue
		}
		threads := []*Thread{}
		if name == autoHome {
			autohome := NewAutoHome()
			threads = autohome.Fetch(maxPages, maxThreads)
		} else if name == bitAuto {
			bitauto := NewBitAuto()
			threads = bitauto.Fetch(maxPages, maxThreads)
		}
		//log.Println("total get threads", len(threads))
		if len(threads) > 0 {
			t := threads[0]
			log.Println(t.Author, t.Title, t.PubTime, t.From)
		}
	}
}
