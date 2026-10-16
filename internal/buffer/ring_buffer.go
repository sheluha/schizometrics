package buffer

type Ring[T any] struct {
	pointer int
	length  int
	buffer  []T
}

func NewRingBuffer[T any](length int) *Ring[T] {
	return &Ring[T]{
		length: length,
		buffer: make([]T, length),
	}
}

func (r *Ring[T]) Add(value T) {
	r.buffer[r.pointer] = value
	r.pointer++
	if r.pointer == r.length {
		r.pointer = 0
	}
}

func (r *Ring[T]) AllValues() []T {
	values := make([]T, r.length)
	for i := 0; i < r.length; i++ {
		idx := (r.pointer + i) % r.length
		values[i] = r.buffer[idx]
	}
	return values
}
