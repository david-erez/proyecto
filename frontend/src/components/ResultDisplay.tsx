import type { ConversionRecord } from "../types/conversion";

interface ResultDisplayProps {
    record: ConversionRecord;
}

export function ResultDisplay( {record}: ResultDisplayProps){
    return(
        <div>
            <h3>{record.filename}</h3>

            {record.type === "ascii" ? (
                <pre
                    style={{
                        fontFamily: "monospace",
                        fontSize: "6px",
                        lineHeight: "6px",
                        whiteSpace: "pre",
                        overflow: "auto",
                        background: "#000",
                        color: "#0f0",
                        padding: "8px",
                    }}
                >
                    {record.result}
                </pre>
      ) : (
        <img
          src={`data:image/png;base64,${record.result}`}
          alt={record.filename}
          style={{ imageRendering: "pixelated", maxWidth: "100%" }}
        />
      )}
        </div>

    );
}