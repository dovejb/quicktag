package quicktag

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPascalToUnderline(t *testing.T) {
	s := "MyGPUIsTheBest"
	if PascalToUnderline(s) != "my_gpu_is_the_best" {
		t.Error("error PascalToUnderline")
	}
}

type Person struct {
	Name       string
	Age        int
	MyChildren []Person
}

var p = Person{
	Name: "dovejb",
	Age:  6,
	MyChildren: []Person{
		Person{
			Name: "baby",
			Age:  3,
		},
	},
}

var pbuf = `{"name":"dovejb","age":6,"my_children":[{"name":"baby","age":3,"my_children":null}]}`

func TestMarshal(t *testing.T) {
	buf, _ := json.Marshal(Q(p))
	if string(buf) != pbuf {
		t.Error("error Marshal actual " + string(buf))
	}
}

func TestUnmarshal(t *testing.T) {
	var p2 Person
	json.Unmarshal([]byte(pbuf), Q(&p2))
	if !reflect.DeepEqual(p, p2) {
		t.Error("error Unmarshal")
	}
}

func BenchmarkQMarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(Q(p))
	}
}

func BenchmarkJsonMarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(p)
	}
}
