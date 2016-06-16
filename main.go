package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ignatov/go-reactjs-example/modules"
	"gopkg.in/macaron.v1"
)

type Configure struct {
	port    int
	root    string
	private bool
}

var gcfg = Configure{}

var m *macaron.Macaron

func init() {
	m = macaron.Classic()
	m.Use(modules.Public)
	m.Use(modules.Renderer)

	flag.IntVar(&gcfg.port, "port", 8000, "Which port to listen")
	flag.BoolVar(&gcfg.private, "private", false, "Only listen on lookback interface, otherwise listen on all interface")
}

func initRouters() {
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "homepage", nil)
	})

	m.Get("/comments", func(c *macaron.Context) {
		type Comment struct {
			Id     int    `json:"id,omitempty"`
			Author string `json:"author,omitempty"`
			Text   string `json:"text,omitempty"`
		}

		comments := [...]Comment{
			Comment{Id: 1, Author: "Pete Hunt", Text: "This is one comment"},
			Comment{Id: 2, Author: "Jordan Walke", Text: "This is *another* comment"},
			Comment{Id: 3, Author: "Jordan Walke", Text: "This is comment"},
			Comment{Id: 4, Author: "Jordan W1234alkeasdad", Text: "This is comment"},
		}

		c.JSON(200, comments)
	})

	ReloadProxy := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Debug, Hot reload", r.Host)
		resp, err := http.Get("http://localhost:3000" + r.RequestURI)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	}
	m.Get("/-/:rand(.*).hot-update.:ext(.*)", ReloadProxy)
	m.Get("/-/bundle.js", ReloadProxy)
}

func main() {
	flag.Parse()
	initRouters()

	http.Handle("/", m)

	i := ":" + strconv.Itoa(gcfg.port)
	p := strconv.Itoa(gcfg.port)
	mesg := "; please visit http://127.0.0.1:" + p
	if gcfg.private {
		i = "localhost" + i
		log.Printf("listens on 127.0.0.1@" + p + mesg)
	} else {
		log.Printf("listens on 0.0.0.0@" + p + mesg)
	}
	if err := http.ListenAndServe(i, nil); err != nil {
		log.Fatal(err)
	}
}
