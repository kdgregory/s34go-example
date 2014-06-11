s34go-example
=====

An program that uses the s34go library to upload and downoad files from S3. While useful, it's
primarily intended as an example of how to use the library.

## Building

Assuming that you've correctly set up your Go workspace, and set `GOPATH` to its root:

    go get github.com/kdgregory/s34go-example
    
After making changes:

    go install github.com/kdgregory/s34go-example


## Running

List all buckets belonging to the current user:

    bin/s34go-example [-config FILE] [-secret KEY] [-public KEY] list 

List all objects within a specified bucket:

    bin/s34go-example [-config FILE] [-secret KEY] [-public KEY] list BUCKET

Retrieve an object, storing it in the specified destination:

    bin/s34go-example [-config FILE] [-secret KEY] [-public KEY] get S3SRC LCLDEST

Stores a local file in the specified object, overwriting it if it already exists:

    bin/s34go-example [-config FILE] [-secret KEY] [-public KEY] put LCLSRC S3DEST


where:
- `S3SRC` and `S3DEST` are S3 object locations, in the form `bucket:path`
- `LCLSRC` and `LCLDEST` are local filesystem paths

- `-config` specifies the name of a configuration file (defaults to `$HOME/.s34go.ini`)
- `-public` specifies the user's public access key, overriding the value in the config file (if any)
- `-secret` specifies the user's secret access key, overriding the value in the config file (if any)
