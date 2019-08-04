# quicktag
auto generate json/bson tag in custom field style. high performance, easy to use

```golang
package main

import (
    "encoding/json"
    "fmt"
    . "github.com/dovejb/quicktag"
    "reflect"
)

type Person struct {
    Name       string
    Age        int 
    MyChildren []Person
}

func main() {
    p := Person{
        Name: "dovejb",
        Age:  6,  
        MyChildren: []Person{
            Person{
                Name: "baby",
                Age:  3,  
            },  
        },  
    }   

    var p2 Person

    buf, _ := json.Marshal(Q(p))
    fmt.Println(string(buf))
    // {"name":"dovejb","age":6,"my_children":[{"name":"baby","age":3,"my_children":null}]}

    json.Unmarshal(buf, Q(&p2))
    fmt.Println(reflect.DeepEqual(p, p2))
    // true
}
```

```golang
import "github.com/dovejb/quicktag"
import "time"

func init() {
    // 自定义转换风格, 默认quicktag.PascalToUnderline, 无omitempty
    quicktag.StyleConvert = MyStyleConvertFunc // func(string) string
    // 自定义受影响业务tag, 默认 []string{"json","bson"}
    quicktag.TagNames = []string{"json", "bson"}
    // 自定义自引用最大层级, 默认5
    quicktag.MaxSelfRefLevel = 3
    
    // 注意！！！
    // 如果某类型自己包含了MarshalJSON/UnmarshalJSON等方法，如time.Time，请在字段后手动添加quicktag:"-"来跳过
    // 如
    data := struct {
        ID string `bson:"_id"`
        CreatedTime time.Time `quicktag:"-"`
    }
    
    // struct中原有的tag, 均会保留
}
```

```
root@dev:/w/try# go test github.com/dovejb/quicktag -bench=.
goos: linux
goarch: amd64
pkg: github.com/dovejb/quicktag
BenchmarkQMarshal-4              1000000              1343 ns/op
BenchmarkJsonMarshal-4           1000000              1565 ns/op
PASS
ok      github.com/dovejb/quicktag      3.936s
root@dev:/w/try# go test github.com/dovejb/quicktag -bench=.
goos: linux
goarch: amd64
pkg: github.com/dovejb/quicktag
BenchmarkQMarshal-4              1000000              1635 ns/op
BenchmarkJsonMarshal-4           1000000              2024 ns/op
PASS
ok      github.com/dovejb/quicktag      3.714s
```
