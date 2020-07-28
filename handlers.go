
package main

import (
   "github.com/go-yaml/yaml"
   "encoding/json"
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
            http.Error(w,"Internal Server Error", 500)
            return
         }

         err = ts.Execute(w,nil)
         if err != nil {
            logError("Failed to execute template(s)"+fmt.Sprintf("%v",err))
            http.Error(w,"Internal Server Error", 500)
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

         marsh,err := json.Marshal(s)
         if err != nil {
            logError("Failed to deserialize")
            http.Error(w,"Error processing data", 422)
            return
         }

         fmt.Fprintf(w,string(marsh))
         break
   }
}

