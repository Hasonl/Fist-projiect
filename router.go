
Attension,Please:
                  When you run the next go in the Centos System,you should be Install git go golang  --yum -y install git go golang
                  And then you should  export GOPATH=$PATH.
                  And the last,you should be     git clone github.com/julienschmidt/httprouter  on your project Directory.
                

// src Directory will be save your clone file,In this directory,you will see  github.com/julienschmidt/httprouter/ directory and some other's go.


Shell:
        mkdir -p hason/src                
        cd hason
        vim 8080.go

Go:

package  main
import (
        "fmt"
        "log"
        "net/http"
        "github.com/julienschmidt/httprouter"
)

func hason(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome! Keep moving, Here we go\n")
}
func main() {
    router := httprouter.New()
    router.GET("/", hason)
    log.Fatal(http.ListenAndServe(":8080", router))
}

http://localhost:8080

PS:  根据路由请求8080端口的hason内容
