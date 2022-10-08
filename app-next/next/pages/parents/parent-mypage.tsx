/* eslint react-hooks/exhaustive-deps: off */ 
import type { NextPage } from 'next'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { auth } from '../../firebaseConfig'
import axios from "axios";
import Layout from "./parent-layout";

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