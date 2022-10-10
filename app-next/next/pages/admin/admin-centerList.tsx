import type { NextPage } from 'next'
import axios from "axios";
import Layout from './admin-layout';
import test from 'node:test';
import Add from './add-base-Modal';
import Lift from './lift-base-Modal';

export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/centerGet`, {
  });

  console.log("test1");
  const data = await res.data;
  console.log("data: ", data);
  
  return { 
    props: {
      data: data 
    }
  };
}

const CenterGet: NextPage = ({data}:any) => {
  console.log("test2");
  console.log(data[0].Name);

  let n = 1;
  function counter(){
    console.log("test3");
    return n++;
  };
  
  return (
    <Layout>
      <div>
        <Add />
      </div>
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
             if(item.Status === 1){
               {console.log(item)}
                return (
                    <>
                    <tr  key={i}>
                     <td>{counter()}</td>
                     <td>{item.Name}</td>
                     <td><Lift/></td>
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
    
    export default CenterGet
 