package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
)

func main(){
	gateway, username, password := "https://my.smscentral.com.au/wrapper/sms", "username", "password"
	action, originator, recipient := "send", "shared", "0400xxxxxx"
	messagetext := "Hello from SMSCentral!"
	
	post_param := make(url.Values)
	post_param.Set("USERNAME", username)
	post_param.Set("PASSWORD", password)
	post_param.Set("ACTION", action)
	post_param.Set("ORIGINATOR", originator)
	post_param.Set("RECIPIENT", recipient)
	post_param.Set("MESSAGE_TEXT", messagetext)
	
	resp, err := http.PostForm(gateway, post_param)

	if err != nil{
		fmt.Printf("error: %s",err);
		return
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s",body)
}
