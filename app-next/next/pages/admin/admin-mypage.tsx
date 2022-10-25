/* eslint-disable */
import axios from 'axios';
import Router, { useRouter } from 'next/router'
import { useEffect, useState } from 'react';
import Layout from './admin-layout';
import { auth } from '../../firebaseConfig'
import type { NextPage } from 'next'

import AddCircleTwoToneIcon from '@mui/icons-material/AddCircleTwoTone';
import LinkTwoToneIcon from '@mui/icons-material/LinkTwoTone';
import SupervisorAccountTwoToneIcon from '@mui/icons-material/SupervisorAccountTwoTone';
import MapsHomeWorkTwoToneIcon from '@mui/icons-material/MapsHomeWorkTwoTone';

interface staffObj {
  [key: string]: Array<Arr>
}
interface Arr {
  Array : Object
}

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/getStaffs`, {
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

  // console.log(staffs)

  return (
    <Layout>
{/*       <h3>管理者HOME</h3> */}
<div>
        <div>
          <button className='buttona1' onClick={() => Router.push('/admin/admin-centerList', '/admin/admin-centerList', { shallow: true})}><MapsHomeWorkTwoToneIcon
          style={{color: "white" , fontSize: 64 }}/></button><p className='pa1'>学童施設一覧</p>
        </div>
        <div>
          <button className='buttona2' onClick={() => Router.push('/admin/admin-yeacherList', '/admin/admin-teacherList', { shallow: true})}><SupervisorAccountTwoToneIcon
          style={{color: "white" , fontSize: 64 }}/></button><p className='pa2'>職員一覧</p>
        </div>
        <div>
          <button className='buttona3' onClick={() => Router.push('/admin/admin-url', '/admin/admin-url', { shallow: true})}><LinkTwoToneIcon
          style={{color: "white" , fontSize: 64 }}/></button><p className='pa3'>URL</p>
        </div>
        <div>      

          <button className='buttona4' onClick={() => Router.push('https://forms.gle/UWVFgSBWAFJT271M9', 'https://forms.gle/UWVFgSBWAFJT271M9',{ shallow: true})}><AddCircleTwoToneIcon
          style={{color: "white" , fontSize: 64 }}/></button><p className='pa4'>新規施設登録</p>        
     

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