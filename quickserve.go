package main

import (
	"github.com/codegangsta/cli"
	"log"
	"net/http"
	"os"
	"strconv"
)

func httpLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %q\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

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

		log.Println("HTTP server listening on localhost:" + port)
		http.ListenAndServe(":"+port, httpLog(http.DefaultServeMux))
	}

	app.Run(os.Args)

}
