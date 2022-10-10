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

interface studentObj {
  [key: string]: Array<Arr>
}
interface Arr {
  Array : Object
}

export async function getServerSideProps() { //ssg
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


const List: NextPage = (students: studentObj) => {
  const router = useRouter()
  const [currentUser, setCurrentUser] = useState<null | object>(null)

  useEffect(() => {
    auth.onAuthStateChanged((user) => {
      user ? setCurrentUser(user) : router.push('/parents/parent-login')
    })
  }, [])
  console.log(students)

  return (
    <>
    <Layout>
      <div className='myname'>
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