package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const HashMultiplier = 31

type HashTable struct {
	Buckets [][]HashPair
}

func NewHashTable(bucketCount int) *HashTable {
	ht := &HashTable{}
	ht.Buckets = make([][]HashPair, bucketCount)
	for i, _ := range ht.Buckets {
		ht.Buckets[i] = make([]HashPair, 0)
	}

	return ht
}

func (ht *HashTable) Set(key, value string) {
	bucketNum := hash(key) % len(ht.Buckets)
	pair := HashPair{Key: key, Value: value}
	ht.Buckets[bucketNum] = append(ht.Buckets[bucketNum], pair)
}

func (ht *HashTable) Get(key string) string {
	bucketNum := hash(key) % len(ht.Buckets)
	for _, pair := range ht.Buckets[bucketNum] {
		if pair.Key == key {
			return pair.Value
		}
	}

	return ""
}

func (ht *HashTable) String() string {
	var buf bytes.Buffer

	for i, b := range ht.Buckets {
		fmt.Fprintln(&buf, "Bucket", i)
		fmt.Fprintln(&buf, "--------")
		for _, s := range b {
			fmt.Fprintln(&buf, s)
		}

		fmt.Fprintln(&buf)
	}

	return buf.String()
}

func hash(s string) (h int) {
	for _, b := range s {
		h = HashMultiplier*h + int(b)
	}

	return h
}

type HashPair struct {
	Key   string
	Value string
}

func (p HashPair) String() string {
	return p.Key + " => " + p.Value
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("How many buckets in hash table? ")
	var n int
	fmt.Scanln(&n)

	ht := NewHashTable(n)

	for {
		if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}

		words := strings.Split(string(line[:len(line)-1]), " ")
		switch words[0] {
		case "print":
			fmt.Println(ht)
		case "set":
			ht.Set(words[1], words[2])
		case "get":
			fmt.Println(ht.Get(words[1]))
		default:
			fmt.Println("Don't know how to:", words[0])
		}
	}
}
