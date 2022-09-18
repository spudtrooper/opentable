// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package proxy

import "time"

//go:generate genopts --prefix=Request --outfile=requestoptions.go "timeout:time.Duration" "timeouts:[]time.Duration"

type RequestOption func(*requestOptionImpl)

type RequestOptions interface {
	Timeout() time.Duration
	Timeouts() []time.Duration
}

func RequestTimeout(timeout time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.timeout = timeout
	}
}
func RequestTimeoutFlag(timeout *time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.timeout = *timeout
	}
}

func RequestTimeouts(timeouts []time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.timeouts = timeouts
	}
}
func RequestTimeoutsFlag(timeouts *[]time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.timeouts = *timeouts
	}
}

type requestOptionImpl struct {
	timeout  time.Duration
	timeouts []time.Duration
}

func (r *requestOptionImpl) Timeout() time.Duration    { return r.timeout }
func (r *requestOptionImpl) Timeouts() []time.Duration { return r.timeouts }

func makeRequestOptionImpl(opts ...RequestOption) *requestOptionImpl {
	res := &requestOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRequestOptions(opts ...RequestOption) RequestOptions {
	return makeRequestOptionImpl(opts...)
}
