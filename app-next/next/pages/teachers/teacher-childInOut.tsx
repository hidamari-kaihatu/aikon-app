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
{/*       <h2 className='center'>子どもの入退室</h2> */}
<div className='inoutblueback'>
      <p className='kensaku'>検索</p><input className='inputnewer88' type="text" value={searchWord} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setSearchWord(e.currentTarget.value)}/></div>
      <div className='inoutsearchblueback'><p className='kensaku' >検索結果</p>
      {
        inOutData.map(data => {
          if(searchWord === ""){
            return 
          }
          else if(((data.Name).includes(searchWord)) || (data.Datetime).includes(searchWord)){
            return (
              <div>
              {data.Name}<span> :  </span>
              {data.Datetime}<br></br>
              </div>
            )
            }else{
              return 
            }
          })
      }
</div>
      </div>
      <Layout>
        <div className='inoutoutput'>
          {inOutData.map((inOut, i) => {
            return (
              <div key={i} className='center'>
                {inOut.Name}<span> :  </span>
                {inOut.Datetime}
              </div>
            )
          })}
                  </div>
{/*           <div className='inoutstaffname'>
            {staffs[0].CenterName}
            <br></br>
            {staffs[0].Name}
          </div> */}
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

