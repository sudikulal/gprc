'use client';

import React from "react";
import Link from 'next/link';


type Folder = {
  folderName: string;
  folderId: string;
};

type FolderComponentProps = {
  folders: Folder[];
};

const FolderComponent: React.FC<FolderComponentProps> = ({ folders }) => {
  return (
    <ul className="menu bg-base-200">
      {folders.map((folder, i) => (
        <li key={i}>
          <Link href={`folder/${folder.folderId}`}>{folder.folderName}</Link>
        </li>
      ))}
    </ul>
  );
};

export default FolderComponent;
