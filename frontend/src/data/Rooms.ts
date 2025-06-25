import type { NewRoom } from "../types";
import { AddNew, FetchAndCache } from "./Utils";

export async function GETAllRooms() {
  try {
    const route = "/rooms"
    const url = import.meta.env.VITE_BASE_URL + route
    const location = import.meta.env.VITE_ALL_ROOMS_LOCATION
    return FetchAndCache(url.toString(), location)
  } catch (err) {
    return err
  }
}

export async function ADDNewRoom(newRoom: NewRoom) {
  try {
    const formData = new URLSearchParams();
    formData.append('name', newRoom.Name);
    formData.append('colour', newRoom.Colour);

    const route = "/rooms/add"
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    return err
  }
}
