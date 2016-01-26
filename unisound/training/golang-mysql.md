1. download golang toolset(compiler&runtime) from http://www.golangtc.com/download
1. extract and add bin\ path to PATH
1. make working directory, and a `src` directory in working directory
1. plase following source in `src` directory, like `main.go`
1. go to working directory in bash and run ``env GOPATH=`pwd` go run src/main.go -ma="oralsvc:oralsvc@tcp(192.168.5.17:3306)/oralsvc" -q1='insert into OralLogDay_021 values("ip","sessionid",?,"server_addr","server_serial","server_type",null,100,"opt_resformat","opt_imei","opt_key","task_type","oral_textoral_textoral_textoral_textoral_textoral_textoral_textoral_textoral_text", "", 100,100,100,100,100,"fileid","attr")' -q2='insert into OralVoiceDay_021 (time_stamp,create_time,voice,result,sessionid) values(null,?,?,"resultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresultresult","sessionid")' -cc=200 -aa=10``

```golang
package main
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"flag"
	"log"
	"time"
)

func main(){
	log.SetFlags(log.Lshortfile)
	bin:= make([]byte, 200 * 1024)
	msqlAddr := flag.String("ma","","MySQL address, like mysql://user:pass@tcp(localhost:3306)/OralEval")
	querySQL1 := flag.String("q1","","SQL template to execute, plase a `?` for current unix-time in nanoseconds")
	querySQL2 := flag.String("q2","","SQL template to execute, plase first `?` for current unix-time in nanoseconds, second `?` for 200KB binary")
	//多少个线程一起跑，每个线程一个独立的MySQL连接
	conCount := flag.Int("cc", 500, "concurrent count, threads and connections to MySQL")
	//每个线程插入多少条以后统计一下平均值，打印出插入耗时（每次插入成功都打印会导致log太多）
	avgAcc := flag.Int("aa", 2000, "make avrage of aa sql consume times to print one message")

	flag.Parse()
	db,err:= sql.Open("mysql", *msqlAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(*conCount + 10)

	for i:=0;i<*conCount;i++{
		go func(i int){
			c := 0
			avg := int64(0)
			for{
				start := time.Now()
				_, err := db.Exec(*querySQL1, start.UnixNano())
				if err != nil{
					log.Fatalln(err)
				}
				end := time.Now()
				avg += int64(end.Sub(start).Nanoseconds()) / 1000 / 1000

				start = time.Now()
				_, err = db.Exec(*querySQL2, start.UnixNano(), bin)
				if err != nil{
					log.Fatalln(err)
				}
				end = time.Now()
				avg += int64(end.Sub(start).Nanoseconds()) / 1000 / 1000

				c++
				if c >= *avgAcc{
					log.Println(i,":",avg / int64(*avgAcc))
					avg = 0
					c = 0
				}
			}

		}(i)
	}
	time.Sleep(time.Hour * time.Duration(100000))
}
```
