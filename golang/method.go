package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "runtime"
)
/*method继承*/
// type Human struct {
// 	name string
// 	age int
// 	phone string
// }
// type Student struct {
// 	Human
// 	school string
// }
// type Employee struct {
// 	Human
// 	company string
// }
// func (h *Human) SayHi() {
// 	fmt.Printf("Hi, Iam %s you can call me on %s\n", h.name, h.phone)
// }
// func (e *Employee) SayHi() {
// 	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
// }

/*goroutine并发*/
// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		runtime.Gosched()
// 		fmt.Println(s)
// 	}
// }

/*channels*/
// func sum(a []int, c chan int) {
// 	total :=0
// 	for _, v := range a {
// 		total += v
// 	}
// 	c <- total
// }


/*建立web服务器*/
// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	fmt.Println(r.Form)
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello astaxie!")
// }

func get() {
	
}
func main() {
/*method继承*/
	// mark := Student{Human{"Mark", 25, "222-222-YYY"}, "MIT"}
	// sam := Employee{Human{"Sam", 45, "111-888-XXX"}, "Golang Inc"}
	// mark.SayHi()
	// sam.SayHi()

/*goroutine并发*/
	// go say("world")
	// say("hello")

/*channels*/
	// a := []int{7, 2, 8, -9, 4, 0}
	// c := make(chan int)
	// go sum(a[:len(a)/2], c)
	// go sum(a[len(a)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)

/*建立web服务器*/
	// http.HandleFunc("/", sayhelloName)
	// err := http.ListenAndServe(":9090", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	// client := &http.Client{}
	// request, _ := http.NewRequest("GET", "https://api.coinmarketcap.com/v1/ticker/ethereum/", nil)
	// request.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	// request.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	// request.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	// request.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	// request.Header.Set("Cache-Control","max-age=0")
	// request.Header.Set("Connection","keep-alive")
	// request.Header.Set("User-Agent","chrome 100")
	// response, _ := client.Do(request)
	// if response.StatusCode == 200 {
	// 	body, _ := ioutil.ReadAll(response.Body)
	// 	bodystr := string(body);
	// 	fmt.Println(bodystr)
	// }
	get()
}