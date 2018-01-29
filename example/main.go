package main

import (
	"github.com/ohko/hst"
	"github.com/ohko/tplWWW"
)

var tt *tplWWW.TplWWW

func main() {
	tt = tplWWW.New()
	tt.Static("/static/", "./static")
	tt.HandleFunc("/", pageIndex)
	tt.RunHTTP("0.0.0.0:8080")
}

func pageIndex(c *hst.Context) {
	tt.RenderFilesByTplDefault(c, nil, "./static/index.html")
}
