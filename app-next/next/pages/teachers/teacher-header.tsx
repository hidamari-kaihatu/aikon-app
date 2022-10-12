import Link from 'next/link';

export default function Header() {
  return (
    <>
        <div className='adminleft'>
      <p className='smagaku'>スマートGAKUDO</p>
      <Link href="/teachers/teacher-mypage">
          <a className='home'>ホーム</a>
        </Link>
        <p className='adminhomeborder'></p>
        <div className='zidou'>児童</div>
        <Link href="/teachers/teacher-todayKids">
          <a className='a1'>今日の児童一覧を見る</a>
        </Link>
      
        <Link href="/teachers/teacher-kidsinfo">
          <a className='a2'>児童名簿一覧を見る</a>
        </Link>
        
        <div className='msg'>メッセージ</div>
      
        <Link href="/teachers/teacher-sendMsg">
          <a className='a3'>メッセージを送る</a>
        </Link>
      
        <Link href="/teachers/teacher-msgList">
          <a className='a4'>メッセージを見る</a>
        </Link>

        <div className='inout'>勤怠管理</div>
      
        <Link href="/teachers/teacher-teacherInOut">
          <a className='a5'>勤怠を確認する</a>
        </Link>
      
        <Link href="/teachers/teacher-childInOut">
          <a className='a6'>入退室を確認する</a>
        </Link>
        </div>
        <p className='sumagakuipad'>スマートGAKUDO</p>
    </>
  );
}