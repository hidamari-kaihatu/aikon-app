import axios from 'axios';
import type { NextPage } from 'next'
import { useState } from 'react';
import Layout from './teacher-layout';

export default function Mypage({students, staffs}:any) {
  const [searchWord, setSearchWord] = useState("");
  // {console.log(students)}
  const inOutData = students.filter(student => student.Center_id === staffs[0].Center_id).reverse()
  //  {console.log(inOutData)}

    return (
      <>
      <div className='center'>
      <h2 className='center'>子どもの入退室</h2>
      <h3>検索</h3><input type="text" value={searchWord} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setSearchWord(e.currentTarget.value)}/>
      <h3 >検索結果</h3>
      {
        inOutData.map(data => {
          if(searchWord === ""){
            return 
          }
          else if(((data.Name).includes(searchWord)) || (data.Datetime).includes(searchWord)){
            return (
              <>
              {data.Name}<span> :  </span>
              {data.Datetime}<br></br>
              </>
            )
            }else{
              return 
            }
          })
      }
      </div>
      <Layout>
        <div>
          {inOutData.map((inOut, i) => {
            return (
              <div key={i} className='center'>
                {inOut.Name}<span> :  </span>
                {inOut.Datetime}
              </div>
            )
          })}
          <div>
            {staffs[0].CenterName}
            <br></br>
            {staffs[0].Name}
          </div>
        </div>
        </Layout>
      </>
    );
} 

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/getStudentInAndOut`, {
  });
  const students = await res.data;

  const stassRes = await axios.get(`${process.env.API}/getStaffAndMiddleAndCenter`, {
  });
  const staffs = await stassRes.data;
  {console.log(staffs)}


  return { 
      props: {
        students: students,
        staffs: staffs
      },
  };
}
