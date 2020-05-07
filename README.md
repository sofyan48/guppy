# GUPPY PARAMETER STORE ENVIRONMENT
Guppy is a storage environment so that the entire environment is concentrated in one place, KV Store based guppies use ETCD as its base

## GETTING STARTED
Installing ETCD depend to use this software

## GUPPY
Client Library for access etcd project

go get github.com/sofyan48/guppy

```golang
package main 

import "github.com/sofyan48/guppy/guppy"

func main(){
    config := config.NewConfig()
    config.DialTimeOut = 10
    config.Urls = $URLS
    guppy.Client(config).New()
}

```
## GUPPY COMMAND LINE (CLI)
```
$./guppy -h                                                                     
NAME:
   guppy - guppy [command]

USAGE:
   guppy [global options] command [command options] [arguments...]

VERSION:
   0.1.1

AUTHOR:
   sofyan48 <meongbego@gmail.com>

COMMANDS:
   put      put [command]
   get      get [option]
   rm       rm [option]
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE  Load environtment config from FILE
   --help, -h              show help
   --version, -v           print the version
```
### PUT Parameters
```
NAME:
   guppy put - put [command]

USAGE:
   guppy put [command options] [arguments...]

OPTIONS:
   --file value, -f value   File Template Path
   --path value, -p value   Key for path
   --value value, -v value  Value for key
   --encryption             Set Encryption For Value
```
### GET Parameters
```
NAME:
   guppy get - get [option]

USAGE:
   guppy get [command options] [arguments...]

OPTIONS:
   --path value, -p value  Path Key
   --with-decryption       Set Decryption For Value
```
### RM Parameters
```
NAME:
   guppy rm - rm [option]

USAGE:
   guppy rm [command options] [arguments...]

OPTIONS:
   --path value, -p value  Path Key
```
## GUPPY REST API