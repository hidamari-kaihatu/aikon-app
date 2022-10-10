import Layout from './teacher-layout';
import axios from 'axios';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/getTeacherMessageForTeacher`, {
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

//ここに絞り込みの.filter(取得したteacher messageのうち、自分が所属するcenter_idのものだけを抜き出す処理。)を書く。


export default function mypage({data, staffs}:any) {
  const info = data
  .filter(obj =>obj.Center_id === staffs[0].Center_id )
  console.log(data)
  console.log(info)
    return (
      <>
      <Layout>
        <h2>保護者に送ったメッセージの一覧</h2>
        {info.map((i) => (
        <ul key={i.Id}>
            日付: {i.Datetime}<br/><br/>宛先：{i.Student_name}<br/><br/>メッセージ: {i.Message}<br/><br/><hr/>
        </ul>
        ))}
        {staffs[0].CenterName}
        <br></br>
        {staffs[0].Name}
        </Layout>
      </>
    );
} 