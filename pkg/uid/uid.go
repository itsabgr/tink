package uid

import (
	"strconv"
	"sync/atomic"
	"time"
)

type UID int64

var start = time.Now().UnixNano()

func (uid UID) String() string {
	return strconv.FormatInt(int64(uid), 32)
}

func Gen() UID {
	return UID(atomic.AddInt64(&start, 1))
}
