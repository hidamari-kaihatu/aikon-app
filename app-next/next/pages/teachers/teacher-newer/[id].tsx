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
    confirmPassword: string
  };
  const schema = yup.object().shape({
    name: yup
      .string()
      .required('必須項目です')
      .matches(
        /^[ぁ-んァ-ヶｱ-ﾝﾞﾟ一-龠]*$/,
        'スペースなしで全角のひらがな、カタカナ、漢字で入力してください'
      ),
    email: yup
      .string()
      .required('必須項目です')
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
  
  })
  .required();//これがないとコンソールにdataが表示されなかった！
  
  export default function App() {
    const { register, handleSubmit, watch, formState: { errors } } = useForm<Inputs>({
      resolver: yupResolver(schema)
    });
      const router = useRouter();
//   console.log(router.query.id);
  const center_id = router.query.id;
//   console.log(center_id);
    const onSubmit: SubmitHandler<Inputs> = data => {
      console.log("data : ",data);
      console.log("firebase");
            createUserWithEmailAndPassword(auth, data.email, data.password)
            
            //DBのスタッフテーブルに職員のName ,EmailをPOST
            const postData = {
                "Name":data.name,
                "Email":data.email,
            }
            
            Axios.post(`api/proxy/staffPost`, postData)
            .then((res) => {
                console.log(res.data);
                // 登録完了したら確認画面へ(新規登録職員の名前を渡す)
                Router.push({pathname: "/teachers/teacher-newer", query: query} ,"/teachers/teacher-newer")
            })
            .catch((error) => {
                console.log(error);
            });

            const query = {
                                name: data.name,
                                center_id: center_id
                            }
            //登録完了したらログイン画面へ
            // Router.push("/admin/admin-login")
  
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
    <div className="centerblack"></div>
      <div>
        <p className="youkoso111">ようこそ！スマートGAKUDOへ</p>
        {/* <h3>〇〇学童</h3> */}
      </div>
      <div className="centerblack22">
       <h3>新規登録</h3>
      名前<span style={{'color': 'red', 'fontSize': 'small'}}> ※必須</span><br></br>
      <input 
        type='text'
        placeholder='全角文字で入力'
        {...register("name")} /><br></br>
      <span style={{'color': 'red', 'fontSize': 'small'}}>{errors.name?.message}</span><br></br>

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

      <input type="submit" />
      <p style={{'fontSize': 'small'}}>会員登録には、<Link href='../all/all-termsOfService'><a>利用規約</a></Link >と<Link href='../all/all-privacyPolicy'><a>プライバシーポリシーへ</a></Link>の同意が必要です。</p>

    </div>
    </form>
  );
}