import React,{ useState } from "react";
import { Axios } from '../../lib/api';
import Layout from "./parent-layout";

export default function dailyReportPost() {
    const [attend, setattend] = useState("");
    const [temperature, settemperature] = useState("");
    const [someToPickup, setsomeToPickup] = useState("");
    const [timeToPickup, settimeToPickup] = useState("");
    const [message, setmessage] = useState("");

    const today = new Date();
    const formatted = today.toLocaleDateString("ja-JP", {
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
    })
    .split("/")
    .join("-");

    //ここの今日の日時と曜日を出すところは、コンポーネント化できる。
    const year = today.getFullYear()
    const month = today.getMonth() + 1
    const day = today.getDate()
    const week = today.getDay()
    const weekItems = ["日", "月", "火", "水", "木", "金", "土"]
    const dayOfWeek = weekItems[week]

    //未完：Student_idでidをGETする必要がある
    function handleSubmit(e:any) {
        e.preventDefault();
        const data = {
            "Date":formatted,
            "Student_id":2,
            "Attend":JSON.parse(attend),
            "Temperature":temperature,
            "SomeoneToPickUp":someToPickup,
            "TimeToPickUp":timeToPickup,
            "Message":message
        }
        Axios.post(`api/proxy/dailyReportPost`, data)
        .then((res) => {
          console.log(res);
        })
        .catch((error) => {
          console.log(error);
        });
        //window.location.reload()
      }

    return(
      <>
      <Layout>
        <div>
         <h2>日々の出欠報告</h2>
            <br></br>
            <label>出欠: </label>
            <select value={attend} onChange={(e) => setattend(e.target.value)}>
                <option value="A">出欠</option>
                <option value={"true"}>学童に行きます</option>
                <option value={"false"}>学童に行きません</option>
            </select>
            <br></br>
            <br></br>
            <label>体温: </label>
            <select value={temperature} onChange={(e) => settemperature(e.target.value)}>
                <option value="A">体温</option>
                <option value="36.0">36.0</option>
                <option value="36.1">36.1</option>
                <option value="36.2">36.2</option>
                <option value="36.3">36.3</option>
                <option value="36.4">36.4</option>
                <option value="36.5">36.5</option>
                <option value="36.6">36.6</option>
                <option value="36.7">36.7</option>
                <option value="36.8">36.8</option>
                <option value="36.9">36.9</option>
                <option value="37.0">37.0</option>
                <option value="37.1">37.1</option>
                <option value="37.2">37.2</option>
                <option value="37.3">37.3</option>
                <option value="37.4">37.4</option>
                <option value="37.5">37.5</option>
            </select>
            <br></br>
            <br></br>
            <label>お迎えの人: </label>
            <input type="text" value={someToPickup} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setsomeToPickup(e.currentTarget.value)}/>
            <br></br>
            <br></br>
            <label>お迎えの時間: </label>
            <select value={timeToPickup} onChange={(e) => settimeToPickup(e.target.value)}>
                <option value="A">時間</option>
                <option value={"16:00"}>16:00</option>
                <option value={"16:30"}>16:30</option>
                <option value={"17:00"}>17:00</option>
                <option value={"17:30"}>17:30</option>
                <option value={"18:00"}>18:00</option>
                <option value={"18:30"}>18:30</option>
                <option value={"19:00"}>19:00</option>
            </select>
            <br></br>
            <br></br>
            <label>メッセージ：</label>
            <input type="text" value={message} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setmessage(e.currentTarget.value)}/>
            <br></br>
            <br></br>
            <button onClick = {handleSubmit}>送信する！</button>
        </div> 
      </Layout>    
      </>  
    )
}