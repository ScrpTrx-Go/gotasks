package main

import "time"

type intMap map[string]int

func (i intMap) Put(key string, value int) {
	i[key] = value
}

func (i intMap) Get(key string) (value int) {
	val, ok := i[key]
	if ok {
		return val
	}
	return 0
}

func main() {
	// MutexFunc()
	SyncMapFunc()
	/* i := make(intMap)
	go func() {
		time.Sleep(100 * time.Millisecond)
		i.Put("stroka", 1)
		time.Sleep(100 * time.Millisecond)
		i.Put("stroka1", 2)
		i.Put("stroka2", 3)
		i.Put("stroka3", 4)
		i.Put("stroka4", 5)
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		i.Put("stroka5", 1)
		i.Put("stroka6", 2)
		i.Put("stroka7", 3)
		i.Put("stroka8", 4)
		i.Put("stroka9", 5)
		time.Sleep(100 * time.Millisecond)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(i) */
	time.Sleep(1 * time.Second)
	// паника возникает потому что обе горутины в одно и то же время кладут значения в мапу: "concurrent map writes"
}
