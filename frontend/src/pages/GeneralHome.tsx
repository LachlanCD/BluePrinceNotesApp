import Card from "../components/Card";
import { GETAllGenerals } from "../data/General";
import { useWorkspaceIDRedirect } from "../data/Utils";
import { type GeneralCard } from "../types";
import { useState, useEffect } from "react";

export default function GeneralHomePage() {
  const workspaceID = useWorkspaceIDRedirect("/generals/")
  const [notes, setNotes] = useState<GeneralCard[] | null>(null);
  const [error, setError] = useState<String | null>(null);


  useEffect(() => {
    async function getNotes() {
      if (!workspaceID) return;
      setError(null);
      try {
        const notes = await GETAllGenerals(workspaceID);
        setNotes(notes);
      } catch (err) {
        console.error(error)
        setError("Failed to retrieve notes.");

      }
    }
    getNotes()
  }, [workspaceID]);

  if (!workspaceID) return<div className="mt-20">Retrieving workspce id</div>;
 
  if (notes === null && error === null) return <div className="mt-20">No Notes Yet.</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      <div className="mt-5">
        <hr className="border-stone-700 dark:border-zinc-200" />
        <div className="mt-8 ml-8 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
          {notes?.map((note) => (
            <Card key={note.Id} workspaceID={workspaceID} room={note}/>
          ))}
        </div>
      </div>
    </div>
  )

}

