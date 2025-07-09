import type { GeneralCard, NewGeneral, UpdateNoteProps } from "../types";
import { AddNew, DeleteItem, FetchAndCache, FetchData } from "./Utils";

export async function GETAllGenerals() {
  try {
    const route = "/general"
    const url = import.meta.env.VITE_BASE_URL + route
    const location = import.meta.env.VITE_ALL_GENERALS_LOCATION
    return FetchAndCache(url, location)
  } catch (err) {
    return err
  }
}

export async function GETGeneralDetails(id: string|undefined) {
  try {
    const route = "/general/" + id
    const url = import.meta.env.VITE_BASE_URL + route
    return FetchData(url)
  } catch (err) {
    return err
  }
}

export async function ADDNewGeneral(newGeneral: NewGeneral) {
  try {
    const formData = new URLSearchParams();
    formData.append('name', newGeneral.Name);

    const route = "/general/add"
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

export async function FormatNewGeneral({ name, navigate }: FormatNewGeneralProps) {
  const formData = {
    Name: name,
  };

  try {
    await ADDNewGeneral(formData)
    navigate('/generals')
  } catch (err) {
    throw err;
  }
}

export async function UpdateGeneral(newGeneral: GeneralCard) {
  try {
    const formData = new URLSearchParams();
    formData.append('name', newGeneral.Name);

    const route = `/general/${newGeneral.Id}/update`
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    throw err;
  }
}

export async function UpdateGeneralNote({ id, note }: UpdateNoteProps) {
  try {
    const formData = new URLSearchParams();
    formData.append('notes', note);

    const route = `/general/${id}/update/note`
    const url = import.meta.env.VITE_BASE_URL + route
    return AddNew(url, formData)
  } catch (err) {
    throw err;
  }
}

export async function DeleteGeneral(id: string|undefined) {
  try {
    const route = `/general/${id}/remove`
    const url = import.meta.env.VITE_BASE_URL + route
    return DeleteItem(url)
  } catch (err) {
    return err
  }
}
