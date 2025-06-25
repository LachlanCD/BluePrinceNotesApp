import { useState } from 'react';
import SelectInput from '../components/SelectInput';

export default function CreationForm() {
  const [name, setName] = useState('');
  const [colour, setColour] = useState('');

  const colours = [
    "Blue",
    "Purple",
    "Orange",
    "Yellow",
    "Green",
    "Red",
    "Black",
  ]

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const formData = {
      name,
      colour,
    };

    console.log('Form submitted:', formData);

    // Clear the form
    setName('');
    setColour('');
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 rounded-lg shadow-md">
      <h2 className="text-2xl font-semibold mb-4">Create Item</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
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

        <SelectInput 
          label="Colour"
          options={colours}
          value={colour}
          onChange={setColour}
          placeholder="Select a colour"
          required
        />

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

