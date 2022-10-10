/* eslint-disable */
import React,{ useEffect, useState } from "react";
import { Axios } from '../../lib/api';
import Link from 'next/link';
import { auth } from '../../firebaseConfig'
import { useRouter } from 'next/router';
import Layout from "./parent-layout";
import axios from "axios";
import SendTwoToneIcon from '@mui/icons-material/SendTwoTone';

export default function dailyReportPost(students:any) {
    const [attend, setattend] = useState("");
    const [temperature, settemperature] = useState("");
    const [someToPickup, setsomeToPickup] = useState("");
    const [timeToPickup, settimeToPickup] = useState("");
    const [message, setmessage] = useState("");

    const router = useRouter()
    const [currentUser, setCurrentUser] = useState<null | object>(null)
  

  const studentId = students.students[0].Id
  //{console.log(studentId)}

    useEffect(() => {
      auth.onAuthStateChanged((user) => {
        user ? setCurrentUser(user) : router.push('/parents/parent-login')
      })
    }, [])
    const logOut = async () => {
      try {
        await auth.signOut()
        router.push('/parents/parent-login')
      } catch (error) {
        router.push('/parents/parent-login')
      }
    }

    const today = new Date();
    const formatted = today.toLocaleDateString("ja-JP", {
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
    })
    .split("/")
    .join("-");


    //未完：Student_idでidをGETする必要がある。今はベタ打ちの「２」
    function handleSubmit(e:any) {
        e.preventDefault();
        const data = {
            "Date":formatted,
            "Student_id":studentId,
            "Attend":JSON.parse(attend),
            "Temperature":String(temperature),
            "SomeoneToPickUp":someToPickup,
            "TimeToPickUp":timeToPickup,
            "Message":message
        }
        Axios.post(`api/proxy/dailyReportPost`, data/*, { withCredentials: true }*/)
        .then((res) => {
          console.log(res);
        })
        .catch((error) => {
          console.log(error);
        });
        window.location.reload()
      }

    return(
      <>
      <Layout>
      <p className='fromcentertitle'>出欠を連絡する</p>
        <div className="bluebackpd">
        <div>
{/*          <h2>日々の出欠報告</h2> */}
            <br></br>
            <label className="sisetsu">出欠: </label>
            <select value={attend} onChange={(e) => setattend(e.target.value)}>
                <option value="A">出欠</option>
                <option value={"true"}>学童に行きます</option>
                <option value={"false"}>学童に行きません</option>
            </select>
            <br></br>
            <br></br>
            <label className="sisetsu">体温: </label>
            <input type="text" value={temperature} onChange={(e: React.ChangeEvent<HTMLInputElement>) => settemperature(e.currentTarget.value)}/>
            <br></br>
            <br></br>
            <label className="sisetsu">お迎えの人: </label>
            <input type="text" value={someToPickup} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setsomeToPickup(e.currentTarget.value)}/>
            <br></br>
            <br></br>
            <label className="sisetsu">帰り時間: </label>
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
            <label className="sisetsu">メッセージ：</label>
            <input type="text" value={message} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setmessage(e.currentTarget.value)}/>
            <br></br>
            <br></br>
{/*             <button className="btn" onClick={logOut}>Logout</button> */}
            </div>
            </div>
            <button className="buttonpdlog" onClick={logOut}>Logout</button>
            <button className="buttonpdsub" onClick = {handleSubmit}>< SendTwoToneIcon style={{ color: "white" , fontSize: 40 }} /></button><p className="pdsub">送信</p>
          </Layout>
        </>       
    )
}

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/studentsGet`, {
  });
  const students = await res.data;
  {console.log(students)}

  return { 
      props: {
        students
      },
  };
}
