//go:generate go-bindata -pkg ctstatic -ignore .../.DS_Store -o files.go files/...

package ctstatic

import (
	"log"
	"net/http"
	"os"

	"github.com/elazarl/go-bindata-assetfs"
)

// all static/ files embedded as a Go library
func FileSystemHandler() http.Handler {
	var h http.Handler
	if info, err := os.Stat("static/files/"); err == nil && info.IsDir() {
		log.Println("Using local static files")
		h = http.FileServer(http.Dir("static/files/"))
	} else if info, err := os.Stat("/go/src/github.com/jpillora/cloud-torrent/static/files/"); err == nil && info.IsDir() {
		log.Println("Using local absolute static files")
		h = http.FileServer(http.Dir("/go/src/github.com/jpillora/cloud-torrent/static/files/"))
	} else {
		log.Println("Using packed static files")
		h = http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "files"})
	}
	return h
}
