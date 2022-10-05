import Link from 'next/link';

export default function Header() {
  return (
    <ul>
      <li>
        <Link href="/">
          <a>Home</a>
        </Link>
      </li>
      <li>
        <Link href="/teachers/teacher-mypage">
          <a>先生HOME</a>
        </Link>
      </li>
      <li>
        <Link href="/teachers/teacher-sendMsg">
          <a>メッセージを送る</a>
        </Link>
      </li>
      <li>
        <Link href="/teachers/teacher-msgList">
          <a>保護者に送ったメッセージを見る</a>
        </Link>
      </li>
      <li>
        <Link href="/teachers/teacher-kidsinfo">
          <a>児童基本情報</a>
        </Link>
      </li>
      <li>
        <Link href="/teachers/teacher-todayKids">
          <a>今日の児童</a>
        </Link>
      </li>
      <li>
        <Link href="/teachers/teacher-childInOut">
          <a>子供の入退室</a>
        </Link>
      </li>
      <li>
        <Link href="/teachers/teacher-teacherInOut">
          <a>先生の勤怠</a>
        </Link>
      </li>
    </ul>
  );
}