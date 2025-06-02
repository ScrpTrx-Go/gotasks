package main

import (
	"fmt"
	"sync"
	"time"
)

type MyMap struct {
	m sync.Map
}

func (m *MyMap) TestSyncMapPut(key string, val int) {
	m.m.Store(key, val)
}

func (m *MyMap) TestSyncMapGet(key string) (val interface{}) {
	value, ok := m.m.Load(key)
	if ok {
		return value
	}
	return 0
}

func SyncMapFunc() {
	var wg sync.WaitGroup
	var MyMap MyMap
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		MyMap.TestSyncMapPut("str1", 1)
		time.Sleep(100 * time.Millisecond)
		MyMap.TestSyncMapPut("str2", 2)
		MyMap.TestSyncMapPut("str3", 3)
		MyMap.TestSyncMapPut("str4", 4)
		MyMap.TestSyncMapPut("str5", 5)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		MyMap.TestSyncMapPut("str1", 1)
		MyMap.TestSyncMapPut("str2", 2)
		MyMap.TestSyncMapPut("str3", 3)
		MyMap.TestSyncMapPut("str4", 4)
		MyMap.TestSyncMapPut("str5", 5)
		time.Sleep(100 * time.Millisecond)
	}()

	go func() {
		wg.Wait()
		res := MyMap.TestSyncMapGet("str1")
		res1 := MyMap.TestSyncMapGet("str2")
		res2 := MyMap.TestSyncMapGet("str3")
		fmt.Println(res, res1, res2)
	}()

	time.Sleep(1 * time.Second)
	MyMap.m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
