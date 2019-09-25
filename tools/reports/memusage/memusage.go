package main

import (
	"fmt"

	"github.com/openacid/low/size"
	"github.com/openacid/slim/benchhelper"
	"github.com/openacid/slim/encode"
	"github.com/openacid/slim/trie"
)

func main() {
	compareTrieMapMemUse()
}

func compareTrieMapMemUse() {

	writeTableHeader()

	for _, n := range []int{1000, 2000, 5000} {
		for _, k := range []int{64, 256, 1024} {

			trieSize := getTrieMem(n, k)
			mapSize := getMapMem(n, k)
			// make key + value as a value in trie
			// kvTrieSize := getKVTrieMem(cnt, l)
			kvTrieSize := getKVTrieMem2(n, k)

			mapAvg := float64(mapSize) / float64(n)
			trieAvg := float64(trieSize) / float64(n)
			kvTrieAvg := float64(kvTrieSize) / float64(n)

			writeTableRow(n, k, 2, trieAvg, mapAvg, kvTrieAvg)
		}
	}
}

func writeTableHeader() {

	fmt.Printf("| %s | %s | %s | %s | %s | %s |\n",
		"Key Count", "Key Length", "Value Size", "Trie Size (Byte/key)", "Map Size (Byte/key)",
		"KV Trie Size (Byte/key)")

	fmt.Printf("| --- | --- | --- | --- | --- | --- |\n")
}

func writeTableRow(cnt, kLen, vLen int, trieAvg, mapAvg, kvTrieAvg float64) {

	fmt.Printf("| %5d | %5d | %5d | %6.1f | %6.1f | %6.1f |\n",
		cnt, kLen, vLen, trieAvg, mapAvg, kvTrieAvg)
}

func getTrieMem(keyCnt, keyLen int) int64 {

	keys := benchhelper.RandSortedStrings(keyCnt, keyLen, nil)
	vals := make([]uint16, keyCnt)
	for i := range vals {
		vals[i] = uint16(i)
	}

	t, err := trie.NewSlimTrie(encode.U16{}, keys, vals)
	if err != nil {
		panic(err)
	}

	return int64(size.Of(t))
}

func getKVTrieMem2(keyCnt, keyLen int) int64 {
	// make key + value as a value in trie

	keys := benchhelper.RandSortedStrings(keyCnt, keyLen, nil)
	indexes := make([]uint32, keyCnt)
	for i := 0; i < len(keys); i++ {
		indexes[i] = uint32(i)
	}

	t, err := trie.NewSlimTrie(encode.U32{}, keys, indexes)
	if err != nil {
		panic(err)
	}

	return int64(size.Of(t))
}

func getMapMem(keyCnt, keyLen int) int64 {

	keys := benchhelper.RandSortedStrings(keyCnt, keyLen, nil)
	vals := make([]uint16, keyCnt)

	m := make(map[string]uint16, len(keys))

	for i := 0; i < len(keys); i++ {
		m[keys[i]] = vals[i]
	}

	return int64(size.Of(m))
}
