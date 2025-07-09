import { HandleNoteUpdate } from "../data/Utils";
import type { UpdateNoteProps } from "../types";
import MarkdownRenderer from "./MarkdownRenderer";

export type noteEditorProps = {
  setEditing: (value: boolean) => void;
  editing: boolean;
  setMarkdown: (value: string) => void;
  markdown: string;
  id: number|undefined;
  handleSubmit: ({id, note}: UpdateNoteProps) => void;
}

const NoteEditor = ({ setEditing, editing, setMarkdown, markdown, id, handleSubmit }: noteEditorProps) => {
  const handleTab = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Tab") {
      e.preventDefault();
      const target = e.target as HTMLTextAreaElement;
      const start = target.selectionStart;
      const end = target.selectionEnd;

      // Insert tab character at caret
      const newValue =
        markdown.substring(0, start) + "\t" + markdown.substring(end);
      setMarkdown(newValue);

      // Move caret
      setTimeout(() => {
        target.selectionStart = target.selectionEnd = start + 1;
      }, 0);
    }
  };

  return (
    <div>
      <div style={{ marginTop: '1rem' }}>
        {editing ? (
          <textarea
            onClick={() => setEditing(true)}
            onBlur={() => { setEditing(false); HandleNoteUpdate({id, note:markdown, updateNote:handleSubmit}) }}
            autoFocus
            value={markdown}
            onChange={(e) => setMarkdown(e.target.value)}
            onKeyDown={handleTab}
            className="cursor-pointer w-full h-64 p-3 border rounded shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        ) : (
          <div onClick={() => setEditing(true)} className="cursor-pointer">
            <MarkdownRenderer content={markdown} />
          </div>
        )}
      </div>
    </div>
  );
};


export default NoteEditor
