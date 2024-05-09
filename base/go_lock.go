package base

import (
	"fmt"
	"math/rand"
	"sync"
)

var lockCount int
var waitGroup sync.WaitGroup
var readWriteLock sync.RWMutex

func LockExecute() {
	waitGroup.Add(10)
	for i := 0; i < 5; i++ {
		go Read(i)
	}
	for i := 0; i < 5; i++ {
		go Write(i)
	}
	waitGroup.Wait()
}

func Read(n int) {
	readWriteLock.RLock()
	fmt.Printf("goroutine read %d\n", n)
	v := lockCount
	fmt.Printf("goroutine %d read finish,value is %d\n", n, v)
	waitGroup.Done()
	readWriteLock.RUnlock()
}

func Write(n int) {
	readWriteLock.Lock()
	fmt.Printf("goroutine write %d\n", n)
	v := rand.Intn(100)
	lockCount = v
	fmt.Printf("goroutine %d write finish,value is %d\n", n, v)
	waitGroup.Done()
	readWriteLock.Unlock()
}

type SynchronizedMap struct {
	rw   sync.RWMutex
	data map[interface{}]interface{}
}

func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	sm.data[k] = v
}

func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	return sm.data[k]
}

func (sm *SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	delete(sm.data, k)
}

func (sm *SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	for k, v := range sm.data {
		cb(k, v)
	}
}

func NewSyncMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   sync.RWMutex{},
		data: make(map[interface{}]interface{}),
	}
}

func SyncMapExecute() {
	synchronizedMap := NewSyncMap()
	synchronizedMap.Put("name", "alone")
	synchronizedMap.Put("age", 18)
	synchronizedMap.Each(PrintKV)
	fmt.Println(synchronizedMap.Get("name"))
	synchronizedMap.Delete("name")
	synchronizedMap.Each(PrintKV)

}

func PrintKV(k interface{}, v interface{}) {
	fmt.Printf("key:%v value:%v \n", k, v)
}
