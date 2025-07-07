import MarkdownRenderer from "./MarkdownRenderer";

export type noteEditorProps = {
  setEditing: (value: boolean) => void;
  editing: boolean;
  setMarkdown: (value: string) => void;
  markdown: string;
  handleSubmit: () => void;
}

const NoteEditor = ({ setEditing, editing, setMarkdown, markdown, handleSubmit }: noteEditorProps) => {
  return (
    <div>
      <button onClick={() => { setEditing(!editing); handleSubmit(); }}>
        {editing ? 'Preview' : 'Edit'}
      </button>
      <div style={{ marginTop: '1rem' }}>
        {editing ? (
          <textarea
            style={{ width: '100%', height: '200px' }}
            value={markdown}
            onChange={(e) => setMarkdown(e.target.value)}
          />
        ) : (
          <MarkdownRenderer content={markdown} />
        )}
      </div>
    </div>
  );
};

export default NoteEditor
