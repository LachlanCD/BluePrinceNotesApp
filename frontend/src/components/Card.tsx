import { GetHexCode } from "../data/Utils";
import { type Room } from "../types";

export type CardProps = {
  room: Room;
}

export default function Card({ room }: CardProps) {
  if (room.Colour != null) {
    return roomCard(room)
  }
  return generalCard(room)
};

function roomCard({ Id, Name, Colour }: Room) {
  const bg = GetHexCode(Colour);
  return (
    <div>
      <a href={`/rooms/${Id}`}>
        <div className="text-lg text-center text-gray-300 font-bold shadow sm:rounded-lg border-5 max-w-55 h-22 place-content-center text-wrap"
          style={{ borderColor: bg }}
        >
          <h3 className="p-2">{Name}</h3>
        </div>
      </a>
    </div>
  );
}

function generalCard({ Id, Name }: Room) {
  return (
    <div>
      <a href={`/general/${Id}`}>
        <div className="text-lg text-center bg-white font-bold shadow overflow-hidden sm:rounded-lg mb-10">
          <h1 className="py-2">{Name}</h1>
        </div>
      </a>
    </div>
  );
}
