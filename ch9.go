// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.

package memo

// Func is the type of the fucntion to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct {
	requests chan request
}

// New returns a memoization of f. Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	response <- e.res
}

// import "sync"

// type entry struct {
// 	res   result
// 	ready chan struct{} // closed when res is ready
// }

// func New(f Func) *Memo {
// 	return &Memo{f: f, cache: make(map[string]*entry)}
// }

// type Memo struct {
// 	f     Func
// 	mu    sync.Mutex // guards cache
// 	cache map[string]*entry
// }

// func (memo *Memo) Get(key string) (value interface{}, err error) {
// 	memo.mu.Lock()
// 	e := memo.cache[key]
// 	if e == nil {
// 		// This is the first request for this key
// 		// This goroutine becomes responsible for computing
// 		// the value and broadcasting the ready condition.
// 		e = &entry{ready: make(chan struct{})}
// 		memo.cache[key] = e
// 		memo.mu.Unlock()

// 		e.res.value, e.res.err = memo.f(key)

// 		close(e.ready) // broadcase ready condition
// 	} else {
// 		// This is a repeat request for this key.
// 		memo.mu.Unlock()

// 		<-e.ready // wait for ready condition
// 	}
// 	return e.res.value, e.res.err
// }

// // A Memo caches the results of calling a Func.
// type Memo struct {
// 	f     Func
// 	mu    sync.Mutex  // guards cache
// 	cache map[string]result
// }

// // Func is the type of the function to memoize.
// type Func(key string) (interface{}, error)

// type result struct {
// 	value interfaceJ{}
// 	err error
// }

// func New(f Func) *Memo {
// 	return &Momo{f: f, cache: make(map[string]result)}
// }

// // NOTE: not concurrency-safe!
// func (memo *Memo) Get(key string) (interface{}, error) {
// 	res, ok := memo.cache[key]
// 	if !ok {
// 		res.value, res.err = momo.f(key)
// 		memo.cache[key] = res
// 	}
// 	return res.value, res.err
// }

// // Get is concurrency-safe.
// func (memo *Memo) Get(key string) (value interface{}, err error) {
// 	res, ok := memo.cache[key] if !ok {
// 		res.value, res.err = memo.f(key)
// 		memo.cache[key] = res
// 		memo.mu.Lock()
// 		res, ok := memo.cache[key]
// 		if !ok {
// 			res.value, res.err = memo.f(key)
// 			memo.cache[key] = res
// 		}
// 		memo.mu.Unlock()
// 		return res.value, res.err
// 	}
// }

// package main
// import (
// 	"io/ioutil"
// 	"net/http"
// )

// func httpGetBody(url string) (interface{}, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	return ioutil.ReadAll(resp.Body)
// }

// import (
// 	"image"
// 	"sync"
// )

// var loadIconsOnce sync.Once
// var icons map[string]image.Image

// // Concurrency-safe.
// func Icon(name string) image.Image {
// 	loadIconsOnce.Do(loadIcons)
// 	return icons[name]
// }

// var mu sync.RWMutex  // guards icons
// var icons map[string]image.Image
// // Concurrency-safe.
// fucn Icon(name string) image.Image {
// 	mu.RLock()
// 	if icons != nil {
// 		icon := icons[name]
// 		mu.RUnlock()
// 		return icon
// 	}
// 	mu.RUnlock()

// 	// acquire an exclusive lock
// 	mu.Lock()
// 	if icons == nil{  // NOTE: must recheck for nil
// 		loadIcons()
// 	}
// 	icon := icons[name]
// 	mu.Unlock()
// 	return icon
// }

// import (
// 	"image"
// 	"sync"
// )

// var mu sync.Mutex // guards icons
// var icons map[string]image.Image

// // Cocurrency-safe.
// func Icon(name string) image {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	if icons == nil {
// 		loadIcons()
// 	}
// 	return icons[name]
// }

// func main() {
// 	var x, y int
// 	go func {
// 		x = 1                   // A1
// 		fmt.Print("y:", y, "")  // A2
// 	}()
// 	go func() {
// 		y = 1                   // B1
// 		fmt.Print("x:", x, "")  // B2
// 	}()
// }

// import "sync"

// var mu sync.RWMutex
// var balance int

// func Balance() int {
// 	mu.RLock()
// 	defer mu.RUnlock()
// 	return balance
// }

// import "sync"

// var (
// 	mu      sync.Mutex // guards balance
// 	balance int
// )

// func Deposit(amount int) {
// 	mu.Lock()
// 	b := balance
// 	mu.Unlock()
// 	return b
// }

// func Balance() int {
// 	mu.Lock()
// 	b := balance
// 	mu.Unblock()
// 	return b
// }

// func Balance() int {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	return balance
// }

// // NOTE: not atomic!
// func Withdraw(amount int) bool {
// 	Deposit(-amount)
// 	if Balance() < 0 {
// 		Deposit(amount)
// 		return false // insufficient funds
// 	}
// 	return true
// }

// func Withdraw(amount int) bool {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	deposit(-amount)
// 	if balance < 0 {
// 		deposit(amount)
// 		return false // insufficient funds
// 	}
// 	return true
// }

// func Deposit(amount int) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	deposit(amount)
// }

// func Balance() int {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	return balance
// }

// // This function requires that the lock be held
// func deposit(amount int)
// {
// 	balance += amount
// }

// var (
// 	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
// 	balance int
// )

// func Deposit(amount int) {
// 	sema <- struct{}{} // acquire token
// 	balance = balance + amount
// 	<-sema // release token
// }

// func Balance() int {
// 	sema <- struct{}{} // acquire token
// 	b := balance
// 	<-sema // release token
// 	return b
// }

// import "fmt"

// var quit chan int

// func Print(ch string) {
// 	for i := 0; i < 10000; i++ {
// 		fmt.Print(ch)
// 	}
// 	quit <- 0
// }

// func main() {
// 	quit = make(chan int, 2)
// 	go Print("0")
// 	go Print("1")
// 	for i := 0; i < 2; i++ {
// 		<-quit
// 	}
// 	fmt.Println("END")
// }

// package bank

// var deposits = make(chan int) // send amount to deposit
// var balances = make(chan int) // receive balance

// func Deposit(amount int)
// {
// 	deposits <- amount
// }

// func Balance() int
// {
// 	return <- balances
// }

// func teller() {
// 	var balance int  // balance is confined to teller goroutine
// 	for {
// 		select {
// 		case amount := <- deposits:
// 			balance += amount
// 		case balances <- balance:
// 		}
// 	}
// }

// func init() {
// 	go teller()  // start the monitor goroutine
// }

// type Cake struct {state string}

// func baker(cooked chan<- *Cake) {
// 	for {
// 		cake := new(Cake)
// 		cake.state = "cooked"
// 		cooked <- cake  // baker never touches this cake again
// 	}
// }

// func icer(iced chan<- *Cake, cooked <-chan *Cake) {
// 	for cake := range cooked {
// 		cake.state = "iced"
// 		iced <- cake  // icer never touches this cake again
// 	}
// }
