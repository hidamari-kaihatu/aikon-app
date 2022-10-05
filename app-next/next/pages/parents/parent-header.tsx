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
        <Link href="/parents/parent-mypage">
          <a>保護者HOME</a>
        </Link>
      </li>
      <li>
        <Link href="/parents/parent-daily">
          <a>日々の出欠報告</a>
        </Link>
      </li>
      <li>
        <Link href="/parents/parent-fromCenter">
          <a>学童からの連絡</a>
        </Link>
      </li>
    </ul>
  );
}