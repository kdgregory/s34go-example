package main

import "fmt"
import "log"
import "os"

import "github.com/kdgregory/s34go"

func main() {
    config, action, args := ParseCommandLine();

    service,err := s34go.NewS3Service(config.AccessKey, config.SecretKey)
    checkError("create service", err)

    switch action {
        case ACTION_LIST  : if len(args) == 0 {
                                listBuckets(service) 
                            } else {
                                listObjects(service, args[0]) 
                            }
        case ACTION_GET   : getObject(service, args[0], args[1]) 
        case ACTION_PUT   : putObject(service, args[0], args[1]) 
    }
}


func listBuckets(service *s34go.S3Service) {
    buckets,err := service.ListBuckets()
    checkError("retrieve bucket list", err)

    for _,bucket := range buckets {
        fmt.Println(bucket)
    }
}


func listObjects(service *s34go.S3Service, bucketName string) {
    bucket,err := service.GetBucket(bucketName)
    checkError("retrieve bucket", err)
    if bucket == nil {
        fmt.Fprintln(os.Stderr, "bucket " + bucketName + " does not exist")
        os.Exit(2)
    }

    objects,err := bucket.ListObjects()
    checkError("retrieve bucket contents", err)
    for _,object := range objects {
        fmt.Println(object)
    }
}


func getObject(service *s34go.S3Service, src string, dst string) {
    log.Println("getObject(" + src + " " + dst + ")")
}


func putObject(service *s34go.S3Service, src string, dst string) {
    log.Println("putObject(" + src + " " + dst + ")")
}


func checkError(msg string, err error) {
    if err != nil {
        fmt.Fprintln(os.Stderr, "failed to " + msg + ":", err)
        os.Exit(2)
    }
}
