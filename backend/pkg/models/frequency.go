package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type Frequency struct {
	Freq map[string]int32
}

type Pair struct {
	Char  string `json:"char"`
	Count int32  `json:"count"`

	left  *Pair
	right *Pair
}

func NewFrequency() *Frequency {
	ret := &Frequency{}
	ret.init()

	return ret
}

func (f *Frequency) init() {
	f.Freq = make(map[string]int32)
}

func (f *Frequency) AddCount(char string) {
	count := f.Freq[char]

	count = count + 1
	f.Freq[char] = count
}

func (f *Frequency) GetCount(char string) int32 {
	return f.Freq[char]
}

func (f *Frequency) TotalCount() (total int32) {

	for _, count := range f.Freq {
		total = total + count
	}

	return
}

func (f *Frequency) Sort() (ret []Pair) {
	var first *Pair
	startTime := time.Now()
	if len(f.Freq) > 0 {
		first = nil

		for char, count := range f.Freq {

			newPair := &Pair{
				Char:  char,
				Count: count,
				left:  nil,
				right: nil,
			}

			if first == nil {
				first = newPair
			} else {
				endWhile := false
				actual := first

				for !endWhile {
					if actual.Count == count {
						newPair.left = actual.left
						actual.left = newPair
						newPair.right = actual
						if newPair.left == nil {
							first = newPair
						} else {
							newPair.left.right = newPair
						}

						endWhile = true
					} else if actual.Count > count {
						if actual.right != nil && actual.right.Count > count {
							actual = actual.right
						} else {
							newPair.right = actual.right
							actual.right = newPair
							newPair.left = actual
							endWhile = true
						}
					} else {
						if actual.left != nil {
							actual = actual.left
						} else {
							newPair.left = actual.left
							actual.left = newPair
							newPair.right = actual
							if newPair.left == nil {
								first = newPair
							} else {
								newPair.left.right = newPair
							}
							endWhile = true
						}
					}
				}
			}
		}
	}

	actual := first
	if actual != nil {
		ret = append(ret, *actual)
		for actual.right != nil {
			actual = actual.right
			ret = append(ret, *actual)
		}
	}

	log.Info(time.Now().Sub(startTime))

	return
}

func (f *Frequency) printListHelper(lista *Pair) {
	if lista != nil {
		texto := fmt.Sprintf("%v", lista)
		primero := lista
		for primero.right != nil {
			primero = primero.right
			texto = fmt.Sprintf("%s -- %v", texto, primero)

		}
		log.Info(texto)
	} else {
		log.Info("vacio")
	}
}
