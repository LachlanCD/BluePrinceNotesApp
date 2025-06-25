import { FetchAndCache } from "./Utils";

export async function GETAllGenerals() {
  try {
    const route = "/general"
    const url = import.meta.env.VITE_BACKEND_URL + route
    const location = import.meta.env.VITE_ALL_GENERALS_LOCATION
    return FetchAndCache(url, location)
  } catch (err) {
    return err
  }
}
