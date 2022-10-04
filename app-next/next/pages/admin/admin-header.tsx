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
        <Link href="/admin/admin-mypage">
          <a>管理者HOME</a>
        </Link>
      </li>
      <li>
        <Link href="/admin/admin-centerList">
          <a>学童施設一覧</a>
        </Link>
      </li>
      <li>
        <Link href="/admin/admin-teacherList">
          <a>職員一覧</a>
        </Link>
      </li>
      <li>
        <Link href="/admin/admin-url">
          <a>URL</a>
        </Link>
      </li>
      <li>
        <Link href="/admin/admin-centerPost">
          <a>新規施設登録</a>
        </Link>
      </li>
    </ul>
  );
}