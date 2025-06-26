import Card from "../components/Card";
import { GETAllGenerals } from "../data/General";
import { type GeneralCard } from "../types";
import { useState, useEffect } from "react";

export default function GeneralHomePage() {
  const [notes, setNotes] = useState<GeneralCard[] | null>(null);
  const [error, setError] = useState<String | null>(null);


  useEffect(() => {
    async function getNotes() {
      setError(null);
      try {
        const notes = await GETAllGenerals();
        setNotes(notes);
      } catch (err) {
        console.error(error)
        setError("Failed to retrieve notes.");

      }
    }
    getNotes()
  }, []);


  if (notes === null && error === null) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      <div className="mt-5">
        <hr className="border-white" />
        <div className="mt-8 ml-8 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
          {notes?.map((note) => (
            <Card key={note.Id} {...note} />
          ))}
        </div>
      </div>
    </div>
  )

}

