import Layout from './admin-layout';
import React,{ useState } from "react";
import { Axios } from '../../lib/api';
import Link from 'next/link';

export default function addCenter() {
  const [newCenter, setNewCenter] = useState("");
  //TODO:Stripeの決済画面へページ遷移
  function handleSubmit(e:any) {
      e.preventDefault();
      const data = {
          "Name":newCenter,
      }
      Axios.post(`api/proxy/postCenter`, data)
      .then((res) => {
        console.log(res);
      })
      .catch((error) => {
        console.log(error);
      });
  }


  return(
    <Layout>
      <h1>新規施設登録画面</h1>
          <div>
          <label>施設の入力：</label>
          <input type="text" value={newCenter} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setNewCenter(e.currentTarget.value)}/>
          <br></br>
          <button onClick = {handleSubmit}>施設名を登録する</button>
          </div>
  </Layout>
         
  )
}
