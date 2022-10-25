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
    "strconv"

    "firebase.google.com/go/v4"
    "github.com/k-washi/jwt-decode/jwtdecode"
    "github.com/dgrijalva/jwt-go"

    _ "github.com/go-sql-driver/mysql"
)

func getDailyReportsRows(db *sql.DB) *sql.Rows {
    rows, err := db.Query("SELECT dailyReports.id, dailyReports.date, dailyReports.student_id, dailyReports.attend, dailyReports.temperature, dailyReports.someoneToPickUp, dailyReports.timeToPickUp, dailyReports.message, students.center_id, students.name FROM dailyReports INNER JOIN students ON students.center_id WHERE dailyReports.student_id = students.id")
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
        error := rows.Scan(&dailyReports.Id, &dailyReports.Date, &dailyReports.Student_id, &dailyReports.Attend, &dailyReports.Temperature, &dailyReports.SomeoneToPickUp, &dailyReports.TimeToPickUp, &dailyReports.Message, &dailyReports.Center_id, &dailyReports.Student_name)
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

type Middle struct {
    Id int `json:id`
    Staff_id int `json:staff_id`
    Center_id *int `json:center_id`
    Role_id int `json:role_id`
}

//admin
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
    var buf bytes.Buffer //バッファを作成　
    enc := json.NewEncoder(&buf) //書き込み先を指定、 &buf のように必ずポインタを渡すことに注意
    if err := enc.Encode(&resultMiddle);/*jsonにエンコード*/ err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) // json を返却
    if err != nil {
        return
    }
}

//teacher
func postMiddle (w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
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

//teacher
func postTeacherMessage (w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
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
    _, err = db.Exec("INSERT INTO teacherMessage (staff_id, message, student_id) VALUES (?, ?, ?)", data.Staff_id, data.Message, data.Student_id)
    log.Printf("postTeacherMessage end")
}

//admin
type Centers struct {
    Id int `json:id`
    Name string`json:name`
    Status int `json:status`
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
    var buf bytes.Buffer 
    enc := json.NewEncoder(&buf) 
    if err := enc.Encode(&resultCenter); err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }

    log.Printf("trace: this is a trace log getCenter end.")
}

func postCenter(w http.ResponseWriter, r *http.Request) {
    log.Printf("trace: this is a trace log postCenter start.")
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
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
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }

    jsonBytes := ([]byte)(b)
    data := new(Students)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("UPDATE centers SET status = 0 where id = ?", data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}
func putCenterProductId(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }

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

//teacher
func putStuStatus(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }

    jsonBytes := ([]byte)(b)
    data := new(Students)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("UPDATE students SET status = 0 where id = ?", data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}

func putStuRfid(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }

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
    Status int `json:status`
    Rfid *string `json:rfid`
}


//admin
func getRowsAllSta(db *sql.DB) *sql.Rows { 
    log.Printf("getRowsAllSta")
    rows, err := db.Query("SELECT * FROM staffs")
    if err != nil {
        fmt.Println("Err2 a")
        panic(err.Error())
    }
    return rows
}
func getAllStaffs(w http.ResponseWriter, r *http.Request) {
    log.Printf("getAllStaffs")
    db := connectionDB()
    defer db.Close()
    rows := getRowsAllSta(db) 
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
    var buf bytes.Buffer 
    enc := json.NewEncoder(&buf) 
    if err := enc.Encode(&resultStaffs); err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }
}

func postStaff(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }

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
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }

    jsonBytes := ([]byte)(b)
    data := new(Staffs)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
    _, err = db.Exec("UPDATE staffs SET status = 0 where id = ?", data.Id)
    if err != nil {
        fmt.Println("update error!")
    }
}

func putStaRfid(w http.ResponseWriter, r *http.Request) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("db connect error!")
    }
    defer db.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("io error")
        return
    }

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


//teacher
type StaffAndMiddleAndCenter struct {
    Id int `json:id`
    Name string `json;name`
    Email string `json:email`
    Status int `json;status`
    Rfid *string `json:rfid`
    Center_id *int `json:center_id`
    Role_id int `json:role_id`
    CenterName *string `json:centerName`
}

func getRowsStaffAndMiddleAndCenter(db *sql.DB) *sql.Rows {
    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(setCookie, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })
    fmt.Printf("%v\n", token)

	if err != nil {
        fmt.Println("verifyToken error")
	}
    for _, val := range claims {
        log.Printf("Verified matchId val: %v\n", val)
    }
    id := claims["sutudent"]

    rows, err := db.Query(`SELECT staffs.id, staffs.name, staffs.email, staffs.status, staffs.rfid, middle.center_id, middle.role_id, centers.name from staffs INNER JOIN middle ON staffs.id = middle.staff_id INNER JOIN centers ON middle.center_id = centers.id WHERE staffs.id = ?`,id)
    if err != nil {
        fmt.Println("getRowsStaffAndMiddle")
        panic(err.Error())
    }
    return rows
}

func getStaffAndMiddleAndCenter(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsStaffAndMiddleAndCenter(db)
    staffAndMiddleAndCenter := StaffAndMiddleAndCenter{}
    var resultStaffAndMiddleAndCenter [] StaffAndMiddleAndCenter
    for rows.Next() {
        error := rows.Scan(&staffAndMiddleAndCenter.Id, &staffAndMiddleAndCenter.Name, &staffAndMiddleAndCenter.Email, &staffAndMiddleAndCenter.Status, &staffAndMiddleAndCenter.Rfid, &staffAndMiddleAndCenter.Center_id, &staffAndMiddleAndCenter.Role_id, &staffAndMiddleAndCenter.CenterName)
        if error != nil {
            fmt.Println("scan error getStaffAndMiddleAndCenter")
        } else {
            resultStaffAndMiddleAndCenter = append(resultStaffAndMiddleAndCenter, staffAndMiddleAndCenter)
        }
    }
    var buf bytes.Buffer
    enc := json.NewEncoder(&buf)
    if err := enc.Encode(&resultStaffAndMiddleAndCenter); err != nil {
        log.Fatal(err)
    }
    log.Printf(buf.String())

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }
}

//teacher
func getRowsAllStu(db *sql.DB) *sql.Rows { 
    rows, err := db.Query("SELECT students.id, students.center_id, students.name, students.contactTell, students.grade, students.email, students.status, students.rfid, centers.name from students INNER JOIN centers ON  students.center_id = centers.id")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsAllStu(db) 
    students := Students{}
    var resultStudents [] Students
    for rows.Next() {
        error := rows.Scan(&students.Id, &students.Center_id, &students.Name, &students.ContactTell, &students.Grade, &students.Email, &students.Status, &students.Rfid, &students.CenterName)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultStudents = append(resultStudents, students)
        }
    }
    var buf bytes.Buffer 
    enc := json.NewEncoder(&buf) 
    if err := enc.Encode(&resultStudents); err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }
}

//teacher
type StudentInAndOut struct {
    Id int  `json:id`
    Center_id int `json:cencer_id`
    Datetime string  `json:datetime`
    Rfid string  `json:rfid`
    Sensor_id int  `json:sensor_id`
    Name string  `json:name`
    Place string  `json:place`
}

func getRowsStudentInAndOut(db *sql.DB) *sql.Rows {
    rows, err := db.Query(`SELECT students.id, students.center_id, inAndOut.datetime, students.rfid, inAndOut.sensor_id, students.name, sensors.place FROM students INNER JOIN inAndOut ON students.rfid = inAndOut.rfid INNER JOIN sensors ON inAndOut.sensor_id = sensors.id`)
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

func  getStudentInAndOut(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsStudentInAndOut(db)
    studentInAndOut := StudentInAndOut{}
    var resultStudentInAndOut [] StudentInAndOut
    for rows.Next() {
        error := rows.Scan(&studentInAndOut.Id, &studentInAndOut.Center_id, &studentInAndOut.Datetime, &studentInAndOut.Rfid, &studentInAndOut.Sensor_id, &studentInAndOut.Name, &studentInAndOut.Place)
        if error != nil {
            fmt.Println("scan error")
        } else {
            resultStudentInAndOut = append(resultStudentInAndOut, studentInAndOut)
        }
    }
    var buf bytes.Buffer
    enc := json.NewEncoder(&buf)
    if err := enc.Encode(&resultStudentInAndOut); err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }
}

//teacher
type TeacherMessageForTeacher struct {
    Id int `json:id`
    Staff_id int `json:staff_id`
    Message *string `json:message`
    Datetime string `json:datetime`
    Student_id int `json:student_id`
    Student_name string `json:student_name`
    Center_id int `json:center_id`
}

func getTeacherMessageRowsForTeacher(db *sql.DB) *sql.Rows { //先生ページで使う用
    rows, err := db.Query("SELECT teacherMessage.id, teacherMessage.staff_id, teacherMessage.message, teacherMessage.datetime, teacherMessage.student_id, students.name, students.center_id FROM teacherMessage INNER JOIN students ON teacherMessage.student_id = students.id")
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

func getTeacherMessageForTeacher(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getTeacherMessageRowsForTeacher(db) // 行データ取得
    teacherMessageForTeacher := TeacherMessageForTeacher{}
    var resultTeacherMessageForTeacher [] TeacherMessageForTeacher
    for rows.Next() {
        error := rows.Scan(&teacherMessageForTeacher.Id, &teacherMessageForTeacher.Staff_id, &teacherMessageForTeacher.Message, &teacherMessageForTeacher.Datetime, &teacherMessageForTeacher.Student_id, &teacherMessageForTeacher.Student_name, &teacherMessageForTeacher.Center_id)
        if error != nil {
            log.Printf("scan error")
        } else {
            resultTeacherMessageForTeacher = append(resultTeacherMessageForTeacher, teacherMessageForTeacher)
        }
    }
    var buf bytes.Buffer 
    enc := json.NewEncoder(&buf) 
    if err := enc.Encode(&resultTeacherMessageForTeacher); err != nil {
        log.Fatal(err)
    }
    log.Printf(buf.String())

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }
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
func staffIsLogin(w http.ResponseWriter, r *http.Request){ 
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
authHeader := r.Header.Get("Authorization")
idToken := strings.Replace(authHeader, "Bearer ", "", 1)

log.Printf("trace: this is a trace JWT の検証.")
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

user := payload.Subject
emailInToken := payload.Email
log.Printf("User ID: " + user + " ,Email: " + emailInToken + "\n")

//staffテーブルのid取得用
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
toString := strconv.Itoa(matchId) //文字列に変換

//jwtにのせる　認証情報が入ったJsonを加工（電子署名を加える等）し、JWTにしたのち、それを認証Tokenとしてクッキーに渡す
afterAuthJwt := CreateToken(toString)

setCookie = afterAuthJwt

cookie := &http.Cookie{
    Name:   "studentID",
    Value:  afterAuthJwt,
    HttpOnly: true,
}

http.SetCookie(w, cookie)
log.Printf("cookie: ", cookie) 
}
func getRowsSta(db *sql.DB) *sql.Rows { 
    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(setCookie, claims, func(token *jwt.Token) (interface{}, error) {
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
    rows := getRowsSta(db) 
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
    var buf bytes.Buffer 
    enc := json.NewEncoder(&buf) 
    if err := enc.Encode(&resultStaffs); err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }
}