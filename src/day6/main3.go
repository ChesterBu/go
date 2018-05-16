package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:"company"`
	Lang []string `json:"-"` //不输出
}
type IT2 struct {
	Company string
	Lang []string
}

func main()  {

	//json化
	s:=IT{"alibaba",[]string{"java","js","python"}}
	if buf,err:=json.Marshal(s);err==nil{
		fmt.Println(string(buf))
	}
	m:=make(map[string]interface{})
	m["company"]="Tencent"
	m["lang"]=[]string{"c++","js","c"}
	if res,err:=json.Marshal(m);err==nil{
		fmt.Println(string(res))
	}

	//json解析
	var tmp IT2
	jsonBuf := `{"company":"bytedance","lang":["go","python","js"]}`
	if err:=json.Unmarshal([]byte(jsonBuf),&tmp);err ==nil{
		fmt.Printf("%+v\n",tmp)
	}
	n:=make(map[string]interface{})
	if err:=json.Unmarshal([]byte(jsonBuf),&n);err ==nil{
		fmt.Printf("%+v\n",n)
	}
}

