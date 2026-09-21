# Especificación de Requisitos de Software (SRS)

**Proyecto:** Conversor de imágenes ASCII / Pixel Art  
**Versión:** 1.0  
**Fecha:** 21 de septiembre de 2026  
**Estado:** Basado en la implementación actual; los requisitos marcados como pendientes requieren validación del equipo.

## 1. Propósito

Este documento define los requisitos del sistema web que permite convertir imágenes a arte ASCII o Pixel Art, consultar el historial de conversiones y administrar sus metadatos.

Está dirigido a desarrollo, pruebas, despliegue y personas interesadas en validar el alcance funcional del producto.

## 2. Alcance

El sistema ofrece una interfaz web para cargar una imagen, elegir un tipo de conversión y visualizar el resultado. El backend procesa la imagen, guarda el resultado y sus metadatos en MongoDB, y expone una API HTTP para consultar, actualizar o eliminar conversiones.

El alcance actual incluye:

- Conversión a arte ASCII.
- Conversión a Pixel Art en PNG codificado en Base64.
- Historial persistente de conversiones.
- Cambio de nombre y tipo de una conversión.
- Eliminación de conversiones.

No forman parte del alcance actual: cuentas de usuario, recuperación de archivos originales, compartición social, descarga explícita de resultados, procesamiento por lotes, cuotas por usuario ni administración de roles.

## 3. Descripción general

### 3.1 Actores

| Actor | Descripción |
|---|---|
| Usuario | Persona que carga una imagen, consulta resultados y administra conversiones. |
| Cliente web | Aplicación React/TypeScript que consume la API. |
| API de conversión | Servicio Go que valida, convierte y persiste datos. |
| Base de datos | MongoDB, responsable de persistir los registros de conversión. |

### 3.2 Arquitectura y dependencias

```text
Usuario → Frontend React → Nginx (/api) → Backend Go → MongoDB
                                  │
                                  └────────→ respuesta JSON
```

- El frontend consume la API bajo la ruta relativa `/api`.
- Nginx dirige `/` al frontend y `/api/` al backend.
- Docker Compose orquesta Nginx, frontend, backend y MongoDB.
- Un túnel ngrok puede publicar el proxy para acceso externo si se configura `NGROK_AUTHTOKEN`.

## 4. Requisitos funcionales

### RF-01. Carga de imagen

El sistema deberá permitir al usuario seleccionar una imagen desde el navegador y enviarla al servicio para conversión.

- El campo multipart se denominará `image`.
- La interfaz deberá limitar la selección a tipos de archivo de imagen.
- El servicio deberá rechazar solicitudes que no incluyan el archivo con un error HTTP 400.
- El servicio deberá verificar que el contenido pueda decodificarse como imagen, sin confiar únicamente en la extensión del archivo.

### RF-02. Formatos de entrada

El sistema deberá aceptar imágenes PNG, JPEG, GIF y WebP que puedan ser decodificadas por el backend.

Las imágenes no válidas deberán recibir una respuesta HTTP 400 con un mensaje de error comprensible.

### RF-03. Límite de carga

El sistema deberá admitir solicitudes de carga de hasta 10 MiB, según el límite configurado actualmente en el backend.

Las cargas que excedan el límite o no puedan procesarse deberán recibir una respuesta HTTP 400.

> Pendiente de decisión: establecer un límite máximo de resolución/píxeles, además del límite de tamaño en bytes.

### RF-04. Selección de conversión

El usuario deberá poder elegir uno de los siguientes tipos:

| Valor de API | Nombre mostrado | Resultado |
|---|---|---|
| `ascii` | ASCII Art | Cadena de texto con caracteres ASCII y saltos de línea. |
| `pixelart` | Pixel Art | Imagen PNG codificada en Base64. |

Si no se envía `type`, el backend deberá usar `ascii` como valor predeterminado. Si el tipo no está soportado, deberá devolver un error.

### RF-05. Conversión a ASCII Art

Para el tipo `ascii`, el sistema deberá:

1. Calcular la luminosidad aproximada de los píxeles de origen.
2. Mapearla al conjunto de caracteres `@%#*+=-:. `.
3. Generar un texto de ancho fijo de 100 caracteres y conservar la proporción visual mediante un factor vertical.
4. Devolver y almacenar el resultado como texto.

### RF-06. Conversión a Pixel Art

Para el tipo `pixelart`, el sistema deberá:

1. Agrupar píxeles en bloques de 8 × 8.
2. Asignar a cada bloque el color promedio de sus píxeles.
3. Generar una imagen PNG.
4. Codificar el PNG en Base64 para devolverlo y almacenarlo.

### RF-07. Persistencia de conversión

Después de una conversión correcta, el sistema deberá guardar un registro con los siguientes campos:

| Campo | Tipo | Descripción |
|---|---|---|
| `id` | cadena | Identificador único del registro. |
| `filename` | cadena | Nombre original del archivo cargado. |
| `type` | `ascii` o `pixelart` | Tipo de conversión aplicada. |
| `result` | cadena | Resultado ASCII o PNG Base64. |
| `createdAt` | fecha/hora | Momento de creación del registro. |

### RF-08. Consulta de historial

El sistema deberá permitir obtener todas las conversiones persistidas mediante `GET /api/conversions`.

Si no existen conversiones, la interfaz deberá informar que todavía no hay resultados. Si ocurre un error de persistencia, la API deberá responder con HTTP 500 y un objeto de error JSON.

> Pendiente de decisión: definir orden de resultados, paginación y cantidad máxima de registros retornados.

### RF-09. Visualización de resultados

La interfaz deberá mostrar cada registro del historial con su nombre y resultado.

- Para `ascii`, deberá mostrar el contenido en un bloque de texto preformateado y desplazable.
- Para `pixelart`, deberá mostrar el resultado como imagen con origen `data:image/png;base64,...` y texto alternativo igual al nombre del archivo.

### RF-10. Actualización completa

El sistema deberá permitir sustituir `filename` y `type` de un registro mediante `PUT /api/conversions/{id}`.

- La solicitud deberá contener JSON con `filename` y `type`.
- La API deberá devolver el registro actualizado.
- La interfaz deberá actualizar el elemento modificado sin recargar toda la página.

> Pendiente de decisión: al modificar `type`, se debe volver a convertir el archivo original o impedir el cambio. El sistema actual solo altera el metadato, por lo que puede no coincidir con el resultado guardado.

### RF-11. Actualización parcial

El sistema deberá permitir cambiar uno o ambos campos `filename` y `type` mediante `PATCH /api/conversions/{id}`.

- La solicitud deberá contener al menos uno de los campos actualizables.
- Una solicitud sin campos deberá devolver HTTP 400.
- La API deberá devolver el registro actualizado.

### RF-12. Eliminación

El sistema deberá permitir eliminar una conversión mediante `DELETE /api/conversions/{id}`.

- La interfaz deberá pedir confirmación antes de la eliminación.
- La API deberá responder HTTP 204 si la eliminación es correcta.
- Si el identificador no existe o es inválido, deberá informar el error correspondiente.
- La interfaz deberá retirar el registro eliminado del historial sin recargar la página.

### RF-13. Respuestas de error

Las respuestas de error de la API deberán usar JSON con la forma:

```json
{ "error": "mensaje descriptivo" }
```

La interfaz deberá mostrar un mensaje comprensible al usuario ante errores de carga, actualización o eliminación.

## 5. Interfaces externas

### 5.1 API HTTP

| Método | Ruta | Cuerpo | Respuesta correcta | Propósito |
|---|---|---|---|---|
| `POST` | `/api/upload` | `multipart/form-data`: `image`, `type` | `200` + conversión | Cargar y convertir. |
| `GET` | `/api/conversions` | — | `200` + lista | Consultar historial. |
| `PUT` | `/api/conversions/{id}` | JSON: `filename`, `type` | `200` + conversión | Actualización completa. |
| `PATCH` | `/api/conversions/{id}` | JSON: `filename?`, `type?` | `200` + conversión | Actualización parcial. |
| `DELETE` | `/api/conversions/{id}` | — | `204` | Eliminar conversión. |

### 5.2 Interfaz de usuario

La pantalla principal deberá contener:

- Título identificando el conversor.
- Selector de archivo de imagen.
- Selector de tipo de conversión.
- Botón para iniciar la conversión.
- Indicador de procesamiento y errores.
- Lista de resultados, con controles para editar y eliminar cada uno.

## 6. Reglas de negocio y validación

| ID | Regla |
|---|---|
| RN-01 | Solo se aceptan los tipos de conversión `ascii` y `pixelart`. |
| RN-02 | Una conversión debe tener un nombre de archivo no vacío. |
| RN-03 | Cada resultado se almacena como una cadena: texto para ASCII y Base64 para Pixel Art. |
| RN-04 | El identificador de una conversión deberá ser un ObjectID válido de MongoDB para operaciones individuales. |
| RN-05 | La eliminación no conserva el registro ni el resultado asociado. |
| RN-06 | Los errores de validación de entrada deberán usar HTTP 400; los errores no controlados de procesamiento o persistencia, HTTP 500. |

## 7. Requisitos no funcionales

### RNF-01. Rendimiento y capacidad

- El sistema deberá rechazar cargas mayores de 10 MiB.
- Se deberán definir y medir objetivos de latencia p95 para carga, conversión y consulta antes de producción.
- El historial deberá incorporar paginación antes de que el volumen de registros comprometa la respuesta o la memoria.

### RNF-02. Seguridad

- Antes de exponer el sistema públicamente, toda operación de consulta o mutación deberá requerir autenticación y autorización por propietario o rol.
- Los orígenes permitidos por CORS deberán configurarse explícitamente; no se deberá usar `*` en producción.
- El servicio público deberá usar HTTPS, limitar solicitudes y validar tamaño y dimensiones de imagen.
- Las credenciales y tokens deberán llegar por variables de entorno o un gestor de secretos; no se almacenarán en el repositorio.

### RNF-03. Fiabilidad

- El sistema deberá disponer de pruebas unitarias para conversión y servicios, e integración para la API y persistencia.
- Deberá ofrecer una comprobación de salud para backend y base de datos antes de producción.
- Las operaciones que dependan de MongoDB deberán tener *timeouts* definidos y errores controlados.
- Se deberá definir y probar un procedimiento de copia y restauración de la base de datos.

### RNF-04. Usabilidad y accesibilidad

- Todos los controles deberán tener etiquetas programáticas asociadas.
- Los mensajes de estado y error deberán anunciarse a tecnologías de asistencia.
- La navegación completa deberá ser posible con teclado y el foco deberá ser visible.
- La interfaz deberá cumplir WCAG 2.2 nivel AA en los flujos de carga, edición y eliminación.

### RNF-05. Mantenibilidad

- El proyecto deberá compilar y pasar pruebas desde un clon limpio con las dependencias declaradas en sus manifiestos.
- El código deberá mantener separación entre presentación, API, negocio y persistencia.
- Deberá existir documentación de instalación, variables de entorno, API y operación.
- Un proceso de CI deberá ejecutar construcción, lint, pruebas y análisis de dependencias en cada cambio.

### RNF-06. Portabilidad y despliegue

- El sistema deberá poder desplegarse con Docker Compose.
- Las imágenes de contenedor deberán fijar versiones concretas en producción para obtener construcciones reproducibles.
- La configuración de puertos, origen CORS, URI de MongoDB, nombre de base de datos y límites deberá poder variar por entorno sin modificar código.

## 8. Criterios de aceptación de alto nivel

| ID | Escenario | Resultado esperado |
|---|---|---|
| CA-01 | Cargar un PNG válido y elegir ASCII | Se muestra y persiste texto ASCII de 100 columnas. |
| CA-02 | Cargar un JPEG válido y elegir Pixel Art | Se muestra y persiste una imagen PNG pixelada. |
| CA-03 | Enviar un archivo que no es imagen | La API devuelve `400` y la interfaz informa el error. |
| CA-04 | Enviar una carga superior al límite | La API la rechaza con `400`. |
| CA-05 | Consultar el historial con registros | Se recibe y muestra cada conversión disponible. |
| CA-06 | Cambiar el nombre de un registro | Se persiste el nombre y la lista se actualiza sin recarga completa. |
| CA-07 | Eliminar una conversión confirmada | La API devuelve `204` y el elemento desaparece de la lista. |
| CA-08 | Compilar desde el repositorio | `npm run build` y `go test ./...` finalizan correctamente. |
| CA-09 | Acceder sin permiso en producción | Las rutas de datos devuelven `401` o `403`. |

## 9. Suposiciones y decisiones pendientes

- Cada registro pertenece actualmente a un historial común, porque no existe identidad de usuario. Si se incorporan cuentas, se debe añadir propietario a `ConversionRecord` y filtrar cada consulta/operación por esa propiedad.
- No se guarda el archivo original; por ello no es posible recalcular una conversión al cambiar su tipo. Se recomienda prohibir ese cambio o conservar el original en almacenamiento controlado.
- El producto está diseñado para conversión de imágenes individuales y de tamaño acotado, no para procesamiento masivo.
- Se debe confirmar la política de conservación de resultados, el volumen esperado de conversiones y los objetivos de rendimiento antes de producción.

## 10. Trazabilidad inicial

| Requisito | Componentes relacionados |
|---|---|
| RF-01 a RF-06 | `frontend/src/components/UploadForm.tsx`, `backend/handlers/conversion_handler.go`, `backend/converters/` |
| RF-07 a RF-08 | `backend/services/conversion_service.go`, `backend/repository/conversion_repository.go`, `backend/models/conversion.go` |
| RF-09 a RF-12 | `frontend/src/components/ConversionList.tsx`, `frontend/src/components/ResultDisplay.tsx`, `frontend/src/api/conversionApi.ts` |
| RNF-06 | `docker-compose.yml`, `nginx/nginx.conf`, Dockerfiles |

