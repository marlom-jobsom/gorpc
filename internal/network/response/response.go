package response

import (
	"time"
)

// Response wrappers a client request's answer of a remote method execution
type Response struct {
	Content  []interface{}
	Duration time.Duration
}
