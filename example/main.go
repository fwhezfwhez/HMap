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
	hm.Set("key","sub_key","value")

	//get the value with two keys
	v,er:=hm.Get("key","sub_key")
	if er!=nil{
		panic(er.Error())
	}
	fmt.Println(v)

	//get values array by main key
	vArr,er:=hm.GetByMainKey("key")
	if er!=nil{
		panic(er.Error())
	}
	fmt.Println(vArr)

	//delete values whose main key is "key"
	hm.DeleteByMainKey("key")
	//delete value whose main key is "key" and sub key is "sub_key"
	hm.Delete("key","sub_key")
	//delete all datas
	hm.DeleteAll()

	//the same
	hm.Clear()

	//smart format the datas view
	hm.Print()

	//about concurrently safe
	l:=sync.Mutex{}
	l.Lock()
	//do sth with hm
	l.Unlock()
}
