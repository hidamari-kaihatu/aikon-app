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
    // "github.com/dgrijalva/jwt-go"
    
    _ "github.com/go-sql-driver/mysql"
)

func postDailyReport(w http.ResponseWriter, r *http.Request) {
    log.Printf("setCookie:",setCookie)

    id := resolveJWT()
    fmt.Printf("%v\n", id) //作成されたjwt確認

    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err1")
    }
    defer db.Close()
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

func getTeacherMessageRows(db *sql.DB) *sql.Rows { //保護者ページで使う用
    id := resolveJWT()

    rows, err := db.Query("SELECT * FROM teacherMessage where teacherMessage.student_id = ?", id)
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
        error := rows.Scan(&teacherMessage.Id, &teacherMessage.Staff_id, &teacherMessage.Message, &teacherMessage.Datetime, &teacherMessage.Student_id)
        if error != nil {
            log.Printf("scan error")
        } else {
            resultTeacherMessage = append(resultTeacherMessage, teacherMessage)
        }
    }
    var buf bytes.Buffer 
    enc := json.NewEncoder(&buf) 
    if err := enc.Encode(&resultTeacherMessage); err != nil {
        log.Fatal(err)
    }
    log.Printf(buf.String())

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }
}


func getRowsStu(db *sql.DB) *sql.Rows { 
    // claims := jwt.MapClaims{}

    // token, err := jwt.ParseWithClaims(setCookie, claims, func(token *jwt.Token) (interface{}, error) {
    //     return []byte("secret"), nil
    // })
    // fmt.Printf("%v\n", token)

	// if err != nil {
    //     fmt.Println("verifyToken error")
	// }
    // for key, val := range claims {
    //     log.Printf("Key: %v, value: %v\n", key, val)
    //     fmt.Printf("%T\n", val)
    //     log.Printf("Verified matchId val: %v\n", val)
    // }
    // id := claims["sutudent"]
    id := resolveJWT()

    rows, err := db.Query("SELECT students.id, students.center_id, students.name, students.contactTell, students.grade, students.email, students.status, students.rfid, centers.name from students INNER JOIN centers ON  students.center_id = centers.id WHERE students.id=?", id)
    if err != nil {
        fmt.Println("Err2")
        panic(err.Error())
    }
    return rows
}

func getStudents(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()
    defer db.Close()
    rows := getRowsStu(db) 
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

func postStudent(w http.ResponseWriter, r *http.Request) {
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
    _, err = db.Exec("INSERT INTO students (center_id, name, contactTell, grade, email, rfid) VALUES (?, ?, ?, ?, ?, ?)", data.Center_id, data.Name, data.ContactTell, data.Grade, data.Email, data.Rfid)
    if err != nil {
        fmt.Println("insert error!")
    }
}


func getRowsStuInAndOutSensors(db *sql.DB) *sql.Rows {
    id := resolveJWT()

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
    rows := getRowsStuInAndOutSensors(db)
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

    _, err := fmt.Fprint(w, buf.String()) 
    if err != nil {
        return
    }
}

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


func parentIsLogin(w http.ResponseWriter, r *http.Request){ 
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
log.Printf("Verified authHeader: %T\n", authHeader)
idToken := strings.Replace(authHeader, "Bearer ", "", 1)

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
toString := strconv.Itoa(matchId) //文字列に変換

//jwtにのせる　認証情報が入ったJsonを加工（電子署名を加える等）し、JWTにしたのち、それを認証Tokenとしてクッキーに渡す
afterAuthJwt := CreateToken(toString)
log.Printf("作成されたjwt:", afterAuthJwt) 

setCookie = afterAuthJwt

cookie := &http.Cookie{
	Name:   "studentID",
	Value:  afterAuthJwt,
	HttpOnly: true,
}

http.SetCookie(w, cookie)
log.Printf("cookie: ", cookie) 
}