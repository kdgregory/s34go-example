s34go-example
=====

An program that uses the s34go library to upload and downoad files from S3. While useful, it's
primarily intended as an example of how to use the library.

    s3get [-r] S3PATH LOCALPATH
    s3put [-r] LOCALPATH S3PATH

*where:*

* `S3PATH` identifies an object on S3. Format is `BUCKET:PATH`, where `PATH` might identify an actual object or the leading path to that object. The latter is only valid for recursive retrievals.
* `LOCALPATH` identifies a local filesystem path. May be either a directory or local file.
* `-r` indicates a recursive operation; only valid for 
