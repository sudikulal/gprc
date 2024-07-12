'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Cookies from 'js-cookie';

export default function Logout() {
  const router = useRouter();

  useEffect(() => {
    const logout = () => {
      Cookies.remove('access_token');
      router.push('/user/signIn');
    };

    logout();
  }, [router]);

  return <div>Logging out...</div>;
}
