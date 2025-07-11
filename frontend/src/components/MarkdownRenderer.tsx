import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import rehypeRaw from 'rehype-raw';

const MarkdownRenderer = ({ content }: { content: string }) => {
  return (
    <div className="bg-white p-8 prose text-left min-h-200 min-w-full overflow-auto">
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        rehypePlugins={[rehypeRaw]}
        remarkRehypeOptions={{ passThrough: ['link'] }}

      >
        {content}</ReactMarkdown>
    </div>
  );
};

export default MarkdownRenderer
