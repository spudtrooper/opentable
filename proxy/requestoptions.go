// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package proxy

import "time"

//go:generate genopts --prefix=Request --outfile=requestoptions.go "timeout:time.Duration" "timeouts:[]time.Duration"

type RequestOption func(*requestOptionImpl)

type RequestOptions interface {
	Timeout() time.Duration
	HasTimeout() bool
	Timeouts() []time.Duration
	HasTimeouts() bool
}

func RequestTimeout(timeout time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.has_timeout = true
		opts.timeout = timeout
	}
}
func RequestTimeoutFlag(timeout *time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		if timeout == nil {
			return
		}
		opts.has_timeout = true
		opts.timeout = *timeout
	}
}

func RequestTimeouts(timeouts []time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.has_timeouts = true
		opts.timeouts = timeouts
	}
}
func RequestTimeoutsFlag(timeouts *[]time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		if timeouts == nil {
			return
		}
		opts.has_timeouts = true
		opts.timeouts = *timeouts
	}
}

type requestOptionImpl struct {
	timeout      time.Duration
	has_timeout  bool
	timeouts     []time.Duration
	has_timeouts bool
}

func (r *requestOptionImpl) Timeout() time.Duration    { return r.timeout }
func (r *requestOptionImpl) HasTimeout() bool          { return r.has_timeout }
func (r *requestOptionImpl) Timeouts() []time.Duration { return r.timeouts }
func (r *requestOptionImpl) HasTimeouts() bool         { return r.has_timeouts }

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
