
package main
import (
   "github.com/jinzhu/gorm"
   "net/http"
   "log"
   "fmt"
)

func main(){
   db,err := gorm.Open("sqlite3","db.sqlite")
   if err != nil {
      panic("Could not open DB!" + fmt.Sprintf("%v",err))
   }
   defer db.Close()

   db.AutoMigrate(&Variable{})
   db.AutoMigrate(&Set{})
   db.AutoMigrate(&Theme{})

   http.HandleFunc("/update/",update)
   http.HandleFunc("/",show)
   if err := http.ListenAndServe(":8080",nil); err != nil {
      log.Fatal(err)
   }
}
