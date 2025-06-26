import { type RoomCard } from "../types";
import Card from "./Card";
import { GetHexCode } from "./Utils";

export type SecitonProps = {
  colour: string;
  rooms: RoomCard[];
}

export default function ColouredSection({ colour, rooms }: SecitonProps) {
  const bgColour = GetHexCode(colour)
  return (
    <div className="mt-5">
      <hr style={{ borderColor: bgColour }} />
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
        {rooms.map((room, id) => (
          <div className="card" key={id}>
            <Card {...room} />
          </div>
        )
        )}
      </div>
      <hr style={{ borderColor: bgColour }} />
    </div>
  )

}
