package main 

import "fmt" 

import "os"

import "io/ioutil"
// import "bufio"
import "encoding/json"
// import "reflect"

func check(e error) {
    if e != nil {
        panic(e)
    }
}
//to get mongodump
func getMongodump () {
}
//to get sql dump
func getMysqlDump () {
}
func compressFile () {
}

func main() { 

	var configPath = os.Args[1]
 	jsonFile, err:= os.Open(configPath)
    check(err)
   // fmt.Println(jsonFile)

   config,err := ioutil.ReadAll(jsonFile)
   var data interface{}
   errr:=json.Unmarshal(config,&data)
fmt.Println(errr)
fmt.Println(data)
}