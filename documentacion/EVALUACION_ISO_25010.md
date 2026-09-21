# Evaluación de calidad — ISO/IEC 25010

**Producto evaluado:** Conversor de imágenes ASCII / Pixel Art  
**Versión de la evaluación:** 1.1  
**Fecha de evaluación:** 21 de septiembre de 2026  
**Línea base de requisitos:** [SRS.md](SRS.md), versión 1.0  
**Alcance:** código fuente disponible de frontend React/TypeScript, API Go, MongoDB y despliegue Docker/Nginx, contrastado contra el SRS. No se realizaron pruebas de carga, penetración ni pruebas con usuarios.

## Resultado ejecutivo

El producto tiene una base funcional clara: recibe una imagen, la convierte a ASCII o Pixel Art, persiste el resultado y ofrece operaciones de consulta, edición y eliminación. La separación frontend/API/repositorio y el despliegue contenedorizado son buenas bases.

Sin embargo, **no está listo para una liberación confiable**. Frente al SRS, los requisitos básicos de carga, conversión, persistencia y consulta tienen implementación parcial o verificable estáticamente, pero los criterios CA-06 a CA-09 no se satisfacen: la aplicación no compila, la actualización/eliminación no se propagan al estado de la pantalla y no existe control de acceso. La revisión detectó además dos bloqueos reproducibles de construcción y ausencia de pruebas automatizadas y controles operativos. La calificación indicativa es **2,1/5** sobre las ocho características aplicables; *Protección* no aplica como característica crítica porque el producto no controla un proceso con riesgo para la vida, salud, propiedad o medio ambiente.

> Esta es una evaluación técnica orientativa basada en evidencia del repositorio, no una certificación ISO/IEC 25010.

## Método y escala

Se aplicaron las nueve características del modelo descrito en [ISO/IEC 25010](https://iso25000.com/index.php/normas-iso-25000/iso-25010). El modelo evalúa la calidad del producto mediante características y subcaracterísticas; por ejemplo, adecuación funcional, eficiencia, compatibilidad e interacción, y también fiabilidad, seguridad, mantenibilidad, flexibilidad y protección. El SRS se usa como criterio de conformidad: un requisito solo se marca como **cumplido** si la implementación o una prueba disponible lo demuestra; **parcial** indica implementación incompleta o no verificable; **no cumplido** indica ausencia o contradicción.

Escala: **1 = deficiente**, **2 = parcial**, **3 = aceptable**, **4 = bueno**, **5 = muy bueno**. Una puntuación refleja la evidencia actual, no requisitos no documentados.

| Característica | Resultado | Nivel | Evidencia principal |
|---|---:|---|---|
| Adecuación funcional | 2/5 | Parcial | Conversión ASCII/Pixel Art y CRUD están implementados, pero el frontend no compila y no hay validación consistente de tipos. |
| Eficiencia de desempeño | 2/5 | Parcial | Límite de formulario de 10 MiB; faltan límites de píxeles, *timeouts*, paginación y medición. |
| Compatibilidad | 3/5 | Aceptable | API HTTP/JSON, proxy Nginx, Compose y decodificación PNG/JPEG/GIF/WebP; no hay contrato versionado. |
| Capacidad de interacción | 2/5 | Parcial | Flujo simple y estados de carga/error; faltan accesibilidad, ayudas, recuperación y la interfaz está bloqueada por compilación. |
| Fiabilidad | 2/5 | Parcial | Errores HTTP básicos y volumen persistente de Mongo; sin pruebas, *health checks*, reintentos ni recuperación controlada. |
| Seguridad | 1/5 | Deficiente | Sin autenticación/autorización; CORS abierto; túnel público; sin límites de abuso ni cabeceras defensivas. |
| Mantenibilidad | 2/5 | Parcial | Capas Go separadas y tipos TS; no hay pruebas/CI, README es plantilla y existen errores/convenciones inconsistentes. |
| Flexibilidad | 3/5 | Aceptable | Contenedores, variables de entorno y patrón de convertidores; falta configuración de límites y escalamiento/observabilidad. |
| Protección | No aplicable | — | El caso de uso no es de seguridad física. Debe reevaluarse si se integra en un dominio con riesgo operativo. |

## Hallazgos por característica

### 1. Adecuación funcional — 2/5

**Evidencia favorable**

- `POST /api/upload` procesa imágenes y permite `ascii` y `pixelart`; `GET`, `PUT`, `PATCH` y `DELETE /api/conversions/{id}` están definidos.
- Los convertidores están aislados tras `ImageConverter`, lo que respalda la pertinencia de las dos funciones ofrecidas.

**Brechas**

- `npm run build` falla con `TS2739`: `App.tsx` usa `ConversionList` sin los props obligatorios `onUpdate` y `onDelete`. Por tanto, el cliente no se puede generar en su estado actual.
- El backend registra `PUT`, `PATCH` y `DELETE`, pero `withCORS` solo anuncia `GET, POST, OPTIONS`; los navegadores bloquearían las mutaciones por la preflight CORS cuando se use un origen distinto.
- `UpdateConversion` valida `tipo == " "`, no una cadena vacía ni el conjunto permitido; los cambios de tipo no vuelven a convertir el resultado. Se pueden guardar registros semánticamente incoherentes.
- El SRS ya define los formatos, casos de error y criterios de aceptación; aún quedan como decisiones pendientes el orden/paginación de historial, el límite de píxeles y la regla de negocio al cambiar `type`.

### 2. Eficiencia de desempeño — 2/5

**Evidencia favorable:** `ParseMultipartForm(10 << 20)` limita el cuerpo del formulario aproximadamente a 10 MiB y los algoritmos trabajan en memoria.

**Brechas y riesgo:** no se limita la dimensión ni el número de píxeles tras decodificar una imagen (riesgo de consumo excesivo de CPU/RAM), no hay paginación en `FindAll`, índices ni orden explícito de MongoDB, ni *timeouts* HTTP/Mongo por solicitud. Tampoco hay presupuesto de latencia, concurrencia, tamaño de resultado o métricas para demostrar comportamiento temporal, uso de recursos y capacidad.

### 3. Compatibilidad — 3/5

**Evidencia favorable:** React se comunica mediante HTTP/JSON; Nginx enruta `/api/`; Docker Compose integra frontend, backend y Mongo; Go admite PNG, JPEG, GIF y WebP para entrada.

**Brechas:** no existe especificación OpenAPI, versionado de API, pruebas de interoperabilidad ni compatibilidad de navegadores. Las imágenes Docker usan etiquetas flotantes (`nginx:latest`, `alpine:latest`), reduciendo la reproducibilidad entre entornos.

### 4. Capacidad de interacción — 2/5

**Evidencia favorable:** el flujo de carga es directo, limita la selección a imágenes en el navegador, muestra procesamiento/errores y pide confirmación antes de eliminar.

**Brechas:** faltan `label` asociados a los controles, mensajes anunciados a tecnologías asistivas (`aria-live`), foco gestionado y una alternativa accesible a `window.confirm`. El resultado ASCII usa texto de 6 px y no hay guía de uso, estados vacíos recuperables ni notificación de error al fallar la carga inicial (solo se escribe en consola). La inclusión y asistencia no están demostradas.

### 5. Fiabilidad — 2/5

**Evidencia favorable:** las respuestas usan códigos HTTP y JSON de error, se cierra el archivo recibido y Mongo tiene volumen persistente.

**Brechas:** `go test ./...` falla porque `golang.org/x/image/webp` se importa pero no está declarado en `go.mod`; el Dockerfile lo instala durante la construcción, haciendo que el repositorio no sea autocontenido para pruebas locales. No existen pruebas automatizadas, *health checks*, manejo de apagado ordenado, reintentos/backoff, copias de seguridad, objetivos de disponibilidad ni procedimiento de recuperación.

### 6. Seguridad — 1/5

**Hallazgos críticos**

- No hay autenticación ni autorización: cualquier cliente que alcance la API puede leer, renombrar o eliminar todas las conversiones.
- `Access-Control-Allow-Origin: *` permite invocación desde cualquier origen y el túnel ngrok está configurado para exponer el servicio públicamente.
- No se aplican limitación de tasa, límite explícito con `http.MaxBytesReader`, cabeceras de seguridad, TLS en Nginx, auditoría ni análisis de dependencias.

**Aspectos existentes:** se decodifica la imagen en lugar de confiar en su extensión y Mongo se configura mediante variables de entorno. Esto no compensa las brechas de confidencialidad, integridad, responsabilidad, autenticidad y resistencia.

### 7. Mantenibilidad — 2/5

**Evidencia favorable:** backend dividido en `handlers`, `services`, `repository`, `converters`, `models` y `config`; interfaz `ImageConverter`; frontend con TypeScript y componentes separados.

**Brechas:** no hay archivos de prueba ni integración continua; el README del frontend sigue siendo el de Vite y no documenta el proyecto; hay nombres y textos inconsistentes (`UpdateTan`, `origHeidh`, `dbNmae`, `Convertion`), errores mezclados en inglés/español y lógica de estado incompleta entre `App` y los componentes de edición. El resultado es difícil de analizar, modificar y probar con seguridad.

### 8. Flexibilidad — 3/5

**Evidencia favorable:** Docker separa responsabilidades, Mongo URI/puerto/base de datos se parametrizan y añadir conversores sigue un punto de extensión definido.

**Brechas:** el entorno y límites operativos están codificados parcialmente; no hay configuración por entorno para CORS, almacenamiento, trazabilidad o umbrales; tampoco réplicas, colas ni estrategia de escalado para conversión costosa. Las etiquetas no fijadas de imágenes perjudican la instalabilidad reproducible.

### 9. Protección — no aplicable

La característica se mantiene fuera de la puntuación: no se identificó interacción del sistema con un proceso físico o de alto riesgo. Como medida preventiva, documentar que imágenes maliciosas o cargas excesivas se deben rechazar protege la infraestructura, pero no convierte el producto en un sistema de seguridad funcional.

## Conformidad con el SRS

Esta matriz convierte los requisitos de [SRS.md](SRS.md) en la evidencia de evaluación ISO/IEC 25010. Los requisitos no funcionales son objetivos de liberación definidos por el SRS; que estén documentados no equivale a que estén cumplidos.

| Requisito SRS | Estado | Evidencia o desviación | Característica ISO/IEC 25010 |
|---|---|---|---|
| RF-01 Carga | Parcial | El formulario y `POST /api/upload` existen; no se verifica la ejecución integral por los bloqueos de compilación. | Adecuación funcional, interacción |
| RF-02 Formatos | Parcial | Se importan decodificadores PNG/JPEG/GIF/WebP, pero WebP no está declarado en `go.mod`, por lo que la construcción local falla. | Adecuación funcional, compatibilidad, fiabilidad |
| RF-03 Límite 10 MiB | Parcial | Se usa `ParseMultipartForm(10 << 20)`; falta `http.MaxBytesReader` y límite de resolución/píxeles. | Desempeño, seguridad |
| RF-04 Tipos admitidos | Parcial | `NewConverter` acepta `ascii` y `pixelart`; las actualizaciones no validan el conjunto permitido. | Adecuación funcional |
| RF-05 ASCII | Parcial | Algoritmo y ancho de 100 caracteres implementados; sin pruebas que demuestren exactitud. | Adecuación funcional, fiabilidad |
| RF-06 Pixel Art | Parcial | Promedio en bloques 8×8 y PNG Base64 implementados; sin pruebas de resultado. | Adecuación funcional, fiabilidad |
| RF-07 Persistencia | Parcial | Modelo y `InsertOne` presentes; no hay prueba de integración con MongoDB. | Adecuación funcional, fiabilidad |
| RF-08 Historial | Parcial | `GET` y estado vacío existen; sin orden, paginación ni presentación del error inicial al usuario. | Adecuación funcional, desempeño, interacción |
| RF-09 Visualización | Parcial | `<pre>` e `<img>` con `alt` están presentes; interfaz no compilable y tamaño ASCII poco legible. | Adecuación funcional, interacción |
| RF-10 PUT | No cumplido | API existe, pero el cambio de `type` no reconvierte y el estado padre no implementa `onUpdate`; CA-06 falla. | Adecuación funcional, fiabilidad |
| RF-11 PATCH | No cumplido | API existe, pero el estado padre no implementa `onUpdate`; `type` no se valida y no se reconvierte. | Adecuación funcional, fiabilidad |
| RF-12 DELETE | No cumplido | Confirmación y endpoint existen, pero falta `onDelete` en `App`; no se puede compilar ni cumplir CA-07. | Adecuación funcional, interacción |
| RF-13 Errores | Parcial | Backend devuelve `{ "error": ... }`; el fallo inicial de lista solo se registra en consola. | Adecuación funcional, interacción |
| RNF-01 Rendimiento | No cumplido | Sin p95, pruebas de carga, paginación ni límite de píxeles. | Desempeño |
| RNF-02 Seguridad | No cumplido | Sin autenticación/autorización, CORS `*`, sin TLS en Nginx ni limitación de tasa. | Seguridad |
| RNF-03 Fiabilidad | No cumplido | Sin pruebas, health check, timeouts operativos ni respaldo/restauración. | Fiabilidad |
| RNF-04 Accesibilidad | No cumplido | Sin etiquetas programáticas, `aria-live` ni verificación WCAG 2.2 AA. | Capacidad de interacción |
| RNF-05 Mantenibilidad | No cumplido | Build y pruebas fallan; no hay CI y la guía del frontend sigue siendo plantilla. | Mantenibilidad |
| RNF-06 Portabilidad | Parcial | Docker Compose y variables de entorno existen; imágenes con etiquetas flotantes y CORS/límites no configurables por entorno. | Flexibilidad, compatibilidad |

### Cobertura de criterios de aceptación

| Criterio | Estado | Motivo |
|---|---|---|
| CA-01 y CA-02 | Pendiente de prueba | Los algoritmos existen, pero no hay prueba automatizada ni ejecución E2E válida. |
| CA-03 | Parcial | Hay rechazo de imagen inválida; falta una prueba repetible. |
| CA-04 | Parcial | Existe límite del formulario; no se valida con `MaxBytesReader` ni prueba. |
| CA-05 | Parcial | API y vista de lista existen; la vista no compila. |
| CA-06 y CA-07 | No cumplido | Callbacks obligatorios de actualización y eliminación no se pasan desde `App.tsx`. |
| CA-08 | No cumplido | Fallan `npm run build` y `go test ./...`. |
| CA-09 | No cumplido | No existe autenticación ni autorización. |

## Plan de mejora priorizado

| Prioridad | Acción verificable | Características impactadas | Criterio de salida |
|---|---|---|---|
| P0 | Corregir el estado del cliente: pasar desde `App` callbacks que actualicen/eliminar registros; ejecutar build. Declarar `golang.org/x/image` en `go.mod` y asegurar compilación fuera de Docker. | Adecuación, fiabilidad, mantenibilidad | `npm run build` y `go test ./...` terminan con código 0. |
| P0 | Proteger la API: autenticación, autorización por propietario, CORS por lista permitida, TLS en el borde, limitar tasa/tamaño/píxeles y desactivar o proteger el túnel público. | Seguridad, fiabilidad, desempeño | Pruebas demuestran 401/403 sin permisos y rechazo de cargas que superen límites. |
| P1 | Definir contrato OpenAPI y validación: tipos permitidos, esquema de errores, paginación/orden y que al cambiar tipo se reconvierta o se prohíba explícitamente. | Adecuación, compatibilidad | Contrato publicado y pruebas de integración cubren todos los endpoints y errores. |
| P1 | Añadir pruebas unitarias (convertidores/servicios), integración (API+Mongo) y E2E accesible; ejecutarlas en CI junto con lint y análisis de dependencias. | Fiabilidad, mantenibilidad, interacción | Cobertura objetivo acordada (sugerido ≥80 % del código crítico) y pipeline obligatorio en cambios. |
| P1 | Mejorar UX y accesibilidad: etiquetas, foco, `aria-live`, mensajes de carga inicial, contraste/tamaño legible y ayuda sobre formatos/límites. | Capacidad de interacción | Revisión automatizada y manual WCAG 2.2 AA de los flujos subir, editar y eliminar. |
| P2 | Añadir observabilidad y operación: `/health`, logs estructurados sin datos sensibles, métricas de latencia/errores/recursos, *timeouts*, backups y restauración probada. | Desempeño, fiabilidad, flexibilidad | Tablero de métricas, alerta básica y simulacro de restauración documentado. |
| P2 | Fijar versiones de imágenes y dependencias, completar README de instalación/arquitectura/variables y normalizar nomenclatura y mensajes. | Compatibilidad, mantenibilidad, flexibilidad | Construcción repetible y guía de puesta en marcha validada desde un entorno limpio. |

## Métricas propuestas para la siguiente evaluación

- **Construcción y fiabilidad:** 100 % de builds y pruebas de la rama principal exitosos; tasa de errores 5xx < 1 %.
- **Desempeño:** p95 de conversión y de listado bajo el objetivo que se acuerde; memoria/CPU y tamaño de resultado medidos por tipo y resolución.
- **Seguridad:** 100 % de endpoints de datos requieren identidad y autorización; 0 vulnerabilidades críticas conocidas sin mitigación; auditoría de mutaciones habilitada.
- **Interacción:** éxito de tareas de subir/editar/eliminar ≥ 95 % en prueba con usuarios y cero incidencias WCAG AA críticas.
- **Mantenibilidad:** cobertura de rutas críticas ≥ 80 %, análisis estático sin errores bloqueantes y documentación de API/operación vigente.

## Evidencia reproducible

Ejecutado sobre el estado del repositorio evaluado:

```text
cd frontend && npm run build
# Falla: TS2739 en src/App.tsx; faltan onUpdate y onDelete en ConversionList.

cd backend && go test ./...
# Falla: no required module provides package golang.org/x/image/webp.
```

La evaluación debe repetirse después de implementar P0 y con resultados de pruebas de carga, seguridad y accesibilidad; esas evidencias permitirán sustituir las puntuaciones indicativas por una medición objetiva.
