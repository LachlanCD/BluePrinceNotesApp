import { type RoomNote } from "../types";
import { DeleteRoom, GETRoomDetails, UpdateRoom } from "../data/Rooms";
import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import NoteEditor from "../components/NoteEditor";
import DeleteButton from "../components/DeleteButton";

export default function RoomNotePage() {
  const navigate = useNavigate()
  const { id } = useParams()
  const [room, setRoom] = useState<RoomNote | null>(null);
  const [error, setError] = useState<String | null>(null);
  const [editing, setEditing] = useState<boolean>(false);
  const [markdown, setMarkdown] = useState<string>(`No notes yet`);

  useEffect(() => {
    async function getRooms() {
      setError(null);
      try {
        const roomData = await GETRoomDetails(id);
        setRoom(roomData);
        if (roomData.Notes) setMarkdown(roomData.Notes)
      } catch (err) {
        console.error(err)
        setError("Failed to retrieve rooms.");
      }
    }
    getRooms()
  }, []);

  const handleUpdate = async () => {
    if (!room?.Name || !room.Colour || !editing) return;

    const newRoom: RoomNote = {
      Id: room.Id,
      Name: room.Name,
      Colour: room.Colour,
      Notes: markdown,
    };

    UpdateRoom(newRoom);
  }

  const handleDelete = async () => {
    try {
      await DeleteRoom(id)
      navigate('/')
    } catch (err) {
      console.error(err)
      setError("Failed to retrieve rooms.");
    }
  }

  if (room === null && error === null) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      <h1>{room?.Name}</h1>
      <NoteEditor
        setEditing={setEditing}
        editing={editing}
        setMarkdown={setMarkdown}
        markdown={markdown}
        handleSubmit={handleUpdate}
      />
      <DeleteButton
        handleSubmit={handleDelete}
      />
    </div>
  )
}
