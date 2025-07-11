import type { NewRoom, RoomCard, UpdateNoteProps } from "../types";
import { AddNew, DeleteItem, FetchAndCache, FetchData } from "./Utils";

export async function GETAllRooms(workspaceID: string) {
  if (!workspaceID) return
  try {
    const route = `/rooms/${workspaceID}`
    const url = import.meta.env.VITE_BASE_URL + route
    const location = import.meta.env.VITE_ALL_ROOMS_LOCATION
    return FetchAndCache(url, location)
  } catch (err) {
    return err
  }
}

export async function GETRoomDetails(workspaceID: string, id: string | undefined) {
  try {
    const route = `/rooms/${workspaceID}/${id}`
    const url = import.meta.env.VITE_BASE_URL + route
    return FetchData(url)
  } catch (err) {
    return err
  }
}

export async function ADDNewRoom(workspaceID: string, newRoom: NewRoom) {
  try {
    const formData = new URLSearchParams();
    formData.append('name', newRoom.Name);
    formData.append('colour', newRoom.Colour);

    const route = `/rooms/${workspaceID}/add`
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

export async function FormatNewRoom(workspaceID: string, { name, colour, navigate }: FormatNewRoomProps) {
  const formData = {
    Name: name,
    Colour: colour,
  };

  try {
    await ADDNewRoom(workspaceID, formData)
    navigate(`/${workspaceID}/rooms`)
  } catch (err) {
    throw err;
  }
}

export async function UpdateRoom(workspaceID: string, newRoom: RoomCard) {
  try {
    const formData = new URLSearchParams();
    formData.append('name', newRoom.Name);
    formData.append('colour', newRoom.Colour);

    const route = `/rooms/${workspaceID}/${newRoom.Id}/update`
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    throw err;
  }
}


export async function UpdateRoomNote({ id, workspaceID, note }: UpdateNoteProps) {
  try {
    const formData = new URLSearchParams();
    formData.append('notes', note);

    const route = `/rooms/${workspaceID}/${id}/update/note`
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    throw err;
  }
}

export async function DeleteRoom(workspaceID: string, id: string | undefined) {
  try {
    const route = `/rooms/${workspaceID}/${id}/remove`
    const url = import.meta.env.VITE_BASE_URL + route
    return DeleteItem(url)
  } catch (err) {
    return err
  }
}
