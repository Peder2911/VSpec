
package main

import (
   "github.com/go-yaml/yaml"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/sqlite"
   _ "github.com/mattn/go-sqlite3"
   "net/http"
   "html/template"
   "fmt"
   "log"
   "bytes"
   "io"
)

func update(w http.ResponseWriter, r *http.Request){
   switch r.Method {

      case "GET":
         templateFiles := []string {
            "./templates/update.html",
            "./templates/base.html",
         }

         ts, err := template.ParseFiles(templateFiles...)

         if err != nil {
            logError("Failed to parse template(s)"+fmt.Sprintf("%v",err))
            httpISE(w)
            return
         }

         err = ts.Execute(w,nil)

         if err != nil {
            logError("Failed to execute template(s)"+fmt.Sprintf("%v",err))
            httpISE(w)
            return
         }
         break

      case "POST":
         s := ModelSpec{}

         err := r.ParseForm()

         if err != nil {
            logError("Failed to parse form")
            return
         }

         file, handler, err := r.FormFile("data")

         if err != nil {
            logError(fmt.Sprintf("%v",err))
            http.Error(w,"Error processing form",400)
            return
         }

         defer file.Close()

         buf := bytes.NewBuffer(nil)
         n,err := io.Copy(buf,file)

         if err != nil {
            logError("Failed to read file")
            http.Error(w,"Error processing data", 422)
            return
         } else {
            fname := handler.Header["Content-Disposition"][0]
            log.Printf("Read %v bytes from %s", n, fname)
         }

         yaml.Unmarshal(buf.Bytes(),&s)

         if err != nil {
            logError("Failed to parse yaml")
            http.Error(w,"Error processing data", 422)
            return
         }

         // Populate DB

         db, err := gorm.Open("sqlite3", "db.sqlite")
         if err != nil {
            logError("Failed to open DB")
            httpISE(w)
         }
         db.Exec("DROP TABLE variables;")
         db.Exec("DROP TABLE themes;")
         db.Exec("DROP TABLE sets;")

         db.AutoMigrate(&Variable{},&Set{},&Theme{})

         tx := db.Begin()

         // Add variables, along with their variable sets 
         for set ,variables := range s.Colsets {
            setObj := Set{Name: set, Variables: make([]*Variable,len(variables))}
            for idx,v := range variables {
               varObj := Variable{Name: v}
               setObj.Variables[idx] = &varObj
               // Do not duplicate variables
               tx.FirstOrCreate(&varObj,Variable{Name:v})
            }
            tx.Create(&setObj)
         }

         // Create themes; collections of variable sets
         for theme, sets := range s.Themes {
            themeObj := Theme{Name: theme, Sets: make([]*Set,len(sets))}
            for idx,set := range sets {
               setObj := Set{}
               tx.FirstOrCreate(&setObj,Set{Name: set})
               themeObj.Sets[idx] = &setObj
            }
            tx.Create(&themeObj)
         }

         tx.Commit()
         db.Close()

         fmt.Fprintf(w,"Success!")
         break
   }
}

func show(w http.ResponseWriter, r *http.Request){
   templateFiles := []string {
      "./templates/list.html",
      "./templates/base.html",
   }
   ts, err := template.ParseFiles(templateFiles...)

   if err != nil {
      logError("Failed to parse template(s)"+fmt.Sprintf("%v",err))
      httpISE(w)
      return
   }

   db, err := gorm.Open("sqlite3", "db.sqlite")
   if err != nil {
      logError("Failed to open DB")
      httpISE(w)
   }
   defer db.Close()
   variables := make([]*Variable,0)
   db.Find(&variables)
   fmt.Printf("%v",variables)
   ts.Execute(w,variables)
}
