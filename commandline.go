package main

import "flag"
import "fmt"
import "log"
import "os"
import "strings"


type Action struct {
    Name string
    NumArgs int
}

var ACTION_LIST = Action{ "list", 1}
var ACTION_GET = Action{ "get", 2 }
var ACTION_PUT = Action{ "put", 2 }
var VALID_ACTIONS = [...]Action{ ACTION_LIST, ACTION_GET, ACTION_PUT }


func ParseCommandLine() (config Config, action string, args []string) {
    flag.Usage = func() { showUsageAndExit("") }

    var configPath string
    var overrides Config

    flag.StringVar(&configPath,            "config", DEFAULT_CONFIG_FILE, "configuration file for S3 connection")
    flag.StringVar(&(overrides.PublicKey), "public", "", "public key for S3 connection; overrides value in .s340go.ini")
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
    fmt.Fprintln(os.Stderr, "usage: s3tool [-config FILE] [-secret KEY] [-public KEY] (list | get | put) SRC [DEST]")
    flag.PrintDefaults()
    os.Exit(1)
}
