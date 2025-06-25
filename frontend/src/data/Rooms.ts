import { FetchAndCache } from "./Utils";

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
