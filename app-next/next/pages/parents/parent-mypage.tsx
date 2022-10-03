/* eslint react-hooks/exhaustive-deps: off */ 
import type { GetServerSideProps, NextPage } from 'next'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { auth } from '../../firebaseConfig'
import Link from 'next/link'

const List: NextPage = ({ list }: any) => {
  const router = useRouter()
  const [currentUser, setCurrentUser] = useState<null | object>(null)

  useEffect(() => {
    auth.onAuthStateChanged((user) => {
      user ? setCurrentUser(user) : router.push('/parents/parent-login')
    })
  }, [])
  const logOut = async () => {
    try {
      await auth.signOut()
      router.push('/parents/parent-login')
    } catch (error) {
      router.push('/parents/parent-login')
    }
  }

  return (
    <>
<div className='list'><p>ご祝儀登録者リスト</p></div>    
<Link href={"/parents/parent-daily"}>
            <h2>イベントページをみる</h2>
          </Link>

      <div>
        {/* <table className='list-table'>
          <thead>
            <tr>
              <th>名前</th>
              <th>住所</th>
              <th>電話番号</th>
              <th>金額</th>
              <th>お返しタイプ</th>
            </tr>
          </thead>
          <tbody>
            {list.map((element: any) => {
              return (
                <tr key={element.id}>
                  <td>{element.name}</td>
                  <td>{element.adress}</td>
                  <td>{element.tel}</td>
                  <td>{element.payAmount}</td>
                  <td>タイプ</td>
                </tr>
              )
            })}
          </tbody>
        </table> */}
        <h1>login test</h1>
      </div>
      <button className="btn" onClick={logOut}>Logout</button>
    </>
  )
}
// export const getServerSideProps: GetServerSideProps = async () => {
//   const resList = await fetch(`http://express-type:4000`) // TODO:データを持ってくるexpressのURLいれる
//   const list = await resList.json()

//   return {
//     props: {
//       list
//     }
//   }
// }

export default List
