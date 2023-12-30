package singleton_service

type CounterSingleton struct {
	Counter *Counter
}
type Counter struct {
	Count int
}

func (cs *CounterSingleton) CounterSingletonService() *Counter {
	if cs.Counter == nil {
		cs.Counter = counterService()
	}
	return cs.Counter
}

func counterService() *Counter {
	return &Counter{
		Count: 0,
	}
}

// increment count by 1
func (s *Counter) Next() {
	s.Count++
}

// decrement count by 1
func (s *Counter) Previous() {
	s.Count--
}
