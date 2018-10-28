package pool

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestPoolWorkRequest struct {
	Counter   *int
	RequestWg *sync.WaitGroup
}

func (r *TestPoolWorkRequest) Execute() {
	defer r.RequestWg.Done()
	*r.Counter++
}

func TestPoolWork(t *testing.T) {
	poolObj := NewPool(5)

	var counter int
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		poolObj.Exec(&TestPoolWorkRequest{Counter: &counter, RequestWg: &wg})
	}

	wg.Wait()
	assert.Equal(t, counter, 100)
}
