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

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/staffsGet`, {
  });
  const resCenter = await axios.get(`${process.env.API}/centerGet`, {
  });
  const resMiddle = await axios.get(`${process.env.API}/middleGet`, {
});
  console.log("test1");
  const data = await res.data;
  console.log("data: ", data);
  const centerData = await resCenter.data
  console.log("centerData: ", centerData);
  const middleData = await resMiddle.data
  console.log("middleData: ", middleData);

  return { 
      props: {
        data: data,
        centerData: data,
        middleData: data
      },
  };
}

const TeacherGet: NextPage = ({data}:any) => {
console.log("test2");
// console.log(staffsData[0].Name);
// const [centerName, setCenterName] = useState("");

let n = 1;
function counter(){
  console.log("test3");
  return n++;
};
// function selectCenter({centerData}:any) {
//     Axios.get(`api/proxy/centerPut`)
//         .then((res) => {
//           console.log("test5");
//           console.log("res : ",res);
//         })
//         .catch((error) => {
//           console.log("test6");
//           console.log(error);
          
//         });
//       }
//     {centerData.map((item:any, i:number) => {
//        console.log("item:", item);
//        if(item.Status === true){
//           {console.log(item)}
//            return (
//                <>
//                <label>施設名: </label>
//             <select value={centerName} onChange={(e) => setCenterName(e.target.value)}>
//                 <option value="A">出欠</option>
//                 <option value={"true"}>学童に行きます</option>
//                 <option value={"false"}>学童に行きません</option>
//             </select>
//                </>
//            )
//        }   
//     }
//     )}
// };
//TODO:ボタンクリックで警告文のポップアップにする
//TODO:除籍ボタンを押すとリロードされるようにする
function liftButton (id:any) {
     console.log("id :", id);
     console.log("test4");
      // e.preventDefault();
      const teacher = {id: id};
      console.log("teacher: ", teacher);
      Axios.put(`api/proxy/staStatustPut`, teacher)
      .then((res) => {
        console.log("teacher: ", teacher);
        console.log("test5");
        console.log("res : ",res);
      })
      .catch((error) => {
        console.log("test6");
        console.log(error,teacher);
      });
      window.location.reload()
    }

return (
  <Layout>
    {/* {selectCenter}
    <div><p>{centerData}</p></div> */}
    <table className='list-table'>
      <thead>
        <tr>
          <th>No.</th>
          <th>職員名</th>
          <th>除籍</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item:any, i:number) => {
           console.log("item:", item);
           if(item.Status === true){
             {console.log(item)}
              return (
                  <>
                  <tr  key={i}>
                   <td>{counter()}</td>
                   <td>{item.Name}</td>
                   <td> <button onClick = {() => {liftButton(item.Id)}}>解除</button></td>
                  </tr>
                  </>   
              )
           }
      }
    )}

      </tbody>
    </table>
  </Layout>
    )
  }
  
  export default TeacherGet
