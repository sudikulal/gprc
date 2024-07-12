"use client"
import React from 'react';

interface Props {
  accessToken: string;
  response:any
}

const Page: React.FC<Props> = ({ accessToken, response }) => {
  if (response.error || response.message) {
    return <p>Error: {response}</p>;
  }

  return (
    <ul className="menu bg-base-200 rounded-box w-56">
      {response.map((folder:{folderName:string}, i:number) => (
        <li key={i}>
          <a>{folder.folderName}</a>
        </li>
      ))}
    </ul>
  );
};

export async function getServerSideProps(context: any) {
  const { accessToken } = context.query;

  try {
    const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
    const res = await fetch(`${apiBaseUrl}/user/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        access_token: accessToken,
      }),
    });

    if (!res.ok) {
      throw new Error('Failed to fetch data');
    }

    const response = await res.json();

    return {
      props: {
        accessToken: accessToken || '',
        response,
      },
    };
  } catch (error) {
    return {
      props: {
        accessToken: accessToken || '',
        folders: [],
        error: error || 'Failed to fetch data',
      },
    };
  }
}

export default Page;
