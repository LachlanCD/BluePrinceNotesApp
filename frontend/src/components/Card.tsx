import { Link } from "react-router-dom";
import { GetHexCode } from "./Utils";
import { type RoomCard } from "../types";

export type CardProps = {
  room: RoomCard;
}

export default function Card({ room }: CardProps) {
  if (room.Colour != null) {
    return roomCard(room)
  }
  return generalCard(room)
};

function roomCard({ Id, Name, Colour }: RoomCard) {
  const bc = GetHexCode(Colour);
  return (
    <div>
      <Link to={`/rooms/${Id}`}>
        <div className="text-md text-center text-gray-300 font-bold shadow sm:rounded-lg border-2 w-35 h-35 place-content-center text-wrap transform hover:scale-115"
          style={{ borderColor: bc }}
        >
          <h3 className="p-3">{Name}</h3>
        </div>
      </Link>
    </div>
  );
}

function generalCard({ Id, Name }: RoomCard) {
  return (
    <div>
      <Link to={`/general/${Id}`}>
        <div className="text-lg text-center bg-white font-bold shadow overflow-hidden sm:rounded-lg mb-10 transform hover:scale-115">
          <h1 className="py-2">{Name}</h1>
        </div>
      </Link>
    </div>
  );
}
