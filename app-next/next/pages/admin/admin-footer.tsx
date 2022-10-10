import axios from 'axios';
import Link from 'next/link';
import { useRouter } from 'next/router'
import { auth } from '../../firebaseConfig';
import Today from '../../components/date';

  
export default function Footer(staffs:any) {
    const router = useRouter()
    const logOut = async () => {
        try {
          await auth.signOut()
          router.push('/admin/admin-login')
        } catch (error) {
          router.push('/admin/admin-login')
        }
      }
    
    return (
        <>
        <Today />
          <button className="btn" onClick={logOut}>Logout</button>
        </>
    );
  }