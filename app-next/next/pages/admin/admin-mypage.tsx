import Router from 'next/router'
import Layout from './admin-layout';

export default function adminHome () {
  return (
    <Layout>
      <h3>管理者HOME</h3>
      <div>
        <div>
          <button onClick={() => Router.push('/admin/admin-centerList', '/admin/admin-centerList', { shallow: true})}>施設一覧</button>
        </div>
        <div>
          <button onClick={() => Router.push('/admin/admin-yeacherList', '/admin/admin-teacherList', { shallow: true})}>職員一覧</button> 
        </div>
        <div>
          <button onClick={() => Router.push('/admin/url', '/admin/url', { shallow: true})}>URL</button> 
        </div>
        <div>
          <button onClick={() => Router.push('/admin/admin-newer', '/admin/admin-newer', { shallow: true})}>新規施設登録</button>         
        </div>
      </div>
    </Layout>
  );
}

