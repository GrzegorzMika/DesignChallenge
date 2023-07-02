package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestCountFruits(t *testing.T) {
	fruits := []string{
		"apple",
		"banana",
		"apple",
		"cherry",
		"banana",
		"cherry",
		"apple",
		"apple",
		"cherry",
		"banana",
		"cherry",
	}
	count := CountFruits(fruits)
	expected := map[string]int{"apple": 4, "banana": 3, "cherry": 4}
	if eq := reflect.DeepEqual(expected, count); !eq {
		t.Errorf("expected %v, got %v", expected, count)
	}
}

func TestCountFruitsEmpty(t *testing.T) {
	fruits := []string{}
	count := CountFruits(fruits)
	expected := map[string]int{}
	if eq := reflect.DeepEqual(expected, count); !eq {
		t.Errorf("expected %v, got %v", expected, count)
	}
}

func BenchmarkCountFruits(b *testing.B) {
	fruits := []string{"banana", "apple", "cherry"}
	sliceLength := 1000000
	allFruits := make([]string, sliceLength)
	for i := 0; i < sliceLength; i++ {
		allFruits[i] = fruits[rand.Intn(len(fruits))]
	}
	for n := 0; n < b.N; n++ {
		CountFruits(allFruits)
	}
}
