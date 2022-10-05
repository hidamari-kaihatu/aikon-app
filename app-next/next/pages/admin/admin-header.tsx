import Link from 'next/link';

export default function Header() {
  return (
    <>
    <div>
      <p>スマートGAKUDO</p>
      <p>管理者</p>
    </div>
    <div>
      <ul>
        <Link href="/admin/admin-mypage">
          <a>管理者HOME</a>
        </Link>
      </ul>
    </div>
    <div>
      <ul>
        <Link href="/admin/admin-centerList">
          <a>学童施設一覧</a>
        </Link>
      </ul>
      <ul>
        <Link href="/admin/admin-teacherList">
          <a>職員一覧</a>
        </Link>
      </ul>
      <ul>
        <Link href="/admin/admin-url">
          <a>URL</a>
        </Link>
      </ul>
      <ul>
        <Link href="/admin/admin-centerPost">
          <a>新規施設登録</a>
        </Link>
      </ul>
    </div>
    </>
  );
}