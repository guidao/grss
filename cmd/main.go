package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/guidao/grss/config"
	"github.com/guidao/grss/pkg/service"
)

func main() {
	path := flag.String("c", "config.yaml", "-c config.yaml")
	flag.Parse()
	err := config.Init(*path)
	if err != nil {
		panic(err)
	}

	grss := service.NewService()
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/atom", func(w http.ResponseWriter, r *http.Request) {
			rss, err := grss.FetchGithub()
			if err != nil {
				log.Printf("fetch github err:%v", err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(rss))
		})
		http.ListenAndServe(":9999", mux)
	}()

	ch := make(chan struct{})
	<-ch
}
