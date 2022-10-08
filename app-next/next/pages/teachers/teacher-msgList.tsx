import Layout from './teacher-layout';
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

//ここに絞り込みの.filter(取得したteacher messageのうち、自分が所属するcenter_idのものだけを抜き出す処理。)を書く。

export default function mypage({data}:any) {
    return (
      <>
      <Layout>
        <h2>保護者に送ったメッセージの一覧</h2>
        {data.map((d) => (
        <ul key={d.Id}>
            日付: {d.Datetime}<br/><br/>メッセージ: {d.Message}<br/><br/><hr/>
        </ul>
        ))}
        </Layout>
      </>
    );
} 