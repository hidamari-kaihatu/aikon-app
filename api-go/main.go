package main
import (
    "fmt"
    "net/http"
    "bytes"
	"encoding/json"
    "log"
    "database/sql"
    "io/ioutil"
    "os"
    "context"
    "strings"
    "time"
	// "crypto/aes"
	// "crypto/rand"
    // "crypto/cipher"
    // "encoding/hex"
    // "encoding/base64"
    "strconv"

    "github.com/joho/godotenv"
    "github.com/comail/colog"
    "firebase.google.com/go/v4"
    "github.com/k-washi/jwt-decode/jwtdecode"
    "github.com/dgrijalva/jwt-go"


    _ "github.com/go-sql-driver/mysql"
)

var setCookie string


func envLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env target")
	}
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello, World</h1>")//固定値を返してる
}

//受け取るデータ構造の定義
type DailyReports struct {
    Id    int `json:id`
	Date  string `json:data`
	Student_id  int `json:student_id`
	Attend bool `json:attend`
	Temperature *string `json:temperature`
	SomeoneToPickUp *string `json:someoneToPickUp`
	TimeToPickUp *string `json:timeToPickUp`
	Message *string `json:message`
}

type Middle struct {
    Id int `json:id`
    Staff_id int `json:staff_id`
    Center_id *int `json:center_id`
    Role_id int `json:role_id`
}

type TeacherMessage struct {
    Id int `json:id`
    Staff_id int `json:staff_id`
    Message *string `json:message`
    Datetime string `json:datetime`
    Student_id int `json:student_id`
    Voice string `json:voice`
    //音声データ追加
}

//データベースに接続する部分
func connectionDB() *sql.DB {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    fmt.Println(dsn)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err1")
    }
    return db
}

// db.Queryはクエリを実行しRows型で返す。Rows型はクエリの実行結果
func getDailyReportsRows(db *sql.DB) *sql.Rows {
    rows, err := db.Query("SELECT * FROM dailyReports")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

//rows.Next()でループを回す。rows.Nextは、Scanメソッドを実行しているレコードの次が存在するかどうかでtrueかfalseを返す。Scanメソッドは引数に与えられた変数にフィールドの値を代入。今回は8つのフィールドがあるので8つの引数を与える。また引数に渡した変数に書き込むのでポインタで渡す。
func getDailyReport(w http.ResponseWriter, r *http.Request) {
    log.Printf("trace: this is getDailyReport log.")
    cookie := r.Header.Get("set-cookie")
    log.Printf("cookie@daily:", cookie)
    db := connectionDB()
    defer db.Close()
    rows := getDailyReportsRows(db) // 行データ取得
    dailyReports := DailyReports{}
    var resultDailyReport [] DailyReports
    for rows.Next() {
        error := rows.Scan(&dailyReports.Id, &dailyReports.Date, &dailyReports.Student_id, &dailyReports.Attend, &dailyReports.Temperature, &dailyReports.SomeoneToPickUp, &dailyReports.TimeToPickUp, &dailyReports.Message)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultDailyReport = append(resultDailyReport, dailyReports)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultDailyReport);/*jsonにエンコード*/ err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

func postDailyReport(w http.ResponseWriter, r *http.Request) {
    log.Printf("trace: this is a trace log postDailyReport start.")
    log.Printf("trace: this is getDailyReport log.")
    log.Printf(setCookie)

    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(setCookie, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })
    fmt.Printf("%v\n", token)

	if err != nil {
        fmt.Println("verifyToken error")
	}
    for key, val := range claims {
        fmt.Printf("Key: %v, value: %v\n", key, val)
        fmt.Printf("%T\n", val)
    }
    id := claims["sutudent"]
    log.Printf("trace: this is studentID log.")
    fmt.Printf("%v\n", id/*claims["sutudent"]*/) //作成されたjwt確認

	//return token, nil

    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err1")
    }
    defer db.Close()
    // request bodyの読み取り
    log.Printf("trace: this is a trace log postDailyReport request body.")
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }
    jsonBytes := ([]byte)(b)
    data := new(DailyReports)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO dailyReports (student_id, attend, temperature, someoneToPickUp, timeToPickUp, message) VALUES (?, ?, ?, ?, ?, ?)", data.Student_id, data.Attend, data.Temperature, data.SomeoneToPickUp, data.TimeToPickUp, data.Message)
    log.Printf("trace: this is a trace log postDailyReport end.")

}

func getMiddleRows(db *sql.DB) *sql.Rows {
    log.Printf("trace: this is a trace log getMiddleRows start.")
    rows, err := db.Query("SELECT * FROM middle")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    log.Printf("trace: this is a trace log getMiddleRows end.")
    return rows
}

func getMiddle(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getMiddleRows(db) // 行データ取得
    middle := Middle{}
    var resultMiddle [] Middle
    for rows.Next() {
        error := rows.Scan(&middle.Id, &middle.Staff_id, &middle.Center_id, &middle.Role_id)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultMiddle = append(resultMiddle, middle)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultMiddle);/*jsonにエンコード*/ err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

func postMiddle (w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err1")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }
    jsonBytes := ([]byte)(b)
    data := new(Middle)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO middle (staff_id, center_id, role_id) VALUES (?, ?, ?)", data.Staff_id, data.Center_id, data.Role_id)
}

func getTeacherMessageRows(db *sql.DB) *sql.Rows {
    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(setCookie/*入る？*/, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })
    fmt.Printf("%v\n", token)

	if err != nil {
        fmt.Println("verifyToken error")
	}
    for key, val := range claims {
        log.Printf("Key: %v, value: %v\n", key, val)
        fmt.Printf("%T\n", val)
        log.Printf("Verified matchId val: %v\n", val)
    }
    id := claims["sutudent"] 

    rows, err := db.Query("SELECT * FROM dailyReports where staffs.staff_id = ?", id)
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

func getTeacherMessage(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getTeacherMessageRows(db) // 行データ取得
    teacherMessage := TeacherMessage{}
    var resultTeacherMessage [] TeacherMessage
    for rows.Next() {
        error := rows.Scan(&teacherMessage.Id, &teacherMessage.Staff_id, &teacherMessage.Message, &teacherMessage.Datetime, &teacherMessage.Student_id, &teacherMessage.Voice)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultTeacherMessage = append(resultTeacherMessage, teacherMessage)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultTeacherMessage);/*jsonにエンコード*/ err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

func postTeacherMessage (w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err1")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }
    jsonBytes := ([]byte)(b)
    data := new(TeacherMessage)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO teacherMessage (staff_id, message, datetime, student_id, voice) VALUES (?, ?, ?, ?, ?)", data.Staff_id, data.Message, data.Datetime, data.Student_id, data.Voice)
}

type Centers struct {
    Id int `json:id`
    Name string`json:name`
    Status *bool `json:status`
    ProductId *string `json:productId`
}

func getRows(db *sql.DB) *sql.Rows { //mysqlからcenterの情報取得
    rows, err := db.Query("SELECT * FROM centers")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

func getCenter(w http.ResponseWriter, r *http.Request) {
    log.Printf("trace: this is a trace log getCenter start.")
    db := connectionDB()
    defer db.Close()
    rows := getRows(db) // 行データ取得
    centers := Centers{}
    var resultCenter [] Centers
    for rows.Next() {
        error := rows.Scan(&centers.Id, &centers.Name, &centers.Status, &centers.ProductId)
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

    log.Printf("trace: this is a trace log getCenter end.")
}

func postCenter(w http.ResponseWriter, r *http.Request) {
    log.Printf("trace: this is a trace log postCenter start.")
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
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
    _, err = db.Exec("INSERT INTO centers (name, productId) VALUES (?,?)", data.Name, data.ProductId)
    log.Printf("trace: this is a trace log postCenter end.")
}

func putCenterStatus(w http.ResponseWriter, r *http.Request) {
    log.Printf("trace: this is a trace log putCenterStatus start.")
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(Students)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("UPDATE centers SET status = false where id = ?", data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}
func putCenterProductId(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    _, err = db.Exec("UPDATE centers SET productId = ? where id = ?", data.ProductId, data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}

type Payments struct {
    Id int `json:id`
    Center_id int `json:center_id`
    Date string`json:date`
    PayAmount int `json:payments`
}

func getRowsP(db *sql.DB) *sql.Rows { //mysqlからcenterの情報取得
    rows, err := db.Query("SELECT * FROM payment")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}


type Students struct {
    Id int `json:id`
    Center_id int `json:center_id`
    Name string`json:name`
    ContactTell *string `json:contactTell`
    Grade int `json:grade`
    Email *string `json:email`
    Status bool `json:status`
    Rfid *string `json:rfid`
}

func getRowsStu(db *sql.DB) *sql.Rows { //mysqlからcenterの情報取得
    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(setCookie/*入る？*/, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })
    fmt.Printf("%v\n", token)

	if err != nil {
        fmt.Println("verifyToken error")
	}
    for key, val := range claims {
        log.Printf("Key: %v, value: %v\n", key, val)
        fmt.Printf("%T\n", val)
        log.Printf("Verified matchId val: %v\n", val)
    }
    id := claims["sutudent"]

    rows, err := db.Query("SELECT * FROM students where students.Id = ?", id)
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

func getStudents(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsStu(db) // 行データ取得
    students := Students{}
    var resultStudents [] Students
    for rows.Next() {
        error := rows.Scan(&students.Id, &students.Center_id, &students.Name, &students.ContactTell, &students.Grade, &students.Email, &students.Status, &students.Rfid)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultStudents = append(resultStudents, students)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultStudents);/*jsonにエンコード*/ err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

func postStudent(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(Students)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO students (center_id, name, contactTell, grade, email, status, rfid) VALUES (?, ?, ?, ?, ?, ?, ?)", data.Center_id, data.Name, data.ContactTell, data.Grade, data.Email, data.Status, data.Rfid)
    if err != nil {
        fmt.Println("insert error!")
    }
}

func putStuStatus(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(Students)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("UPDATE students SET status = false where id = ?", data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}

func putStuRfid(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(Students)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("UPDATE students SET rfid = ? where id = ?", data.Rfid, data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}

type Staffs struct {
    Id int `json:id`
    Name string`json:name`
    Email string `json:email`
    Status bool `json:status`
    Rfid string `json:rfid`
}
func getRowsSta(db *sql.DB) *sql.Rows { //mysqlからcenterの情報取得
    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(setCookie/*入る？*/, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })
    fmt.Printf("%v\n", token)
    log.Printf(setCookie)

	if err != nil {
        fmt.Println("verifyToken error")
	}
    for key, val := range claims {
        log.Printf("Key: %v, value: %v\n", key, val)
        fmt.Printf("%T\n", val)
        log.Printf("Verified matchId val: %v\n", val)
    }
    id := claims["sutudent"]

    rows, err := db.Query("SELECT * FROM staffs where staffs.id = ?", id)
    if err != nil {
        fmt.Println("Err2 a")
        panic(err.Error())
    }
    return rows
}
func getStaffs(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsSta(db) // 行データ取得
    staffs := Staffs{}
    var resultStaffs [] Staffs
    for rows.Next() {
        error := rows.Scan(&staffs.Id, &staffs.Name, &staffs.Email, &staffs.Status, &staffs.Rfid)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultStaffs = append(resultStaffs, staffs)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultStaffs);/*jsonにエンコード*/ err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

func postStaff(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(Staffs)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO staffs (name, email, status, rfid) VALUES (?, ?, ?, ?)", data.Name, data.Email, data.Status, data.Rfid)
    if err != nil {
        fmt.Println("insert error!")
    }
}

func putStaStatus(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(Staffs)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("UPDATE staffs SET status = false where id = ?", data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}

func putStaRfid(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(Staffs)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("UPDATE staffs SET rfid = ? where id = ?", data.Rfid, data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}

type Sensors struct {
    Id int `json:id`
    Place string`json:place`
    SerialNumber string `json:serialNumber`
    Center_id int `json:center_id`
}
func getRowsSensors(db *sql.DB) *sql.Rows {
    rows, err := db.Query("SELECT * FROM sensors")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}
func getSensors(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsSensors(db) // 行データ取得
    sensors := Sensors{}
    var resultSensors [] Sensors
    for rows.Next() {
        error := rows.Scan(&sensors.Id, &sensors.Place, &sensors.SerialNumber, &sensors.Center_id)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultSensors = append(resultSensors, sensors)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultSensors); err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

func postSensor(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(Sensors)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO sensors (place, serialNumber, center_id) VALUES (?, ?, ?)", data.Place, data.SerialNumber, data.Center_id)
    if err != nil {
        fmt.Println("insert error!")
    }
}

type InAndOut struct {
    Id int `json:id`
    Rfid string`json:rfid`
    Sensor_id int `json:sensor_id`
    Datetime string `json:datetime`
}
func getRowsInAndOut(db *sql.DB) *sql.Rows {
    rows, err := db.Query("SELECT * FROM inAndOut")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}
func getInAndOut(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsInAndOut(db) // 行データ取得
    inAndOut := InAndOut{}
    var resultInAndOut [] InAndOut
    for rows.Next() {
        error := rows.Scan(&inAndOut.Id, &inAndOut.Rfid, &inAndOut.Sensor_id, &inAndOut.Datetime)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultInAndOut = append(resultInAndOut, inAndOut)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultInAndOut); err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

func postInAndOut(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB-HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
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
    data := new(InAndOut)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO inAndOut (rfid, sensor_id, datetime) VALUES (?, ?, ?)", data.Rfid, data.Sensor_id, data.Datetime)
    if err != nil {
        fmt.Println("insert error!")
    }
}

type StuInAndOutSensors struct {
    Id int  `json:id`
    Datetime string  `json:datetime`
    Rfid string  `json:rfid`
    Sensor_id int  `json:sensor_id`
    Name string  `json:name`
    Place string  `json:place`
}

func getRowsStuInAndOutSensors(db *sql.DB, id int) *sql.Rows {
    rows, err := db.Query(`SELECT students.id, inAndOut.datetime, students.rfid, inAndOut.sensor_id, students.name, sensors.place FROM students INNER JOIN inAndOut ON students.rfid = inAndOut.rfid INNER JOIN sensors ON inAndOut.sensor_id = sensors.id WHERE students.id = ? AND inAndOut.datetime > CURDATE() AND sensors.place = "入口・出口"`,id)
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

func getStuInAndOutSensors(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsStuInAndOutSensors(db, 1)
    stuInAndOutSensors := StuInAndOutSensors{}
    var resultStuInAndOutSensors [] StuInAndOutSensors
    for rows.Next() {
        error := rows.Scan(&stuInAndOutSensors.Id, &stuInAndOutSensors.Datetime, &stuInAndOutSensors.Rfid, &stuInAndOutSensors.Sensor_id, &stuInAndOutSensors.Name, &stuInAndOutSensors.Place)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultStuInAndOutSensors = append(resultStuInAndOutSensors, stuInAndOutSensors)
        }
    }
    var buf bytes.Buffer
    enc := json.NewEncoder(&buf)
    if err := enc.Encode(&resultStuInAndOutSensors); err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

// // //暗号化スタート
// func GenerateIV() ([]byte, error) { //１番初めの暗号文ブロック作成
// 	iv := make([]byte, aes.BlockSize) // ランダムなiv作成　BlockSize は 16byte サイズ16のスライス作成
// 	if _, err := rand.Read(iv); err != nil { // Note that err == nil only if we read len(b) bytes.
// 		return nil, err
// 	}
// 	return iv, nil
// }

// func Pkcs7Pad(data []byte) []byte { //16bytesに足らない分を追加  平文[]byte の長さが16の倍数ではない可能性がある場合、16の倍数にするためにパディングする
//     length := aes.BlockSize - (len(data) % aes.BlockSize)
//     trailing := bytes.Repeat([]byte{byte(length)}, length)
//     return append(data, trailing...)
// }

// var (
//     enc string
//     outIv []byte
// )

// func Encrypt(dataString string) (iv []byte, encrypted []byte, err error) {
//     log.Printf("trace: this is a 暗号化関数.")
//     key, _ := hex.DecodeString(os.Getenv("KEY")) //keyを引数ではなく、関数の中で固定値で持つ
//     data, _ := hex.DecodeString(dataString)

//     iv, err = GenerateIV()
//     if err != nil {
//         return nil, nil, err
//     }
//     block, err := aes.NewCipher(key)// 暗号オブジェクトを作る
//     if err != nil {
//         return nil, nil, err
//     }
//     padded := Pkcs7Pad(data)
//     encrypted = make([]byte, len(padded))
//     cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
//     cbcEncrypter.CryptBlocks(encrypted, padded)

//     log.Printf("trace: byte型をエンコード")
//     enc = base64.StdEncoding.EncodeToString(encrypted)
//     outIv = iv
//     log.Printf("Encrypted:" + enc + "\n")

//     // log.Printf("trace: base64をデコード")
//     // dec, err := base64.StdEncoding.DecodeString(enc)
//     // if err != nil {
// 	// 	log.Fatal(err)
// 	// }
//     // fmt.Println("decode:", dec)

//     return iv, encrypted, nil
// }       

// //複合化　スタート　cookieからivと暗号化してエンコードされたuser_idを取ってくる処理追加 ←にするなら引数は文字列？
// func Pkcs7Unpad(data []byte) []byte {
//     dataLength := len(data)
//     padLength := int(data[dataLength-1])
//     return data[:dataLength-padLength]
// }

// func Decrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
//     block, err := aes.NewCipher(key)
//     if err != nil {
//         return nil, err
//     }
//     decrypted := make([]byte, len(data))
//     cbcDecrypter := cipher.NewCBCDecrypter(block, iv)
//     cbcDecrypter.CryptBlocks(decrypted, data)
//     return Pkcs7Unpad(decrypted), nil
// }

//トークン作成
func CreateToken(studentID string) (string) {
    // tokenの作成
    token := jwt.New(jwt.GetSigningMethod("HS256"))

    // claimsの設定
    token.Claims = jwt.MapClaims{
        "sutudent": studentID, //1つめのjwtで取れたstudent_idが入る
        "exp":  time.Now().Add(time.Hour * 1).Unix(), // 有効期限を指定 １時間は適切 ???
    }

    // 署名
    var secretKey = "secret" // 任意の文字列
    tokenString, err := token.SignedString([]byte(secretKey)) 
    if err != nil {
        log.Printf("signiture error") //return "", err
    }
    return tokenString
}

//cookieに乗ってきたjwtを解析する用？
// func VerifyToken(tokenString string) (*jwt.Token, error) {

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("secret"), nil
// 	})
// 	if err != nil {
// 		return  nil, err
// 	}

// 	return token, nil
//}

type StuId struct {
    Id int `json:id`
}
type LoginResponse struct {
    Token string `json:"token"`
}
//studentテーブルのid取得用関数宣言
func getRowsStuId(db *sql.DB, email string) *sql.Rows {
    rows, err := db.Query("SELECT id FROM aikon_db.students where email = ?", email)
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    fmt.Println(rows)
    return rows
}      

type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}


func parentIsLogin(w http.ResponseWriter, r *http.Request){ //保護者ページ用
        app, err := firebase.NewApp(context.Background(), nil)
        if err != nil {
            fmt.Printf("error: %v\n", err)
            os.Exit(1)
        }
        auth, err := app.Auth(context.Background())
        if err != nil {
            fmt.Printf("error: %v\n", err)
            os.Exit(1)
        }
    //log.Printf("trace: this is a trace headers test.")
	authHeader := r.Header.Get("Authorization")
    log.Printf("Verified authHeader: %T\n", authHeader)
    idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	// fmt.Fprintln(w, idToken)
    // fmt.Println(w, idToken)

    log.Printf("trace: this is a trace JWT の検証.")
    // JWT の検証
    token, err := auth.VerifyIDToken(context.Background(), idToken)
    if err != nil {
    // JWT が無効なら Handler に進まず別処理
        fmt.Printf("error verifying ID token: %v\n", err)
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("error verifying ID token\n"))
        return
    }
    log.Printf("Verified ID token: %v\n", token)

    jwt := authHeader
    hCS, err := jwtdecode.JwtDecode.DecomposeFB(jwt)
    if err != nil {
        log.Fatalln("Error : ", err)
    }
    payload, err := jwtdecode.JwtDecode.DecodeClaimFB(hCS[1])
    if err != nil {
        log.Fatalln("Error :", err)
    }

    //ユーザーIDと, メールアドレスを表示
    user := payload.Subject
    emailInToken := payload.Email
    log.Printf("User ID: " + user + " ,Email: " + emailInToken + "\n")

    //studentテーブルのid取得用
    db := connectionDB()
    defer db.Close()
    rows := getRowsStuId(db, emailInToken)
    stuId := StuId{}
    var resultStuId [] StuId
    for rows.Next() {
        error := rows.Scan(&stuId.Id)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultStuId = append(resultStuId, stuId)
        }
    }
    matchId := resultStuId[0].Id //トークンで取れたemailと同じアドレスの人のidを変数に入れた！
    log.Printf("Verified matchId: %v\n", matchId)
    fmt.Printf("%T\n", matchId)//数値型
    toString := strconv.Itoa(matchId) //文字列に変換
    log.Printf("Verified matchId type: %T\n", toString)

    // //暗号化処理
    // Encrypt(toString)
    // log.Printf("Verified enc: %v\n", enc)
    // log.Printf("Verified outIv: %v\n", outIv)

    //jwtにのせる　認証情報が入ったJsonを加工（電子署名を加える等）し、JWTにしたのち、それを認証Tokenとして ×クライアントに渡す　〇クッキーに渡す
    afterAuthJwt := CreateToken(toString)
    log.Printf(afterAuthJwt) //作成されたjwt確認

    setCookie = afterAuthJwt

    cookie := &http.Cookie{
        Name:   "studentID",
        Value:  afterAuthJwt,
        HttpOnly: true,
    }

    http.SetCookie(w, cookie)
    log.Printf("cookie: ", cookie) 
}

type StaId struct {
    Id int `json:id`
}

func getRowsStaId(db *sql.DB, email string) *sql.Rows {
    rows, err := db.Query("SELECT id FROM aikon_db.staffs where email = ?", email)
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    fmt.Println(rows)
    return rows
}      


//staff用ログイン関数
func staffIsLogin(w http.ResponseWriter, r *http.Request){ //保護者ページ用
    app, err := firebase.NewApp(context.Background(), nil)
    if err != nil {
        fmt.Printf("error: %v\n", err)
        os.Exit(1)
    }
    auth, err := app.Auth(context.Background())
    if err != nil {
        fmt.Printf("error: %v\n", err)
        os.Exit(1)
    }
log.Printf("trace: this is a trace headers test.")
authHeader := r.Header.Get("Authorization")
log.Printf("Verified authHeader: %T\n", authHeader)
idToken := strings.Replace(authHeader, "Bearer ", "", 1)
// fmt.Fprintln(w, idToken)
// fmt.Println(w, idToken)

log.Printf("trace: this is a trace JWT の検証.")
// JWT の検証
token, err := auth.VerifyIDToken(context.Background(), idToken)
if err != nil {
// JWT が無効なら Handler に進まず別処理
    fmt.Printf("error verifying ID token: %v\n", err)
    w.WriteHeader(http.StatusUnauthorized)
    w.Write([]byte("error verifying ID token\n"))
    return
}
log.Printf("Verified ID token: %v\n", token)

jwt := authHeader
hCS, err := jwtdecode.JwtDecode.DecomposeFB(jwt)
if err != nil {
    log.Fatalln("Error : ", err)
}
payload, err := jwtdecode.JwtDecode.DecodeClaimFB(hCS[1])
if err != nil {
    log.Fatalln("Error :", err)
}

//ユーザーIDと, メールアドレスを表示
user := payload.Subject
emailInToken := payload.Email
log.Printf("User ID: " + user + " ,Email: " + emailInToken + "\n")

//studentテーブルのid取得用
db := connectionDB()
defer db.Close()
rows := getRowsStaId(db, emailInToken)
staId := StaId{}
var resultStaId [] StaId
for rows.Next() {
    error := rows.Scan(&staId.Id)
    if error != nil {
        fmt.Println("scan error")
    } else {
        resultStaId = append(resultStaId, staId)
    }
}
matchId := resultStaId[0].Id //トークンで取れたemailと同じアドレスの人のidを変数に入れた！
log.Printf("Verified matchId: %v\n", matchId)
fmt.Printf("%T\n", matchId)//数値型
toString := strconv.Itoa(matchId) //文字列に変換
log.Printf("Verified matchId type: %T\n", toString)

// //暗号化処理
// Encrypt(toString)
// log.Printf("Verified enc: %v\n", enc)
// log.Printf("Verified outIv: %v\n", outIv)

//jwtにのせる　認証情報が入ったJsonを加工（電子署名を加える等）し、JWTにしたのち、それを認証Tokenとして ×クライアントに渡す　〇クッキーに渡す
afterAuthJwt := CreateToken(toString)
log.Printf(afterAuthJwt) //作成されたjwt確認

setCookie = afterAuthJwt

cookie := &http.Cookie{
    Name:   "studentID",
    Value:  afterAuthJwt,
    HttpOnly: true,
}

http.SetCookie(w, cookie)
log.Printf("cookie: ", cookie) 
}


//cookieに乗ってきたjwtを解析する用
// func verifyToken(tokenString string) (*jwt.Token, error){
//     claims := jwt.MapClaims{}
//     token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
//         return []byte("secret"), nil
//     })
    
// 	if err != nil {
//         fmt.Println("verifyToken error")
// 	}
//     for key, val := range claims {
//         fmt.Printf("Key: %v, value: %v\n", key, val)
//         fmt.Printf("%T\n", val)
//     }
//     fmt.Printf("%v\n", claims["sutudent"]) //作成されたjwt確認

// 	return token, nil

// }


func main() {
    envLoad()
    colog.SetDefaultLevel(colog.LDebug)
    colog.SetMinLevel(colog.LTrace)
    colog.SetFormatter(&colog.StdFormatter{
        Colors: true,
        Flag:   log.Ldate | log.Ltime | log.Lshortfile,
    })
    colog.Register()
    log.Printf("trace: this is a trace log test.")
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/dailyReportGet", getDailyReport) 
    http.HandleFunc("/dailyReportPost", postDailyReport)
    http.HandleFunc("/middleGet", getMiddle)
    http.HandleFunc("/middlePost", postMiddle)
    http.HandleFunc("/teacherMessageGet", getTeacherMessage)//その先生が送ったメッセージをgetする必要あり　staff_idカラムあり
    http.HandleFunc("/teacherMessagePost", postTeacherMessage)
	http.HandleFunc("/centerGet", getCenter)
    http.HandleFunc("/centerPost", postCenter)
    http.HandleFunc("/centerPut", putCenterStatus)
    http.HandleFunc("/centerProductIdPut", putCenterProductId)
    http.HandleFunc("/studentsGet", getStudents)
    http.HandleFunc("/studentPost", postStudent)
    http.HandleFunc("/stuStatustPut", putStuStatus)
    http.HandleFunc("/stuRfidPut", putStuRfid)
    http.HandleFunc("/staffsGet", getStaffs)
    http.HandleFunc("/staffPost", postStaff)
    http.HandleFunc("/staStatustPut", putStaStatus)
    http.HandleFunc("/staRfidPut", putStaRfid)
    http.HandleFunc("/sensorsGet", getSensors)
    http.HandleFunc("/sensorPost", postSensor)
    http.HandleFunc("/inAndOutGet", getInAndOut)
    http.HandleFunc("/inAndOutPost", postInAndOut)
    http.HandleFunc("/stuInAndOutSensorsGet", getStuInAndOutSensors)
    http.HandleFunc("/parentIsLogin", parentIsLogin)
    http.HandleFunc("/staffIsLogin", staffIsLogin)
    // fmt.Println("Server Start")
    // http.ListenAndServe(":8080", handler)
    http.ListenAndServe(":8080", nil)
    
}