import { useState } from 'react';
import SelectInput from '../components/SelectInput';
import { FormatNewRoom } from '../data/Rooms';
import { FormatNewGeneral } from '../data/General';
import { useNavigate } from 'react-router-dom';

export default function CreationForm() {
  const [name, setName] = useState('');
  const [colour, setColour] = useState('');
  const [noteType, setNoteType] = useState('');

  const navigate = useNavigate();

  const colours = [
    "Blue",
    "Purple",
    "Orange",
    "Yellow",
    "Green",
    "Red",
    "Black",
  ]

  const noteTypes = [
    "Room",
    "General",
  ]

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      if (noteType === "Room") return await FormatNewRoom({ name, colour, navigate })
      return await FormatNewGeneral({ name, navigate })
    } catch (err) {
      console.error(err)
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 rounded-lg shadow-md">
      <h2 className="text-2xl font-semibold mb-4">Create Item</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <SelectInput
          label="noteType"
          options={noteTypes}
          value={noteType}
          onChange={setNoteType}
          placeholder="Select a note type"
          required
        />
        <div>
          <label className="block mb-1 font-medium">Name</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:ring-blue-300"
            placeholder="Enter name"
          />
        </div>

        {noteType === "Room" && (
          <SelectInput
            label="Colour"
            options={colours}
            value={colour}
            onChange={setColour}
            placeholder="Select a colour"
            required
          />) 
        }

        <button
          type="submit"
          className="w-full bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition"
        >
          Submit
        </button>
      </form>
    </div>
  );
}

