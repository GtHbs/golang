package base

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoRoutines = 5
	poolSize      = 3
)

type Pool struct {
	m       sync.Mutex                // 互斥锁，保证资源访问安全
	res     chan io.Closer            // 资源池，实现了io.Closer的接口都可以放进资源池
	factory func() (io.Closer, error) // 资源生成工厂
	closed  bool                      // 资源池是否被关闭
}

type DbConnection struct {
	ID int32 // 连接标志
}

var ErrPoolClosed = errors.New("pool closed")

func NewPool(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}
	return &Pool{factory: fn, res: make(chan io.Closer, size)}, nil
}

func (pool *Pool) Acquire() (io.Closer, error) {
	if pool.closed {
		return nil, ErrPoolClosed
	}
	select {
	case source, err := <-pool.res:
		log.Println("Acquire source")
		if !err {
			return nil, ErrPoolClosed
		}
		return source, nil
	default:
		log.Println("generate new source")
		return pool.factory()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	// 关闭通道
	close(p.res)
	// 关闭通道里的资源
	for re := range p.res {
		re.Close()
	}
}

func (p *Pool) Release(closer io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		closer.Close()
		return
	}
	select {
	case p.res <- closer:
		log.Println("resource put into pool")
	default:
		log.Println("pool is full")
		closer.Close()
	}
}

var idCounter int32

func (db *DbConnection) Close() error {
	log.Printf("close connection, ID:%d", db.ID)
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &DbConnection{id}, nil
}

func dbQuery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用ID为%d的数据库连接", query, conn.(*DbConnection).ID)
}

func PoolExecute() {
	var wg sync.WaitGroup
	wg.Add(maxGoRoutines)
	pool, err := NewPool(createConnection, poolSize)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < maxGoRoutines; i++ {
		go func(q int) {
			dbQuery(q, pool)
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("close pool")
	pool.Close()
}

type SystemConnection struct {
	ID int32
}

func systemQuery(query int, pool *sync.Pool) {
	conn := pool.Get().(*SystemConnection)
	defer pool.Put(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用ID为%d的数据库连接", query, conn.ID)
}

func (conn *SystemConnection) Close() error {
	log.Printf("close system connection, ID:%d", conn.ID)
	return nil
}

func SystemPoolExecute() {
	var wg sync.WaitGroup
	wg.Add(maxGoRoutines * 10)
	pool := &sync.Pool{
		// New表示一个返回对象的方法
		New: createSystemConnection,
	}
	for i := 0; i < maxGoRoutines*10; i++ {
		go func(q int) {
			systemQuery(q, pool)
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("close pool")
}

func createSystemConnection() interface{} {
	id := atomic.AddInt32(&idCounter, 1)
	return &SystemConnection{id}
}
