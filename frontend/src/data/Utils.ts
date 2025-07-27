import type { UpdateNoteProps } from "../types";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

export async function FetchData(url: string) {
  const response = await fetch(url);
  if (!response.ok) throw new Error('Network response was not ok');
  const data = await response.json();
  return data
}

export async function FetchAndCache(url: string) {
  const data = FetchData(url)
  return data;
}

export async function AddNew(url: string, data: URLSearchParams) {
  const response = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: data.toString(),
  })

  if (!response.ok) throw new Error('Unable to add new item');
  const rData = await response.json();
  return rData;
}

export async function DeleteItem(url: string) {
  const response = await fetch(url, {
    method: 'DELETE',
  })

  if (!response.ok) throw new Error('Unable to remove item');
}

export type HandleNoteUpdateProps = {
  id: number | undefined;
  workspaceID: string;
  note: string;
  updateNote: (updateNoteProps: UpdateNoteProps) => void;
}

export async function HandleNoteUpdate({ id, workspaceID, note, updateNote }: HandleNoteUpdateProps) {
  if (!note || !id) return;
  const noteProps: UpdateNoteProps = {
    id: id,
    workspaceID: workspaceID,
    note: note,
  };

  updateNote(noteProps);
}

export function useWorkspaceIDRedirect(currentUrl: string) {
  const navigate = useNavigate();
  const { workspaceID } = useParams<{ workspaceID?: string }>();
  const [currentWorkspaceID, setCurrentWorkspaceID] = useState<string | null>(null);

  useEffect(() => {
    async function getID() {
      if (!workspaceID) {
        try {
          const route = `/api/create-workspace`
          const url = import.meta.env.VITE_BASE_URL + route
          const res = await fetch(url, { method: "POST" });
          const newID = await res.text();
          navigate(`${currentUrl}${newID}`, { replace: true });

          setCurrentWorkspaceID(newID)
        } catch (err) {
          console.error(err)
        }
      } else {
        setCurrentWorkspaceID(workspaceID);
      }
    }
    getID()
  }, [workspaceID, navigate]);

  return currentWorkspaceID;
}
