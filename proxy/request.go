package proxy

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	goutilerrors "github.com/spudtrooper/goutil/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/request"
)

var (
	numScraperAPIProxies = flags.Int("num_scraper_api_proxies", "number of scraperapi.com proxies to use")
	numBrightdataProxies = flags.Int("num_brightdata_proxies", "number of brightdata.com proxies to use")
	requestRetries       = flags.Int("request_retries", "number of HTTP retries, assuming we aren't using a proxy")
	proxyConfigFile      = flag.String("proxy_config_file", ".proxy.json", "proxy config file")
)

type respFn struct {
	fn  func() (*Response, error)
	str func() string
}

// redundantProxy fans out `respFn` `n` times and uses the first valid result
func redundantProxy(uri string, fns []respFn) (*Response, error) {
	type resp struct {
		res *Response
		err error
	}
	respCh := make(chan resp)

	go func() {
		var wg sync.WaitGroup
		var wgCnt int32
		var wgMu sync.Mutex
		var isDone bool
		add := func() {
			atomic.AddInt32(&wgCnt, 1)
			wg.Add(1)
		}
		done := func() bool {
			if atomic.LoadInt32(&wgCnt) > 0 {
				wg.Done()
				atomic.AddInt32(&wgCnt, -1)
				return false
			}
			return true
		}
		for i, fn := range fns {
			i := i
			fn := fn
			add()
			go func() {
				res, err := fn.fn()
				if cnt := atomic.LoadInt32(&wgCnt); cnt == 0 {
					log.Printf("%s from proxy[%d: %s]", color.HiYellowString("THROW AWWAY"), i, fn.str())
					return
				}
				if err == nil {
					if strings.Contains(string(res.Data), `You're sending requests a bit too fast!`) {
						err = errors.Errorf("Application error: %s", string(res.Data))
					}
				}
				wgMu.Lock()
				defer wgMu.Unlock()
				// Use double-lock checking for anyone waiting on this lock
				if isDone {
					return
				}
				isDone = true
				respCh <- resp{
					res: res,
					err: err,
				}
				done()
				if err != nil {
					log.Printf("%s from proxy[%d: %s]: %v", color.HiRedString("ERROR"), i, fn.str(), err)
				} else {
					log.Printf("%s from proxy[%d: %s]", color.HiGreenString("SUCCESS"), i, fn.str())
					for {
						if done() {
							break
						}
					}
				}
			}()
		}
		wg.Wait()
		close(respCh)
	}()

	// Return the first valid result
	var err error
	for r := range respCh {
		if r.err == nil && r.res != nil {
			return r.res, nil
		}
		err = r.err
	}
	return nil, err
}

type Response struct {
	Resp    *http.Response
	Data    []byte
	Cookies []request.Cookie
	Payload interface{}
}

func Get(uri string, payloadFn func() interface{}, headers map[string]string, opts ...request.RequestOption) (*Response, error) {
	return get(uri, payloadFn, headers, opts...)
}

func ReadURL(uri string, headers map[string]string, rOpts ...RequestOption) ([]byte, error) {
	opts := MakeRequestOptions(rOpts...)

	timeouts := opts.Timeouts()
	if len(opts.Timeouts()) == 0 {
		timeout := or.Duration(opts.Timeout(), time.Second*30)
		timeouts = []time.Duration{timeout}
	}

	ec := goutilerrors.MakeErrorCollector()
	for i, to := range timeouts {
		if i > 0 {
			log.Printf("retrying %s", uri)
		}
		resp, err := get(uri, nil, headers, request.RequestTimeout(to))
		if err != nil {
			ec.Add(err)
			continue
		}
		return resp.Data, nil
	}
	return nil, ec.Build()
}

func get(uri string, payloadFn func() interface{}, headers map[string]string, opts ...request.RequestOption) (*Response, error) {
	requesters := makeRequesters()

	if payloadFn == nil {
		payloadFn = func() interface{} { return nil }
	}

	var fns []respFn
	for _, req := range requesters {
		req := req
		fn := func() (*Response, error) {
			payload := payloadFn()
			resp, err := req.Get(uri, payload, request.RequestExtraHeaders(headers))
			if err != nil {
				return nil, err
			}
			res := &Response{
				Resp:    resp.Resp,
				Data:    resp.Data,
				Cookies: resp.Cookies,
				Payload: payload,
			}
			return res, nil
		}
		fns = append(fns, respFn{
			fn: fn,
			str: func() string {
				return req.String()
			},
		})
	}

	if len(fns) == 1 {
		retries := or.Int(*requestRetries, 1)
		fn := fns[0].fn
		if retries == 1 {
			return fn()
		}
		return requestWithRetries(fn, retries)
	}

	return redundantProxy(uri, fns)
}

func requestWithRetries(fn func() (*Response, error), retries int) (*Response, error) {
	var retErr error
	for i := 0; i < retries; i++ {
		res, err := fn()
		if err != nil {
			retErr = err
			continue
		}
		if res != nil {
			// if b, m := isBlocked(res); b {
			sleep := time.Millisecond * time.Duration(100*(i+1))
			time.Sleep(sleep)
			m := "TODO"
			retErr = errors.Errorf("blocked by %v after %d tries", m, i)
			continue
			// }
		}
		if res != nil {
			return res, nil
		}
	}
	return nil, retErr
}

func Post(uri string, payload interface{}, body io.Reader, headers map[string]string, opts ...request.RequestOption) (*request.Response, error) {
	reqs := makeRequesters()

	var retErr error
	for i, req := range reqs {
		if i > 0 {
			log.Printf("retry[%d] %s", i, uri)
		}
		resp, err := req.Post(uri, payload, body, request.RequestExtraHeaders(headers))
		if err != nil {
			retErr = err
			continue
		}
		return resp, nil
	}
	return nil, retErr
}

var proxyConfig = &ProxyConfig{}

func makeRequesters() []requester {
	if err := proxyConfig.MaybeRead(*proxyConfigFile); err != nil {
		check.Err(err)
	}

	var res []requester
	for i := 0; i < *numScraperAPIProxies; i++ {
		res = append(res, makeURIRewriter("scraperapi", scraperAPIProxyURL))
	}
	for i := 0; i < *numBrightdataProxies; i++ {
		proxyURL := brightdataProxyURL()
		res = append(res, makeDedicatedProxy(proxyURL))
	}
	if len(res) == 0 {
		res = append(res, makePassthrough())
	}
	return res
}
