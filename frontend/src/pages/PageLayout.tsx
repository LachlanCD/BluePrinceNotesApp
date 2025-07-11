import { Outlet, useParams } from "react-router-dom";
import Navbar from "../components/Navbar";

export default function WorkspaceLayout() {
  const { workspaceID } = useParams<{ workspaceID: string }>();

  return (
    <div>
      <Navbar workspaceID={workspaceID!} />
      <main>
        <Outlet />
      </main>
    </div>
  );
}
