import { HandleNoteUpdate } from "../data/Utils";
import type { UpdateNoteProps } from "../types";
import MarkdownRenderer from "./MarkdownRenderer";

export type noteEditorProps = {
  setEditing: (value: boolean) => void;
  editing: boolean;
  setMarkdown: (value: string) => void;
  markdown: string;
  id: number | undefined;
  workspaceID: string | null;
  handleSubmit: ({ id, workspaceID, note }: UpdateNoteProps) => void;
}

const NoteEditor = ({ setEditing, editing, setMarkdown, markdown, id, workspaceID, handleSubmit }: noteEditorProps) => {
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

  if (!workspaceID) return;

  return (
    <div className="mx-20">
      <div className="pt-4 bg-slate-50/90 rounded-lg">
        <div className="p-5">
          {editing ? (
            <textarea
              onClick={() => setEditing(true)}
              onBlur={() => { setEditing(false); HandleNoteUpdate({ id, workspaceID, note: markdown, updateNote: handleSubmit }) }}
              autoFocus
              value={markdown}
              onChange={(e) => setMarkdown(e.target.value)}
              onKeyDown={handleTab}
              className="cursor-pointer bg-white w-full h-200 p-3 border rounded shadow-sm focus:outline-none border-slate-50 text-black"
            />
          ) : (

            <div onClick={() => setEditing(true)} className="cursor-pointer">
              <MarkdownRenderer content={markdown || "Click to enter some notes."} />
            </div>
          )}
        </div>
      </div>
    </div>
  );
};


export default NoteEditor
