# Resultados de la aplicación

**Proyecto:** Conversor de imágenes ASCII / Pixel Art  
**Fecha de verificación:** 21 de septiembre de 2026  
**Versión evaluada:** estado actual del repositorio (sin etiqueta de release)  
**Alcance:** verificación de construcción y revisión estática del flujo funcional. No se obtuvo una ejecución E2E exitosa porque los bloqueos de construcción impiden levantar una versión verificable de la aplicación.

## 1. Resumen ejecutivo

La aplicación implementa el flujo previsto de carga de imagen, conversión ASCII o Pixel Art, persistencia en MongoDB, consulta de historial y operaciones de edición/eliminación. Sin embargo, el estado actual **no es liberable**: el frontend no compila y las pruebas del backend no se ejecutan desde el repositorio limpio.

| Resultado global | Estado |
|---|---|
| Construcción del frontend | Fallida |
| Pruebas/compilación del backend | Fallidas |
| Ejecución funcional E2E | No ejecutada: bloqueada |
| Seguridad de acceso | No conforme para exposición pública |
| Recomendación | No liberar hasta cerrar los bloqueos P0. |

## 2. Funcionalidades implementadas observadas

| Funcionalidad | Componentes | Resultado de revisión |
|---|---|---|
| Carga de imágenes | `UploadForm.tsx`, `POST /api/upload` | Implementada; el navegador limita la selección a imágenes y el backend decodifica el contenido. |
| Conversión ASCII | `converters/ascii.go` | Implementada con ancho de 100 caracteres y mapa de luminosidad. |
| Conversión Pixel Art | `converters/pixelart.go` | Implementada en bloques de 8×8, genera PNG Base64. |
| Persistencia | servicio, repositorio y MongoDB | Implementada mediante `ConversionRecord` y colección `conversions`. |
| Historial | `GET /api/conversions`, `ConversionList.tsx` | Implementada estáticamente; no verificada en ejecución E2E. |
| Edición y eliminación | API, `ResultDisplay.tsx` | Implementadas parcialmente; no funcionan integradas porque faltan callbacks obligatorios en `App.tsx`. |

## 3. Resultados de verificación reproducible

| ID | Verificación | Comando / condición | Resultado | Evidencia |
|---|---|---|---|---|
| R-01 | Build frontend | `cd frontend && npm run build` | Falló | `TS2739`: `ConversionList` requiere `onUpdate` y `onDelete`, que `App.tsx` no entrega. |
| R-02 | Pruebas backend | `cd backend && go test ./...` | Falló | Falta declarar `golang.org/x/image/webp` en `go.mod`. |
| R-03 | Pruebas unitarias | Búsqueda de archivos de prueba | No disponible | No se encontraron pruebas de frontend o backend en el repositorio. |
| R-04 | Flujo E2E | Frontend + API + MongoDB | Bloqueado | Depende de R-01; no es posible validar la interfaz compilada. |
| R-05 | Control de acceso | Revisión de middleware/API | Falló | No hay autenticación/autorización y CORS permite cualquier origen. |

## 4. Estado frente a los criterios de aceptación del SRS

| Criterio | Resultado | Justificación |
|---|---|---|
| CA-01: PNG a ASCII | Pendiente | Algoritmo presente, sin ejecución/prueba registrada. |
| CA-02: JPEG a Pixel Art | Pendiente | Algoritmo presente, sin ejecución/prueba registrada. |
| CA-03: archivo no imagen | Pendiente | Existe manejo HTTP 400, sin prueba ejecutada. |
| CA-04: carga superior al límite | Pendiente | Existe límite de formulario, sin prueba ejecutada. |
| CA-05: historial | Bloqueado | La UI no compila. |
| CA-06: cambio de nombre | Fallido/bloqueado | Falta sincronización del estado mediante `onUpdate`. |
| CA-07: eliminación | Fallido/bloqueado | Falta sincronización del estado mediante `onDelete`. |
| CA-08: compilación en clon limpio | Fallido | R-01 y R-02 fallan. |
| CA-09: acceso con permisos | Fallido | El sistema no implementa identidad ni autorización. |

## 5. Evidencias disponibles

- [SRS.md](SRS.md): requisitos y criterios de aceptación.
- [PLAN_PRUEBAS_IEEE_829.md](PLAN_PRUEBAS_IEEE_829.md): casos, incidencias y registro de ejecución.
- [PLAN_PRUEBAS_ISTQB.md](PLAN_PRUEBAS_ISTQB.md): estrategia basada en riesgo y decisión de no liberar.
- [EVALUACION_ISO_25010.md](EVALUACION_ISO_25010.md): calidad de producto y plan técnico.

## 6. Resultado esperado de la siguiente ejecución

Después de implementar los cambios P0 de [CAMBIOS_SUGERIDOS.md](CAMBIOS_SUGERIDOS.md), se debe repetir:

```text
cd frontend && npm run build
cd backend && go test ./...
```

Si ambos comandos finalizan correctamente, se deben ejecutar los casos P0 y P1 del plan IEEE 829, adjuntar capturas/salidas al registro de pruebas y actualizar esta tabla con resultados reales de la aplicación.
