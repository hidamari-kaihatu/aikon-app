import type { NextPage } from 'next'
import axios from 'axios';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { auth } from '../../firebaseConfig';
import Layout from './teacher-layout';
import Router from 'next/router';

import AccessTimeTwoToneIcon from '@mui/icons-material/AccessTimeTwoTone';
import EmailTwoToneIcon from '@mui/icons-material/EmailTwoTone';
import MeetingRoomTwoToneIcon from '@mui/icons-material/MeetingRoomTwoTone';
import FaceTwoToneIcon from '@mui/icons-material/FaceTwoTone';
import ArticleTwoToneIcon from '@mui/icons-material/ArticleTwoTone';
import SendTwoToneIcon from '@mui/icons-material/SendTwoTone';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/getStaffAndMiddleAndCenter`, {
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
        <div className='teachername'>
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
        <div>
{/*          <p className='teachermypagetitle'>ホーム</p>  */}
        <div>
          <button className='button' onClick={() => Router.push('/teachers/teacher-todayKids', '/teachers/teacher-todayKids', { shallow: true})}><FaceTwoToneIcon style={{ color: "white", fontSize: 64  }}/></button>
          <p className='p1'>今日の児童一覧を見る</p>
        </div>
        <div>
          <button className='button2' onClick={() => Router.push('/teachers/teacher-kidsinfo', '/teachers/teacher-kidsinfo', { shallow: true})}><ArticleTwoToneIcon
          style={{color: "white" , fontSize: 64 }}/></button> 
          <p className='p2'>児童名簿一覧を見る</p>
        </div>
        <div>
          <button className='button3' onClick={() => Router.push('/teachers/teacher-sendMsg', '/teachers/teacher-sendMsg', { shallow: true})}><SendTwoToneIcon
          style={{color: "white" , fontSize: 64 }}/></button>
          <p className='p3'>メッセージを送る</p> 
        </div>
        <div>
          <button className='button4' onClick={() => Router.push('/teachers/teacher-msgList', '/teachers/teacher-msgList', { shallow: true})}>< EmailTwoToneIcon style={{ color: "white" , fontSize: 64 }} /></button>         
          <p className='p4'>メッセージを見る</p>
        </div>
        <div>
          <button className='button5' onClick={() => Router.push('/teachers/teacher-childInOut', '/teachers/teacher-childInOut', { shallow: true})}><MeetingRoomTwoToneIcon style={{ color: "white", fontSize: 64  }} /></button>
          <p className='p5'>子供の入退室を確認する</p>        
        </div>
        <div>
          <button className='button6' onClick={() => Router.push('/teachers/teacher-teacherInOut', '/teachers/teacher-teacherInOut', { shallow: true})}><AccessTimeTwoToneIcon style={{ color: "white", fontSize: 64  }}/></button>
          <p className='p6'>先生の出退勤を確認する</p>         
        </div>
      </div>
      </Layout>
    );
}
export default Mypage