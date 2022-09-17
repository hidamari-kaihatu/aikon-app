package main
import (
    "fmt"
    "net/http"
    "bytes"
	"encoding/json"
    "log"
    "database/sql"
    "io/ioutil"

    _ "github.com/go-sql-driver/mysql"
)
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello, World</h1>")//固定値を返してる
}

type Centers struct {
    Id int `json:id`
    Name string`json:name`
}

func connectionDB() *sql.DB { //dbと接続
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err1")
    }
    return db
}

func getRows(db *sql.DB) *sql.Rows { //mysqlからcenterの情報取得
    rows, err := db.Query("SELECT * FROM center")
    if err != nil {
        fmt.Println("Err2")
         panic(err.Error())
    }
    return rows
}

func getCenter(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRows(db) // 行データ取得
    centers := Centers{}
    var resultCenter [] Centers
    for rows.Next() {
        error := rows.Scan(&centers.Id, &centers.Name)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultCenter = append(resultCenter, centers)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultCenter);/*jsonにエンコード*/ err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }

    //fmt.Println(resultCenter)
}

func postCenter(w http.ResponseWriter, r *http.Request) {
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err1")
    }
    defer db.Close()
    // request bodyの読み取り
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }

    // jsonのdecode
    jsonBytes := ([]byte)(b)
    data := new(Centers)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
     _, err = db.Exec("INSERT INTO center (id, name) VALUES (?, ?)", data.Id, data.Name)
     //TODO ここにどうやってCORS対策を入れられるか？？

}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/centerGet", getCenter)
    http.HandleFunc("/centerPost", postCenter)
    // fmt.Println("Server Start")
    http.ListenAndServe(":8080", nil)
    // fmt.Println("Hello, World!!")
}