# gogridclient

Golang client for [Grid](https://github.com/guardian/grid) services

---

This is a work in progress and only currently supports the usage reindex endpoint.

##Example Usage
```go
package main

import (
	"github.com/guardian/gogridclient"
  "log"
)

func main() {
  usageReindexRaw()
  usageReindex()
}

func usageReindexRaw(){
  //A raw request
  path := "/usages/digital/content/" + id + "/reindex"
  url := "https://usage.example.com" + path

  req, err := http.NewRequest("GET", url, nil)

  gridService := gogridclient.NewGridService("my-api-key")
  response, err := gridService.Do(req)

  if err != nil {
    log.Fatal(err)
  }

  log.Print(response)
}

func usageReindex(){
  //Reindex image usages by content ID
  usageService := gogridclient.NewUsageService(
    "https://usage.example.com", "my-api-key")

  response, err := usageService.Reindex("some/content/id")
  if err != nil {
    log.Fatal(err)
  }

  log.Print(response)
}
```
