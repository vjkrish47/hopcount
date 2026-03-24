package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Request struct {
	id   int
	hops int
}

type Router struct {
	next *Router
}

func (r *Router) forward(req *Request) {
	req.hops++
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3)))
	if r.next != nil {
		r.next.forward(req)
	}
}

type Node struct {
	id int
}

func (n *Node) handle(req *Request) {
	req.hops++
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
}

type StaticPath struct {
	routers []*Router
	node    *Node
}

func (p *StaticPath) route(req *Request) {
	if len(p.routers) > 0 {
		p.routers[0].forward(req)
	}
	p.node.handle(req)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	node := &Node{id: 1}

	r1 := &Router{}
	r2 := &Router{}
	r3 := &Router{}

	r1.next = r2
	r2.next = r3

	path := &StaticPath{
		routers: []*Router{r1, r2, r3},
		node:    node,
	}

	var wg sync.WaitGroup

	totalRequests := 100
	totalHops := 0
	var mu sync.Mutex

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			req := &Request{id: id}
			path.route(req)
			mu.Lock()
			totalHops += req.hops
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Requests:", totalRequests)
	fmt.Println("Total Hops:", totalHops)
	fmt.Println("Average Hops:", float64(totalHops)/float64(totalRequests))
}
