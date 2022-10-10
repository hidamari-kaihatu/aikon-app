import { useForm, SubmitHandler } from "react-hook-form";
import isEmail from 'validator/lib/isEmail';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { ErrorMessage } from "@hookform/error-message";
import React,{ useState, useEffect } from "react";
import{ FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {faEye, faEyeSlash} from "@fortawesome/free-solid-svg-icons";
import { Axios } from '../../../lib/api';
import Router from 'next/router';
import { auth } from "../../../firebaseConfig";
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth";
import Link from 'next/link';
import { useRouter } from 'next/router';
import axios from "axios";

type Inputs = {
  name: string,
  email: string,
  password: string,
  confirmPassword: string,
  tell: string,
  grade: number,
};
const schema = yup.object().shape({
  name: yup
    .string()
    .required('入力してください')
    .matches(
      /^[ぁ-んァ-ヶｱ-ﾝﾞﾟ一-龠]*$/,
      'スペースなしで全角のひらがな、カタカナ、漢字で入力してください'
    ),
  email: yup
    .string()
    .required('入力してください')
    .email('メールアドレスの形式が不正です'),
  password: yup
    .string()
    .required('入力してください')
    .matches(
      /^(?=.*?[a-zA-Z])(?=.*?\d)[a-zA-Z\d]{8,}$/,
      'アルファベットと数字を組み合わせて8文字以上で入力してください'
    ),
  confirmPassword: yup
    .string()
    .required('入力してください')
    .matches(
      /^(?=.*?[a-zA-Z])(?=.*?\d)[a-zA-Z\d]{8,}$/,
      'アルファベットと数字を組み合わせて8文字以上で入力してください'
    )
    .oneOf([yup.ref('password'), null], '確認用パスワードが一致していません'),
  tell: yup
    .string()
    .required('入力してください')
    .matches(
        /^(0[5-9]0-[0-9]{4}-[0-9]{4}|0[0-9]{3}-[0-9]{2}-[0-9]{4}|0[0-9]{2}-[0-9]{3}-[0-9]{4}|0[0-9]{1}-[0-9]{4}-[0-9]{4})$/,
        'ハイフン付きで入力してください'
    ),
  grade: yup
    .string()
    .required('入力してください')
    .matches(
      /[1-6]/,
      '学年を数字で入力してください'
    )
    

})
.required();//これがないとコンソールにdataが表示されなかった！


export default function App() {
  const { register, handleSubmit, watch, formState: { errors }} = useForm<Inputs>({
    resolver: yupResolver(schema)
  });
  const router = useRouter();
  // console.log(router.query.id);//1
  const center_id = Number(router.query.id);
  //   console.log(center_id);
  const onSubmit: SubmitHandler<Inputs> = data => {
    console.log("data : ",data);
    console.log("firebase");
    const grade = Number(data.grade);
    createUserWithEmailAndPassword(auth, data.email, data.password)
    
    //DBのstudentテーブルに保護者のCenter_id, Name ,Email,ContactTell,GradeをPOST
    const postData = {
      "Center_id":center_id,
      "Name":data.name,
      "ContactTell":data.tell,
      "Email":data.email,
      "Grade":grade,
      "Status":true//これがないと、statusがfalseでDBに入る。
    }
    
    Axios.post(`api/proxy/studentPost`, postData)
    .then((res) => {
      console.log("Post Start!");
      console.log(res.data);
      // 登録完了したらログイン画面へ ※POSTが走る！goでエラーが出ていてもここではエラーが出ない！！
      Router.push("/parents/parent-login")
    })
    .catch((error) => {
      console.log("post Miss!");
      console.log(error);
    });
    
    
  }
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
    /* "handleSubmit" will validate your inputs before invoking "onSubmit" */
    
    <form onSubmit={handleSubmit(onSubmit)} noValidate>
              <p className="youkoso1">ようこそ！スマートGAKUDOへ</p>
      <div className="centerblack">
        {/* <h3>〇〇学童</h3> */}
        <h3></h3>
      </div>
      <div className="centerblack2">
       <h3>新規登録</h3>
      名前<span style={{'color': 'red', 'fontSize': 'small'}}> ※必須</span><br></br>
      <input 
        type='text'
        placeholder='全角文字で入力'
        {...register("name")} /><br></br>
      <span style={{'color': 'red', 'fontSize': 'small'}}>{errors.name?.message}</span><br></br>

      学年<span style={{'color': 'red', 'fontSize': 'small'}}> ※必須</span><br></br>
      <input 
        type='text'
        placeholder='例）1'
        {...register("grade")} /><br></br>
      <span style={{'color': 'red', 'fontSize': 'small'}}>{errors.grade?.message}</span><br></br>
      メールアドレス<span style={{'color': 'red', 'fontSize': 'small'}}> ※必須</span><br></br>
      <input 
        type="email"
        placeholder='mail@example.com' 
        {...register("email")} /><br></br>
      <span style={{'color': 'red', 'fontSize': 'small'}}>{errors.email?.message}</span><br></br>
      
      パスワード<span style={{'color': 'red', 'fontSize': 'small'}}> ※必須</span><br></br>
      <input 
        placeholder="半角英数字で８文字以上"
        {...register('password')}
        type={isRevealPassword ? 'text' : 'password'}
      />
      <span
	          onClick={togglePassword}
            role="presentation"
            style={{'color': 'black','fontSize': 'small'}}
            >
                {isRevealPassword ? (
            <button> <FontAwesomeIcon icon={faEye} /></button>
             ) : (
              <button> <FontAwesomeIcon icon={faEyeSlash} /></button>
               )}
          </span><br></br>
      <span style={{'color': 'red', 'fontSize': 'small'}}>{errors.password?.message}</span><br></br>
      
      パスワード確認<span style={{'color': 'red', 'fontSize': 'small'}}> ※必須</span><br></br>
      <input 
        placeholder="半角英数字で８文字以上"
        {...register('confirmPassword')}
        type={isRevealRe_Password ? 'text' : 'password'}
      />
      <span
	          onClick={toggleRe_Password}
            role="presentation"
            style={{'color': 'black', 'fontSize': 'small'}}
            >
           {isRevealRe_Password ? (
             <button><FontAwesomeIcon icon={faEye} /></button>
             ) : (
               <button><FontAwesomeIcon icon={faEyeSlash} /></button>
               )}
          </span><br></br>
      <span style={{'color': 'red', 'fontSize': 'small'}}>{errors.confirmPassword?.message}</span><br></br>

      緊急連絡先<span style={{'color': 'red', 'fontSize': 'small'}}> ※必須</span><br></br>
      <input 
        type="tell"
        placeholder='090-1111-2222' 
        {...register("tell")} /><br></br>
      <span style={{'color': 'red', 'fontSize': 'small'}}>{errors.tell?.message}</span><br></br>
      <input type="submit" />
      <p style={{'fontSize': 'small'}}>会員登録には、<Link href='../all/all-termsOfService'><a>利用規約</a></Link >と<Link href='../all/all-privacyPolicy'><a>プライバシーポリシーへ</a></Link>の同意が必要です。</p>

    </div>
    </form>
  );
}