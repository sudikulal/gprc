// app/[userId]/page.tsx
import React from 'react';
import JournalCard from "@/components/journalCard"
import { cookies } from "next/headers";


type Journal = {
    "journalId": string,
    "userId": string,
    "folderId": string,
    "title": string,
    "createdAt": string,
    "dayRating": number
}


async function fetchJournalData(journalId: string) {
    const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

    const cookieStore = cookies();
    const accessToken = cookieStore.get("access_token")?.value;

    if (!accessToken) {
        throw new Error("Access token is missing");
    }

    const res = await fetch(`${apiBaseUrl}/journals/${journalId}`, {
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
        journalId: string;
    }
}) {
    const journalDetail: Journal[] = await fetchJournalData(params.journalId);

    return <JournalCard />;
};

