import type { NextPage } from 'next'
import React, { useEffect, useState, FC } from 'react'
import { useRouter } from 'next/router'
import { signInWithEmailAndPassword } from "firebase/auth";
import { auth } from '../../firebaseConfig';
import { getAuth, getIdToken } from "firebase/auth";
import { Axios } from '../../lib/api';

const Login: FC = () => {
  const router = useRouter()
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')

  const logIn = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    try {
      await signInWithEmailAndPassword(auth, email, password)
      const firebaseAuth = getAuth()
      const user:any = firebaseAuth.currentUser
      const idToken = await getIdToken(user, true)
      //console.log(idToken)
        Axios.post(`api/proxy/parentIsLogin`,idToken, {
        headers:{
          // 'Content-Type':'application/x-www-form-urlencoded',
          Authorization: `Bearer ${idToken}` //ヘッダにトークン付与
          },
      })
      // .then((res) => {
      //   console.log(res);
      // })
      .catch((error) => {
        console.log(error);
      });

      //localStorage.setItem('idToken', idToken)
      router.push('/parents/parent-mypage')
    } catch (err) {
       alert("メールアドレスまたはパスワードが間違っています")
      router.push('/404')
    }
  }

  return (
    <div>
{/*       ayaka@ayaka.com */}
      <form onSubmit={logIn}>
      <p className="youkoso2">ようこそ!スマートGAKUDOへ</p>
      <div>
        <div className='parentloginback'>
        <p className='palog'>ログイン</p>
        <div>
          <label htmlFor="email" className='leftname'>
            <p className='leftname2'>メールアドレス</p>
          </label>
          <input
                className="inputnewerpl"
            id="email"
            type="email"
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <br></br>
        <div>
          <label htmlFor="password">
          <p className='leftname2'>パスワード</p>
          </label>
          <input
                className="inputnewerpl"
            id="password"
            type="password"
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <br></br>
        <button className='orangebutton2' type="submit">
          送信
        </button>
        <br></br>
        <br></br>
        </div>
        </div>
      </form>
    </div>
  )
}

export default Login