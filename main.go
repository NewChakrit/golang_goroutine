package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m sync.Mutex

var n = 10

func main() {
	//go helloSet1() // เขียน go ข้างหน้า จะเป็น goroutine ทันที
	//helloSet2()
	//
	//time.Sleep(time.Second)

	// ================= Channel ================= //
	// Create New Channel
	//ch := make(chan int)

	//go func() {
	//	time.Sleep(time.Second * 2)
	//	fmt.Println("Test 1")
	//	ch <- 20
	//	ch <- 30
	//}() // จบงาน ปิด thread
	//// sent data to channel
	////ch <- 10

	//fmt.Println("Test 2")
	//
	//// get data from channel
	//v := <-ch
	//fmt.Println(v)
	//
	//v = <-ch
	//fmt.Println(v)

	// Loop Channel
	//go func() {
	//	ch <- 10
	//	ch <- 20
	//	ch <- 30
	//	ch <- 40
	//	close(ch)
	//}()
	//
	//for v := range ch {
	//	fmt.Println(v)
	//}

	// Select  Channel
	//
	//channel1 := make(chan int)
	//channel2 := make(chan int)
	//
	//go func() {
	//	channel1 <- 10
	//	close(channel1) // เกิด infinity loop set channel1 = 0 ในครั้งต่อไป
	//}()
	//
	//go func() {
	//	channel2 <- 20
	//	close(channel2) // เกิด infinity loop set channel2 = 0 ในครั้งต่อไป
	//}()
	//
	//closedChannel1, closedChannel2 := false, false
	//
	//for {
	//	if closedChannel1 && closedChannel2 {
	//		break
	//	}
	//	select {
	//	case v, ok := <-channel1:
	//		if !ok {
	//			closedChannel1 = true
	//			continue
	//		}
	//		fmt.Println("Channel1", v)
	//	case v, ok := <-channel2:
	//		if !ok {
	//			closedChannel2 = true
	//			continue
	//		}
	//		fmt.Println("Channel2", v)
	//	}
	//}
	//
	//// =============== Sync =============== //
	//// WaitGroup
	//
	//var wg sync.WaitGroup
	//
	//// Launch several goroutines and increment the WaitGroup counter for each
	//wg.Add(5)
	//for i := 1; i <= 5; i++ {
	//	go worker(i, &wg)
	//}
	//
	//wg.Wait() // Block until the WaitGroup counter goes back to 0; all workers are done
	//
	//fmt.Println("All workers completed")
	//
	//// Mutex
	//
	//fmt.Println("FIRST")
	//go p()
	//fmt.Println("SECOND")
	//p()
	//fmt.Println("THIRD")
	//time.Sleep(3 * time.Second)
	//fmt.Println("DONE")
	//
	////var wg sync.WaitGroup
	//counter := Counter{}
	//
	//// Start 10 goroutines
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for j := 0; j < 100; j++ {
	//			counter.Increment()
	//		}
	//	}()
	//}
	//
	//wg.Wait() // Wait for all goroutines to finish
	//fmt.Println("Final counter value:", counter.Value())
	//
	//// Once
	//var once sync.Once
	////var wg sync.WaitGroup
	//
	//initialize := func() {
	//	fmt.Println("Initializing only once")
	//}
	//
	//doWork := func(workerId int) {
	//	defer wg.Done()
	//	fmt.Printf("Worker %d started\n", workerId)
	//	once.Do(initialize) // This will only be executed once
	//	fmt.Printf("Worker %d done\n", workerId)
	//}
	//
	//numWorkers := 5
	//wg.Add(numWorkers)
	//
	//// Launch several goroutines
	//for i := 0; i < numWorkers; i++ {
	//	go doWork(i)
	//}
	//
	//// Wait for all goroutines to complete
	//wg.Wait()
	//fmt.Println("All workers completed")

	// Cond
	//var once sync.Once
	//var wg sync.WaitGroup

	//initialize := func() {
	//	fmt.Println("Initializing only once")
	//}
	//
	//doWork := func(workerId int) {
	//	defer wg.Done()
	//	fmt.Printf("Worker %d started\n", workerId)
	//	once.Do(initialize) // This will only be executed once
	//	fmt.Printf("Worker %d done\n", workerId)
	//}
	//
	//numWorkers := 5
	//wg.Add(numWorkers)
	//
	//// Launch several goroutines
	//for i := 0; i < numWorkers; i++ {
	//	go doWork(i)
	//}
	//
	//// Wait for all goroutines to complete
	//wg.Wait()
	//fmt.Println("All workers completed")

	// ================ Pub/Sub ================ //

	// สร้าง channel เพื่อส่งข้อความ
	//ch := make(chan string)

	// Publisher //create goroutine for sent message to channel
	//go func() {
	//	for i := 0; i < 10; i++ {
	//	}
	//	ch <- fmt.Sprintf("Hello, world! %d", i)
	//	time.Sleep(1 * time.Second)
	//}()
	//
	//// Subscriber // create goroutine for get message from chanel
	//go func() {
	//	for {
	//		msg := <-ch
	//		fmt.Println(msg)
	//	}
	//}()
	//
	//// wait for goroutine working done
	//time.Sleep(5 * time.Second)

	Pubsub()
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // D≠ecrement the counter when the goroutine completes

	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work by sleeping
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d done\n", id)
}

func p() {
	m.Lock()
	fmt.Println("LOCK")

	fmt.Println(n)

	time.Sleep(1 * time.Second)
	m.Unlock()

	fmt.Println("UNLOCK")
}

// Counter struct holds a value and a mutex
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment method increments the counter's value safely using the mutex
func (c *Counter) Increment() {
	c.mu.Lock()   // Lock the mutex before accessing the value
	c.value++     // Increment the value
	c.mu.Unlock() // Unlock the mutex after accessing the value
}

// Value method returns the current value of the counter
func (c *Counter) Value() int {
	return c.value
}

//func helloSet1() {
//	fmt.Println("Hello 1")
//	fmt.Println("Hello 2")
//}
//
//func helloSet2() {
//	fmt.Println("Hello 3")
//	fmt.Println("Hello 4")
//}
