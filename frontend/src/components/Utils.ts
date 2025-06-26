export function GetHexCode(color: string): string | undefined {
  const colorMap: { [key: string]: string } = {
    blue: "#0000FF",
    purple: "#800080",
    orange: "#FFA500",
    yellow: "#FFFF00",
    green: "#008000",
    red: "#FF0000",
    black: "#000000",
  };

  return colorMap[color.toLowerCase()];
}
