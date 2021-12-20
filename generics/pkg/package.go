package pkg

type cache[K comparable, V any] struct {
	items map[K]V
}

type Optional[V any] struct {
	value V
	valid bool
}

type CacheV2[K comparable, V any] interface {
	Get(key K) Optional[V]
	Set(key K, item V)
}

func NewGenericsStruct[K comparable, V any]() *cache[K, V] {
	s := cache[K, V]{
		items: make(map[K]V),
	}
	return &s
}

func (c *cache[K, V]) Get(key K) Optional[V] {
	v, ok := c.items[key]
	if ok {
		return Optional[V]{
			value: v,
			valid: true,
		}
	}
	return Optional[V]{
		valid: false,
	}
}

func (c *cache[K, V]) Set(key K, item V) {
	c.items[key] = item
}
