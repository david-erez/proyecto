import type { ConversionRecord } from "../types/conversion";
import { ResultDisplay } from "./ResultDisplay";

interface ConversionListProps {
  conversions: ConversionRecord[];
}

export function ConversionList({ conversions }: ConversionListProps) {
  if (!conversions || conversions.length === 0) {
    return <p>Todavía no hay conversiones.</p>;
  }

  return (
    <div>
      {conversions.map((record) => (
        <ResultDisplay key={record.id} record={record} />
      ))}
    </div>
  );
}