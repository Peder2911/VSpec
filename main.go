
package main

import (
   "net/http"
   "log"
)

func main(){
   http.HandleFunc("/",update)
   if err := http.ListenAndServe(":8080",nil); err != nil {
      log.Fatal(err)
   }
}

