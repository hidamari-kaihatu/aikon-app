import type { NextPage } from 'next'
import axios from "axios";
import { Axios } from '../../lib/api';
import Link from 'next/link';
import Layout from './admin-layout';
import React,{ useState } from "react";

// export interface staffs {
//     data: staff[];
// }

// export interface staff {
//     Id : number
//     Name : string
//     Email : string 
//     Status : boolean
//     Rfid : string 
// }

//職員情報と施設情報と中間情報をDBから取得
export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/staffsGet`, {});
  const resCenter = await axios.get(`${process.env.API}/centerGet`, {});
  const resMiddle = await axios.get(`${process.env.API}/middleGet`, {});
  //単なるdataは、staffData。名前を変えるとエラーになるので、現状dataのまま
  // console.log("test1");
  const data = await res.data;
  // console.log("data: ", data);
  const centerData = await resCenter.data
  // console.log("centerData: ", centerData);
  const middleData = await resMiddle.data
  // console.log("middleData: ", middleData);

  return { 
      props: {
        data: data,
        centerData: centerData,
        middleData: middleData
      },
  };
}

interface TeacherGetProps {
  data: any,//typeof dataにするとdataに赤波線が入る
  centerData: any,//typeof dataにするとdataに赤波線が入る
  middleData: any,//typeof dataにするとdataに赤波線が入る
  onChangeCenter: (event: any) => void;
}

// const TeacherGet: NextPage = (props: {
  //   data: any,//typeof dataにするとdataに赤波線が入る
  //   centerData: typeof data,
  //   middleData: typeof data
  // }) => {  
    const TeacherGet: NextPage = ({data,centerData,middleData,onChangeCenter}:TeacherGetProps) => {
      // const {data} = props
      // const {centerData} = props
      // const {middleData} = props
      // console.log("test2");
      const [newStaffDatas, setNewStaffDatas] = useState([]);
      
      // console.log("newStaffDatas : ", newStaffDatas);
      //選んだcenterData[i].Idと同じCenter_idを持ち、Role＿Idが１の職員を表示させる。
      onChangeCenter = (event:any) => {
        const center_id = Number(event.target.value);
        console.log("center_id :", center_id);//2
        console.log("center_id type:", typeof center_id);//number
        console.log("middleData: ", middleData);//OK
        //center_idでmiddleDataの絞り込み
        const newMiddleData = middleData.filter((obj: any) => {
          if(obj["Center_id"] === center_id && obj["Role_id"] === 1) {
            return obj;
          }
        });
        console.log("newmiddleData: ",newMiddleData);//OK
        //sfatt_idでdata(staffData)の絞り込み
        console.log("data: ", data);
        console.log("data type:", typeof data[0].Id);//number
        const result:any = [];
        for(const key in newMiddleData){
            console.log(newMiddleData[key]);
            const newStaffData = data.filter((obj:any) => {
            if(obj.Id === newMiddleData[key]["Staff_id"] && obj.Status === true) {
              result.push(obj);
              setNewStaffDatas(result);
              return obj;
            }
          });
          console.log("newStaffData : ", newStaffData);//1つのobjが入った配列
        }
        console.log("newStaffDatas : ", newStaffDatas);//該当するstaffのobjが全て入った配列

       }

let n = 1;
const counter = () => {
  // console.log("test3");
  return n++;
};
//セレクトボタンでSatusがtrueの施設名だけを表示させる
const selectCenter = centerData.map((item:any, i:number) => {
  if(item.Status === true){
    //  {console.log("item.Name : ",item.Name)}
      return(
        <option key={i} value={item.Id}>{item.Name}</option>
      )}
});
//TODO:アラート後、除籍が即時に画面に反映されるようにする
//NOW:アラートでOKボタンを押すと一覧が一度消える。そのあと、再び同じ施設を選ぶと除籍が反映されている。
const liftButton = (e:any,Id:any) => {
  e.preventDefault();
  window.alert('本当に除籍しますか？この職員の情報はこの画面から削除され、この職員がサービスにアクセスできなくなります。')
  const teacher = {Id: Id};
  console.log("teacher: ", teacher);
  Axios.put(`api/proxy/staStatustPut`, teacher)
  .then((res) => {
    console.log("res : ",res);
    window.alert('除籍しました。再度施設を選択し、確認してください')
    window.location.reload()

  })
  .catch((error) => {
    console.log(error,teacher);
  });

}

const showTable = newStaffDatas.map((item:any, i:number) => {
  console.log("item:", item);
    // if(newStaffDatas.length !== 0){
      return (
          <>
           <tr  key={i}>
            <td>{counter()}</td>
            <td>{item.Name}</td>
            <td>{item.Email}</td>
            {/* <td> <button onClick = {() => {liftButton(item.Id)}}>解除</button></td> */}
            <td> <button onClick = {(e:any) => {liftButton(e,item.Id)}}>解除</button></td>
           </tr>
          </>   
       )
    // }
  });

//TODO:ボタンクリックで警告文のポップアップにする
//TODO:除籍ボタンを押すとリロードされるようにする
// const liftButton = (id:any) => {
//       // e.preventDefault();
//       const teacher = {id: id};
//       console.log("teacher: ", teacher);
//       Axios.put(`api/proxy/staStatustPut`, teacher)
//       .then((res) => {
//         console.log("res : ",res);
//       })
//       .catch((error) => {
//         console.log(error,teacher);
//       });
//       window.location.reload()
//     }
return (
  <Layout>
    <div>
      <label>施設名: </label>
        <select name="centerData" id="centerData" onChange={(e) => onChangeCenter(e)}>
        <option  value="example">施設を選択してください</option>
        {selectCenter}
        </select>
    </div>
    <table className='list-table'>
      <thead>
        <tr>
          <th>No.</th>
          <th>職員名</th>
          <th>Email</th>
          <th>除籍</th>
        </tr>
      </thead>
      <tbody>
        {showTable}
      </tbody>
    </table>
  </Layout>
    )
  }
  
  export default TeacherGet