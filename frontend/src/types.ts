export interface NewRoom extends NewGeneral {
  Colour: string
};

export interface NewGeneral {
  Name: string
};

export interface GeneralCard {
  Id: number
  Name: string
};

export interface RoomCard extends GeneralCard {
  Colour: string
};

export interface Card extends GeneralCard {
  Colour?: string
};

export interface GeneralNote {
  Id: number
  Name: string
  Notes: string
};

export interface RoomNote extends GeneralNote {
  Colour: string
};

export type UpdateNoteProps = {
  id: number;
  workspaceID: string;
  note: string;
}
