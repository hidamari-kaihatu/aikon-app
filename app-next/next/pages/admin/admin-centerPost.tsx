import Layout from './admin-layout';

export default function addCenter () {
  return (
    <Layout>
      <h1>新規施設登録画面</h1>
    </Layout>
  );
}
// import React,{ useState } from "react";
// import { Axios } from '../../lib/api';
// import Link from 'next/link';


// export default function ItemPost() {
//     const [item, setitem] = useState("");

//     function handleSubmit(e:any) {
//         e.preventDefault();
//         const data = {
//             "Name":item,
//         }
//         Axios.post(`api/proxy/centerPost`, data)
//         .then((res) => {
//           console.log(res);
//         })
//         .catch((error) => {
//           console.log(error);
//         });
//     }


//     return(
//         <div>
//             <h1>学童の追加</h1>
//             <br></br>
//             <div>
//             <label>学童名</label>
//             <input type="text" value={item} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setitem(e.currentTarget.value)}/>
//             <br></br>
//             <button onClick = {handleSubmit}>追加！</button>
//             <Link href={"/"}>
//             <h2>HOME</h2>
//             </Link>
//             </div>
//         </div>
//     )
// }