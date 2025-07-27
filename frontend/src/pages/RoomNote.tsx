import { type RoomCard, type RoomNote } from "../types";
import { DeleteRoom, GETRoomDetails, UpdateRoom, UpdateRoomNote } from "../data/Rooms";
import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import NoteEditor from "../components/NoteEditor";
import DeleteButton from "../components/DeleteButton";
import NoteTitle from "../components/NoteTitle";
import { useWorkspaceIDRedirect } from "../data/Utils";

export default function RoomNotePage() {
  const workspaceID = useWorkspaceIDRedirect("/rooms")
  const navigate = useNavigate()
  const { id } = useParams()
  const [room, setRoom] = useState<RoomNote | null>(null);
  const [error, setError] = useState<String | null>(null);
  const [editingNote, setEditingNote] = useState<boolean>(false);
  const [markdown, setMarkdown] = useState<string>("");
  const [name, setName] = useState<string>(`Something Went Wrong`);
  const [editingName, setEditingName] = useState<boolean>(false);
  const [colour, setColour] = useState<string>(`White`);

  const colours = [
    "Blue",
    "Purple",
    "Orange",
    "Yellow",
    "Green",
    "Red",
    "Black",
  ]

  useEffect(() => {
    async function getRooms() {
      if (!workspaceID) return;
      setError(null);
      try {
        const roomData = await GETRoomDetails(workspaceID, id);
        setRoom(roomData);
        if (roomData.Name) setName(roomData.Name)
        if (roomData.Notes) setMarkdown(roomData.Notes)
        if (roomData.Colour) setColour(roomData.Colour)
      } catch (err) {
        console.error(err)
        setError("Failed to retrieve rooms.");
      }
    }
    getRooms()
  }, [workspaceID]);

  const handleUpdate = async () => {
    if (!workspaceID) return;
    try {
      if (!room?.Name || !room.Colour) return;

      const newRoom: RoomCard = {
        Id: room.Id,
        Name: name,
        Colour: colour,
      };

      UpdateRoom(workspaceID, newRoom);
    } catch (err) {
      console.error(err);
    }
  }

  const handleDelete = async () => {
    if (!workspaceID) return;
    try {
      await DeleteRoom(workspaceID, id)
      navigate(`/${workspaceID}/rooms`)
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
          colour={colour}
          setColour={setColour}
          colours={colours}
        />
        <NoteEditor
          setEditing={setEditingNote}
          editing={editingNote}
          setMarkdown={setMarkdown}
          markdown={markdown}
          id={room?.Id}
          workspaceID={workspaceID}
          handleSubmit={UpdateRoomNote}
        />
      </div>
      <DeleteButton
        handleSubmit={handleDelete}
      />
    </div>
  )
}
