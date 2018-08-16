package main

func main() {
	inter201()
}



//实现简单的GC
func inter201() {
	var p Pool
	for i := 0; i < 10000000; i++ {
		_ = p.alloc(100)
	}
}
type Pool struct {
	buf []byte
}
func (p *Pool) alloc(size int) []byte {
	if len(p.buf) < size {
		l := 1024 * 1024
		for l < size {
			l += l
			if l <= 0 {
				panic("out of memory")
			}
		}
		p.buf = make([]byte, l)
	}
	buff := p.buf[:size]
	p.buf = p.buf[size:]
	return buff
}
