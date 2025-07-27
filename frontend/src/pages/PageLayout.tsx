import { Outlet, Navigate, useParams } from "react-router-dom";
import Navbar from "../components/Navbar";

export default function WorkspaceLayout() {
  const { workspaceID } = useParams<{ workspaceID: string }>();

  const reservedRoutes = ['rooms', 'generals', 'add-new'];

  if (workspaceID && reservedRoutes.includes(workspaceID)) {
     return <Navigate to='/' replace />;
  }

  return (
    <div>
      <Navbar workspaceID={workspaceID!} />
      <main>
        <Outlet />
      </main>
    </div>
  );
}
