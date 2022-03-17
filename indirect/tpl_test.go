package main

import (
	"strconv"
	"testing"
	"text/template"
)

var benchTempl *template.Template
var benchData interface{}
var benchWriter MockWriter

func init() {
	var err error
	benchTempl, err = template.ParseFiles("./tpl.gohtml")
	if err != nil {
		panic(err)
	}

	type Data struct {
		Var1   int
		Var2   int
		Values []string
	}

	type Foo struct {
		Values *Data
		ID     int
	}

	type State struct {
		Foo *Foo
	}

	values := make([]string, 100)
	for i, _ := range values {
		values[i] = "Var" + strconv.Itoa(i+1)
	}

	benchData = State{
		Foo: &Foo{
			Values: &Data{
				Var1:   10,
				Var2:   5,
				Values: values,
			},
		},
	}
}

func BenchmarkIndirectTempl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchTempl.Execute(benchWriter, benchData)
	}
}

type MockWriter struct {
}

func (m MockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
