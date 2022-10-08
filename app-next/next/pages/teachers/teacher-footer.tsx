import Link from 'next/link';
import Today from '../../components/date';

export default function Footer() {
    return (
        <>
        <Today />
        <p>MsEクラブ（ここにログインした人が所属する学童名が入る予定）</p>
        <p>☆❤△先生のMYページ（ここにログインした人のnameが入る予定）</p>
        </>
    );
}