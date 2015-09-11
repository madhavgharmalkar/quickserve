package main

import (
	"github.com/codegangsta/cli"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	app := cli.NewApp()
	app.Name = "quickserve"
	app.Usage = "quick webserver for testing"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 3000,
			Usage: "http port to listen on",
		},
	}

	app.Action = func(c *cli.Context) {

		port := strconv.Itoa(c.Int("port"))

		fs := http.FileServer(http.Dir("/"))
		http.Handle("/", fs)

		log.Println("http server listening on 127.0.0.1:" + port)
		http.ListenAndServe(":"+port, nil)
	}

	app.Run(os.Args)

}
