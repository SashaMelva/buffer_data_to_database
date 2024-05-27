package buffer

import "github.com/SashaMelva/buffer_data_to_database/internal/entity"

type Buffer struct {
	Len int
	Buf chan *entity.Fact
}

func New() *Buffer {
	ch := make(chan *entity.Fact, 10000)
	return &Buffer{
		Len: 0,
		Buf: ch,
	}
}

func (b *Buffer) Add(fact *entity.Fact) {
	b.Len += 1
	b.Buf <- fact
}

func (b *Buffer) Read() *entity.Fact {
	b.Len -= 1
	fact := <-b.Buf
	return fact
}

func (b *Buffer) Reader() *entity.Fact {
	b.Len -= 1
	fact := <-b.Buf
	return fact
}
