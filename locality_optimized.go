import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Request struct {
	id int
}

type ServiceInstance struct {
	id   int
	hops int
}

func (s *ServiceInstance) handle(r Request) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(4)))
}

type HopAwareGateway struct {
	services []*ServiceInstance
	mu       sync.Mutex
}

func NewGateway(services []*ServiceInstance) *HopAwareGateway {
	return &HopAwareGateway{services: services}
}

func (g *HopAwareGateway) selectInstance() *ServiceInstance {
	g.mu.Lock()
	defer g.mu.Unlock()
	min := g.services[0]
	for _, s := range g.services {
		if s.hops < min.hops {
			min = s
		}
	}
	return min
}

func (g *HopAwareGateway) route(r Request) int {
	instance := g.selectInstance()
	instance.handle(r)
	return instance.hops
}

func main() {
	rand.Seed(time.Now().UnixNano())

	clusterSizes := []int{3, 5, 7, 9, 11}

	for _, size := range clusterSizes {
		var services []*ServiceInstance
		for i := 0; i < size; i++ {
			services = append(services, &ServiceInstance{
				id:   i,
				hops: rand.Intn(4) + 1,
			})
		}

		gateway := NewGateway(services)

		totalHops := 0
		requests := 200
		var wg sync.WaitGroup
		var mu sync.Mutex

		for i := 0; i < requests; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				h := gateway.route(Request{id: id})
				mu.Lock()
				totalHops += h
				mu.Unlock()
			}(i)
		}

		wg.Wait()
		fmt.Println("Cluster Size:", size, "Average Hops:", float64(totalHops)/float64(requests))
	}
}
