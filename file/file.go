package file


import (
	"io"
)

type Reader interface {
	Read(p []byte)(n int, err error)
}

type Write interface {
	Write(r []byte)(n int, err error)
}

type NotHowReaderIsdefined interface{
	Read()(n []byte, err error)
}

func CountLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}

	}
}


