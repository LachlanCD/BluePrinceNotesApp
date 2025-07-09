import { type GeneralCard, type GeneralNote } from "../types";
import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { DeleteGeneral, GETGeneralDetails, UpdateGeneral } from "../data/General";
import NoteEditor from "../components/NoteEditor";
import DeleteButton from "../components/DeleteButton";
import NoteTitle from "../components/NoteTitle";

export default function RoomNotePage() {
  const navigate = useNavigate()
  const { id } = useParams()
  const [general, setGeneral] = useState<GeneralNote | null>(null);
  const [error, setError] = useState<String | null>(null);
  const [editingNote, setEditingNote] = useState<boolean>(false);
  const [markdown, setMarkdown] = useState<string>(`No notes yet`);
  const [name, setName] = useState<string>(`Something Went Wrong`);
  const [editingName, setEditingName] = useState<boolean>(false);

  useEffect(() => {
    async function getRooms() {
      setError(null);
      try {
        const genData = await GETGeneralDetails(id);
        setGeneral(genData);
        if (genData.Name) setName(genData.Name)
        if (genData.Notes) setMarkdown(genData.Notes)
      } catch (err) {
        console.error(error)
        setError("Failed to retrieve rooms.");
      }
    }
    getRooms()
  }, []);


  const handleUpdate = async () => {

    try {
      if (!general?.Name) return;
      const newGeneral: GeneralCard = {
        Id: general.Id,
        Name: name,
      };

      UpdateGeneral(newGeneral);
    } catch (err) {
      console.error(err)
      setError("Failed to retrieve rooms.");
    }
  }

  const handleDelete = async () => {
    try {
      await DeleteGeneral(id)
      navigate('/generals')
    } catch (err) {
      console.error(err)
      setError("Failed to retrieve rooms.");
    }
  }

  if (general === null && error === null) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      <NoteTitle
        editing={editingName}
        setEditing={setEditingName}
        name={name}
        setName={setName}
        handleSubmit={handleUpdate}
      />
      <NoteEditor
        editing={editingNote}
        setEditing={setEditingNote}
        markdown={markdown}
        setMarkdown={setMarkdown}
        id={general?.Id}
        handleSubmit={handleUpdate}
      />
      <DeleteButton
        handleSubmit={handleDelete}
      />

    </div>
  )
}
