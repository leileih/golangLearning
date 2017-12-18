// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
// )

// // 监听指定信号
// func main() {
// 	//合建chan
// 	c := make(chan os.Signal)
// 	//监听指定信号 ctrl+c kill
// 	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGHUP)
// 输出当前进程的pid
// fmt.Println("pid is: ", os.Getpid())
// 	//阻塞直到有信号传入
// 	fmt.Println("启动")

// 	//阻塞直至有信号传入
// 	s := <-c
// 	fmt.Println("退出信号", s)
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
// )

// func main() {
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan,
// 		syscall.SIGHUP,
// 		syscall.SIGINT,
// 		syscall.SIGTERM,
// 		syscall.SIGQUIT)

// 	exitChan := make(chan int)
// 	go func() {
// 		for {
// 			s := <-signalChan
// 			switch s {
// 			// kill -SIGHUP XXXX
// 			case syscall.SIGHUP:
// 				fmt.Println("Got signal: hungup")

// 			// kill -SIGINT XXXX or Ctrl+c
// 			case syscall.SIGINT:
// 				fmt.Println("Warikomi")

// 			// kill -SIGTERM XXXX
// 			case syscall.SIGTERM:
// 				fmt.Println("force stop")
// 				exitChan <- 0

// 			// kill -SIGQUIT XXXX
// 			case syscall.SIGQUIT:
// 				fmt.Println("stop and core dump")
// 				exitChan <- 0

// 			default:
// 				fmt.Println("Unknown signal.")
// 				exitChan <- 1
// 			}
// 		}
// 	}()

// 	code := <-exitChan
// 	os.Exit(code)
// }

// example from golang.org
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
