package HashMap

import (
	"testing"
	"fmt"
)

var hm *HMap

func Init() {
	hm = New()
	hm.Set("app_list", "cdd", "100.0.9.0")
	hm.Set("app_list", "cdd2", "100.0.9.1")
	hm.Set("app_list", "cdd3", "100.0.9.2")
	hm.Set("app_list2", "cdd4", "100.0.9.2")
	hm.Set("app_list3", "cdd5", "100.0.9.2")
}
func TestNew(t *testing.T) {
	Init()
	hm.Print()
}

func TestHMap_Set(t *testing.T) {
	Init()
	hm.Set("app_list", "cdd", "100.0.9.0")
	hm.Set("app_list", "cdd2", "100.0.9.11")
	hm.Set("app_list2", "cdd2", "100.0.9.11")
	hm.Print()
}

func TestHMap_Get(t *testing.T) {
	Init()
	t.Log(hm.Get("app_list", "cdd3"))
	//fmt.Println(hm)
}

func TestHMap_GetByMainKey(t *testing.T) {
	Init()
	rs, er := hm.GetByMainKey("app_list")
	if er != nil {
		t.Fatal(er.Error())
	}
	t.Log(rs)
}

func TestHMap_Print(t *testing.T) {
	Init()
	hm.Print()
}

func TestHMap_Delete(t *testing.T) {
	Init()
	fmt.Println("before delete")
	hm.Print()
	hm.Delete("app_list","cdd3")
	fmt.Println("after delete")
	hm.Print()
}

func TestHMap_DeleteByMainKey(t *testing.T) {
	Init()
	fmt.Println("before delete")
	hm.Print()
	hm.DeleteByMainKey("app_list")
	fmt.Println("after delete")
	hm.Print()
}

func TestHMap_DeleteAll(t *testing.T) {
	Init()
	fmt.Println("before delete")
	hm.Print()
	hm.DeleteAll()
	fmt.Println("after delete")
	hm.Print()
}