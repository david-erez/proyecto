# Plan y especificación de pruebas — IEEE 829

**Proyecto:** Conversor de imágenes ASCII / Pixel Art  
**Identificador:** PTP-001  
**Versión:** 1.0  
**Fecha:** 21 de septiembre de 2026  
**Base de requisitos:** [SRS.md](SRS.md) v1.0  
**Estado:** Preparado; ejecución bloqueada parcialmente por fallos de construcción identificados.

## 1. Referencia, objetivo y alcance

IEEE 829-2008 definió la forma y contenido de la documentación de pruebas de software y sistemas. Es un estándar retirado y sustituido por la serie ISO/IEC/IEEE 29119; esta última mantiene la documentación de pruebas como salida de los procesos de prueba. [IEEE 829-2008](https://www.nsi.ie/images/uploads/standards/IEEE-829-2008.pdf) fue reemplazado por la [serie ISO/IEC/IEEE 29119](https://committee.iso.org/sites/jtc1sc7/home/projects/flagship-standards/isoiecieee-29119-series.html). Se adopta IEEE 829 aquí por la solicitud del proyecto y por su estructura práctica de documentos.

Este documento reúne, de forma proporcional a un proyecto pequeño, los artefactos de IEEE 829: plan maestro de pruebas, plan de nivel, diseño, casos, procedimientos, registro, informe de incidencias, transferencia de ítems e informe resumen.

El objetivo es verificar que el sistema cumpla los requisitos RF-01 a RF-13, RNF relevantes y CA-01 a CA-09 del SRS para carga, conversión, persistencia, consulta, actualización y eliminación de conversiones.

**Incluido:** backend Go, frontend React/TypeScript, comunicación HTTP, MongoDB y despliegue local con Docker Compose.  
**Excluido en esta iteración:** prueba de carga a gran escala, prueba de penetración certificada, compatibilidad exhaustiva de navegadores, túnel ngrok público y certificación de accesibilidad; se documentan como riesgos y pruebas futuras.

## 2. Artefactos IEEE 829 y estado

| Artefacto IEEE 829 | Identificador | Sección | Estado |
|---|---|---|---|
| Plan maestro de pruebas | MTP-001 | 3 | Preparado |
| Plan de pruebas de nivel | LTP-001 | 4 | Preparado |
| Especificación de diseño | TDS-001 | 5 | Preparado |
| Especificación de casos | TCS-001 | 6 | Preparado |
| Especificación de procedimientos | TPS-001 | 7 | Preparado |
| Registro de pruebas | TLG-001 | 8 | Inicializado |
| Informe de incidencias | TIR-001 | 9 | Inicializado con dos incidencias |
| Informe de transferencia | TTR-001 | 10 | Pendiente de ejecución |
| Informe resumen de pruebas | TSR-001 | 11 | Preliminar |

## 3. MTP-001 — Plan maestro de pruebas

### 3.1 Ítems bajo prueba

| Ítem | Componentes | Riesgo principal |
|---|---|---|
| Carga y validación | `UploadForm.tsx`, `conversion_handler.go` | Archivo inválido, tamaño excesivo, consumo de recursos. |
| Conversión ASCII | `converters/ascii.go` | Salida incorrecta, proporción o límites de imagen. |
| Conversión Pixel Art | `converters/pixelart.go` | Bloques/color/Base64/PNG incorrectos. |
| Persistencia e historial | servicios, repositorio, MongoDB | Pérdida, listado incorrecto, registros incoherentes. |
| Actualización/eliminación | API, `ResultDisplay.tsx`, `App.tsx` | Estado UI no sincronizado, mutación no autorizada. |
| Entrega y calidad | Dockerfiles, Compose, manifiestos | Construcción no reproducible y dependencias faltantes. |

### 3.2 Características a probar y no probar

| Área | Se prueba | No se prueba en esta iteración |
|---|---|---|
| Funcional | Flujos RF-01 a RF-13 y criterios CA-01 a CA-07. | Procesamiento por lotes, descargas y cuentas, fuera del SRS. |
| Calidad técnica | Build, pruebas unitarias/API, errores, límite de carga, resultado y persistencia. | Benchmarks masivos y comparación visual subjetiva. |
| Seguridad | Revisión de acceso no autenticado, CORS y límites como hallazgos. | Pentest, análisis DAST completo y gestión de identidad aún inexistente. |
| Interfaz | Carga, estados, resultado, editar, eliminar y error visible. | Matriz completa de navegadores y validación WCAG formal. |

### 3.3 Enfoque y niveles

| Nivel | Objetivo | Técnica | Responsable propuesto |
|---|---|---|---|
| Unitario | Validar convertidores y reglas de servicio. | Valores límite, particiones de equivalencia, pruebas de tabla. | Desarrollo |
| Integración | Validar handlers, API y MongoDB. | Solicitudes HTTP controladas y base efímera. | Desarrollo/QA |
| Sistema/E2E | Validar los flujos de usuario definidos por CA. | Escenarios manuales y automatizados en navegador. | QA |
| Aceptación | Confirmar que el SRS responde a la necesidad del usuario. | Ejecución de CA-01 a CA-09 con responsable de producto. | Cliente/responsable |

### 3.4 Criterios de inicio, suspensión y salida

| Tipo | Criterio |
|---|---|
| Inicio | Dependencias instaladas, MongoDB disponible, build de frontend y compilación/pruebas Go exitosas, datos de prueba disponibles. |
| Suspensión | Un ítem bloquea el entorno o impide ejecutar el caso; se crea incidente y se continúa con casos independientes. |
| Reanudación | Incidente resuelto, corrección desplegada en entorno de prueba y verificada por repetición del caso bloqueado. |
| Salida de sistema | 100 % de casos P0/P1 ejecutados; 0 incidencias críticas/altas abiertas; CA aplicables aprobados; informe resumen y aceptación registrados. |
| Salida de release | Además de lo anterior, `npm run build` y `go test ./...` terminan en código 0, y no hay defectos críticos abiertos. |

### 3.5 Riesgos de prueba

| Riesgo | Impacto | Mitigación |
|---|---|---|
| Frontend no compila. | Bloquea E2E y aceptación. | Resolver TIR-001 antes de iniciar UI. |
| Dependencia WebP no declarada. | Bloquea pruebas Go en clon limpio. | Resolver TIR-002 y ejecutar en entorno limpio. |
| No hay datos/imágenes de referencia. | Resultados no reproducibles. | Añadir activos de prueba pequeños, válidos e inválidos. |
| MongoDB compartido. | Pruebas no deterministas o borrado de datos. | Usar base aislada `imgapp_test` y limpiar solo datos identificados. |
| Sin autenticación. | No puede validarse CA-09 ni seguridad de datos. | Tratar como defecto de diseño; implementar identidad antes de prueba de autorización. |

## 4. LTP-001 — Plan de pruebas de nivel

### 4.1 Entorno

| Componente | Configuración requerida |
|---|---|
| SO y contenedores | Docker y Docker Compose disponibles. |
| Frontend | Node 22, dependencias bloqueadas por `package-lock.json`, puerto mediante Nginx. |
| Backend | Go declarado por `go.mod`, variables `MONGO_URI`, `PORT`, `DATABASE_NAME`. |
| Base de datos | MongoDB 8 con base exclusiva de pruebas. |
| Navegador | Chromium o Firefox actual; registrar versión usada. |
| Datos | PNG/JPEG/GIF/WebP válidos, archivo no imagen, archivo mayor de 10 MiB y registros semilla. |

### 4.2 Datos de prueba mínimos

| ID | Dato | Uso |
|---|---|---|
| D-01 | PNG 16×16 de alto contraste. | ASCII, Pixel Art y persistencia. |
| D-02 | JPEG 32×24. | Formato admitido y Pixel Art. |
| D-03 | GIF o WebP válido pequeño. | Decodificación de formato. |
| D-04 | Archivo `.txt` renombrado a `.png`. | Rechazo por contenido inválido. |
| D-05 | Imagen válida de más de 10 MiB. | Límite de carga. |
| D-06 | ObjectID MongoDB válido e inexistente. | Errores de actualización/eliminación. |
| D-07 | ObjectID malformado. | Validación de identificador. |

### 4.3 Responsabilidades propuestas

| Rol | Responsabilidad |
|---|---|
| Responsable de pruebas | Mantener este documento, coordinar ejecución y emitir TSR-001. |
| Desarrollador | Corregir incidencias, añadir pruebas automatizadas y evidencias. |
| Responsable de producto | Aclarar requisitos pendientes y aceptar/rechazar CA. |
| Operaciones | Proporcionar entorno, configuración segura y respaldo de la base de pruebas. |

Los nombres de las personas deben asignarse antes de ejecutar pruebas.

## 5. TDS-001 — Especificación de diseño de pruebas

| Diseño | Requisitos cubiertos | Técnica de diseño | Casos asociados |
|---|---|---|---|
| TD-01 Carga y validación | RF-01, RF-02, RF-03, RF-04, RF-13; CA-03, CA-04 | Partición: imagen válida/inválida; límites: ausencia y >10 MiB. | TC-01 a TC-06 |
| TD-02 Conversión ASCII | RF-05, RF-07, RF-09; CA-01 | Valores de frontera y propiedades de forma. | TC-07, TC-08 |
| TD-03 Pixel Art | RF-06, RF-07, RF-09; CA-02 | Valores de frontera y validación de formato PNG/Base64. | TC-09, TC-10 |
| TD-04 Historial | RF-08, RF-13; CA-05 | Estado vacío, estado con datos y error. | TC-11, TC-12 |
| TD-05 Mutaciones | RF-10, RF-11, RF-12; CA-06, CA-07 | Transición de estado y valores inválidos. | TC-13 a TC-17 |
| TD-06 Construcción y acceso | RNF-02, RNF-05, RNF-06; CA-08, CA-09 | Verificación de configuración y construcción. | TC-18 a TC-20 |

## 6. TCS-001 — Especificación de casos de prueba

**Prioridad:** P0 bloquea release; P1 requerida para aceptación; P2 mejora de cobertura.  
**Precondición común:** entorno de prueba disponible, base `imgapp_test` y datos D-01 a D-07 preparados, salvo que el caso diga lo contrario.

| ID | Prioridad | Requisito | Precondición / entrada | Resultado esperado |
|---|---|---|---|---|
| TC-01 | P0 | RF-01, RF-04 | D-01, tipo `ascii`. | HTTP 200; registro con `type=ascii`, nombre, `id`, `createdAt` y resultado no vacío. |
| TC-02 | P0 | RF-01, RF-04 | D-01, tipo `pixelart`. | HTTP 200; `result` se decodifica como PNG válido. |
| TC-03 | P1 | RF-02 | D-02 y D-03, ambos tipos de conversión aplicables. | Cada formato admitido se procesa o se informa con error explícito documentado. |
| TC-04 | P0 | RF-01, RF-13; CA-03 | Sin parte multipart `image`. | HTTP 400 y `{ "error": ... }`. |
| TC-05 | P0 | RF-02, RF-13; CA-03 | D-04. | HTTP 400 y error de imagen inválida. |
| TC-06 | P0 | RF-03; CA-04 | D-05. | HTTP 400; no se persiste registro ni se agota el servicio. |
| TC-07 | P1 | RF-05; CA-01 | D-01 en ASCII. | Cada fila completa tiene 100 caracteres; caracteres pertenecen a `@%#*+=-:. `; se almacena el texto. |
| TC-08 | P2 | RF-05 | Imagen de ancho 0/no válida o convertidor con ancho no positivo en unitario. | No hay pánico; se devuelve resultado/error definido. |
| TC-09 | P1 | RF-06; CA-02 | D-01 en Pixel Art. | Base64 decodificable, firma PNG válida y dimensiones iguales al origen. |
| TC-10 | P2 | RF-06 | Imagen cuyos bordes no son múltiplo de 8. | No hay píxeles omitidos ni error; se genera PNG completo. |
| TC-11 | P1 | RF-08; CA-05 | Base sin registros. | HTTP 200 con arreglo vacío; UI muestra estado vacío. |
| TC-12 | P1 | RF-08, RF-09 | Dos registros: ASCII y Pixel Art. | HTTP 200; UI lista nombre y resultado correcto de cada tipo. |
| TC-13 | P0 | RF-10; CA-06 | Registro existente; PUT con nombre y tipo válidos. | HTTP 200, datos actualizados y UI sincronizada sin recarga. Si cambia `type`, aplicar regla de negocio aprobada (reconvertir o rechazar). |
| TC-14 | P1 | RF-11 | PATCH con `filename` válido. | HTTP 200, solo nombre actualizado, resultado y tipo conservados. |
| TC-15 | P1 | RF-11 | PATCH `{}` o tipo no admitido. | HTTP 400 y ningún cambio persistido. |
| TC-16 | P0 | RF-12; CA-07 | Registro existente; confirmar eliminación. | HTTP 204; registro no aparece en GET ni UI. |
| TC-17 | P1 | RF-12 | D-06 y D-07 en DELETE. | Error HTTP coherente; no se elimina otro registro. |
| TC-18 | P0 | RNF-05; CA-08 | Clon limpio. | `npm run build` y `go test ./...` finalizan con código 0. |
| TC-19 | P0 | RNF-02; CA-09 | Cliente no autenticado contra rutas de datos en configuración de producción. | HTTP 401 o 403. |
| TC-20 | P1 | RNF-04 | Navegación por teclado y lector de pantalla básico. | Controles tienen nombre accesible, foco visible y errores anunciados. |

## 7. TPS-001 — Procedimientos de prueba

### TP-01. Preparación y verificación de build

1. Crear una copia limpia del repositorio y registrar el hash Git.
2. Configurar una base de pruebas aislada y las variables del backend.
3. Ejecutar `npm ci` y `npm run build` en `frontend`.
4. Ejecutar `go test ./...` en `backend`.
5. Registrar comandos, versión de herramientas, hora, salida y resultado en TLG-001.
6. Si algún comando falla, abrir TIR y suspender casos E2E dependientes.

### TP-02. Pruebas de API

1. Levantar MongoDB y backend contra la base de pruebas.
2. Ejecutar TC-01 a TC-11 y TC-13 a TC-17 por HTTP, conservando solicitudes y respuestas sin datos sensibles.
3. Consultar MongoDB solo para validar persistencia del caso y limpiar datos por identificador de ejecución.
4. Comparar resultado real con el esperado y registrar Pasa/Falla/Bloqueado.

### TP-03. Pruebas de interfaz y aceptación

1. Solo después de TP-01 exitosa, levantar frontend, backend y Nginx.
2. Ejecutar TC-01, TC-02, TC-11 a TC-16 y TC-20 desde el navegador.
3. Capturar evidencia de pantalla y registrar navegador/versión.
4. Pedir al responsable de producto que ejecute o valide CA-01 a CA-09.
5. Emitir TSR-001 al cerrar la iteración.

## 8. TLG-001 — Registro de pruebas

### 8.1 Registro de ejecución inicial

| Ejecución | Fecha | Caso/procedimiento | Resultado | Evidencia | Observación |
|---|---|---|---|---|---|
| LOG-001 | 2026-09-21 | TP-01 / build frontend | Falló | `npm run build`: TS2739 en `App.tsx`. | Bloquea pruebas E2E y CA-06/CA-07. |
| LOG-002 | 2026-09-21 | TP-01 / pruebas backend | Falló | `go test ./...`: falta `golang.org/x/image/webp`. | Bloquea pruebas backend desde clon limpio. |

### 8.2 Plantilla para futuras ejecuciones

| Ejecución | Fecha/hora | Ambiente y versión | Caso | Resultado (Pasa/Falla/Bloqueado) | Evidencia | Ejecutante |
|---|---|---|---|---|---|---|
| LOG-XXX | AAAA-MM-DD hh:mm | Hash Git, navegador/contenedor | TC-XX | — | Ruta/enlace a salida o captura | — |

## 9. TIR-001 — Informes de incidencias

| ID | Severidad | Estado | Caso afectado | Descripción | Resultado esperado | Acción requerida |
|---|---|---|---|---|---|---|
| INC-001 | Crítica | Abierta | TC-18, TC-13, TC-16 | El frontend no compila: `ConversionList` exige `onUpdate` y `onDelete`, pero `App.tsx` no los entrega. | Build exitoso; actualizar y eliminar deben sincronizar la lista. | Implementar callbacks de estado en `App.tsx`, añadir pruebas y repetir TP-01/TP-03. |
| INC-002 | Alta | Abierta | TC-18, TC-03 | El backend importa `golang.org/x/image/webp`, pero la dependencia no está declarada en `go.mod`. | Pruebas Go ejecutables desde clon limpio. | Declarar/versionar dependencia, ejecutar `go mod tidy`, añadir prueba de decodificación y repetir TP-01. |
| INC-003 | Alta | Abierta | TC-19 | No existe autenticación/autorización; las rutas de datos son accesibles y CORS permite `*`. | Rutas protegidas deben devolver 401/403 sin permisos. | Diseñar e implementar identidad, autorización, CORS restringido y pruebas de seguridad. |
| INC-004 | Media | Abierta | TC-13, TC-15 | PUT/PATCH no validan correctamente `type`; cambiar tipo no reconvierte el resultado. | Regla del SRS aplicada: reconvertir o rechazar cambio. | Acordar regla, validarla en servicio y actualizar SRS/pruebas. |

### Plantilla de incidente

```text
ID: INC-XXX
Fecha / reportante:
Ambiente y hash Git:
Caso(s) afectado(s):
Severidad / prioridad:
Descripción y pasos de reproducción:
Resultado esperado / resultado real:
Evidencia:
Análisis de causa:
Acción correctiva / responsable / fecha:
Verificación de corrección y cierre:
```

## 10. TTR-001 — Informe de transferencia de ítems de prueba

**Estado:** pendiente. Debe emitirse antes de cada ciclo de ejecución.

| Campo | Contenido requerido |
|---|---|
| Build o release | Etiqueta Git, hash de commit, imagen Docker y fecha. |
| Ítems entregados | Frontend, backend, configuración Nginx/Compose, migraciones/datos de prueba. |
| Dependencias | Versiones Node, Go, MongoDB y módulos. |
| Configuración | Variables no secretas, URL, puertos y base de datos de pruebas. |
| Integridad | Hashes/etiquetas de imágenes y resultado de build. |
| Aprobación | Quién entrega, quién recibe y restricciones conocidas. |

No se debe transferir un ítem a aceptación mientras TC-18 esté fallando.

## 11. TSR-001 — Informe resumen de pruebas

### 11.1 Resumen preliminar

| Métrica | Resultado actual |
|---|---|
| Casos planificados | 20 |
| Casos ejecutados completamente | 0 |
| Casos bloqueados | Al menos los dependientes de build/frontend y pruebas Go locales. |
| Incidencias abiertas | 4 (1 crítica, 2 altas, 1 media). |
| Criterios de salida cumplidos | No. |
| Recomendación de release | **No liberar.** |

### 11.2 Evaluación frente a criterios de aceptación del SRS

| Criterio | Estado | Razón |
|---|---|---|
| CA-01 y CA-02 | Pendiente | Requieren ejecución de conversiones y pruebas de resultado. |
| CA-03 y CA-04 | Pendiente | Hay lógica estática, sin evidencia de ejecución repetible. |
| CA-05 | Bloqueado | UI no compila. |
| CA-06 y CA-07 | Fallido/Bloqueado | INC-001: faltan callbacks en el estado padre. |
| CA-08 | Fallido | TP-01 documenta ambos comandos fallidos. |
| CA-09 | Fallido | INC-003: no hay autenticación/autorización. |

### 11.3 Cierre requerido

Para sustituir este informe preliminar por uno final se debe: cerrar INC-001 e INC-002, repetir TP-01, ejecutar los casos P0/P1, adjuntar resultados y evidencias, resolver o aceptar formalmente los defectos restantes, y obtener aceptación sobre los criterios del SRS.

## 12. Matriz de trazabilidad

| Requisito / criterio | Diseño | Casos |
|---|---|---|
| RF-01, RF-02, RF-03, RF-04, RF-13; CA-03, CA-04 | TD-01 | TC-01 a TC-06 |
| RF-05, RF-07, RF-09; CA-01 | TD-02 | TC-07, TC-08 |
| RF-06, RF-07, RF-09; CA-02 | TD-03 | TC-09, TC-10 |
| RF-08, RF-13; CA-05 | TD-04 | TC-11, TC-12 |
| RF-10, RF-11, RF-12; CA-06, CA-07 | TD-05 | TC-13 a TC-17 |
| RNF-02, RNF-04, RNF-05, RNF-06; CA-08, CA-09 | TD-06 | TC-18 a TC-20 |
