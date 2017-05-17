package main

import (
	"io/ioutil"
        "net/http"
	"testing"
        "fmt"
)

func TestUnit(t *testing.T) {

  //resp, err := http.Get("http://HOSTIP:8080/test")
  resp, err := http.Get("http://192.168.225.121:8080/test")
  if err != nil {
    t.Error(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  fmt.Println("http get response was:", string(body))

  if err != nil {
    t.Error(err)
  } else if string(body) != "ok" {
    t.Error("Test Failed: expected reply of 'ok', got reply: ", string(body))
  }
}
