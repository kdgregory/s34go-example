package main

import "flag"
import "fmt"
import "log"
import "os"
import "strings"


type Action struct {
    Name string
    NumArgs int
    UsageDesc string
    UsageFlags string
    UsageArgs string
}

var ACTION_LIST = "list"
var ACTION_GET = "get"
var ACTION_PUT = "put"

var VALID_ACTIONS = [...]Action{
    Action{ ACTION_LIST, 0, "lists all buckets",             "[-config FILE] [-secret KEY] [-public KEY]", "" },
    Action{ ACTION_LIST, 1, "lists all objects in bucket",   "[-config FILE] [-secret KEY] [-public KEY]", "BUCKET" },
    Action{ ACTION_GET,  2, "retrieves objects",             "[-config FILE] [-secret KEY] [-public KEY]", "SRC DEST" },
    Action{ ACTION_PUT,  2, "stores objects",                "[-config FILE] [-secret KEY] [-public KEY]", "SRC DEST" },
}


func ParseCommandLine() (config Config, action string, args []string) {
    flag.Usage = func() { showUsageAndExit("") }

    var configPath string
    var overrides Config

    flag.StringVar(&configPath,            "config", DEFAULT_CONFIG_FILE, "configuration file for S3 connection")
    flag.StringVar(&(overrides.AccessKey), "public", "", "public key for S3 connection; overrides value in .s340go.ini")
    flag.StringVar(&(overrides.SecretKey), "secret", "", "secret key for S3 connection; overrides value in .s340go.ini")
    flag.Parse()

    config, err := ReadConfigFile(configPath)
    if err != nil {
        log.Fatal("unable to process config file:", err) 
    }
    config.Merge(overrides)

    action = validateAction()
    args = flag.Args()[1:]
    return 
}


func validateAction() string {
    if flag.NArg() == 0 {
        showUsageAndExit("must specify an action")
    }

    action := strings.ToLower(strings.TrimSpace(flag.Args()[0]))
    for _,allowed := range VALID_ACTIONS {
        if allowed.Name == action {
            if allowed.NumArgs == flag.NArg() - 1 {
                return action
            } else {
                showUsageAndExit("incorrect number of arguments for action " + action)
            }
        }
    }

    showUsageAndExit("invalid action: " + flag.Args()[0])
    return ""   // never called
}


func showUsageAndExit(message string) {
    if (len(message) > 0) {
        fmt.Fprintln(os.Stderr, message)
    }
    fmt.Fprintln(os.Stderr, "usage:")
    for _, action := range VALID_ACTIONS {
        fmt.Fprintln(os.Stderr, "    s3tool", action.UsageFlags, action.Name, action.UsageArgs)
        fmt.Fprintln(os.Stderr, "        ", action.UsageDesc)
    }
    fmt.Fprintln(os.Stderr, "where:")
    flag.PrintDefaults()
    os.Exit(1)
}
