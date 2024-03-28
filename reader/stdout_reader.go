package reader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

// BlockingReader simply reads from the provided reader until no content is left, at which point it waits for more.
type BlockingReader struct {
	r   io.Reader
	ctx context.Context
}

func (br *BlockingReader) Read(p []byte) (n int, err error) {
	for {
		select {
		case <-br.ctx.Done():
			return 0, fmt.Errorf("timeout before contents")
		default:
			n, err = br.r.Read(p)
			if err != io.EOF {
				return
			}
			readContent := n > 0
			if readContent {
				return
			}
			time.Sleep(time.Millisecond) // dont needlessly hog cpu
		}
	}
}

func TestBlockingReader() {
	buf := bytes.Buffer{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	br := &BlockingReader{
		r:   &buf,
		ctx: ctx,
	}

	p := make([]byte, 1)
	var readErr error
	buf.WriteByte('a')
	// contents already exist in reader before calling Read
	_, readErr = br.Read(p)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(p))

	// contents get added to reader halfway through context lifetime
	func() {
		time.Sleep(500 * time.Millisecond)
		buf.WriteByte('b')
	}()
	_, readErr = br.Read(p)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(p))

	// contents do not get added to reader through context lifetime
	_, readErr = br.Read(p)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(p))
}
