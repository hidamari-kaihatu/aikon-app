import Layout from "./teacher-layout";
import axios from "axios";
import { Axios } from '../../lib/api';
import React,{ useState } from "react";

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/studentsGet`, {
  });
  const data = await res.data;

  return { 
    props: {
      data: data,
    },
  };
}

const StudentGet = ({data}:any) => {

//現状2回ダブルクリックしないと、変更が反映されない
function liftButton (Id:any) {
  const student = {Id: Id};
  Axios.put(`api/proxy/stuStatustPut`, student)
  .then((res) => {
  })
  .catch((error) => {
  });
  window.location.reload()
}

//statusがtrueのものだけ表示
//Center_idが今はベタ打ち。ここに、先生が所属するCenter_idが入るようにしたい
const info = data
.filter(obj => obj.Status.toString() === "true")
.filter(obj =>obj.Center_id === 1 )

return (
<Layout>
<div className="blueback">
    <table className='list-table'>
      <thead>
        <tr >
          <th>センターID</th>
          <th>名前</th>
          <th>学年</th>
          <th>緊急連絡先</th>
          <th>メールアドレス</th>
          <th>status</th>
          <th>除籍</th>
        </tr>
      </thead>
      <tbody>
        {info.map((item:any, i:number) => {
          return (
          <>
          <tr  key={item}>
            <td>{item.Center_id}</td>
            <td>{item.Name}</td>
            <td>{item.Grade}</td>
            <td>{item.ContactTell}</td>
            <td>{item.Email}</td>
            <td>{item.Status.toString()}</td>
            <td> <button onClick = {() => {liftButton(item.Id)}}>除籍</button></td>
            </tr>
            </>   
          )
        }
        )}
      </tbody>
    </table>
    </div>
    </Layout>
    )
  }
  
  export default StudentGet

