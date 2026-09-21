# Cambios sugeridos

**Proyecto:** Conversor de imágenes ASCII / Pixel Art  
**Fecha:** 21 de septiembre de 2026  
**Fuente:** SRS, evaluaciones ISO/IEC 25010, ISO/IEC 29110, ISO 9001 y planes IEEE 829/ISTQB.

## Prioridad de implementación

| Prioridad | Cambio sugerido | Motivo | Criterio de verificación |
|---|---|---|---|
| P0 | En `App.tsx`, implementar y pasar `onUpdate` y `onDelete` a `ConversionList`; actualizar/eliminar el registro en el estado `conversions`. | El frontend no compila y las funciones de editar/eliminar no se sincronizan. | `npm run build` termina correctamente; TC-13, TC-14 y TC-16 pasan. |
| P0 | Declarar `golang.org/x/image/webp` en `backend/go.mod` y actualizar `go.sum`. | `go test ./...` falla fuera de la imagen Docker. | `go test ./...` termina correctamente desde un clon limpio. |
| P0 | Definir la regla de negocio para cambiar `type`: reconvertir usando el archivo original almacenado o impedir el cambio de tipo. Validar `filename` no vacío y tipos permitidos. | Evita registros con un resultado que no corresponde al tipo indicado. | PUT/PATCH con datos inválidos devuelven 400; TC-13 a TC-15 pasan. |
| P0 | Antes de exponer el sistema: implementar autenticación, autorización por propietario o rol, CORS por lista permitida y proteger/desactivar el túnel público. | Cualquier cliente puede leer, modificar o borrar todas las conversiones. | Rutas de datos devuelven 401/403 sin permiso; TC-19 pasa. |
| P0 | Limitar de forma robusta las cargas con `http.MaxBytesReader`, límite de dimensiones/píxeles y limitación de tasa. | Reduce consumo excesivo de CPU/RAM y abuso de la API. | Cargas fuera de límite se rechazan sin degradar el servicio; TC-04 a TC-06 pasan. |
| P1 | Crear pruebas unitarias de ASCII/Pixel Art, servicios y validaciones; pruebas de integración para API y MongoDB. | No existe evidencia automatizada de que los requisitos se cumplan. | Cobertura de rutas críticas y `go test ./...` en CI. |
| P1 | Añadir pruebas E2E de carga, historial, edición y eliminación; ejecutar CA-01 a CA-09. | Convierte el SRS en evidencia de aceptación real. | Matriz requisito–prueba–resultado completa y aprobada. |
| P1 | Mejorar accesibilidad: `label` de controles, foco visible, `aria-live` para errores/estado y mensajes de fallo inicial de lista. | La interfaz no cumple aún los RNF de accesibilidad del SRS. | TC-20 y revisión WCAG 2.2 AA básica pasan. |
| P1 | Crear contrato OpenAPI, estandarizar mensajes/códigos de error y añadir orden/paginación al historial. | Mejora interoperabilidad, mantenibilidad y capacidad. | Contrato publicado y pruebas API aprobadas. |
| P1 | Fijar versiones de imágenes Docker, documentar variables de entorno y reemplazar el README de Vite por documentación del producto. | El despliegue no es plenamente reproducible ni está documentado. | Construcción/puesta en marcha validada desde un entorno limpio. |
| P2 | Agregar `/health`, *timeouts*, logs estructurados, métricas y copias/restauración de MongoDB. | Mejora operación, fiabilidad y diagnóstico. | Health check, tablero mínimo y restauración verificada. |
| P2 | Establecer CI para build, lint, pruebas, análisis de dependencias y publicación de evidencia. | Evita regresiones y da trazabilidad de calidad. | Pipeline obligatorio y verde en rama principal. |
| P2 | Formalizar SGC ligero: responsable, control de cambios del SRS, riesgos, objetivos de calidad, revisión y acciones correctivas. | Cierra las brechas ISO 9001 e ISO/IEC 29110. | Registros de gestión y revisión periódica disponibles. |

## Orden recomendado

1. Resolver los cinco cambios P0 y repetir la construcción.
2. Automatizar las pruebas P0/P1 y ejecutar el plan IEEE 829.
3. Aplicar mejoras de accesibilidad, API, documentación y reproducibilidad.
4. Incorporar observabilidad, CI y gestión de calidad continua.

## Cambios que requieren decisión del responsable de producto

| Decisión | Alternativas |
|---|---|
| Cambio de tipo de conversión | Conservar el archivo original y reconvertir; o permitir solo cambio de nombre y prohibir modificar `type`. |
| Propiedad de conversiones | Historial público temporal; o cuentas/propietario y autorización obligatoria. |
| Historial | Definir orden, paginación, retención y máximo de resultados. |
| Límites operativos | Definir máximo de píxeles, formatos finales, concurrencia, tiempo de conversión y tamaño de salida. |

La implementación de los P0 es condición para generar resultados funcionales verificables y considerar una liberación de prueba.
