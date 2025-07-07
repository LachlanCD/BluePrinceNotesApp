import { type GeneralNote } from "../types";
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import NoteEditor from "../components/NoteEditor";
import { GETGeneralDetails, UpdateGeneral } from "../data/General";

export default function RoomNotePage() {
  const { id } = useParams()
  const [general, setGeneral] = useState<GeneralNote | null>(null);
  const [error, setError] = useState<String | null>(null);
  const [editing, setEditing] = useState<boolean>(false);
  const [markdown, setMarkdown] = useState<string>(`No notes yet`);

  useEffect(() => {
    async function getRooms() {
      setError(null);
      try {
        const genData = await GETGeneralDetails(id);
        setGeneral(genData);
        if (genData.Notes) setMarkdown(genData.Notes)
      } catch (err) {
        console.error(error)
        setError("Failed to retrieve rooms.");

      }
    }
    getRooms()
  }, []);

  const handleSubmit = async () => {
    if (!general?.Name || !editing) return;

    const newGeneral: GeneralNote = {
      Id: general.Id,
      Name: general.Name,
      Notes: markdown,
    };

    UpdateGeneral(newGeneral);
  }

  if (general === null && error === null) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      <h1>{general?.Name}</h1>
      <NoteEditor
        setEditing={setEditing}
        editing={editing}
        setMarkdown={setMarkdown}
        markdown={markdown}
        handleSubmit={handleSubmit}
      />

    </div>
  )
}
