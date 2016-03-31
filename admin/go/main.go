package main
import (
        "net/http"
)
func main() {
        http.Handle("/", http.FileServer(http.Dir("/home/zhijian/doge-dev/src/github.com/cooljiansir/doge/admin/web/")))
        http.ListenAndServe(":8088", nil)
}
