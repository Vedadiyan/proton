package functions

import (
	"strings"

	"github.com/vedadiyan/gocollections/pkg/queue"
)

func getKeys(key string) []string {
	key = fixString(key)
	return strings.Split(key, ".")
}

func fixString(key string) string {
	key = strings.TrimPrefix(key, "\"")
	key = strings.TrimSuffix(key, "\"")
	key = strings.TrimPrefix(key, "`")
	key = strings.TrimSuffix(key, "`")
	return key
}

func reduce[T any](key []string, data map[string]any, function func(current *T) bool) {
	q := queue.New[string]()
	for _, value := range key {
		q.Enqueue(value)
	}

	var ref any = data
	for {
		b := false
		d, ok := ref.(map[string]any)
		if !ok {
			break
		}
		if q.IsEmpty() {
			val := ref.(T)
			function(&val)
			return
		}
		for key, value := range d {
			p, err := q.Peek()
			if err != nil {
				panic(err)
			}
			if p == key {
				q.Dequeue()
				ref = value
				b = true
				break
			}
		}
		if !b {
			return
		}
	}
	if arr, ok := ref.([]any); ok {
		_key := make([]string, 0)
		for !q.IsEmpty() {
			v, err := q.Dequeue()
			if err != nil {
				panic(err)
			}
			_key = append(_key, v)
		}
		for _, item := range arr {
			if v, ok := item.(map[string]any); ok {
				reduce(_key, v, function)
				continue
			}
			if _, ok := item.([]any); ok {
				panic("two dimensional arrays are not supported")
			}
			if v, ok := item.(T); ok {
				cont := function(&v)
				if !cont {
					break
				}
				continue
			}
			panic("type mismatch")
		}
		return
	}
	if v, ok := ref.(T); ok {
		function(&v)
		return
	}
	if ref == nil {
		function(nil)
		return
	}
	panic("type mismatch")
}
