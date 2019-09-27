package http_learing

import (
	"testing"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func TestHttpServer(t *testing.T) {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/projects", projects)
	http.HandleFunc("/services", services)
	http.HandleFunc("/developers", developers)
	http.HandleFunc("/list", list)
	http.HandleFunc("/rmserver/get-history-records", historyRecord)
	http.ListenAndServe(":9097", nil)

}

func respHeader(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
}

func echo(w http.ResponseWriter, r *http.Request) {
	reqBody, reqErr := ioutil.ReadAll(r.Body)
	if reqErr != nil {
		fmt.Println(reqErr)
	}

	fmt.Println(string(reqBody))

	s := `{"name":"网站","num":3,"sites":[{"name":"Google","info":["Android","Google搜索","Google翻译"]},{"name":"Runoob","info":["菜鸟教程","菜鸟工具","菜鸟微信"]},{"name":"Taobao","info":["淘宝","网购"]}]}`
	respHeader(w)
	writeLen, wErr := w.Write([]byte(s))
	if wErr != nil || writeLen != len(s) {
		fmt.Println(wErr)
	}
}

func projects(w http.ResponseWriter, r *http.Request) {
	respHeader(w)
	write(w, `[{"key":"xb","name":"小白"},{"key":"he","name":"琥珀"}]`)
}

func services(w http.ResponseWriter, r *http.Request) {
	respHeader(w)
	write(w, `[{"key":"connector","name":"接入层"},{"key":"yexiu","name":"叶修"}]`)
}

func developers(w http.ResponseWriter, r *http.Request) {
	respHeader(w)
	write(w, `[{"id":"1","name":"Yelo"},{"id":"2","name":"gengyangyang"}]`)
}

func list(w http.ResponseWriter, r *http.Request) {
	respHeader(w)
	write(w, `{"total:10,"list":[{"id":"1","name":"Yelo"},{"id":"2","name":"gengyangyang"}]}`)
}

func historyRecord(w http.ResponseWriter, r *http.Request) {
	respHeader(w)
	defer r.Body.Close()
	reqBodyByte, _ := ioutil.ReadAll(r.Body)
	type reqBodyStruct struct {
		QueryType string `json:"query_type"`
		Count     int    `json:"count"`
		CurPage   int    `json:"cur_page"`
	}

	var reqBody reqBodyStruct
	json.Unmarshal(reqBodyByte, &reqBody)
	type Item struct {
		RecordId int    `json:"record_id"`
		Content  string `json:"content"`
	}
	type respBodyStruct struct {
		Total int    `json:"total"`
		List  []Item `json:"list"`
	}
	respBody := respBodyStruct{Total: 100}
	var list []Item
	for i := 0; i < reqBody.Count; i++ {
		idx := (reqBody.CurPage - 1) * reqBody.Count + i
		list = append(list, Item{RecordId: idx, Content: fmt.Sprintf(reqBody.QueryType+"第%d个", idx+1)})
	}
	respBody.List = list
	respBodyBytes, _ := json.Marshal(respBody)

	write(w, string(respBodyBytes))
}

func write(w http.ResponseWriter, respBody string) (int, error) {
	return w.Write([]byte(respBody))
}
