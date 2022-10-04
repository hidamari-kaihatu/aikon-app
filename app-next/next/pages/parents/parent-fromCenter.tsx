import type { NextPage } from 'next'
import Layout from './parent-layout';
import { useRouter } from 'next/router';

export default function mypage() {
    const router = useRouter();
    return (
      <>
      <Layout>
        <h2>学童からの連絡</h2>
        </Layout>
      </>
    );
} 