/* eslint react-hooks/exhaustive-deps: off */ 
import type { NextPage } from 'next'
import axios from 'axios';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { auth } from '../../firebaseConfig';
import Layout from './teacher-layout';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/getStaffAndMiddle`, {
  });
  const staffs = await res.data;
  {console.log(staffs)}

  return { 
      props: {
        staffs
      },
  };
}

interface staffObj {
  [key: string]: Array<Arr>
}
interface Arr {
  Array : Object
}

const Mypage: NextPage = (staffs:staffObj) => {
  const [currentUser, setCurrentUser] = useState<null | object>(null)
  const router = useRouter()

  useEffect(() => {
    auth.onAuthStateChanged((user) => {
      user ? setCurrentUser(user) : router.push('/teachers/teacher-login')
    })
  }, [])

    return (
      <Layout>
        <h2>TEACHER MY PAGE</h2>
        <div>
            {staffs.staffs.map((s:any, i:number) => {
                return (
                    <div key={i}>
                      {s.CenterName}
                      <br></br>
                      {s.Name}
                    </div>
                )
                })}
        </div>
      </Layout>
    );
}
export default Mypage