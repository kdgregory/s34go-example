package main

import "fmt"
import "github.com/kdgregory/s34go"

func main() {
    service,err := s34go.NewS3Service("http://s3.amazon.com", "ksdflkjsdldsflkdfdlf", "lk09r32r09lskdfjkf")
	fmt.Println("service = ", service, ", err = ", err)

    bucket,err := service.GetBucket("example")
	fmt.Println("bucket = ", bucket, ", err = ", err)
}
