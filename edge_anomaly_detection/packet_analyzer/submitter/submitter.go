package submitter

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Submit_to_detector(features [17]float64)(int){
	var submission_data [17] string
	for i:=0; i<len(features);i++{
		submission_data[i]=strconv.FormatFloat(features[i],'f',-1,32)
	}
	submission:=strings.Join(submission_data[:],",")
	var a string="["
	var b string="]"
	submission=a+submission+b
	resp, err:=http.PostForm("http://172.18.0.5:30033/check",url.Values{"features":{submission}})
	//resp, err=http.Get("http://172.18.0.4:30033")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,_:=io.ReadAll(resp.Body)
	print(string(body))
	i,_:=strconv.Atoi(strings.Trim(string(body),"[]"))
	return i
}
