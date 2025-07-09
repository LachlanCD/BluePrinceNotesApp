import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import rehypeRaw from 'rehype-raw';

const MarkdownRenderer = ({ content }: { content: string }) => {
  return (
    <div className="prose text-left text-white">
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
