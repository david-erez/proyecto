import type { ConversionRecord, ConversionType } from "../types/conversion";
import { useState } from "react";
import { deleteConversion, replaceConversion, updateConversion } from "../api/conversionApi";
import style from "ResultDisplay.module.css"
interface ResultDisplayProps {
    record: ConversionRecord;
    onUpdate: (update:ConversionRecord) => void;
    onDelete: (id: string) => void;
}


export function ResultDisplay( {record,onUpdate ,onDelete}: ResultDisplayProps){
    const [isEnding, setIsEnding] = useState(false);
    const [filenameInput, setFilenameInput] = useState(record.filename);
    const [typeInput, setTypeInput] =  useState<ConversionType>(record.type)
    const [isSaving, setIsSaving] = useState(false)
    const [error, setError] = useState <string | null>(null);
    
    async function handlePatch(){
        setIsSaving(true)
        setError(null)
        try {
            const updated = await updateConversion (record.id, {filename: filenameInput})
            onUpdate(updated);
            setIsEnding(false);
        }
        catch (err){
            setError(err instanceof Error ? err.message : "Error al nombrar")
        }
        finally {
            setIsSaving(false);
        }
    }
    async function handlePut() {
        setIsSaving(true)
        setError(null)
        try {
            const updated = await replaceConversion(record.id,filenameInput,typeInput);
            onUpdate(updated)     
            setIsEnding(false)
    
        }catch (err){
            setError(err instanceof Error ? err.message : "Error al reemplazar");
        }finally {
            setIsSaving(false)
        }
    }

    async function handleDelete() {
        const confirm = window.confirm(`Eliminar "${record.filename}"?`);
        if(!confirm) return;
        try{
            await deleteConversion(record.id);
            onDelete(record.id);
        }catch(err){
            setError(err instanceof Error ? err.message : "Error al eliminar")
        }
    }
    return(
        <div className={style["div-result"]}>
            {isEnding ? (
                    <div>
                    <input value={filenameInput} 
                    onChange={(e) => setFilenameInput(e.target.value)}
                    placeholder="Nombre del archivo"
                    />
                    <select 
                        value={typeInput}
                        onChange={(e)=> setTypeInput(e.target.value as ConversionType)}
                    >
                        <option value="ascii">ASCII art</option>
                        <option value="pixelart">PIXEL art</option>
                    </select>
                    <button onClick={handlePatch} disabled={isSaving}>Guardar solo nombre </button>
                    <button onClick={handlePut} disabled={isSaving}>Reemplazar todo </button>
                    <button onClick={()=> setIsEnding(false)} disabled={isSaving}>Cancelar</button>    
            </div>
            ):(
            <div>
                <h3>{record.filename}</h3>
                <button onClick={() => setIsEnding(true)}>Editar</button>
                <button onClick={handleDelete}>Eliminar</button>
            </div>
            )}

            {error && <p className={style["error-button"]}>{error}</p>}

            {record.type == "ascii" ? ( <pre className={style["result-ascii"]}>{record.result}</pre>
            ):(
                <img
                    src={`data:image/png;base64,${record.result}`}
                    alt={record.filename}
                    style={{ imageRendering: "pixelated", maxWidth: "100%" }}
                />
            )}
        </div>
    )
}


