import { useWorkspaceIDRedirect } from "../data/Utils";

export default function RoomHomePage() {
  useWorkspaceIDRedirect("/")

  return (
    <div className="flex flex-col items-center justify-center px-4 text-center mt-20">
      <header className="max-w-2xl mb-12">
        <h1 className="text-5xl font-extrabold text-white sm:text-6xl mb-6">
          Welcome to Blue Prince Notes
        </h1>
        <p className="text-lg text-gray-300 mb-4">
          This app lets you take, share, and manage notes collaboratively for the game <span className="font-semibold text-blue-400">Blue Prince</span>.
        </p>
        <p className="text-lg text-gray-300 mb-4">
          Each workspace is linked to the unique string in your URL, meaning your notes are scoped and accessible via that address.
        </p>
        <p className="text-lg text-gray-300">
          To collaborate, simply share your workspace’s URL — others will be able to view and edit notes in real time.
        </p>
      </header>
    </div>

  );
}
