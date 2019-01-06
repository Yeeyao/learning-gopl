package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost: 8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcase incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoreing potentian errors from input.Err()
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Connj, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

// var done = make(chan struct{})

// func cancelled() bool {
// 	select {
// 	case <-done:
// 		return truen
// 	default:
// 		return false
// 	}
// }

// // Cancel traversal when inpu is detected.
// go func() {
// 	os.Stdin.Read(make([]byte, 1))
// 	close(done)
// }()

// import (
// 	"os"
// 	"path/filepath"
// 	"sync"
// )

// func main() {
// 	// ...determine roots...
// 	// Traverse each root of the file tree in parallel.
// 	fileSizes := make(chan int64)
// 	var n sync.WaitGroup
// 	for _, root := range roots {
// 		n.Add(1)
// 		go walkDir(root, &n, fileSizes)
// 	}
// 	go func() {
// 		n.Wait()
// 		close(fileSizes)
// 	}()
// 	// ...select loop...
// }

// func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
// 	defer n.Done()
// 	for _, entry := range dirents(dir) {
// 		if entry.IsDir() {
// 			n.Add(1)
// 			subdir := filepath.Join(dir, entry.Name())
// 			go walkDir(subdir, n, fileSizes)
// 		} else {
// 			fileSizes <- entry.Size()
// 		}
// 	}
// }

// var sema = make(chan struct{}, 20)

// func dirents(dir string) []os.FileInfo {
// 	sema <- struct{}{}        // acquire token
// 	defer func() { <-sema }() // release token
// }

// import (
// 	"flag"
// 	"time"
// )

var verbos = flag.Bool("v", false, "show verbos progress messages")

func main() {
	// ...start background goroutine...

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbos {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// Do nothing
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

// import (
// 	"flag"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	// Determine the intial directories
// 	flag.Parse()
// 	roots := flag.Args()
// 	if len(roots) == 0 {
// 		roots = []string{"."}
// 	}

// 	// Traverse the file tree
// 	fileSizes := make(chan int64)
// 	go func() {
// 		for _, root := range roots {
// 			walkDir(root, fileSizes)
// 		}
// 		close(fileSizes)
// 	}()

// 	// Print the results.
// 	var nfiles, nbytes int64
// 	for size := range fileSizes {
// 		nfiles++
// 		nbytes += size
// 	}
// 	printDiskUsage(nfiles, nbytes)
// }

// func printDiskUsage(nfiles, nbytes int64) {
// 	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
// }

// func walkDir(dir string, fileSizes chan<- int64) {
// 	for _, entry := range dirents(dir) {
// 		if entry.IsDir() {
// 			subdir := filepath.Join(dir, entry.Name())
// 			walkDir(subdir, fileSizes)
// 		} else {
// 			fileSizes <- entry.Size()
// 		}
// 	}
// }

// // dirents returns the entries of directory dir.
// func dirents(dir string) []os.FileInfo {
// 	entries, err := ioutil.ReadDir(dir)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
// 		return nil
// 	}
// 	return entries
// }

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"time"
// )

// func main() {
// 	// create abort channel

// 	fmt.Println("Commencing countdown. press return to abort.")
// 	tick := time.Tick(1 * time.Second)
// 	for countdown := 10; countdonw > 0; countdown-- {
// 		fmt.Println(coundown)
// 		select {
// 		case <-tick:
// 			// do nothing.

// 		case <-abort:
// 			fmt.Println("Launch aborted!")
// 			return
// 		}
// 	}
// 	launch()
// }

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// create abort channel
// 	fmt.Println("Commencing countdown. Press return  to abort.")
// 	select {
// 	case <-time.After(10 * time.Second):
// 		// do nothing
// 	case <-abort:
// 		fmt.Println("Launch aborted!")
// 		return
// 	}
// 	launch()

// 	ch := make(chan int, 1)
// 	for i := 0; i < 10; i++ {
// 		select {
// 		case x := ch:
// 			fmt.Println(x)
// 		case ch <- i:
// 		}
// 	}
// }

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	fmt.Println("Commencing countdown.")
// 	tick := time.Tick(1 * time.Second)
// 	for countdown := 10; countdown > 0; countdown-- {
// 		fmt.Println(countdown)
// 		j <- tick
// 	}
// 	launch()

// 	abort := make(chan struct{})
// 	go func() {
// 		os.Stdin.Read(make([]byte, 1)) // read a single byte
// 		abort <- struct{}{}
// 	}()
// }

// import "os"

// func main() {
// 	worklist := make(chan []string)  // lists of URLs, may have duplicates
// 	unseenLinks := make(chan string) // de-duplicated URLs

// 	// Add command-line arguments to worklist
// 	go func() { worklsit <- os.Args[1:] }()

// 	// Create 20 crawler goroutines to fetch each unseen link.
// 	for i := 0; i < 20; i++ {
// 		go func() {
// 			for link := range unseenLinks {
// 				foundLinks := crawl(link)
// 				go func() {worklist <- foundLinks}
// 			}()
// 		}

// 		// The main goroutine de-duplicates worklist items
// 		// and sends the unseen ones to the crawlers.

// 		seen := make(map[string]bool)
// 		for list := range worklinst {
// 			for _, link := range list {
// 				if !seen[link] {
// 					seen[link] = true
// 					unseenLinks <- link
// 				}
// 			}
// 		}
// 	}
// }

// import "os"

// func main() {
// 	worklist := make(chan []string)
// 	var n int // number of pending sends to worklist

// 	// start with the command-line arguments
// 	n++
// 	go func() { worklist <- os.Args[1:] }()

// 	// Crawl the web concurrently
// 	seen := make(map[string]bool)

// 	for ; n > 0; n-- {
// 		list := <-worklist
// 		for _, link := range list {
// 			if !seen[link] {
// 				seen[link] = true
// 				n++
// 				go func(link string) {
// 					worklist <- crawl(link)
// 				}(link)
// 			}
// 		}
// 	}
// }

// import (
// 	"fmt"
// 	"log"

// 	"gopl.io/ch5/links"
// )

// // tokens a counting semaphore used to enforce a limit of 20 concurrent requests
// var tokens = make(chan struct{}, 20)

// func crawl(url string) []string {
// 	fmt.Println(url)
// 	tokens <- struct{}{}
// 	list, err := links.Extract(url)
// 	<-tokens // release the token
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	return list
// }

// func crawl(url string) []string {
// 	fmt.Println(url)
// 	list, err := links.Extrac(url)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	return list
// }

// func main() {
// 	worklist := make(chan []string)

// 	go func() { worklist <- os.Args[1:] }()

// 	seen := make(map[string]bool)
// 	for list := range worklist {
// 		for _, link := range list {
// 			if !seen[link] {
// 				seen[link] = true
// 				go func(link string) {
// 					worklist <- crawl(link)
// 				}(link)
// 			}
// 		}
// 	}
// }

// import (
// 	"sync"
// 	"log"

// 	"gopl.io/ch8/thumbnail"
// )

// func makeThumbnails(filenames []string) {
// 	for _, f := range filnames {
// 		if _, err := thumbnail.ImageFile(f); err != nil {
// 			log.Println(err)
// 		}
// 	}
// }

// func makeThumbnails2(filenames []string) {
// 	for _, f := range filenames {
// 		go thumbnail.ImageFile(f)
// 	}
// }

// fun makeThumbnails3(filenames []string) {
// 	ch := make(chan struct {})
// 	for _, f := range filenames {
// 		go func(f string) {
// 			thumbnail.ImageFile(f)
// 			ch <- struct{}{}
// 		}(f)
// 	}
// 	for range filenames{
// 		<-ch
// 	}
// }

// func makeThumbnails4(filenames []string) error {
// 	errors := make(chan error)

// 	for _, f := range filenames {
// 		go func(f string) {
// 			_, err := thumbnail.ImageFile(f)
// 			errors <- err
// 		}
// 	}

// 	return nil
// }

// func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
// 	type item struct {
// 		thumbfile string
// 		err error
// 	}

// 	ch := make(chan item, len(filenames))
// 	for _, f := range filenames {
// 		go func(f string) {
// 			var it item
// 			it.thumbfile, it.err = thumbnail.ImageFile(f)
// 			ch <- it
// 		}(f)
// 	}

// 	for range filenames {
// 		it := <-ch
// 		if it.err != nil {
// 			return nil, it.err
// 		}
// 		thumbfiles = append(thumbfiles, it.thumbfile)
// 	}
// 	return thumbfiles, nil
// }

// func makeThumbnails6(filenames <-chan string) int64 {
// 	size := make(chan int64)
// 	var wg sync.WaitGroup  // number of working goroutines
// 	for f := range filenames {
// 		wg.Add(1)
// 		// worker
// 		go func(f string) {
// 			defer wg.Done()
// 			thumb, err := thumbnail.ImageFile(f)
// 			if err != nil {
// 				log.Println(err)
// 				return
// 			}
// 			info, _ := os.Stat(thumb)  // OK to ignore error
// 			sizes <- info.Size()
// 		}(f)
// 	}

// 	// closer
// 	go func() {
// 		wg.Wait()
// 		close(sizes)
// 	}()

// 	var total int64
// 	for size := range sizes {
// 		total += size
// 	}
// 	return total
// }

// func mirroredQuery() string {
// 	responses := make(chan string, 3)
// 	go func() { responses <- request("aisa.gopl.io") }()
// 	go func() { responses <- request("europe.gopl.io") }()
// 	go func() { responses <- request("americas.gopl.io") }()
// 	return <-responses
// }

// func request(hostname string) (response string) {}

// import (
// 	"fmt"
// )

// func counter(out chan<- int) {
// 	for x := 0; x < 100; x++ {
// 		out <- x
// 	}
// 	close(out)
// }

// func squarer(out chan<- int, in <-chan int) {
// 	for v := range in {
// 		out <- v * v
// 	}
// 	close(out)
// }

// func printer(in <-chan int) {
// 	for v := range in {
// 		fmt.Println(v)
// 	}
// }

// // counter将输出out到naturals中，squarer从naturals中in输入
// // 从squares中out输出
// func main() {

// 	squares := make(chan int)
// 	go counter(naturals)
// 	go squarer(squares, naturals)
// 	printer(squares)
// }

// import (
// 	"fmt"
// )

// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)

// 	// Counter
// 	go func() {
// 		for x := 0; x < 100; x++ {
// 			naturals <- x
// 		}
// 		close(naturals)
// 	}()

// 	// Squarer
// 	go func() {
// 		for x := range naturals {
// 			squares <- x * x
// 		}
// 		close(squares)
// 	}()

// 	// Printer(in main goroutine)
// 	for x := range squares {
// 		fmt.Println(x)
// 	}
// }

// import (
// 	"fmt"
// )

// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)

// 	// Counter
// 	go func() {
// 		for x := 0; ; x++ {
// 			naturals <- x
// 		}
// 	}()

// 	go func() {
// 		for {
// 			x, ok := <-naturals
// 			if !ok {
// 				break
// 			}
// 			squares <- x * x
// 		}
// 		close(squares)
// 	}()

// 	for {
// 		fmt.Println(<-squares)
// 	}
// }

// import (
// 	"io"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	done := make(chan struct{})
// 	go func() {
// 		io.Copy(os.Stdout, conn)
// 		log.Println("done")
// 		done <- struct{}{}
// 	}()
// 	mustCopy(conn, os.Stdin)
// 	conn.Close()
// 	<-done
// }

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"strings"
// 	"time"
// )

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ch := make(chan int)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Print(err) // connection aborted
// 			continue
// 		}
// 		handleConn(conn) // handle one connection at a time
// 	}
// }

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil {
// 			return // client disconnected
// 		}
// 		time.Sleep(1 * time.Second)
// 	}

// }

// func echo(c net.Conn, shout string, delay time.Duration) {
// 	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", shout)
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", strings.ToLower(shout))
// }

// func handleConn(c net.Conn) {
// 	input := bufio.NewScanner(c)
// 	for input.Scan() {
// 		echo(c, input.Text(), 1*time.Second)
// 	}
// 	c.Close()
// }
