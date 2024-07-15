import React from "react";
import { cookies } from "next/headers";
import FolderComponent from "@/components/folderList";


type Folder = {
  folderName: string;
  folderId: string;
};

export default async function Folder() {
  const cookieStore = cookies();
  const accessToken = cookieStore.get("access_token")?.value;

  if (!accessToken) {
    throw new Error("Access token is missing");
  }

  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

  if (!apiBaseUrl) {
    throw new Error("API base URL is not defined");
  }

  const response = await fetch(`${apiBaseUrl}/folders/`, {
    method:"GET",
    headers: {
      "Content-Type": "application/json",
      access_token: accessToken,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch data");
  }

  let folders: Folder[] = await response.json() || []


  return <FolderComponent folders={folders} />;
}



