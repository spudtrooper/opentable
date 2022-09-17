package api

import "fmt"

func xCSRFToken() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", randomString(8), randomString(4), randomString(4), randomString(4), randomString(12))
}
