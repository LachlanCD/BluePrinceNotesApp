import { Link } from "react-router-dom";
import { GetHexCode } from "./Utils";
import { type Card } from "../types";

export default function BaseCard({ Id, Name, Colour="white" }: Card) {
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
};
