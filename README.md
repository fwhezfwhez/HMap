## HMap
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/fwhezfwhez/HMap)

HMap is suited for a map which has two key.

#### start

`go get github.com/fwhezfwhez/HMap`

#### Example:
```go
package main
import (
	"github.com/fwhezfwhez/HMap"
	"fmt"
	"sync"
)
func main() {
	//init a h-map
	hm := HMap.New()

	//set a value with two keys
	hm.Set("app_list", "cdd", "100.0.9.0")
	hm.Set("app_list", "cdd2", "100.0.9.1")
	hm.Set("app_list", "cdd3", "100.0.9.2")
	hm.Set("app_list2", "cdd4", "100.0.9.2")
	hm.Set("app_list3", "cdd5", "100.0.9.2")
	hm.Print()

	//get the value with two keys
	v,er:=hm.Get("app_list","cdd")
	if er!=nil{
		panic(er.Error())
	}
	fmt.Println(v)

	//get values array by main key
	vArr,er:=hm.GetByMainKey("app_list")
	if er!=nil{
		panic(er.Error())
	}
	fmt.Println(vArr)

	//delete values whose main key is "key"
	hm.DeleteByMainKey("app_list")
	hm.Print()
	//delete value whose main key is "key" and sub key is "sub_key"
	hm.Delete("app_list2","cdd4")
	hm.Print()
	//delete all datas
	hm.DeleteAll()
	hm.Print()
	//the same
	hm.Clear()

	//about concurrently safe
	l:=sync.Mutex{}
	l.Lock()
	//do sth with hm
	l.Unlock()
}

```