import axios from "axios";
import React,{ useState } from "react";
import Layout from './teacher-layout';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/dailyReportGet`, {
  });
  const data = await res.data;

  return { 
      props: {
        data: data,
      },
  };
}

const StudentGet = ({data}:any) => {
  let n = 1;
  function counter(){
    return n++;
  }; 

  let attend;

//ここにfilterの処理を書かないとだめ。すべてのdailyReportの中から、「学童の先生」の施設IDと生徒の施設IDが同じものだけを表示させるようにする
//dailyReportはstudent_idしかgetできないので、その情報をもとに、studentテーブルに行きstudent_idとcenter_idを照合する必要がある。


return (
<Layout>
    <table className='list-table'>
      <thead>
        <tr>
          <th>No.</th>
          <th>日付</th>
          <th>Student_ID</th>
          <th>出欠</th>
          <th>体温</th>
          <th>お迎えの人</th>
          <th>お迎えの時間</th>
          <th>メッセージ</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item:any, i:number) => {
              return (
                  <>
                  <tr  key={item}>
                   <td>{counter()}</td>
                   <td>{item.Date}</td>
                   <td>{item.Student_id}</td>
                   <td>{item.Attend.toString()}</td>
                   <td>{item.Temperature}</td>
                   <td>{item.SomeoneToPickUp}</td>
                   <td>{item.TimeToPickUp}</td>
                   <td>{item.Message}</td>
                  </tr>
                  </>   
              )
      }
    )}
      </tbody>
    </table>
    </Layout>
    )
  }
  
  export default StudentGet