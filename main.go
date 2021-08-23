package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

type A struct {
	 name string
}

func (a A) Name() string {
	a.name = "hi" + a.name
	return a.name
}

func NameOfA(a A) string  {
	a.name = "hi" + a.name
	return a.name
}

//func main() {
//	//a := A{name:"eggo"}
//	//fmt.Println(a.Name())
//	//fmt.Println(A.Name(a))
//
//	t1 := reflect.TypeOf(A.Name)
//	t2 := reflect.TypeOf(NameOfA)
//
//	fmt.Println(t1)
//	fmt.Println(t2)
//	fmt.Println(t1 == t2)
//}

//func main() {
//	ctx := context.Background()
//	process(ctx)
//
//	ctx = context.WithValue(ctx, "traceId","qcrao-2019")
//	process(ctx)
//
//
//}

func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n",traceId)
	} else {
		fmt.Printf("process over.no trace_id\n")
	}
}

//发生协程泄露
//func gen() <-chan int {
//	ch := make(chan int)
//	go func() {
//		var n int
//		for {
//			ch <- n
//			n++
//			time.Sleep(time.Second)
//		}
//	}()
//	return ch
//}

//func main() {
//	for n := range gen() {
//		fmt.Println(n)
//		if n == 5 {
//			break
//		}
//	}
//
//}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <- ctx.Done():
				return
			case ch <- n:
				n++
				time.Sleep(time.Second)
			}
		}
	}()
	return ch
}


func goroutineA(a <- chan int) {
	val := <- a
	fmt.Println("goroutine A received data: ",val)
	return

}

func goroutineB(b <- chan int) {
	val := <- b
	fmt.Println("goroutine B received data: ",val)
	return
}

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "hello, world"
	done <-true
}


//func main() {
//	//ctx, cancel := context.WithCancel(context.Background())
//	//defer cancel()
//	//for n := range gen(ctx) {
//	//	fmt.Println(n)
//	//	if n == 5 {
//	//		cancel()
//	//		break
//	//	}
//	//}
//
//	//ch := make(chan int)
//	//go goroutineA(ch)
//	//go goroutineB(ch)
//	//
//	//ch <- 3
//	//time.Sleep(time.Second)
//
//	go aGoroutine()
//	<- done
//	println(msg)
//
//}

////做定时器
//func worker() {
//	ticker := time.Tick(1 * time.Second)
//	for {
//		select {
//		case <- ticker:
//
//		}
//	}
//
//
//}


//
func worker(taskCh <- chan int) {
	const N = 5

	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <- taskCh
				fmt.Printf("finish task: %d by worker %d\n",task,id)
				time.Sleep(time.Second)
			}
		}(i)
	}

}


type Child struct {
	Name string
	handsome bool
}


//func main() {
//	qcrao := Child{
//		Name : "qcrao",
//		handsome: true,
//	}
//	v := reflect.ValueOf(&qcrao)
//	f := v.Elem().FieldByName("Name")
//	fmt.Println(f.String())
//
//	f.SetString("stefno")
//	fmt.Println(f.String())
//
//	f = v.Elem().FieldByName("handsome")
//	//f.SetBool(true)
//	//fmt.Println(f.Bool())
//
//	golang := Go{}
//	php    := PHP{}
//
//	sayHello(golang)
//	sayHello(php)
//}

//func main() {
//	var m sync.Map
//	m.Store("qcrao", 18)
//	m.Store("stefno", 20)
//
//	age , _ := m.Load("qcrao")
//	fmt.Println(age.(int))
//
//	m.Range(func(key, value interface{}) bool {
//		name := key.(string)
//		age := value.(int)
//		fmt.Println(name, age)
//		return true
//	})
//
//	m.Delete("qcrao")
//	age, ok := m.Load("qcrao")
//	fmt.Println(age, ok)
//
//	m.LoadOrStore("stefno", 100)
//	age, _ = m.Load("stefno")
//	fmt.Println(age)
//
//}

type IGreeting interface {
	sayHello()
}

type Go struct {

}
func (g Go) sayHello() {
	fmt.Println("Hi, I am go")
}

type PHP struct {

}
func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP")
}

func sayHello(i IGreeting) {
	i.sayHello()
}


type Duck interface {
	Quack()
}

type Cat struct {

}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y  =  y, x + y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

//func main() {
//	ch := make(chan int)
//	select {
//	case i := <- ch:
//		println(i)
//
//	default:
//		println("default")
//	}
//}

//func main() {
//	ch := make(chan int)
//	go func() {
//		for range time.Tick(1 * time.Second) {
//			ch <- 0
//		}
//	}()
//
//	for {
//		select {
//		case <- ch:
//			println("case1")
//		case <- ch:
//			println("case2")
//
//		}
//	}
//}


func main() {
	ctx, cancel := context.WithTimeout(context.Background(),10 *time.Second)
	defer cancel()

	go handle(ctx, 500 *time.Millisecond)
	select {
	case <- ctx.Done():
		fmt.Println("main",ctx.Err())
	}

}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <- ctx.Done():
		fmt.Println("handle",ctx.Err())
	case <- time.After(duration):
		fmt.Println("process request with", duration)
	}
	
}

func main() {
	reflect.TypeOf()
}