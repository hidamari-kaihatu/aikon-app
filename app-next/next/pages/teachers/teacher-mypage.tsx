/* eslint react-hooks/exhaustive-deps: off */ 
import type { NextPage } from 'next'
import Layout from './teacher-layout';
import Router from 'next/router';

    return (
      <>
      <Layout>
      <h3>先生HOME</h3>
      <div>
        <div>
          <button onClick={() => Router.push('/teacher/teacher-todayKids', '/teacher/teacher-todayKids', { shallow: true})}>今日の児童一覧を見る</button>
        </div>
        <div>
          <button onClick={() => Router.push('/teacher/teacher-kidsinfo', '/teacher/teacher-kidsinfo', { shallow: true})}>児童名簿一覧を見る</button> 
        </div>
        <div>
          <button onClick={() => Router.push('/teacher/sendMsg', '/teacher/sendMsg', { shallow: true})}>保護者メッセージを送る</button> 
        </div>
        <div>
          <button onClick={() => Router.push('/teacher/teacher-msgList', '/teacher/teacher-msgList', { shallow: true})}>保護者に送ったメッセージを見る</button>         
        </div>
        <div>
          <button onClick={() => Router.push('/teacher/teacher-childInOut', '/teacher/teacher-newer', { shallow: true})}>子供の入退室を確認するS</button>         
        </div>
        <div>
          <button onClick={() => Router.push('/teacher/teacher-teacherInOut', '/teacher/teacher-teacherInOut', { shallow: true})}>先生の出退勤を確認する</button>         
        </div>
      </div>
    </Layout>
      </>
    );
}
export default Mypage