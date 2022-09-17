package main

import (
	"flag"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/opentable/api"
)

func readAuthCke(cookie string) (string, error) {
	for _, part := range strings.Split(cookie, "; ") {
		parts := strings.SplitN(part, "=", 2)
		if len(parts) != 2 {
			return "", errors.Errorf("malformed cookie part: %q", part)
		}
		key, val := parts[0], parts[1]
		if key == "authCke" {
			return val, nil
		}
	}
	return "", nil
}

func useCookie() error {
	cookie := flag.Arg(0)
	if cookie == "" {
		return errors.Errorf("cookie should be the first arg")
	}
	authCke, err := readAuthCke(cookie)
	if err != nil {
		return err
	}
	if authCke == "" {
		return errors.Errorf("no authCke")
	}
	creds, err := api.ReadCredsFromFlags()
	if err != nil {
		return errors.Errorf("api.ReadCredsFromFlags: %v", err)
	}
	creds.AuthCke = authCke
	if err := api.WriteCredsFromFlags(creds); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	check.Err(useCookie())
}
