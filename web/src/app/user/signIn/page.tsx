'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Cookies from 'js-cookie';

export default function SignIn() {
  const [identifier, setIdentifier] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<string>('');
  const router = useRouter();

  const isEmail = (str: string) => {
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailPattern.test(str);
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    let payload:{registerType?:number,password:string,userName?:string,emailId?:string} = {password}

    if(isEmail(identifier)){
      payload.registerType = 1
      payload.emailId = identifier
    }else{
      payload.registerType = 2
      payload.userName = identifier
    }

    const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
    const res = await fetch(`${apiBaseUrl}/user/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    });

    if (res.ok) {
     let data =await res.json()
     Cookies.set("access_token", data.accessToken, { expires: 1 }); 
     router.push('/');
    } else {
      const errorData = await res.json();
      setError(errorData.message || 'Something went wrong');
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-3 bg-white rounded shadow-md">
        <h1 className="text-2xl font-bold text-center">Sign In</h1>
        {error && <p className="text-red-500 text-center">{error}</p>}
        <form onSubmit={handleSubmit}>
          <div className="form-control">
            <label className="label">
              <span className="label-text">Email or Username</span>
            </label>
            <input
              type="text"
              placeholder="Email or Username"
              className="input input-bordered w-full"
              value={identifier}
              onChange={(e) => setIdentifier(e.target.value)}
            />
          </div>
          <div className="form-control">
            <label className="label">
              <span className="label-text">Password</span>
            </label>
            <input
              type="password"
              placeholder="Password"
              className="input input-bordered w-full"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className="form-control mt-6">
            <button type="submit" className="btn btn-primary w-full">Sign In</button>
          </div>
        </form>
        <p className="text-center">
          Don't have an account? <a href="/user/signUp" className="text-primary">Sign Up</a>
        </p>
      </div>
    </div>
  );
}
