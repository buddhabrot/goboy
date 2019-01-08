package main

import (
	"io/ioutil"
	"log"
	"net/url"

	"github.com/zserge/webview"
)

// WebviewVideo draws pixels on a webview
type WebviewVideo struct {
	w webview.WebView
}

// Setup sets up the webview and displays it
func (video *WebviewVideo) Setup() {
	data, err := ioutil.ReadFile("aux/webview.html")
	if err != nil {
		log.Fatal(err)
	}

	video.w = webview.New(webview.Settings{
		URL: `data:text/html,` + url.PathEscape(string(data)),
	})

	video.w.Loop(false)
}

// Step steps the video system
func (video *WebviewVideo) Draw(screen *Screen) {

}
