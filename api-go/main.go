package main
import (
    "fmt"
    "net/http"
    "bytes"
	"encoding/json"
    "log"
    "database/sql"
    "io/ioutil"

    // "github.com/rs/cors"
    _ "github.com/go-sql-driver/mysql"
)
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
    Center_id int `json:center_id`
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
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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
    jsonBytes := ([]byte)(b)
    data := new(DailyReports)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO dailyReports (date, student_id, attend, temperature, someoneToPickUp, timeToPickUp, message) VALUES (?, ?, ?, ?, ?, ?, ?)", data.Date, data.Student_id, data.Attend, data.Temperature, data.SomeoneToPickUp, data.TimeToPickUp, data.Message)
}

func getMiddleRows(db *sql.DB) *sql.Rows {
    rows, err := db.Query("SELECT * FROM middle")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
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
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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
    rows, err := db.Query("SELECT * FROM teacherMessage")
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
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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
    _, err = db.Exec("INSERT INTO centers (name) VALUES (?)", data.Name)
     //TODO ここにどうやってCORS対策を入れられるか？？

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

func getPayment(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsP(db) // 行データ取得
    payments := Payments{}
    var resultPayments [] Payments
    for rows.Next() {
        error := rows.Scan(&payments.Id, &payments.Center_id, &payments.Date, &payments.PayAmount)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultPayments = append(resultPayments, payments)
        }
    }
    var buf bytes.Buffer //バッファを作成　TODO調べる
    enc := json.NewEncoder(&buf) //書き込み先を指定するのですから、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultPayments);/*jsonにエンコード*/ err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

func postPayment(w http.ResponseWriter, r *http.Request) {
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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
    data := new(Payments)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("INSERT INTO payment (center_id, date, payAmount) VALUES (?, ?, ?)", data.Center_id, data.Date, data.PayAmount)
    if err != nil {
        fmt.Println("insert error!")
    }
}

type Students struct {
    Id int `json:id`
    Center_id int `json:center_id`
    Name string`json:name`
    ContactTell *string `json:contactTell`
    Grade int `json:grade`
    Email *string `json:email`
    Status bool `json:status`
    Rfid string `json:rfid`
}
func getRowsStu(db *sql.DB) *sql.Rows { //mysqlからcenterの情報取得
    rows, err := db.Query("SELECT * FROM students")
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
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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


type Staffs struct {
    Id int `json:id`
    Center_id int `json:center_id`
    Name string`json:name`
    Email string `json:email`
    RoleId int `json:roleId`
    Status *bool `json:status`
    Rfid string `json:rfid`
}
func getRowsSta(db *sql.DB) *sql.Rows { //mysqlからcenterの情報取得
    rows, err := db.Query("SELECT * FROM staffs")
    if err != nil {
        fmt.Println("Err2")
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
        error := rows.Scan(&staffs.Id, &staffs.Center_id, &staffs.Name, &staffs.Email, &staffs.RoleId, &staffs.Status, &staffs.Rfid)
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
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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
    _, err = db.Exec("INSERT INTO staffs (center_id, name, email, roleId, status, rfid) VALUES (?, ?, ?, ?, ?, ?)", data.Center_id, data.Name, data.Email, data.RoleId, data.Status, data.Rfid)
    if err != nil {
        fmt.Println("insert error!")
    }
}

func putStaStatus(w http.ResponseWriter, r *http.Request) {
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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
    _, err = db.Exec("UPDATE staffs SET status = false where id = ?", data.Id)
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
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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
    dsn := "user:pass@tcp(mysql-db:3306)/aikon_db"
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

func main() {
    // mux := http.NewServeMux() //登録されたパターンのリストと各受信リクエストのURLをマッチさせ、URLに最も理解パターンのハンドラを呼び出す
    // mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     w.Header().Set("Content-Type", "application/json") //ヘッダーの設定
    //     w.Write([]byte("{\"hello\": \"world\"}"))
    // })
    // handler := cors.Default().Handler(mux)
    // c := cors.New(cors.Options{
    //     AllowedOrigins: []string{"http://localhost:3000", "http://app-next:3000"},
    //     AllowCredentials: true,
    //     // Enable Debugging for testing, consider disabling in production
    //     Debug: true,
    // })
    // // Insert the middleware
    // handler = c.Handler(handler)
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/dailyReportGet", getDailyReport)
    http.HandleFunc("/dailyReportPost", postDailyReport)
    http.HandleFunc("/middleGet", getMiddle)
    http.HandleFunc("/middlePost", postMiddle)
    http.HandleFunc("/teacherMessageGet", getTeacherMessage)
    http.HandleFunc("/teacherMessagePost", postTeacherMessage)
	http.HandleFunc("/centerGet", getCenter)
    http.HandleFunc("/centerPost", postCenter)
    http.HandleFunc("/paymentGet", getPayment)
    http.HandleFunc("/paymentPost", postPayment)
    http.HandleFunc("/studentsGet", getStudents)
    http.HandleFunc("/studentPost", postStudent)
    http.HandleFunc("/stuStatustPut", putStuStatus)
    http.HandleFunc("/staffsGet", getStaffs)
    http.HandleFunc("/staffPost", postStaff)
    http.HandleFunc("/staStatustPut", putStaStatus)
    http.HandleFunc("/sensorsGet", getSensors)
    http.HandleFunc("/sensorPost", postSensor)
    http.HandleFunc("/inAndOutGet", getInAndOut)
    http.HandleFunc("/inAndOutPost", postInAndOut)
    // fmt.Println("Server Start")
    // http.ListenAndServe(":8080", handler)
    http.ListenAndServe(":8080", nil)
    fmt.Println("Hello, World!!")
}