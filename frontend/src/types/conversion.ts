export type ConversionType = "ascii" | "pixelart";

export interface ConversionRecord {
  id: string;
  filename: string;
  type: ConversionType;
  result: string;
  createdAt: string;
}

export interface ConversionUpdate {
  filename?: string;
  type?: ConversionType;
}