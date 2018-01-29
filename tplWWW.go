package tplWWW

import (
	"net/http"
	"time"

	"github.com/ohko/hst"
)

// TplWWW ...
type TplWWW struct {
	s *hst.HST
}

// New ...
func New() *TplWWW {
	o := new(TplWWW)
	o.s = hst.NewHST(nil)
	o.s.Favicon()
	o.s.HandleFunc("/sysStatic/", func(c *hst.Context) {
		gz := hst.NewGzip(c.W)
		http.StripPrefix("/sysStatic", http.FileServer(assetFS())).ServeHTTP(gz, c.R)
		gz.CloseGzip()
	})
	return o
}

// RunHTTP ...
func (o *TplWWW) RunHTTP(addr string) {
	go o.s.ListenHTTP(addr)
	hst.Shutdown([]*hst.HST{o.s}, time.Second*5)
}

// RunHTTPS ...
func (o *TplWWW) RunHTTPS(addr, crt, key string) {
	go o.s.ListenHTTPS(addr, crt, key)
	hst.Shutdown([]*hst.HST{o.s}, time.Second*5)
}

// RunTLS ...
func (o *TplWWW) RunTLS(addr, ca, crt, key string) {
	go o.s.ListenTLS(addr, ca, crt, key)
	hst.Shutdown([]*hst.HST{o.s}, time.Second*5)
}

// Static ...
func (o *TplWWW) Static(pattern, path string) {
	o.s.StaticGzip(pattern, path)
}

// HandleFunc ...
func (o *TplWWW) HandleFunc(pattern string, handler ...hst.HandlerFunc) {
	o.s.HandleFunc(pattern, handler...)
}

// RenderFilesByTplDefault ...
func (o *TplWWW) RenderFilesByTplDefault(c *hst.Context, data interface{}, filePath ...string) {
	c.RenderFiles("<!-- {{", "}} -->", data, append(filePath, "../sysStatic/tplDefault.html")...)
}

// RenderContentByTplDefault ...
func (o *TplWWW) RenderContentByTplDefault(c *hst.Context, data interface{}, htm ...string) {
	c.RenderContent("<!-- {{", "}} -->", data, append(htm, string(_sysstaticTpldefaultHtml))...)
}
