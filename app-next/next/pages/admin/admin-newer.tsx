import type { NextPage } from 'next';
import Link from 'next/link';
import React,{ useState, useEffect } from "react";
import Router from 'next/router';
import { auth } from '../../firebaseConfig';
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {faEye, faEyeSlash} from "@fortawesome/free-solid-svg-icons";
import { Axios } from '../../lib/api';
import { useForm, SubmitHandler }from 'react-hook-form';
import { ErrorMessage } from "@hookform/error-message";
import isEmail from 'validator/lib/isEmail';
import { redirect } from 'next/dist/server/api-utils';

type FormData = {
  staffName: string,
  staffEmail: string,
  password: string,
  passwordCheck: string
}

//TODOリスト
//登録ボタンを押すと、staffテーブルにName ,EmailがPOSTされる。[OK]
//passwordがfirebaseに登録される[OK]
//未記入のものがあれば登録できないようにする。（未記入のものがありますメッセージ）
//メールアドレスに全角があったらエラーを出す
//パスワードが半角英数字で８文字以上でなかったら（半角英数字で８文字以上で入力してくださいメッセージ）
//パスワードとパスワード確認の内容が一致しなかったら（パスワードと同じものを再度入力してください）
//”登録しました”のポップアップ表示


const SignUp: NextPage = () => {
  const { register, setValue,trigger,getValues, handleSubmit, formState: { errors },
 } = useForm<FormData>();
  const onSubmit: SubmitHandler<FormData> = data => console.log(data);
  
  const [staffName, setStaffName] = useState("");
  const [staffEmail, setStaffEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passwordCheck, setPasswordCheck] = useState("");

  function Password (){
    // パスワード表示制御ようのstate
    const [isRevealPassword, setIsRevealPassword] = useState(false);
    const [isRevealRe_Password, setIsRevealRe_Password] = useState(false);

    const togglePassword = () => {
      setIsRevealPassword((prevState) => !prevState);
    }
    const toggleRe_Password = () => {
      setIsRevealRe_Password((prevState) => !prevState);
    }
    return (
      <>
        <div className='newer-form'>
          パスワード<span> ※必須</span><br></br>
         <input 
           placeholder="半角英数字で８文字以上"
           type={isRevealPassword ? 'text' : 'password'}
           name="password"
           value={password}
           onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.currentTarget.value)}
           />
          <span
	          onClick={togglePassword}
            role="presentation"
            style={{'color': 'black'}}
            >
           {isRevealPassword ? (
             <FontAwesomeIcon icon={faEye} />
             ) : (
               <FontAwesomeIcon icon={faEyeSlash} />
               )}
          </span><br></br>
         </div>
         <div className='newer-form'>
          パスワード確認<span> ※必須</span><br></br>
         <input 
           placeholder="半角英数字で８文字以上"
           type={isRevealRe_Password ? 'text' : 'password'}
           name="password"
           value={passwordCheck}
           onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPasswordCheck(e.currentTarget.value)}
           />
          <span
	          onClick={toggleRe_Password}
            role="presentation"
            style={{'color': 'black'}}
            >
           {isRevealRe_Password ? (
             <FontAwesomeIcon icon={faEye} />
             ) : (
               <FontAwesomeIcon icon={faEyeSlash} />
               )}
          </span><br></br>
         </div>
         </>
         )
         
        }
        //バリデーションに通った後、これ（finalSubmit)を実行するボタンを押せるようにする
        //firebaseに保存したemail, passwordを渡し、新規登録
        const finalSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
          e.preventDefault()
          console.log("firebase");
          await createUserWithEmailAndPassword(auth, staffEmail, password)
          Router.push("/admin/admin-login")
          //DBのスタッフテーブルに職員のName ,EmailをPOST
          const staffData = {
            "Name":staffName,
            "Email":staffEmail,
          }
          
          Axios.post(`api/proxy/staffPost`, staffData)
          .then((res) => {
            console.log(res);
          })
          .catch((error) => {
            console.log(error);
          });
        }
        
  return (
    <>
    <h1>ようこそ！スマートGAKUDOへ</h1>
    <div>
      <h3>新規登録</h3>
      <form>
        <div className='newer-form'>
          名前<span> ※必須</span><br></br>
          <input 
           type="text" 
           value={staffName} 
           {...register("staffName", { 
             required: '必須入力です',
             pattern: {
               value: /^[ぁ-んァ-ヶー一-龠]+$/,
               message: '全角文字で入力してください'
             }
            })}
           onChange={(e: React.ChangeEvent<HTMLInputElement>) => setStaffName(e.currentTarget.value)}
          /><br></br>
          {/* 名前入力のバリデーションメッセージ */}
          {errors.staffName && <p style={{'color': 'red'}}>{errors.staffName.message}</p>}<br></br>
        </div>
        <div className='newer-form'>
          メールアドレス<span> ※必須</span><br></br>
          <input 
           type="email"
           placeholder='mail@example.com' 
           value={staffEmail}
           {...register("staffEmail", {
            required: "メールアドレスを入力してください。",
            validate: value => isEmail(value) || '',
          })}
           onChange={(e: React.ChangeEvent<HTMLInputElement>) => setStaffEmail(e.currentTarget.value)}
           /><br></br>
            {/* メールアドレスのバリデーションメッセージ */}
            {errors.staffEmail && <p style={{'color': 'red'}}>{errors.staffEmail.message}</p>}<br></br>
        </div>
        {Password()}
        {/* <div><button onClick = {(e:any) => {finalSubmit(e)}}>登録</button></div> */}
        <div><button onSubmit={handleSubmit(onSubmit)}>登録</button></div>
        <p>会員登録には、<Link href='../all/all-termsOfService'><a>利用規約</a></Link >と<Link href='../all/all-privacyPolicy'><a>プライバシーポリシーへ</a></Link>の同意が必要です。</p>
      </form>

    </div>
   
   </>
  );
}

export default SignUp