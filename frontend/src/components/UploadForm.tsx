import { useState } from "react"
import type { ConversionType, ConversionRecord } from "../types/conversion"
import { uploadImage } from "../api/conversionApi"


interface UploadFormProps{
  onUploadSuccess: (record: ConversionRecord) => void;
}

export function UploadForm({ onUploadSuccess }: UploadFormProps) {
    const [file,  setFile] = useState<File | null>(null)
    const [type, setType] = useState<ConversionType>("ascii")
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null);
    
    function handleFileChange(e:  React.ChangeEvent<HTMLInputElement>) {
        const selected =e.target.files?.[0] ?? null;

        setFile(selected);
        setError(null);
    }

    async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();

    if (!file) {
      setError("Selecciona una imagen primero");
      return;
    }

    setIsLoading(true);
    setError(null);

    try {
      const record = await uploadImage(file, type);
      if (!record || typeof record !== "object") {
        setError("Respuesta inválida del servidor");
        return;
      }

      onUploadSuccess(record);
      setFile(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Error desconocido");
    } finally {
      setIsLoading(false);
    }
  }
  return(
    <form onSubmit={handleSubmit}>
      <input type="file" accept="image/*" onChange={handleFileChange} />

      <select
        value={type}
        onChange={(e) => setType(e.target.value as ConversionType)}
      >
        <option value="ascii">ASCII Art</option>
        <option value="pixelart">Pixel Art</option>
      </select>

      <button type="submit" disabled={isLoading}>
        {isLoading ? "Procesando..." : "Convertir"}
      </button>

      {error && <p style={{ color: "red" }}>{error}</p>}
    </form>
  );
}