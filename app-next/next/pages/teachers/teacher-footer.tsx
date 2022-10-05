import Link from 'next/link';
import Today from '../../components/date';

export default function Footer() {
    return (
        <>
        <Today />
        <p>MsEクラブ（ここに所属する学童名が入る）</p>
        <p>☆❤△先生のMYページ（ここにログインした人のnameが入る）</p>
        </>
    );
}