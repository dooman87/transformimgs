package main

import (
	"flag"
	"github.com/dooman87/transformimgs/health"
	"github.com/dooman87/transformimgs/img"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	http.HandleFunc("/health", health.Health)

	img.CheckImagemagick()
	img := img.Service{
		Processor: &img.ImageMagickProcessor{},
		Reader:    &img.ImgUrlReader{},
	}
	http.HandleFunc("/img/resize", http.HandlerFunc(img.ResizeUrl))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
