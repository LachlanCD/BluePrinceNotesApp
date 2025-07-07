import { type RoomNote } from "../types";
import { GETRoomDetails, UpdateRoom } from "../data/Rooms";
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import NoteEditor from "../components/NoteEditor";

export default function RoomNotePage() {
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
        console.error(error)
        setError("Failed to retrieve rooms.");

      }
    }
    getRooms()
  }, []);

  const handleSubmit = async () => {
    if (!room?.Name || !room.Colour || !editing) return;

    const newRoom: RoomNote = {
      Id: room.Id,
      Name: room.Name,
      Colour: room.Colour,
      Notes: markdown,
    };

    UpdateRoom(newRoom);
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
        handleSubmit={handleSubmit}
      />

    </div>
  )
}
