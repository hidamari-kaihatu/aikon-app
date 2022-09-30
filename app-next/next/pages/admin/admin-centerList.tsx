import type { NextPage } from 'next'
import axios from "axios";
import { Axios } from '../../lib/api';
import Link from 'next/link';
import Layout from './admin-layout';

export interface centers {
    data: center[];
}

export interface center {
    id: number;
    name: string;
    status: boolean;
}
export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/centerGet`, {
  });
//   const res = await axios.get(`http://localhost:8080/centerGet`, {
// });
  console.log("test1");
  const data = await res.data;
  console.log("data: ", data);
  
  return { 
    props: {
      data: data 
    }
  };
  window.location.reload()
}

const CenterGet: NextPage = ({data}:any) => {
  console.log("test2");
  console.log(data[0].Name);

  let n = 1;
  function counter(){
    console.log("test3");
    return n++;
  };
  //TODO:ボタンクリックで警告文のポップアップにする
  //Stripeのサブスク解除の画面にも飛ぶようにする
  function liftButton (id:any) {
       console.log("id :", id);
       console.log("test4");
        // e.preventDefault();
        const center = {id: id};//任意のIdにしたい
        console.log("center: ", center);
        Axios.put(`api/proxy/centerPut`, center)
        .then((res) => {
          console.log("center: ", center);
          console.log("test5");
          console.log("res : ",res);
        })
        // .then(() => {
          //   console.log("testreload");
          //   location.reload()
          // })
          .catch((error) => {
            console.log("test6");
            console.log(error,center);
            
          });
          
        //TODO:リロード処理(現状、２回クリックでしかリロードされない)
        //window.location.reload()

      }

  return (
    <Layout>
      <table className='list-table'>
        <thead>
          <tr>
            <th>No.</th>
            <th>施設名</th>
            <th>登録解除</th>
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
                    <div>
                    <Link href={"/admin/admin-centerPost"}>
                    <a><button>施設の新規登録</button></a>
                    </Link>
                    </div>
    </Layout>
      )
    }
    
    export default CenterGet
 