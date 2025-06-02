package main

import "sync/atomic"

func onlyatomic() {
	u := User{}
	for atomic.LoadInt64(&u.counter) > 100 {

	}
}
