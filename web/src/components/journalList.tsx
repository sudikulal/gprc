'use client';

import React from "react";
import Link from 'next/link';



type Journal = {
    "journalId": string,
    "userId": string,
    "folderId": string,
    "title": string,
    "createdAt": string
}

type JournalComponentProps = {
    journalList: Journal[];
};

const JournalListComponent: React.FC<JournalComponentProps> = ({ journalList }) => {
  console.log("++> ~ journalList:", journalList)
  return (
    <ul className="menu bg-base-200">
      {journalList.map((journal, i) => (
        <li key={i}>
          <Link href={`journal/${journal.journalId}`}>{journal.title}</Link>
        </li>
      ))}
    </ul>
  );
};

export default JournalListComponent;
