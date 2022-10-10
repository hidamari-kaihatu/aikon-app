import Layout from './parent-layout';
import { useRouter } from 'next/router';
import axios from 'axios';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/teacherMessageGet`, {
  });
  const data = await res.data;

  return { 
    props: {
      data: data,
    },
  };
}

//ここに絞り込みの.filter(取得したteacher messageのうち、自分（保護者）の子供のstudent_idのものだけを抜き出す処理。)を書く。

export default function mypage({data}:any) {
    return (
      <>
      <Layout>
{/*         <h2>学童からの連絡一覧</h2> */}
<p className='fromcentertitle'>学童からの連絡を見る</p>
<div className='msgblueback'>
        {data.map((d) => (
          <div className='sendmsg'>
        <ul key={d.Id}>
            日付: {d.Datetime}<br/><br/>メッセージ: {d.Message}<br/><br/>
        </ul>
        </div>
        ))}
        </div>
        </Layout>
      </>
    );
