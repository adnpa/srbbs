package test

import (
	"log"
	"srbbs/src/util/lib/algo"
	"testing"
)

//func TestReddit(t *testing.T) {
//	log.Println(lib.Hot(400, 100, time.Now()))
//}

func TestRedditComment(t *testing.T) {
	log.Println(algo.HotComment(0, 1000))
	log.Println(algo.HotComment(400, 100))
	log.Println(algo.HotComment(1000, 0))
}
