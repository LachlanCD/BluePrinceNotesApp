import type { GeneralCard, NewGeneral, UpdateNoteProps } from "../types";
import { AddNew, DeleteItem, FetchAndCache, FetchData } from "./Utils";

export async function GETAllGenerals(workspaceID: string) {
  try {
    const route = `/general/${workspaceID}`
    const url = import.meta.env.VITE_BASE_URL + route
    const location = import.meta.env.VITE_ALL_GENERALS_LOCATION
    return FetchAndCache(url, location)
  } catch (err) {
    return err
  }
}

export async function GETGeneralDetails(workspaceID: string, id: string | undefined) {
  try {
    const route = `/general/${workspaceID}/${id}`
    const url = import.meta.env.VITE_BASE_URL + route
    return FetchData(url)
  } catch (err) {
    return err
  }
}

export async function ADDNewGeneral(workspaceID: string, newGeneral: NewGeneral) {
  try {
    const formData = new URLSearchParams();
    formData.append('name', newGeneral.Name);

    const route = `/general/${workspaceID}/add`
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    return err
  }
}

export type FormatNewGeneralProps = {
  name: string;
  navigate: (value: string) => void;
}

export async function FormatNewGeneral(workspaceID: string, { name, navigate }: FormatNewGeneralProps) {
  const formData = {
    Name: name,
  };

  try {
    await ADDNewGeneral(workspaceID, formData)
    navigate(`/${workspaceID}/generals`)
  } catch (err) {
    throw err;
  }
}

export async function UpdateGeneral(workspaceID: string, newGeneral: GeneralCard) {
  try {
    const formData = new URLSearchParams();
    formData.append('name', newGeneral.Name);

    const route = `/general/${workspaceID}/${newGeneral.Id}/update`
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    throw err;
  }
}

export async function UpdateGeneralNote({ id, workspaceID, note }: UpdateNoteProps) {
  try {
    const formData = new URLSearchParams();
    formData.append('notes', note);

    const route = `/general/${workspaceID}/${id}/update/note`
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    throw err;
  }
}

export async function DeleteGeneral(workspaceID: string, id: string | undefined) {
  try {
    const route = `/general/${workspaceID}/${id}/remove`
    const url = import.meta.env.VITE_BASE_URL + route
    return DeleteItem(url)
  } catch (err) {
    return err
  }
}
