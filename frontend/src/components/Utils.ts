export function GetHexCode(colour: string|undefined): string {
  if (!colour) return "ffffff";
  const colourMap: { [key: string]: string } = {
    blue: "#1d6fb7",
    purple: "#7b397b",
    orange: "#bf652a",
    yellow: "#e4ca39",
    green: "#3f982a",
    red: "#d01719",
    black: "#000000",
  };

  return colourMap[colour.toLowerCase()];
}
