import type { NewGeneral } from "../types";
import { AddNew, FetchAndCache } from "./Utils";

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
  setName: (value: string) => void;
}

export async function FormatNewGeneral({name, setName}: FormatNewGeneralProps){
    const formData = {
      Name: name,
    };

    try {
      await ADDNewGeneral(formData)
      setName('');
    } catch (err) {
      throw err;
    }
}
