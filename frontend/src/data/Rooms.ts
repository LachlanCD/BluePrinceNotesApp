import type { NewRoom, RoomNote } from "../types";
import { AddNew, DeleteItem, FetchAndCache, FetchData } from "./Utils";

export async function GETAllRooms() {
  try {
    const route = "/rooms"
    const url = import.meta.env.VITE_BASE_URL + route
    const location = import.meta.env.VITE_ALL_ROOMS_LOCATION
    return FetchAndCache(url, location)
  } catch (err) {
    return err
  }
}

export async function GETRoomDetails(id: string|undefined) {
  try {
    const route = "/rooms/" + id
    const url = import.meta.env.VITE_BASE_URL + route
    return FetchData(url)
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

export type FormatNewRoomProps = {
  name: string;
  colour: string;
  navigate: (value: string) => void;
}

export async function FormatNewRoom({ name, colour, navigate }: FormatNewRoomProps) {
  const formData = {
    Name: name,
    Colour: colour,
  };

  try {
    await ADDNewRoom(formData)
    navigate('/')
  } catch (err) {
    throw err;
  }
}

export async function UpdateRoom(newRoom: RoomNote) {
  try {
    const formData = new URLSearchParams();
    formData.append('name', newRoom.Name);
    formData.append('colour', newRoom.Colour);
    formData.append('notes', newRoom.Notes);

    const route = `/rooms/${newRoom.Id}/update`
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    throw err;
  }
}

export async function DeleteRoom(id: string|undefined) {
  try {
    const route = `/rooms/${id}/remove`
    const url = import.meta.env.VITE_BASE_URL + route
    return DeleteItem(url)
  } catch (err) {
    return err
  }
}
