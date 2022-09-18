package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"sync"

	goutiljson "github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/request"
)

type storedProxyConfig struct {
	ScraperAPIKey      string
	BrightdataUsername string
	BrightdataPassword string
	BrightdataPort     int
}

type ProxyConfig struct {
	*storedProxyConfig
	readLock sync.Mutex
}

func (p *ProxyConfig) MaybeRead(infile string) error {
	if p.storedProxyConfig != nil {
		return nil
	}
	p.readLock.Lock()
	defer p.readLock.Unlock()

	if p.storedProxyConfig != nil {
		return nil
	}

	log.Printf("reading proxy config from %q", infile)

	b, err := ioutil.ReadFile(infile)
	if err != nil {
		return err
	}
	var stored storedProxyConfig
	if err := json.Unmarshal(b, &stored); err != nil {
		return err
	}
	p.storedProxyConfig = &stored

	log.Printf("Proxy config: %s", goutiljson.MustColorMarshal(p))

	return nil
}

func randomChars(n int) string {
	const randomCharSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(randomCharSource[rand.Int()%len(randomCharSource)])
	}
	return buf.String()
}

type requester interface {
	Get(route string, result interface{}, opts ...request.RequestOption) (*request.Response, error)
	Post(url string, result interface{}, body io.Reader, opts ...request.RequestOption) (*request.Response, error)
	String() string
}

func makePassthrough() requester {
	return &passthrough{}
}

func makeURIRewriter(name string, rewriteURI func(uri string) string) requester {
	return &uriRewriter{name: name, rewriteURI: rewriteURI}
}

func makeDedicatedProxy(proxyURL string) requester {
	return &dedicatedProxy{proxyURL}
}

type passthrough struct{}

func (*passthrough) Get(uri string, result interface{}, opts ...request.RequestOption) (*request.Response, error) {
	return request.Get(uri, result, opts...)
}

func (*passthrough) Post(uri string, result interface{}, body io.Reader, opts ...request.RequestOption) (*request.Response, error) {
	return request.Post(uri, result, body, opts...)
}

func (*passthrough) String() string {
	return "passthrough"
}

type uriRewriter struct {
	rewriteURI func(uri string) string
	name       string
}

func (r *uriRewriter) Get(uri string, result interface{}, opts ...request.RequestOption) (*request.Response, error) {
	uri = r.rewriteURI(uri)
	return request.Get(uri, result, opts...)
}

func (r *uriRewriter) Post(uri string, result interface{}, body io.Reader, opts ...request.RequestOption) (*request.Response, error) {
	uri = r.rewriteURI(uri)
	return request.Post(uri, result, body, opts...)
}

func (r *uriRewriter) String() string {
	return r.name
}

type dedicatedProxy struct {
	proxyURL string
}

func (d *dedicatedProxy) opts(rOpts ...request.RequestOption) []request.RequestOption {
	var opts []request.RequestOption
	opts = append(opts, rOpts...)
	opts = append(opts, request.RequestProxyURL(d.proxyURL))
	return opts
}

func (d *dedicatedProxy) Get(uri string, result interface{}, rOpts ...request.RequestOption) (*request.Response, error) {
	return request.Get(uri, result, d.opts(rOpts...)...)
}

func (d *dedicatedProxy) Post(uri string, result interface{}, body io.Reader, rOpts ...request.RequestOption) (*request.Response, error) {
	return request.Post(uri, result, body, d.opts(rOpts...)...)
}

func (d *dedicatedProxy) String() string {
	return fmt.Sprintf("dedicatedProxy(%s)", d.proxyURL)
}

func scraperAPIProxyURL(uri string) string {
	if strings.Contains(uri, "?") {
		uri += "&"
	} else {
		uri += "?"
	}
	uri += randomChars(rand.Int()%10) + "=" + randomChars(rand.Int()%10)
	return fmt.Sprintf("http://api.scraperapi.com?api_key=%s&url=%s", proxyConfig.ScraperAPIKey, url.QueryEscape(uri))
}

func brightdataProxyURL() string {
	proxy := fmt.Sprintf("zproxy.lum-superproxy.io:%d", proxyConfig.BrightdataPort)
	session := rand.Int() % 500000
	proxyUser := fmt.Sprintf("%s-country-us-session-%d:%s", proxyConfig.BrightdataUsername, session, proxyConfig.BrightdataPassword)
	return fmt.Sprintf("http://%s@%s", proxyUser, proxy)
}
