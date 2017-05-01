package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func getter (url string,size chan string){
	length, err := getPage(url)
	if(err == nil ){
		size <- fmt.Sprintf("Length of url : %s is :%d",url,length)
	}

}

func getPage(url string) (int,error) {
	resp, err := http.Get(url)
	if(err != nil ){
		fmt.Printf("There was an error : %s",err)
		return 0,err
	}
	
	for k,v := range resp.Header {
		fmt.Printf("Headers are %s,%s\n",k,v)
	}
	body,err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if(err != nil ){
		fmt.Printf("Error reading body : %s\n",err)
		return 0,err
	}
		return len(body),nil
}

func main(){
	urls := [ ]string{"http://www.yahoo.com","http://www.google.com","http://www.bing.com","http://bbc.co.uk","https://en.wikipedia.org"}
	size := make(chan string)
	for _,url := range urls {
		go getter(url,size)
	}
	for i:=0; i< len(urls); i++ {
		fmt.Printf("%s\n", <- size)
	}
}	
