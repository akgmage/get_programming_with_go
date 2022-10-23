// In the year 2038, the number of seconds since January 1, 1970 will exceed two billion, the capacity of int32
// Thankfully int64 can support dates beyond 2038. This is a situation where int32 or int simply wont won't do.
// ONly the int64 and uint64 integer types are able to store numbers well beyond two billions on all platforms.
// The following uses the Unix function from time package.
// It accepts tow int64 parameters, corresponding to the nimber of seconds and
// the number of nanoseconds since January 1, 1970.
// Using a suitably larger value (over 12 billion) demonstrates that
// dates beyond 2038 work just fine in Go
package main

import (
	"fmt"
	"time"
)

func main() {
	future := time.Unix(12622780800, 0)
	fmt.Println(future) //2370-01-01 05:30:00 +0530 IST
}
