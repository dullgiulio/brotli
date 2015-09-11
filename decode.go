package brotli

import "io"

type Decoder struct {
	reader io.Reader
}

func NewDecorder(r io.Reader) *Decoder {
	return &Decoder{reader: r}
}

func (d *Decoder) Decode(buf []byte) error {
	var i int
	// readWindowSide(buf[i])
	for {
		last := isLast(buf[i])
		if last {
			if isLastEmpty(buf[i]) {
				return nil
			}
		}
		nNibbles := readNibbles()
		if nNibbles == 0 {
			// verify reserved bit is zero
			skip := readSkip()
			i = i + skip
			continue
		} else {
			mlen = readMLen()
		}
		if !last {
			if isUncompressed(buf[i]) {
				// copy(buf[i+1:i+mlen+1]) to writer
				i = i + mlen + 2
				continue
			}
		}
		nblTypes := make([]byte, 3)
		blen := make([]int, 3)
		// types L, I, D
		for t := 0; t < 3; t++ {
			btype[t] := readNblTypes(buf[i])
			if btype[t] >= 2 {
				// read prefix (Huffman) code for block types, HTREE_BTYPE[i]
				// read prefix (Huffman) code for block counts, HTREE_BLEN[i]
				blen[t] := blockCount(buf[i])
				btype[t] = 0
				htreeBtype[len(htreeBtype)-1] = 0
				htreeBtype[len(htreeBtype)] = 1
				continue
			}
			blen[t] = 268435456
			btype[t] = 0
		}
		npostfix := readNPostfix(buf[i])
		ndirect := readNDirect(buf[i])
		// read array of literal context modes, CMODE[]
		ntreesL := readNTrees(buf[i])
		cmapl := make([]byte, ntreesL)
		if ntreesL >= 2 {
			// read literal context map, CMAPL[]
		}
		ntreesD := readNTrees(buf[i])
		cmapd := make([]byte, ntreesD)
		if ntreesD >= 2 {
			// read distance context map, CMAPD[]
		}
		// read array of prefix codes for literals, HTREEL[]
		// read array of prefix codes for insert-and-copy, HTREEI[]
		// read array of prefix codes for distances, HTREED[]
		for ; i < blockLen; i++ {
		}
	}
}
