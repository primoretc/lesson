package main
import (
    "fmt"
    "net/http"
)
func main() {
    resp, err := http.Get("https://google.com") 
    if err != nil { 
        fmt.Println(err) 
        return
    } 
    defer resp.Body.Close()
    for true {
             
        bs := make([]byte, 1014)
        n, err := resp.Body.Read(bs)
        fmt.Println(string(bs[:n]))
         
        if n == 0 || err != nil{
            break
        }
    }
}