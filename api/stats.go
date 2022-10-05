package api

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync/atomic"

	"github.com/fatih/color"

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

//go:generate genopts --function IncInt verbose
func (s *stats) IncInt(key string, optss ...IncIntOption) {
	opts := MakeIncIntOptions(optss...)
	s.AddInt(key, 1)
	if opts.Verbose() {
		var keys []string
		for key := range s.intValues {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		var strs []string
		var total int32
		for _, key := range keys {
			val := atomic.LoadInt32(s.intValues[key])
			total += val
			parts := strings.Split(key, ":")
			abbrev := parts[len(parts)-1]
			strs = append(strs, fmt.Sprintf("%s: %s ", abbrev, color.CyanString(fmt.Sprintf("%5d", val))))
		}
		strs = append(strs, fmt.Sprintf("%s: %s ", "total", color.CyanString(fmt.Sprintf("%5d", total))))
		log.Println(strings.Join(strs, " "))
	}
}

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
