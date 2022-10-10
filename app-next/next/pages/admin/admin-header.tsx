import Link from 'next/link';
import Today from '../../components/date';

const today = new Date();
const year = today.getFullYear()
const month = today.getMonth() + 1
const day = today.getDate()
const week = today.getDay()
const weekItems = ["日", "月", "火", "水", "木", "金", "土"]
const dayOfWeek = weekItems[week]

export default function Header() {
  return (
    <>
    <div className='adminleft'>
      <p className='smagaku'>スマートGAKUDO</p>
    <div>
      <ul>
        <Link href="/admin/admin-mypage">
          <a className='adminHome'>ホーム</a>
        </Link>
      </ul>
    </div>
    <p className='adminhomeborder'></p>
    <div>
      <ul>
        <Link href="/admin/admin-centerList">
          <a className='adminCentermypage'>学童施設一覧</a>
        </Link>
      </ul>
      <ul>
        <Link href="/admin/admin-teacherList">
          <a className='adminstaffmypage'>職員一覧</a>
        </Link>
      </ul>
      <ul>
        <Link href="/admin/admin-url">
          <a className='adminurlmypage'>URL</a>
        </Link>
      </ul>
{/*       <ul>
        <Link href="/admin/admin-centerPost">
          <a className='adminnewmypage'>新規施設登録</a>
        </Link>
      </ul> */}
      <div suppressHydrationWarning className='admintoday'>{month}月{day}日（{dayOfWeek}）</div> 
{/*       <p className='admintoday'><Today /></p> */}
    </div>
  </div>
    </>
  );
}