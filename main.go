
package main
import (
   "net/http"
   "log"
)

func main(){
   http.HandleFunc("/update/",update)
   http.HandleFunc("/",show)
   if err := http.ListenAndServe(":8080",nil); err != nil {
      log.Fatal(err)
   }
}
