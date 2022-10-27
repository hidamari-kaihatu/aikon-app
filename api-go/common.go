package main

import (
    "fmt"
    "log"
    "database/sql"
    "os"
    "time"

    "github.com/joho/godotenv"
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

//データベースに接続する部分
func connectionDB() *sql.DB {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
    fmt.Println(dsn)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err1")
    }
    return db
}

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
    var secretKey = "XXXXXXXX" 
    tokenString, err := token.SignedString([]byte(secretKey)) 
    if err != nil {
        log.Printf("signiture error") //return "", err
    }
    return tokenString
}

func resolveJWT() interface{} {
    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(setCookie, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("XXXXXXXX"), nil
    })
     fmt.Printf("%v\n", token)

	if err != nil {
        fmt.Println("verifyToken error")
	}
    for _, val := range claims {
        // log.Printf("Key: %v, value: %v\n", key, val)
        // fmt.Printf("%T\n", val)
        log.Printf("Verified matchId val: %v\n", val)
    }
    return claims["sutudent"]
}