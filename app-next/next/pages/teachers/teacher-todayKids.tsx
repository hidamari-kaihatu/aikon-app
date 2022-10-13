import axios from "axios";
import React,{ useState } from "react";
import Layout from './teacher-layout';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/dailyReportGet`, {
  });
  const data = await res.data;

  const staffsRes = await axios.get(`${process.env.API}/getStaffAndMiddleAndCenter`, {
  });
  const staffs = await staffsRes.data;

  return { 
      props: {
        data: data,
        staffs: staffs
      },
  };
}

const StudentGet = ({data, staffs}:any) => {
  let n = 1;
  function counter(){
    return n++;
  }; 

  let attend;

//ここにfilterの処理を書かないとだめ。すべてのdailyReportの中から、「学童の先生」の施設IDと生徒の施設IDが同じものだけを表示させるようにする →施設名同じでも？
//dailyReportはstudent_idしかgetできないので、その情報をもとに、studentテーブルに行きstudent_idとcenter_idを照合する必要がある。　→dailyReportsにcenter_idつけて、それを先生のcenter_idと照合
// console.log(data)
// console.log(staffs)
const today = new Date();
const year = today.getFullYear()
const month = today.getMonth() + 1
const day = today.getDate()
const schoolDay = [year, month, day];
// console.log(schoolDay)
const todaySchool = schoolDay.join("-");
// console.log(todaySchool)

const info = data
.filter(obj =>obj.Center_id  === staffs[0].Center_id)//idでふるいにかける ＋日付のfilter
.filter(obj =>obj.Date === todaySchool)
// {console.log(info)}
// const newInfo
// .filter(obj =>obj.Center_id === staffs[0].CenterName )//学童名でふるいにかける

function chanteAttendType(int: number) {
  if ( int === 0) {
    return "欠席"
  }
  return "出席"
}
return (
  <>
<Layout>
<div  className="blueback2">
    <table className='list-table'>
      <thead>
        <tr>
          <th>No.</th>
          <th>日付</th>
          <th>名前</th>
          <th>出欠</th>
          <th>体温</th>
          <th>お迎えの人</th>
          <th>お迎えの時間</th>
          <th>メッセージ</th>
        </tr>
      </thead>
      <tbody>
        {info.map((item:any, i:number) => {
              return (
                  <>
                  <tr  key={item}>
                   <td>{counter()}</td>
                   <td>{item.Date}</td>
                   <td>{item.Student_name}</td>
                   <td>{chanteAttendType(item.Attend)}</td>
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
    </div>
{/*     {staffs[0].CenterName}
      <br></br>
    {staffs[0].Name} */}
    </Layout>
    </>
    )
  }
  
  export default StudentGet