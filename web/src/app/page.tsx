'use client';
import { useRouter } from 'next/navigation';
import Cookies from 'js-cookie';

export default function Home() {
  const accessToken = Cookies.get("access_token"); 
  const router = useRouter();
  if(!accessToken) 
    router.push('/user/signIn');
  else
    router.push("/folder")
}
