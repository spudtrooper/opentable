package apihandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	invalidURLParamCharsRE = regexp.MustCompile(`["'<>]`)
)

type errorResponse struct {
	Error string
}

func respondWithError(w http.ResponseWriter, req *http.Request, err error) {
	respondWithJSON(req, w, errorResponse{
		Error: err.Error(),
	})
}

func respondWithErrorString(w http.ResponseWriter, req *http.Request, tmpl string, args ...interface{}) {
	respondWithJSON(req, w, errorResponse{
		Error: fmt.Sprintf(tmpl, args...),
	})
}

func respondWithJSON(req *http.Request, w http.ResponseWriter, obj interface{}) {
	j, err := json.Marshal(obj)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	responseWithJSONBytes(req, w, j)
}

func responseWithJSONBytes(req *http.Request, w http.ResponseWriter, j []byte) {
	fmt.Fprint(w, string(j))
	fmt.Fprint(w, "\n")
	if debug := getBoolURLParam(req, "debug"); debug {
		log.Printf("respondWithJSON: %s", string(j))
	}
}

func sanitizeURLParam(s string) string {
	return invalidURLParamCharsRE.ReplaceAllString(s, "")
}

func getStringURLParam(req *http.Request, key string) string {
	key = sanitizeURLParam(key)
	vals := req.URL.Query()[key]
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func getStringURLParamOrDie(w http.ResponseWriter, req *http.Request, key string) (string, bool) {
	if res := getStringURLParam(req, key); res != "" {
		return res, true
	}
	respondWithErrorString(w, req, fmt.Sprintf("%s required", key))
	return "", false
}

func getStringListURLParam(req *http.Request, key string) []string {
	key = sanitizeURLParam(key)
	p := getStringURLParam(req, key)
	if p == "" {
		return []string{}
	}
	return strings.Split(p, ",")
}

func getIntURLParam(req *http.Request, key string) int {
	key = sanitizeURLParam(key)
	vals := req.URL.Query()[key]
	if len(vals) > 0 {
		v := vals[0]
		if v != "" {
			res, err := strconv.Atoi(v)
			if err != nil {
				log.Printf("getIntURLParam: %v", err)
				return 0
			}
			return res
		}
	}
	return 0
}

func getBoolURLParam(req *http.Request, key string) bool {
	key = sanitizeURLParam(key)
	vals := req.URL.Query()[key]
	if len(vals) > 0 {
		v := vals[0]
		if v == "0" || strings.ToLower(v) == "false" {
			return false
		}
		return true
	}
	return false
}
