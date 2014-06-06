package main

import "log"
// import "github.com/kdgregory/s34go"

func main() {
    config, action, args := ParseCommandLine();

    log.Println("config = ", config)

    switch action {
        case ACTION_LIST.Name : listBucket(args[0]) 
        case ACTION_GET.Name  : getObject(args[0], args[1]) 
        case ACTION_PUT.Name  : putObject(args[0], args[1]) 
    }
}


func listBucket(bucketName string) {
    log.Println("listBucket(" + bucketName + ")")
}


func getObject(src string, dst string) {
    log.Println("getObject(" + src + " " + dst + ")")
}


func putObject(src string, dst string) {
    log.Println("putObject(" + src + " " + dst + ")")
}
