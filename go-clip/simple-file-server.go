package main
import (
	"net/http"
	"io"
	"os"
	"fmt"
	"flag"
	"path"
)

func main() {
	mp3Root := flag.String("mr",".","mp3 files root directory")
	flag.Parse()
	fmt.Println("use mp3 in", *mp3Root)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "audio/mpeg")
		p := path.Join(*mp3Root, r.URL.Path[1:])
		f, err := os.Open(p)
		if err != nil {
			fmt.Fprintf(w, "Error: opening file %s failed:%v", p, err)
			return
		}
		defer f.Close()
		n, err := io.Copy(w, f)
		if err != nil {
			fmt.Println("Error: copy file to http responsefailed:", err)
		}else{
			fmt.Printf("send %v to %v success with %v bytes\n", r.URL.Path, r.RemoteAddr, n)
		}
	})

	fmt.Println(http.ListenAndServe(":8082", nil))
}
