import type { NextPage } from 'next'
import Router from 'next/router'
import { useRouter } from 'next/router';
import axios from "axios";
import { Axios } from '../../lib/api';
// staffテーブル情報Get
export async function getServerSideProps() {
    const res = await axios.get(`${process.env.API}/staffsGet`, {
    });
    const data = await res.data;
    console.log("data:",data);
    return { 
        props: {
          data: data,
        },
    };
  }
  let Id:number ;
function confirm({data}:any) {
    const router = useRouter();
    console.log("staffName :", router.query.name);
    const staffName = router.query.name//OK
    console.log("data :",data);//stafftable

    const staff = data.forEach((obj:any) => {
        console.log(obj["Name"]);//OK
        if(obj["Name"] === staffName) {
           console.log("obj: ", obj["Id"]);//OK
        //新規登録者のstaff_idを取得
           Id = obj["Id"]
           console.log("Id: ",Id);
        }

      });
      console.log("Id: ",Id);
      const centerId = Number(router.query.center_id);
      console.log("router.query.center_id: ",router.query.center_id);//1
    //middleテーブルにpost
    const postData = {
        "Center_id":centerId,
        "Staff_id": Id,
        "Role_id": 1
      };
      console.log("postData :",postData);
      
      Axios.post(`api/proxy/middlePost`, postData)
      .then((res) => {
        console.log("post start");
        console.log(res.data);
      })
      .catch((error) => {
        console.log(error);
      });
    

    return (
        <div>
           <p>登録が完了しました！</p> 
           <p>ログイン画面へお進みください</p>
           <button onClick={() => Router.push('/teachers/teacher-login', '/teachers/teacher-login', { shallow: true})}>ログイン画面へ進む</button> 
        </div>
    );
}

export default confirm;