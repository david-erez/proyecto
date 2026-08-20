import { useEffect, useState } from "react";
import type { ConversionRecord } from "./types/conversion";
import { fetchConversions } from "./api/conversionApi";
import { UploadForm } from "./components/UploadForm";
import { ConversionList } from "./components/ConversionList";

function App() {
  const [conversions, setConversions] = useState<ConversionRecord[]>([]);
  const [isLoadingList, setIsLoadingList] = useState(true);

  useEffect(() => {
    fetchConversions()
      .then(setConversions)
      .catch((err) => console.error("Error cargando conversiones:", err))
      .finally(() => setIsLoadingList(false));
  }, []);

  function handleUploadSuccess(record: ConversionRecord) {
    setConversions((prev) => [record, ...prev]);
  }

  return (
    <div>
      <h1>Conversor de Imágenes ASCII / Pixel Art</h1>

      <UploadForm onUploadSuccess={handleUploadSuccess} />

      <hr />

      {isLoadingList ? (
        <p>Cargando conversiones...</p>
      ) : (
        <ConversionList conversions={conversions} />
      )}
    </div>
  );
}

export default App;