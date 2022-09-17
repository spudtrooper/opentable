package api

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/io"
	goutiljson "github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/request"
)

var (
	token     = flags.String("token", "auth token")
	userCreds = flag.String("user_creds", ".user_creds.json", "file with user credentials")
)

// Client is a client for opentable.com
type Client struct {
	authCke string
}

// NewClientFromFlags creates a new client from command line flags
func NewClientFromFlags() (*Client, error) {
	if *token != "" {
		client := NewClient(*token)
		return client, nil
	}
	if *userCreds != "" {
		client, err := NewClientFromFile(*userCreds)
		if err != nil {
			return nil, err
		}
		return client, nil
	}
	return nil, errors.Errorf("Must set --user & --token or --creds_file")
}

// NewClient creates a new client directly from the API Key
func NewClient(authCke string) *Client {
	return &Client{
		authCke: authCke,
	}
}

// NewClientFromFile creates a new client from a JSON file like `user_creds-example.json`
func NewClientFromFile(credsFile string) (*Client, error) {
	creds, err := readCreds(credsFile)
	if err != nil {
		return nil, err
	}
	return &Client{
		authCke: creds.AuthCke,
	}, nil
}

type Creds struct {
	AuthCke string `json:"authCke"`
}

func ReadCredsFromFlags() (Creds, error) {
	return readCreds(*userCreds)
}

func WriteCredsFromFlags(creds Creds) error {
	b, err := json.Marshal(&creds)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(*userCreds, b, 0755); err != nil {
		return err
	}
	log.Printf("wrote to %s", *userCreds)
	return nil
}

func readCreds(credsFile string) (creds Creds, ret error) {
	if !io.FileExists(credsFile) {
		return
	}
	credsBytes, err := ioutil.ReadFile(credsFile)
	if err != nil {
		ret = err
		return
	}
	if err := json.Unmarshal(credsBytes, &creds); err != nil {
		ret = err
		return
	}
	return
}

func (c *Client) query(operationName string, auth, verbose bool, variables interface{}, sha256Hash string, payload interface{}) error {
	uri := request.MakeURL("https://www.opentable.com/dapi/fe/gql",
		request.MakeParam("optype", `query`),
		request.MakeParam("opname", operationName),
	)

	headers := map[string]string{
		"authority":       `www.opentable.com`,
		"accept":          `*/*`,
		"accept-language": `en-US,en;q=0.9`,
		"cache-control":   `no-cache`,
		"content-type":    `application/json`,
		"user-agent":      `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`,
		"x-csrf-token":    xCSRFToken(),
		"x-query-timeout": `7539`,
	}

	useGpidFromAuth := false
	if auth {
		cookie := [][2]string{
			{"authCke", c.authCke},
		}
		headers["cookie"] = request.CreateCookie(cookie)
		useGpidFromAuth = true
	}

	type PersistedQuery struct {
		Version    int    `json:"version"`
		Sha256Hash string `json:"sha256Hash"`
	}
	type Extensions struct {
		PersistedQuery  PersistedQuery `json:"persistedQuery"`
		UseGpidFromAuth bool           `json:"useGpidFromAuth"`
	}
	type Body struct {
		OperationName string      `json:"operationName"`
		Variables     interface{} `json:"variables"`
		Extensions    Extensions  `json:"extensions"`
	}

	body := Body{
		OperationName: operationName,
		Variables:     variables,
		Extensions: Extensions{
			PersistedQuery: PersistedQuery{
				Version:    1,
				Sha256Hash: sha256Hash,
			},
			UseGpidFromAuth: useGpidFromAuth,
		},
	}

	b, err := json.Marshal(body)
	if err != nil {
		return errors.Errorf("json.Marshal: %v", err)
	}

	if _, err := request.Post(uri, &payload, strings.NewReader(string(b)), request.RequestExtraHeaders(headers)); err != nil {
		return errors.Errorf("request.Post: %v", err)
	}

	if verbose {
		log.Printf("payload: %s", goutiljson.MustColorMarshal(payload))
	}

	return nil
}

func (c *Client) get(uri string, auth, verbose bool) ([]byte, error) {
	headers := map[string]string{
		"authority":                 `www.opentable.com`,
		"accept":                    `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`,
		"accept-language":           `en-US,en;q=0.9`,
		"cache-control":             `no-cache`,
		"dnt":                       `1`,
		"pragma":                    `no-cache`,
		"referer":                   `https://www.opentable.com/`,
		"sec-ch-ua":                 `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`,
		"sec-ch-ua-mobile":          `?0`,
		"sec-ch-ua-platform":        `"macOS"`,
		"sec-fetch-dest":            `document`,
		"sec-fetch-mode":            `navigate`,
		"sec-fetch-site":            `same-origin`,
		"sec-fetch-user":            `?1`,
		"upgrade-insecure-requests": `1`,
		"user-agent":                `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`,
	}

	if auth {
		cookie := [][2]string{
			{"authCke", c.authCke},
		}
		headers["cookie"] = request.CreateCookie(cookie)
	}

	res, err := request.Get(uri, nil, request.RequestExtraHeaders(headers))
	if err != nil {
		return nil, errors.Errorf("request.Get: %v", err)
	}

	payload := res.Data

	if verbose {
		log.Printf("payload: %s", string(payload))
	}

	return payload, nil
}