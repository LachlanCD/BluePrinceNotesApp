import { type RoomCard, type RoomNote } from "../types";
import { DeleteRoom, GETRoomDetails, UpdateRoom, UpdateRoomNote } from "../data/Rooms";
import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import NoteEditor from "../components/NoteEditor";
import DeleteButton from "../components/DeleteButton";
import NoteTitle from "../components/NoteTitle";

export default function RoomNotePage() {
  const navigate = useNavigate()
  const { id } = useParams()
  const [room, setRoom] = useState<RoomNote | null>(null);
  const [error, setError] = useState<String | null>(null);
  const [editingNote, setEditingNote] = useState<boolean>(false);
  const [markdown, setMarkdown] = useState<string>(`No notes yet`);
  const [name, setName] = useState<string>(`Something Went Wrong`);
  const [editingName, setEditingName] = useState<boolean>(false);



  useEffect(() => {
    async function getRooms() {
      setError(null);
      try {
        const roomData = await GETRoomDetails(id);
        setRoom(roomData);
        if (roomData.Name) setName(roomData.Name)
        if (roomData.Notes) setMarkdown(roomData.Notes)
      } catch (err) {
        console.error(err)
        setError("Failed to retrieve rooms.");
      }
    }
    getRooms()
  }, []);

  const handleUpdate = async () => {
    try {
    if (!room?.Name || !room.Colour) return;

    const newRoom: RoomCard = {
      Id: room.Id,
      Name: name,
      Colour: room.Colour,
    };

    UpdateRoom(newRoom);
    } catch (err) {
      console.error(err);
    }
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
      <div 
      >
        <NoteTitle
          editing={editingName}
          setEditing={setEditingName}
          name={name}
          setName={setName}
          handleSubmit={handleUpdate}
        />
        <NoteEditor
          setEditing={setEditingNote}
          editing={editingNote}
          setMarkdown={setMarkdown}
          markdown={markdown}
          id={room?.Id}
          handleSubmit={UpdateRoomNote}
        />
      </div>
      <DeleteButton
        handleSubmit={handleDelete}
      />
    </div>
  )
}
