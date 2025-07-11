import { Link } from "react-router-dom";
import { GetHexCode } from "./Utils";
import { type Card } from "../types";

export type CardProps = {
  workspaceID: string;
  room: Card;
}

export default function BaseCard({ workspaceID, room }: CardProps) {
  console.log(room)
  const bc = GetHexCode(room.Colour);
  const link = getLink(workspaceID, room.Colour, room.Id);
  return (
    <div>
      <Link to={link}>
        <div className="text-md text-center text-stone-700 dark:text-zinc-200 font-bold shadow sm:rounded-lg border-2 w-35 h-35 place-content-center text-wrap transform hover:scale-115"
          style={room.Colour ? { borderColor: bc } : undefined}
        >
          <h3 className="p-3"
            style={room.Colour ? { color: bc } : undefined}
          >
            {room.Name}
          </h3>
        </div>
      </Link>
    </div>
  );
};

function getLink(workspaceID: string, colour: string | undefined, id: number) {
  if (!colour) return `/${workspaceID}/generals/${id}`
  return `/${workspaceID}/rooms/${id}`
}
