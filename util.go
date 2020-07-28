
package main

import (
   "log"
   "net/http"
)

func logError(msg string){
   log.Println("ERROR: "+msg)
}

// HTTP convenience functions 

func httpISE(w http.ResponseWriter){
   http.Error(w,"Internal server error", 500)
}

