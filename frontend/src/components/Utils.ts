export function GetHexCode(colour: string): string | undefined {
  const colourMap: { [key: string]: string } = {
    white: "#FFFFFF",
    blue: "#0000FF",
    purple: "#800080",
    orange: "#FFA500",
    yellow: "#FFFF00",
    green: "#008000",
    red: "#FF0000",
    black: "#000000",
  };

  return colourMap[colour.toLowerCase()];
}
