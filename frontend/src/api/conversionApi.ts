import type  { ConversionRecord, ConversionType, ConversionUpdate } from "../types/conversion";

const BASE_URL= "/api"

export async function uploadImage(
  file: File,
  type: ConversionType
): Promise<ConversionRecord> {
  const formData = new FormData();
  formData.append("image", file);
  formData.append("type", type);

  const response = await fetch(`${BASE_URL}/upload`, {
    method: "POST",
    body: formData,
  });

  if (!response.ok) {
    const errorBody = await response.json().catch(() => null);
    throw new Error(errorBody?.error ?? "Error al subir la imagen");
  }

  return response.json();
}

export async function fetchConversions(): Promise<ConversionRecord[]> {
  const response = await fetch(`${BASE_URL}/conversions`);

  if (!response.ok) {
    throw new Error("Error al obtener las conversiones");
  }

  const body = await response.json().catch(() => null);
  return body ?? [];
}

export async function replaceConversion(id: string, filename: string, type: ConversionType): Promise<ConversionRecord>  {
  const response = await fetch(`${BASE_URL}/conversions/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ filename, type }),
  });

  if (!response.ok) {
    const errorBody = await response.json().catch(()=> null);
    throw new Error(errorBody?.error ?? "Error al actualizar la convercion");
  }

  return response.json();
}


export async function updateConversion(id: string, update: ConversionUpdate): Promise<ConversionRecord> {
  const response = await fetch(`${BASE_URL}/conversions/${id}`,{
    method: "PATCH",
    headers:  { "Content-Type": "application/json" },
    body: JSON.stringify(update),
  })
  if(!response.ok){
    const errorBody = await response.json().catch(() => null);
    throw new Error(errorBody?.error ?? "Error al actualizar la convercion");
  }
  return response.json();

}

export async function deleteConversion (id:string): Promise<void>{
    const response = await fetch(`${BASE_URL}/conversions/${id}`,{
    method: "DELETE",
    })

    if (!response.ok) {
      throw new Error("Error al eliminar la conversion");
      
    }
}