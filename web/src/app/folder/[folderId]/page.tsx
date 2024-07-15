// app/[userId]/page.tsx
import React from 'react';
import JournalList from "@/components/journalList"
import { cookies } from "next/headers";


type Journal = {
  "journalId": string,
  "userId": string,
  "folderId": string,
  "title": string,
  "createdAt": string
}


async function fetchFolderData(folderId: string) {
  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

  const cookieStore = cookies();
  const accessToken = cookieStore.get("access_token")?.value;

  if (!accessToken) {
    throw new Error("Access token is missing");
  }

  const res = await fetch(`${apiBaseUrl}/journals/?folder_id=${folderId}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      access_token: accessToken,
    },
  });

  return await res.json() || [];
}

export default async function FolderDetail({ params }: {
  params: {
    folderId: string;
  }
}) {
  const journalList:Journal[] = await fetchFolderData(params.folderId);

  return <JournalList journalList={journalList} />;
};

