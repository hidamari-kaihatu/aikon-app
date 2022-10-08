/* eslint-disable */
import axios from 'axios';
import Router, { useRouter } from 'next/router'
import { useEffect, useState } from 'react';
import Layout from './admin-layout';
import { auth } from '../../firebaseConfig'
import type { NextPage } from 'next'

interface staffObj {
  [key: string]: Array<Arr>
}
interface Arr {
  Array : Object
}

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

const adminHome: NextPage = (staffs: staffObj) => {
//export function adminHome (): JSX.Element {
  const router = useRouter()
  const [currentUser, setCurrentUser] = useState<null | object>(null)

  useEffect(() => {
    auth.onAuthStateChanged((user) => {
      user ? setCurrentUser(user) : router.push('/admin/admin-login')
    })
  }, [])

  console.log(staffs)

  return (
    <Layout>
      <h3>管理者HOME</h3>
      <div>
        <div>
          <button onClick={() => Router.push('/admin/admin-centerList', '/admin/admin-centerList', { shallow: true})}>施設一覧</button>
        </div>
        <div>
          <button onClick={() => Router.push('/admin/admin-yeacherList', '/admin/admin-teacherList', { shallow: true})}>職員一覧</button> 
        </div>
        <div>
          <button onClick={() => Router.push('/admin/url', '/admin/url', { shallow: true})}>URL</button> 
        </div>
        <div>
          <button onClick={() => Router.push('/admin/admin-newer', '/admin/admin-newer', { shallow: true})}>新規施設登録</button>         
        </div>
      </div>
      <div>
        {staffs.staffs.map((s:any, i:number) => {
          return (
            <div key={i}>
              {s.Name}
            </div>
            )
          })}
            </div>
    </Layout>
  );
}

export default adminHome 