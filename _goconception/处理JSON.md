struct ==> json
[]byte=json.Marshal(struct)
fmt.Println(string([]byte))

json ==> struct
json.Unmarshal([]byte(jsonStr), &rsp)

```
package learngo

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 有用的网站 在线json转struct
// http://json2struct.mervine.net/

// 很重要 调用json.Marshal 结构体中的属性需要大写
type Book struct {
	BookName string `json:"book_name"`
	BookId   string `json:"book_id"`
}

type Request struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Books []Book `json:"books"`
}

func TestJson(t *testing.T) {
	//  struct ===> json
	books := []Book{{"book1", "book_id_1"}, {"book2", "book_id_2"}}
	request := Request{
		Name:  "a book",
		Id:    0,
		Books: books,
	}
	// Marshal返回的是[]byte, 需要转为string(xxx)
	bookBytes, err := json.Marshal(request)
	fmt.Println(bookBytes)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(bookBytes))

	// json ===> struct
	var jsonStr = `{"name":"a book","id":0,"books":[{"book_name":"book1","book_id":"book_id_1"},{"book_name":"book2","book_id":"book_id_2"}]}`
	var rsp Request
	json.Unmarshal([]byte(jsonStr), &rsp)
	fmt.Println(rsp)

	// map ===> json
	var mapInstance []map[string]interface{}
	instance1 := map[string]interface{}{"a": 1, "b": 2}
	instance2 := map[string]interface{}{"c": 1, "d": 2}
	mapInstance = append(mapInstance, instance1, instance2)
	mapInstanceBytes, err := json.Marshal(mapInstance)
	if err != nil {
		return
	}
	fmt.Println(string(mapInstanceBytes))
	// json ===> map
	jsonStr = `[{"a":1,"b":2},{"c":1,"d":2}]`
	err = json.Unmarshal([]byte(jsonStr), &mapInstance)
	if err != nil {
		return
	}
	fmt.Println(mapInstance)
}
```
