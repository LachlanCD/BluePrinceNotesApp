import { useEffect, useRef } from "react"
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

const NoteTitle = ({ setEditing, editing, setName, name, handleSubmit, colour = 'White', setColour, colours }: noteEditorProps) => {
  const wrapperRef = useRef<HTMLDivElement>(null);
  const editingRef = useRef(editing);
  const bc = GetHexCode(colour);
  const firstUpdate = useRef(true);


  useEffect(() => {
    if (firstUpdate.current){
      firstUpdate.current = false;
      return;
    }

    if (!editing) {
      handleSubmit();
    }
    
    editingRef.current = editing;
  }, [editing])

  checkOutsideClick(wrapperRef, () => {
    if (!editingRef.current) return;
    setEditing(false);
  });


  return (
    <div className="p-4">
      {editing ? (
        <div
          ref={wrapperRef}
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

const checkOutsideClick = ((
  ref: React.RefObject<HTMLElement | null>,
  onOutsideClick: (() => void),
) => {
  useEffect(() => {

    function handleClickOutside(e: MouseEvent) {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        onOutsideClick()
      }
    }
    // Bind the event listener
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      // Unbind the event listener on clean up
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [ref]);
});

export default NoteTitle
