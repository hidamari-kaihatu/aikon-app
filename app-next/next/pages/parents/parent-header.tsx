import Link from 'next/link';
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
    {/* <p className='smagaku00'>スマートGAKUDO</p> */}
      <ul>
        <Link href="/parents/parent-mypage">
          <a className='parenthome'>ホーム</a>
        </Link>
      </ul>
      <p className='parenthomeborder'></p>
      <ul>
        <Link href="/parents/parent-daily">
          <a className='parenthomedaily'>出欠を連絡する</a>
        </Link>
      </ul>
      <ul>
        <Link href="/parents/parent-fromCenter">
          <a className='parenthomecenter'>学童からの連絡を見る</a>
        </Link>
      </ul>
      <div suppressHydrationWarning className='admintoday'>{month}月{day}日（{dayOfWeek}）</div> 
      </div>
      <p className='smagakusumaho'>スマートGAKUDO</p>
    </>
  );
}