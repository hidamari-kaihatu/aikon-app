import Layout from "./teacher-layout";
import axios from "axios";
import { Axios } from '../../lib/api';
import React,{ useState } from "react";

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/dailyReportGet`, {
  });
  const data = await res.data;

  return { 
      props: {
        data: data,
      },
  };
}

const StudentGet = ({data}:any) => {
console.log(data)
/* let n = 1;
function counter(){
  return n++;
}; */

function liftButton (id:any) {
      // e.preventDefault();
      const teacher = {id: id};
      Axios.put(`api/proxy/stuStatustPut`, teacher)
      .then((res) => {
      })
      .catch((error) => {
      });
      window.location.reload()
    }

return (
<Layout>
    <table className='list-table'>
      <thead>
        <tr>
          <th>ID</th>
          <th>Date</th>
          <th>Student_ID</th>
          <th>Attend</th>
          <th>Temperature</th>
          <th>topickup</th>
          <th>timepickup</th>
          <th>message</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item:any, i:number) => {
              return (
                  <>
                  <tr  key={item}>
                   <td>{item.Id}</td>
                   <td>{item.Date}</td>
                   <td>{item.Student_id}</td>
                   <td>{item.Attend.toString()}</td>
                   <td>{item.Temperature}</td>
                   <td>{item.SomeoneToPickUp}</td>
                   <td>{item.TimeToPickUp}</td>
                   <td>{item.Message}</td>
                   <td> <button onClick = {() => {liftButton(item.Id)}}>除籍</button></td>
                  </tr>
                  </>   
              )
      }
    )}
      </tbody>
    </table>
    </Layout>
    )
  }
  
  export default StudentGet

