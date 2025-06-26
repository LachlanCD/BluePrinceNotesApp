import type { NewGeneral } from "../types";
import { AddNew, FetchAndCache } from "./Utils";

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
    navigate('/general')
  } catch (err) {
    throw err;
  }
}
