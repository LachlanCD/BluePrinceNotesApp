export type noteEditorProps = {
  setEditing: (value: boolean) => void;
  editing: boolean;
  setName: (value: string) => void;
  name: string;
  handleSubmit: () => void;
}

const NoteTitle = ({ setEditing, editing, setName, name, handleSubmit }: noteEditorProps) => {
  return (
    <div className="p-4">
      {editing ? (
        <input
          type="text"
          onClick={() => setEditing(true)}
          onBlur={() => { setEditing(false); handleSubmit(); }}
          autoFocus
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="text-3xl font-bold border p-1"
        />
      ) : (
        <h1
          className="text-3xl font-bold cursor-pointer"
          onClick={() => setEditing(true)}
        >
          {name}
        </h1>
      )}
    </div>
  );
};

export default NoteTitle
