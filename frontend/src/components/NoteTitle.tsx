import SelectInput from "./SelectInput";
import { GetHexCode } from "./Utils";

export type noteEditorProps = {
  setEditing: (value: boolean) => void;
  editing: boolean;
  setName: (value: string) => void;
  name: string;
  handleSubmit: () => void;
  colour?: string;
  setColour?: ((value: string) => void);
  colours?: string[];
}

const NoteTitle = ({ setEditing, editing, setName, name, handleSubmit, colour='White', setColour, colours }: noteEditorProps) => {
  const bc = GetHexCode(colour);
  return (
    <div className="p-4">
      {editing ? (
        <div
          onBlur={() => { setEditing(false); handleSubmit(); }}
          autoFocus
        >
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="text-3xl font-bold border p-1"
          />
          {(colour && setColour && colours) && (
            <SelectInput
              label="Colour"
              options={colours}
              value={colour}
              onChange={setColour}
              placeholder="Select a colour"
            />
          )}
        </div>
      ) : (
        <h1
          className="text-3xl font-bold cursor-pointer"
          style={{ color: bc }}
          onClick={() => setEditing(true)}
        >
          {name}
        </h1>
      )}
    </div>
  );
};

export default NoteTitle
