import { useRouter } from 'next/router';
import Today from '../../components/date';
import { auth } from '../../firebaseConfig';

export default function Footer() {
    const router = useRouter()
    const logOut = async () => {
        try {
          await auth.signOut()
          router.push('/parents/parent-login')
        } catch (error) {
          router.push('/parents/parent-login')
        }
      }
    return (
        <>
            <Today />
            <button className="btn" onClick={logOut}>Logout</button>
        </>
    );
}
