import type { ConversionRecord } from "../types/conversion";
import { ResultDisplay } from "./ResultDisplay";

interface ConversionListProps {
  conversions: ConversionRecord[];
  onUpdate: (updated: ConversionRecord) => void;
  onDelete: (id: string) => void;
}
export function ConversionList({ conversions, onUpdate, onDelete}: ConversionListProps) {
  if (!conversions || conversions.length === 0) {
    return <p>Todavía no hay conversiones.</p>;
  }

  return (
    <div>
      {conversions.map((record) => (
        <ResultDisplay
          key={record.id}
          record={record}
          onUpdate={onUpdate}
          onDelete={onDelete}
        />
      ))}
    </div>
  );
}