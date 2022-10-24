/* eslint react-hooks/exhaustive-deps: off */ 
import type { NextPage } from 'next'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { auth } from '../../firebaseConfig'
import axios from "axios";
import Layout from "./parent-layout";

import Router from 'next/router';
import EmailTwoToneIcon from '@mui/icons-material/EmailTwoTone';
import AccessAlarmTwoToneIcon from '@mui/icons-material/AccessAlarmTwoTone';

import internal from 'stream';


interface studentObj {
  [key: string]: Array<Arr>
}
interface Arr {
  Datetime(Datetime: any): unknown;
  Array : Object
}

interface inOut {
  Id: number
  Datetime: Date
  Rfid: string
  Sensor_id: number
}

export async function getServerSideProps() { //ssg
  const res = await axios.get(`${process.env.API}/getStudents`, {
  });
  const students = await res.data;
  const inOutRes = await axios.get(`${process.env.API}/getStuInAndOutSensors`, {
  });
  const inOut = await inOutRes.data;

  return { 
      props: {
        students:students,
        inOut: inOut
      },
  };
}


const List: NextPage = (students: studentObj) => {
  const router = useRouter()
  const [currentUser, setCurrentUser] = useState<null | object>(null)

  useEffect(() => {
    auth.onAuthStateChanged((user) => {
      user ? setCurrentUser(user) : router.push('/parents/parent-login')
    })
  }, [])

  function checkStudentIn() { //引数students.inOut[0].Datetimeにすると問答無用で[0]読んで、それはないって怒られるから関数の中で呼んだ
    if(students.inOut !== null){
      return students.inOut[0].Datetime
    }
    return 
  }
  function checkStudentOut() {
    if(!students.inOut){
      return 
    }
    else if(!students.inOut[1]) {
      return 
    } else {
      students.inOut[1].Datetime
    }
  }


  return (
    <>
    {/* <div className='nyusitsu'>
      入室：{checkStudentIn()}
      <br></br>
      退室：{checkStudentOut()}
    </div> */}
    <Layout>
    <p className='parentmypagetitle'>ホーム</p>
      <div className='myname'>
            {students.students.map((d:any, i:number) => {
            return (
                <div className='namecenter' key={i}>
                  {d.CenterName}
                  <br></br>
                  {d.Name}
                </div>
            )
          })}
      </div>
      <div>
        <div>
          <button className='buttonp1' onClick={() => Router.push('/parents/parent-daily', '/parents/parent-daily', { shallow: true})}><AccessAlarmTwoToneIcon style={{ color: "white", fontSize: 64  }}/></button>
          <p className='pp1'>出欠を連絡する</p>
        </div>
        <div>
          <button className='buttonp2' onClick={() => Router.push('/parents/parent-fromCenter', '/parents/parent-fromCenter', { shallow: true})}><EmailTwoToneIcon
          style={{color: "white" , fontSize: 64 }}/></button> 
          <p className='pp2'>学童からの連絡を見る</p>
        </div>
        </div>
    </Layout>
    </>
  )
}

export default List