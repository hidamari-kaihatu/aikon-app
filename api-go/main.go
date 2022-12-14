package main

import (
    "net/http"
    "log"

    "github.com/comail/colog"
)

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
    http.HandleFunc("/getDailyReport", getDailyReport) 
    http.HandleFunc("/dailyReportPost", postDailyReport)
    http.HandleFunc("/getMiddle", getMiddle)
    http.HandleFunc("/postMiddle", postMiddle)
    http.HandleFunc("/getTeacherMessage", getTeacherMessage)
    http.HandleFunc("/postTeacherMessage", postTeacherMessage)
	http.HandleFunc("/centerGet", getCenter)
    http.HandleFunc("/postCenter", postCenter)
    http.HandleFunc("/putCenterStatus", putCenterStatus)
    http.HandleFunc("/putCenterProductId", putCenterProductId)
    http.HandleFunc("/getStudents", getStudents)
    http.HandleFunc("/postStudent", postStudent)
    http.HandleFunc("/putStuStatus", putStuStatus)
    http.HandleFunc("/stuRfidPut", putStuRfid)
    http.HandleFunc("/staffPost", postStaff)
    http.HandleFunc("/getAllStaffs", getAllStaffs)
    http.HandleFunc("/putStaStatus", putStaStatus)
    http.HandleFunc("/putStaRfid", putStaRfid)
    http.HandleFunc("/getStuInAndOutSensors", getStuInAndOutSensors)
    http.HandleFunc("/parentIsLogin", parentIsLogin)
    http.HandleFunc("/staffIsLogin", staffIsLogin)
    http.HandleFunc("/getStaffAndMiddleAndCenter", getStaffAndMiddleAndCenter)
    http.HandleFunc("/getAllStudents", getAllStudents)
    http.HandleFunc("/getStudentInAndOut", getStudentInAndOut)
    http.HandleFunc("/getTeacherMessageForTeacher", getTeacherMessageForTeacher)
    http.HandleFunc("/getStaffs", getStaffs)
    http.ListenAndServe(":8080", nil)
    
}