package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	"unicode/utf8"
)
var wg sync.WaitGroup

func main() {
	//getInterfaces()
	//go sayHello()
	//wg.Add(2)
	//go func(){
	//	fmt.Println("hello, GOLANG")
	//	wg.Done()
	//}()
	//saySomething := func() {
	//	fmt.Println("hello, world2")
	//	wg.Done()
	//}
	//go saySomething()
	//
	//wg.Wait()
	//fmt.Println("Hello, world")
	//time.Sleep(1 *time.Second)
	//PORT := ":8001"
	//arguments := os.Args
	//if len(arguments) == 1 {
	//	fmt.Println("Using default port number: ", PORT)
	//}else{
	//	PORT = ":" + arguments[1]
	//	fmt.Println("Using port number: ", PORT)
	//}
	//r := http.NewServeMux()
	//r.HandleFunc("/time", timeHandler)
	//r.HandleFunc("/", myHandler)
	//r.HandleFunc("/debug/pprof/", pprof.Index)
	//r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	//err := http.ListenAndServe(PORT, r)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//n := flag.Int("n", 20, "Number of goroutines")
	//flag.Parse()
	//count := *n
	//fmt.Printf("Going to create %d goroutines.\n", count)
	//var waitGroup sync.WaitGroup
	//
	//waitGroup.Add(1)
	//
	//fmt.Printf("%#v\n", waitGroup)
	//for i := 0; i < count; i++ {
	//	waitGroup.Add(1)
	//	go func(x int) {
	//		defer waitGroup.Done()
	//		fmt.Printf("%d ", x)
	//	}(i)
	//}
	//
	//fmt.Printf("%#v\n", waitGroup)
	//waitGroup.Wait()
	//fmt.Println("\nExiting...")

	//var test map[string]int
	//test["12"]=1
	//var s []int
	//s = append(s,1)
	//x := [3]int{1,2,3}
	//go func(arr *[3]int){
	//	arr[0] = 7
	//
	//
	//}(&x)
	//time.Sleep(1 * time.Second)
	//getUnicode()
	fmt.Println(INT_MAX)
	fmt.Println(^INT_MAX)

}


func sayHello() {
	const sLiteral= "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x99"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)
	fmt.Printf("sLiteral length: %d\n", len(sLiteral))
}

func getInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range interfaces {
		fmt.Printf("Interface: %v\n", v.Name)
		byName, err := net.InterfaceByName(v.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n",k,v)
		}
		fmt.Println()

	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func generateSlice() {
	x := 3
	y := 4

	table := make([][]int,x)
	for i := range table {
		table[i] = make([]int, y)
	}
}


func getUnicode() {
	data := "♥"
	fmt.Println(utf8.RuneCountInString(data))
}
