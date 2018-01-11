// +build arm64,linux

package raw

import (
	"syscall"
	"time"
)

const SYS_SETSOCKOPT = 208

// newTimeval transforms a duration into a syscall.Timeval struct.
// An error is returned in case of zero time value.
func newTimeval(timeout time.Duration) (*syscall.Timeval, error) {
	if timeout < time.Microsecond {
		return nil, &timeoutError{}
	}
	return &syscall.Timeval{
		Sec:  int64(timeout / time.Second),
		Usec: int64(timeout % time.Second / time.Microsecond),
	}, nil
}
