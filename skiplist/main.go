// MIT License

// Copyright (c) [2018] [Yang Yu]

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.



//UNDER CONSTRUCTION


//Package Skiplist provide a probabilistic data structure with a expected time
//complexity of O()[complement]. The implementation of the data structure follows
//the pusedocode in Pugh's papers (see reference [1] and [2]).

//The data structure supports the operations including operations of insert, search, delete
//References
//[1] Pugh, William. "Skip lists: a probabilistic alternative to balanced trees."
//	  Communications of the ACM 33.6 (1990): 668-676.
//[2] Pugh, William. A skip list cookbook. 1998.

package main

// [IMPORTANT] Determine the API first, which will make it much easier to know which will be
// exported and which will not be
import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

type Node struct {
	Key     int
	Value   string
	Forward []*Node
}

type Skiplist struct {
	MaxLevel int
	Level    int
	P        float64
	Header   *Node
}

//MaxLvl is the maxlevel of the skiplist, which is calculated by Log_(1/p)^N,
//where p is the ratio of number of nodes in level i+1 and those in level i
func MaxLvl(P float64, N float64) int {
	return int(math.Ceil(math.Log(N) / math.Log(1/P)))
}

func newNode(Key, Level int, Value string) *Node {
	return &Node{Key, Value, make([]*Node, Level)}
}

func newSkiplist(MaxLevel, Level int, P float64, N *Node) *Skiplist {
	return &Skiplist{MaxLevel, Level, P, N}
}

//func (s *Skiplist) Search(key) {
//}

func (s *Skiplist) RandomLevel() int {
	lel := 1
	for rand.Float64() < s.P && lel < s.MaxLevel {
		lel++
	}
	return lel
}

func (s *Skiplist) Insert(key int, value string) bool {
	update := make([]*Node, s.MaxLevel)
	x := s.Header
	for i := s.Level; i >= 0; i-- {
		for x.Forward[i] != nil && x.Forward[i].Key < key {

			x = x.Forward[i]
		}
		update[i] = x

	}
	x = x.Forward[0]
	lvl := s.RandomLevel()
	fmt.Printf("randomlevel:%d\n", lvl)
	if x != nil && x.Key == key {
		x.Value = value

	} else {

		if lvl-1 > s.Level {
			for i := s.Level + 1; i < lvl; i++ {
				update[i] = x.Forward[i]
			}
			s.Level = lvl - 1
		}

		x = newNode(key, lvl, value)

		for i := 0; i < lvl; i++ {
			fmt.Println(i)
			x.Forward[i] = update[i].Forward[i]
			update[i].Forward[i] = x
		}
		//fmt.Printf("%#v\n", update)
	}
	return true
}

func main() {
	maxlvl := MaxLvl(.5, 100)
	header := newNode(1, maxlvl, "Patric")
	// fmt.Printf("%#v\n", header)

	fmt.Printf("maxlvl: %d\n", maxlvl)
	sl := newSkiplist(maxlvl, maxlvl-1, 0.5, header)
	fmt.Printf("current_lvl: %d\n", sl.Level)
	//fmt.Printf("%#v\n", sl)
	for i := 0; i < 20; i++ {
		sl.Insert(i, "xx")
	}

}
