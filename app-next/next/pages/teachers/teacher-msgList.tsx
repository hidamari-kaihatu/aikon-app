import Layout from './teacher-layout';
import axios from 'axios';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/teacherMessageGet`, {
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
  // const info = data
  // .filter(obj =>obj.Center_id === staffs[0].Center_id )

    return (
      <>
      <Layout>
{/*         <h2>保護者に送ったメッセージの一覧</h2> */}
        <div className='msgblueback'>
        {data.map((d) => (
                  <div className='sendmsg'>
        <ul key={d.Id}>
            日付: {d.Datetime}<br/><br/>メッセージ: {d.Message}<br/><br/><hr/>
        </ul>
              </div>
        ))}
{/*         {staffs[0].CenterName}
        <br></br>
        {staffs[0].Name} */}
              </div>
        </Layout>
      </>
    );
} 