/* eslint react-hooks/exhaustive-deps: off */ 
import type { NextPage } from 'next'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { auth } from '../../firebaseConfig'
import axios from "axios";
import Layout from "./parent-layout";
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
  const res = await axios.get(`${process.env.API}/studentsGet`, {
  });
  const students = await res.data;
  const inOutRes = await axios.get(`${process.env.API}/stuInAndOutSensorsGet`, {
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
    <div>
      入室：{checkStudentIn()}
      <br></br>
      退室：{checkStudentOut()}
    </div>
    <Layout>
      <div>
            {students.students.map((d:any, i:number) => {
            return (
                <div key={i}>
                  {d.CenterName}
                  <br></br>
                  {d.Name}
                </div>
            )
          })}
      </div>
    </Layout>
    </>
  )
}

export default List