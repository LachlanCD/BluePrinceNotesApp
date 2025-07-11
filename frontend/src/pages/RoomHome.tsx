import { type RoomCard } from "../types";
import { GETAllRooms } from "../data/Rooms";
import { useState, useEffect } from "react";
import ColouredSection from "../components/ColouredSection";

export default function RoomHomePage() {
  const [rooms, setRooms] = useState<RoomCard[] | null>(null);
  const [error, setError] = useState<String | null>(null);


  useEffect(() => {
    async function getRooms() {
      setError(null);
      try {
        const rooms = await GETAllRooms();
        setRooms(rooms);
      } catch (err) {
        console.error(err)
        setError("Failed to retrieve rooms.");

      }
    }
    getRooms()
  }, []);


  const groupedRooms = rooms?.reduce<Record<string, RoomCard[]>>((acc, obj) => {
    if (!acc[obj.Colour]) {
      acc[obj.Colour] = [];
    }
    acc[obj.Colour].push(obj);
    return acc;
  }, {}) || {};

  if (rooms === null && error === null) return <div className="mt-20">No Notes Yet.</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      {Object.entries(groupedRooms).map(([colour, group]) => (
        <ColouredSection key={colour} colour={colour} rooms={group} />
      ))}
    </div>
  )
}

