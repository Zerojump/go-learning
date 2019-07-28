package reflect_learing

import (
	"fmt"
	"testing"
	"reflect"
	"encoding/json"
)

var (
	d         D
	dType     = reflect.TypeOf(d)
	aImpl     AImpl
	aImplType = reflect.TypeOf(aImpl)

	aIntf     AIntf = aImpl
	aIntfType       = reflect.TypeOf(aIntf)

	bIntf     BIntf = aImpl
	bIntfType       = reflect.TypeOf(bIntf)
)

func TestLocalAlign(t *testing.T) {

}

func TestLocalFieldAlign(t *testing.T) {

}

func TestType(t *testing.T) {
	t.Logf("aImpl = %v", aImpl)
	t.Logf("aImplType = %v", aImplType)
	t.Logf("aImplType.Name() = %v", aImplType.Name())
	t.Logf("aImplType.PkgPath() = %v", aImplType.PkgPath())
	t.Logf("aImplType.Size() = %v", aImplType.Size())
	t.Logf("aImplType.String() = %v", aImplType.String())
	t.Log()
}

func TestLocalMethod(t *testing.T) {
	t.Log("aImplType.NumMethod() = ", aImplType.NumMethod())
	for i := 0; i < aImplType.NumMethod(); i++ {
		t.Logf("aImplType.Method(%d) : %v", i, aImplType.Method(i))
		methodName := aImplType.Method(i).Name
		t.Logf("aImplType.%s.Func = %v", methodName, aImplType.Method(i).Func)
		t.Logf("aImplType.%s.Index = %v", methodName, aImplType.Method(i).Index)
		t.Logf("aImplType.%s.PkgPath = %v", methodName, aImplType.Method(i).PkgPath)
		t.Logf("aImplType.%s.Type = %v", methodName, aImplType.Method(i).Type)
	}

	t.Log()

	t.Logf("aIntfType = %v", aIntfType)
	t.Log("aIntfType.NumMethod() = ", aIntfType.NumMethod())
	for i := 0; i < aIntfType.NumMethod(); i++ {
		t.Logf("aIntfType.Method(%d) : %v", i, aIntfType.Method(i))
		methodName := aIntfType.Method(i).Name
		t.Logf("aIntfType.%s.Func = %v", methodName, aIntfType.Method(i).Func)
		t.Logf("aIntfType.%s.Index = %v", methodName, aIntfType.Method(i).Index)
		t.Logf("aIntfType.%s.PkgPath = %v", methodName, aIntfType.Method(i).PkgPath)
		t.Logf("aIntfType.%s.Type = %v", methodName, aIntfType.Method(i).Type)
	}
	t.Log()
}

func TestLocalMethodByName(t *testing.T) {
	method, ok := aImplType.MethodByName("PublicSingleFunc")
	t.Logf("method = %v", method)
	t.Logf("ok = %v", ok)

	methodName := method.Name
	t.Logf("aIntfType.%s.Func = %v", methodName, method.Func)
	t.Logf("aIntfType.%s.Index = %v", methodName, method.Index)
	t.Logf("aIntfType.%s.PkgPath = %v", methodName, method.PkgPath)
	t.Logf("aIntfType.%s.Type = %v", methodName, method.Type)

	t.Log()
}

func TestLocalKind(t *testing.T) {
	t.Logf("aImplType.Kind() = %v", aImplType.Kind())
	t.Logf("aIntfType.Kind() = %v", aIntfType.Kind())
	t.Logf("reflect.ValueOf(aImpl).Kind() = %v", reflect.ValueOf(aImplType).Kind())
	t.Logf("reflect.ValueOf(aIntf).Kind() = %v", reflect.ValueOf(aIntf).Kind())
	t.Logf("aImplType.Field(5).Type.Kind() = %v", aImplType.Field(5).Type.Kind())

	t.Log()
}

func TestLocalImplements(t *testing.T) {

	//t.Logf("aImplType.Implements(aIntfType) = %v", aImplType.Implements(reflect.ValueOf(aIntf)))
	//t.Logf("bIntfType.Implements(aIntfType) = %v", bIntfType.Implements(aIntfType))
	//t.Logf("aImplType.Implements(dType) = %v", aImplType.Implements(dType))

	//aImpl.L = B{}
	//aImpl.M = D{}
	t.Logf("aImplType.Field(3).Type = %v", aImplType.Field(3).Type)
	t.Logf("aImplType.Field(6).Type = %v", aImplType.Field(6).Type)
	t.Log(aImplType.Field(3).Type.Implements(aImplType.Field(6).Type))

	t.Log()
}

func TestLocalAssignableTo(t *testing.T) {
	t.Logf("aImplType.AssignableTo(aIntfType) = %v", aImplType.AssignableTo(aIntfType))
	t.Logf("bIntfType.AssignableTo(aIntfType) = %v", bIntfType.AssignableTo(aIntfType))
	t.Logf("aImplType.AssignableTo(bIntfType) = %v", aImplType.AssignableTo(bIntfType))
	t.Logf("aImplType.AssignableTo(dType) = %v", aImplType.AssignableTo(dType))
	t.Log()
}

func TestLocalConvertibleTo(t *testing.T) {

}

func TestLocalComparable(t *testing.T) {

}

func TestLocalBits(t *testing.T) {

}

func TestLocalChanDir(t *testing.T) {

}

func TestLocalIsVariadic(t *testing.T) {

}

func TestLocalElem(t *testing.T) {
	t.Logf("reflect.TypeOf([]int{}).Elem() = %v", reflect.TypeOf([]int{}).Elem())
	t.Log()
}

func TestLocalFieldByIndex(t *testing.T) {
	t.Logf("aImplType.FieldByIndex([]int{3,0}) = %v", aImplType.FieldByIndex([]int{3, 0,0}))
	t.Log()
}

func TestLocalFieldByName(t *testing.T) {
	t.Logf("aImplType.Method(0) = %v", aImplType.Method(0))
	t.Logf("aImplType.Method(0).Type.In(0) = %v", aImplType.Method(0).Type.In(3))
	t.Logf("aImplType.Method(0).Type.Out(0) = %v", aImplType.Method(0).Type.Out(2))
	t.Logf("aImplType.Method(0).Type.NumIn() = %v", aImplType.Method(0).Type.NumIn())
	t.Logf("aImplType.Method(0).Type.NumOut() = %v", aImplType.Method(0).Type.NumOut())
}

func TestLocalFieldByNameFunc(t *testing.T) {

}

func TestLocalIn(t *testing.T) {

}

func TestLocalKey(t *testing.T) {

}

func TestLocalLen(t *testing.T) {

}

func TestLocalNumField(t *testing.T) {

}

func TestLocalNumIn(t *testing.T) {

}

func TestLocalNumOut(t *testing.T) {

}

func TestLocalOut(t *testing.T) {

}

type AIntf interface {
	singleFunc()
}

type BIntf interface {
	AIntf
	PublicSingleFunc()
	ParamFunc1(a, b int, c string) (h, k int, l bool, e error)
}

type B struct {
	C
}

type C struct {
	G string `json:"g"`
}

type D struct {
	H string `json:"h"`
}

type AImpl struct {
	a  int                                 `json:"a"`
	B  string                              `json:"b"`
	F  func(c int, d string) (bool, error) `json:"-"`
	AB B                                   `json:"ab"`
	D
	L interface{}
	M interface{}
}

func (AImpl) singleFunc() {
	fmt.Println("AImpl exec singleFunc()")
}

func (AImpl) PublicSingleFunc() {
	fmt.Println("AImpl exec PublicSingleFunc()")
}

func (AImpl) ParamFunc1(a, b int, c string) (h, k int, l bool, e error) {
	fmt.Println("AImple exec ParamFunc1()")
	return
}

func (AImpl) ParamFunc2(i interface{}) {
	fmt.Println("AImpl exec ParamFunc2()")
}

func (a AImpl) String() string {
	if bytes, e := json.Marshal(a); e == nil {
		return string(bytes)
	} else {
		return e.Error()
	}
}
