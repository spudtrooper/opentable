package api

import (
	"bytes"
	"fmt"
	"sort"
	"sync/atomic"

	"github.com/olekukonko/tablewriter"
)

type stats struct {
	intValues map[string]*int32
}

func newStats() *stats {
	return &stats{
		intValues: make(map[string]*int32, 0),
	}
}

func (s *stats) AddInt(key string, value int32) {
	if _, ok := s.intValues[key]; !ok {
		s.intValues[key] = new(int32)
	}
	atomic.AddInt32(s.intValues[key], value)
}

func (s *stats) IncInt(key string) { s.AddInt(key, 1) }

func (s *stats) Build() string {
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"Key", "Value"})
	var keys []string
	for key := range s.intValues {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		val := atomic.LoadInt32(s.intValues[key])
		table.Append([]string{key, fmt.Sprintf("%d", val)})
	}
	table.Render()
	return buf.String()
}
