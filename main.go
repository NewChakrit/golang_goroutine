package main

import "fmt"

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

	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		channel1 <- 10
		close(channel1) // เกิด infinity loop set channel1 = 0 ในครั้งต่อไป
	}()

	go func() {
		channel2 <- 20
		close(channel2) // เกิด infinity loop set channel2 = 0 ในครั้งต่อไป
	}()

	closedChannel1, closedChannel2 := false, false

	for {
		if closedChannel1 && closedChannel2 {
			break
		}
		select {
		case v, ok := <-channel1:
			if !ok {
				closedChannel1 = true
				continue
			}
			fmt.Println("Channel1", v)
		case v, ok := <-channel2:
			if !ok {
				closedChannel2 = true
				continue
			}
			fmt.Println("Channel2", v)
		}
	}
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
