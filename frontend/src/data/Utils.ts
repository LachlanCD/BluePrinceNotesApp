export async function FetchData(url: string) {
  const response = await fetch(url);
  if (!response.ok) throw new Error('Network response was not ok');
  const data = await response.json();
  return data
}

export async function FetchAndCache(url: string, location: string) {
  const data = FetchData(url)
  localStorage.setItem(location, JSON.stringify(data));
  return data;
}

export async function AddNew(url: string, data: URLSearchParams) {
  const response = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: data.toString(),
  })

  if (!response.ok) throw new Error('Unable to add new item');
  const rData = await response.json();
  return rData;
}

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

