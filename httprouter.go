/*@Author : Manasvini Banavara Suryanarayana
*SJSU ID : 010102040
*CMPE 273 Lab#2
*/
package main

import (
    "fmt"
    "./httprouter"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type Resp struct {
   Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func hello2(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
   
   //creating struct to read request json 
   type ReqBody struct {
        Name string `json:"name"`
    }  
    var x ReqBody

    //fetch the body from request 
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
       fmt.Println("error occured ")
   }
 
    //converting json body to struct of type ReqBody
    err1 := json.Unmarshal(body, &x)
    if err1 != nil {
       fmt.Println(err1)
   }
  
    m := x.Name
    //constructing struct for sending back response body
    test := Resp{
        Greeting: "Hello, "+m+"!",
    }
    //converting response body struct to json format
   respjson, err2 := json.Marshal(test)
   if err2 != nil {
        fmt.Println("error occured 2")
    }
     
    rw.Header().Set("Content-Type","application/json")
    rw.WriteHeader(200)
    //sending back response
    fmt.Fprintf(rw, "%s", respjson)
     
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", hello2)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }

    server.ListenAndServe()
}