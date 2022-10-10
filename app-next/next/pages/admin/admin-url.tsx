import type { NextPage } from 'next'
import axios from "axios";
import { Axios } from '../../lib/api';
import Link from 'next/link';
import Layout from './admin-layout';
import test from 'node:test';
import { center } from './admin-centerList';

export async function getServerSideProps() {
  const resCenter = await axios.get(`${process.env.API}/centerGet`, {
  });
//   const res = await axios.get(`http://localhost:8080/centerGet`, {
// });
  console.log("test1");
  const centerData = await resCenter.data;
  console.log("centerData: ", centerData);
  
  return { 
    props: {
      centerData: centerData 
    }
  };
}
interface ShowURLProps {
  centerData: any,//typeof dataにするとdataに赤波線が入る
  copyUrl: (text: string) => void;
}

const ShowURL: NextPage = ({centerData,copyUrl}:ShowURLProps) => {
  console.log("test2");
  console.log(centerData[0].Name);

  let n = 1;
  function counter(){
    console.log("test3");
    return n++;
  };
  let num = 1;
  function counter2(){
    return num++;
  }
  copyUrl = (text) => {
    navigator.clipboard.writeText(text)
    .then(function() {
      console.log('Async: Copying to clipboard was successful!');

    }, function(err) {
      console.error('Async: Could not copy text: ', err);
    });
  }
  return (
    <Layout>
    <div>
      <h3>管理者用</h3>
      <p>あなたの使用するURLです</p>
      <table className='list-table'>
         <thead>
           <tr>
             <th>新規登録用URL</th>
             <th>ログインURL</th>
           </tr>
         </thead>
         <tbody>
           <tr>
             <td>
               `http://localhost:3000/register/admin`
                <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/register/admin`)}}></input>
             </td>
             <td>
               `http://localhost:3000/login/admin`
                <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/login/admin`)}}></input>
             </td>
           </tr>
         </tbody>
         </table>
    </div>
    <div>
      <h3>職員用</h3>
      <p>学童の職員が使用するURLです。各施設の職員へこのURLをお知らせください。</p>
      <table className='list-table'>
         <thead>
           <tr>
             <th>No.</th>
             <th>施設名</th>
             <th>新規登録用URL</th>
             <th>ログインURL</th>
           </tr>
         </thead>
         <tbody>
          {centerData.map((item:any, i:number) => {
             console.log("item:", item);
             if(item.Status === 1){
               {console.log(item)}
                return (
                    <>
                    <tr  key={i}>
                      <td>{counter()}</td>
                      <td>{item.Name}</td>
                      <td>
                        `http://localhost:3000/register/staff/{item.Id}`
                        <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/register/staff/${item.Id}`)}}></input>
                      </td>
                      <td>
                        `http://localhost:3000/login/staff/{item.Id}`
                        <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/login/staff/${item.Id}`)}}></input>
                      </td>
                    </tr>
                    </>
                  )
                }
           }
          )}
         </tbody>
       </table>
    </div>
    <div>
      <h3>保護者用</h3>
      <p>学童利用の保護者が使用するURLです。各施設から保護者へこのURLをお知らせください。</p>
      <table className='list-table'>
         <thead>
           <tr>
             <th>No.</th>
             <th>施設名</th>
             <th>新規登録用URL</th>
             <th>ログインURL</th>
           </tr>
         </thead>
         <tbody>
           {centerData.map((item:any, i:number) => {
              console.log("item:", item);
              if(item.Status === 1){
                {console.log(item)}
                  return (
                    <>
                    <tr  key={i}>
                      <td>{counter2()}</td>
                      <td>{item.Name}</td>
                      <td>
                        `http://localhost:3000/parents/parent-newer/{item.Id}`
                        <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/parents/parent-newer/${item.Id}`)}}></input>
                      </td>
                      <td>
                        `http://localhost:3000/parents/parent-login`
                        <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/login/${item.Id}`)}}></input>
                      </td>
                    </tr>
                    </>
                  )
              }
            }
           )}
         </tbody>
      </table>
    </div>
    </Layout>
  );
}

export default ShowURL


// import type { NextPage } from 'next'
// import axios from "axios";
// import Layout from './admin-layout';

// export async function getServerSideProps() {
//   const resCenter = await axios.get(`${process.env.API}/centerGet`, {
//   });
// //   const res = await axios.get(`http://localhost:8080/centerGet`, {
// // });
//   console.log("test1");
//   const centerData = await resCenter.data;
//   console.log("centerData: ", centerData);
  
//   return { 
//     props: {
//       centerData: centerData 
//     }
//   };
// }
// interface ShowURLProps {
//   centerData: any,//typeof dataにするとdataに赤波線が入る
//   copyUrl: (text: string) => void;
// }

// const ShowURL: NextPage = ({centerData,copyUrl}:ShowURLProps) => {
//   console.log("test2");
//   console.log(centerData[0].Name);

//   let n = 1;
//   function counter(){
//     console.log("test3");
//     return n++;
//   };
//   let num = 1;
//   function counter2(){
//     return num++;
//   }
//   copyUrl = (text) => {
//     navigator.clipboard.writeText(text)
//     .then(function() {
//       console.log('Async: Copying to clipboard was successful!');

//     }, function(err) {
//       console.error('Async: Could not copy text: ', err);
//     });
//   }
//   return (
//     <Layout>
//     <div>
//       <h3>管理者用</h3>
//       <p>あなたの使用するURLです</p>
//       <table className='list-table'>
//          <thead>
//            <tr>
//              <th>新規登録用URL</th>
//              <th>ログインURL</th>
//            </tr>
//          </thead>
//          <tbody>
//            <tr>
//              <td>
//                `http://localhost:3000/register/admin`
//                 <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/register/admin`)}}></input>
//              </td>
//              <td>
//                `http://localhost:3000/login/admin`
//                 <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/login/admin`)}}></input>
//              </td>
//            </tr>
//          </tbody>
//          </table>
//     </div>
//     <div>
//       <h3>職員用</h3>
//       <p>学童の職員が使用するURLです。各施設の職員へこのURLをお知らせください。</p>
//       <table className='list-table'>
//          <thead>
//            <tr>
//              <th>No.</th>
//              <th>施設名</th>
//              <th>新規登録用URL</th>
//              <th>ログインURL</th>
//            </tr>
//          </thead>
//          <tbody>
//           {centerData.map((item:any, i:number) => {
//              console.log("item:", item);
//              if(item.Status === true){
//                {console.log(item)}
//                 return (
//                     <>
//                     <tr  key={i}>
//                       <td>{counter()}</td>
//                       <td>{item.Name}</td>
//                       <td>
//                         `http://localhost:3000/register/staff/{item.Id}`
//                         <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/register/staff/${item.Id}`)}}></input>
//                       </td>
//                       <td>
//                         `http://localhost:3000/login/staff/{item.Id}`
//                         <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/login/staff/${item.Id}`)}}></input>
//                       </td>
//                     </tr>
//                     </>
//                   )
//                 }
//            }
//           )}
//          </tbody>
//        </table>
//     </div>
//     <div>
//       <h3>保護者用</h3>
//       <p>学童利用の保護者が使用するURLです。各施設から保護者へこのURLをお知らせください。</p>
//       <table className='list-table'>
//          <thead>
//            <tr>
//              <th>No.</th>
//              <th>施設名</th>
//              <th>新規登録用URL</th>
//              <th>ログインURL</th>
//            </tr>
//          </thead>
//          <tbody>
//            {centerData.map((item:any, i:number) => {
//               console.log("item:", item);
//               if(item.Status === true){
//                 {console.log(item)}
//                   return (
//                     <>
//                     <tr  key={i}>
//                       <td>{counter2()}</td>
//                       <td>{item.Name}</td>
//                       <td>
//                         `http://localhost:3000/register/staff/{item.Id}`
//                         <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/register/${item.Id}`)}}></input>
//                       </td>
//                       <td>
//                         `http://localhost:3000/login/staff/{item.Id}`
//                         <input type="button" value="コピー" onClick={()=>{copyUrl(`http://localhost:3000/login/${item.Id}`)}}></input>
//                       </td>
//                     </tr>
//                     </>
//                   )
//               }
//             }
//            )}
//          </tbody>
//       </table>
//     </div>
//     </Layout>
//   );
// }

// export default ShowURL