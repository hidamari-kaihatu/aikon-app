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
        Axios.post(`api/proxy/staffIsLogin`,idToken, { 
        headers:{
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
      router.push('/admin/admin-mypage')
    } catch (err) {
       alert("メールアドレスまたはパスワードが間違っています")
      router.push('/404')
    }
  }

  return (
    <div>
      takumi@takumi.com 
      <form onSubmit={logIn}>
        <div>
          <label htmlFor="email">
            Email:{' '}
          </label>
          <input
            id="email"
            type="email"
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor="password">
            Password:{' '}
          </label>
          <input
            id="password"
            type="password"
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button type="submit">
          Login
        </button>
      </form>
    </div>
  )
}

export default Login