import Layout from './parent-layout';
import { useRouter } from 'next/router';
import axios from 'axios';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/teacherMessageGet`, {
  });
  const data = await res.data;
  const stuRes = await axios.get(`${process.env.API}/studentsGet`, {
  });
  const students = await stuRes.data;
  //{console.log(students)}

  //console.log(data)

  return { 
    props: {
      data: data,
      students: students
    },
  };
}

//ここに絞り込みの.filter(取得したteacher messageのうち、自分（保護者）の子供のstudent_idのものだけを抜き出す処理。)を書く。

export default function mypage({data, students}:any) {
    return (
      <>
      <Layout>
        <h2>学童からの連絡一覧</h2>
        {data.map((d) => (
        <ul key={d.Id}>
            日付: {d.Datetime}<br/><br/>メッセージ: {d.Message}<br/><br/><hr/>
        </ul>
        ))}
      <div>
        {students.map((d:any, i:number) => {
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
    );
} 

