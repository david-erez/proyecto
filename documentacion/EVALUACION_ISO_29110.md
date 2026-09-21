# Evaluación de proceso — ISO/IEC 29110

**Proyecto:** Conversor de imágenes ASCII / Pixel Art  
**Fecha de evaluación:** 21 de septiembre de 2026  
**Línea base:** [SRS.md](SRS.md) v1.0, código fuente, historial Git y configuración de despliegue  
**Alcance:** prácticas de gestión y de implementación visibles en el repositorio. No incluye entrevistas, registros externos, ejecución con usuarios, auditoría de seguridad ni certificación.

## 1. Objetivo y perfil aplicado

ISO/IEC 29110 establece perfiles de ciclo de vida para entidades muy pequeñas (VSE), definidas como empresa, organización, departamento o proyecto de hasta 25 personas. La serie propone perfiles Entry, Basic, Intermediate y Advanced, y puede emplearse con enfoques ágiles o iterativos. [ISO/IEC 29110-1-1:2024](https://www.iso.org/standard/85337.html) y la [serie ISO/IEC 29110](https://committee.iso.org/sites/jtc1sc7/home/projects/flagship-standards/isoiec-29110-series.html) describen ese marco.

Esta evaluación aplica de manera **provisional** el **perfil genérico Entry para ingeniería de software**, adecuado para una VSE de inicio o un proyecto pequeño de menos de seis persona-meses y sin software crítico para la seguridad, según [ISO/IEC 29110-5-1-1:2025](https://www.iso.org/standard/85420.html). El perfil es coherente con el alcance del conversor, pero no se dispone de evidencia del tamaño real del equipo, duración o criticidad; esas condiciones deben confirmarse.

El perfil Entry se organiza, de forma práctica para esta evaluación, en dos procesos:

| Proceso | Finalidad evaluada |
|---|---|
| Gestión del proyecto (PM) | Acordar, planificar, seguir y cerrar el trabajo con el cliente/equipo. |
| Implementación de software (SI) | Analizar requisitos, diseñar, construir, integrar, probar y entregar el producto. |

> Esta evaluación identifica preparación y brechas de proceso. No demuestra ni concede certificación ISO/IEC 29110.

## 2. Método y escala

Se revisaron los productos de trabajo existentes y se contrastaron con las prácticas mínimas esperables del perfil Entry. Cada práctica recibe uno de estos estados:

| Estado | Significado |
|---|---|
| Cumplido | Existe evidencia suficiente y actual en el repositorio. |
| Parcial | Hay evidencia útil, pero incompleta, sin aprobación, sin trazabilidad o sin verificación. |
| No cumplido | No hay evidencia en el alcance evaluado o la evidencia contradice la práctica. |
| No aplicable | La práctica no corresponde al alcance actual; requiere una justificación explícita. |

La conclusión se limita a evidencia observable. Documentos creados en esta sesión (`SRS.md` y evaluaciones) cuentan como línea base documental, pero no prueban que hayan sido aprobados por cliente o equipo.

## 3. Resultado ejecutivo

El repositorio muestra una implementación técnica inicial: historial de cambios en Git, separación de frontend/backend/persistencia, Docker Compose y un SRS que define requisitos y criterios de aceptación. Esto proporciona una base para la implementación de software.

El proceso, sin embargo, es **incipiente y no conforme todavía con el perfil Entry**. No se observa acuerdo de trabajo, responsable/roles, plan de proyecto, estimaciones, gestión de riesgos, seguimiento, registro de cambios, revisiones formales, pruebas automatizadas, registro de incidencias, aceptación del cliente ni entrega versionada. Además, los dos comandos de verificación definidos por el SRS fallan en el estado evaluado:

```text
cd frontend && npm run build  # falla TS2739 en App.tsx
cd backend && go test ./...   # falla dependencia golang.org/x/image/webp
```

| Proceso | Cumplido | Parcial | No cumplido | Dictamen |
|---|---:|---:|---:|---|
| Gestión del proyecto | 0 | 2 | 7 | No conforme; faltan los controles mínimos de planificación y seguimiento. |
| Implementación de software | 2 | 6 | 5 | Parcial; existe código y requisitos, pero no una cadena verificable de entrega. |
| **Total** | **2** | **8** | **12** | **Preparación baja** para adoptar ISO/IEC 29110 Entry. |

Los conteos son indicadores de preparación, no una fórmula oficial de certificación.

## 4. Evaluación de gestión del proyecto (PM)

| ID | Práctica esperada | Estado | Evidencia o brecha |
|---|---|---|---|
| PM-01 | Identificar cliente, responsable y roles del proyecto. | No cumplido | El SRS identifica actores de sistema, no personas, responsabilidades ni autoridad de aprobación. |
| PM-02 | Definir y acordar alcance, entregables y criterios de aceptación. | Parcial | `SRS.md` define alcance, RF/RNF y CA-01 a CA-09; no existe aprobación, versión acordada ni acuerdo de trabajo. |
| PM-03 | Estimar esfuerzo, calendario, presupuesto y recursos. | No cumplido | No hay plan, *backlog*, estimación, hitos ni asignación de recursos. |
| PM-04 | Mantener un plan de proyecto proporcional al tamaño del trabajo. | No cumplido | No existe `PLAN_PROYECTO.md`, tablero o plan equivalente con tareas, responsables y fechas. |
| PM-05 | Identificar y gestionar riesgos. | No cumplido | Hay riesgos técnicos descritos en la evaluación ISO 25010, pero no un registro con propietario, probabilidad, impacto y respuesta. |
| PM-06 | Dar seguimiento al avance frente al plan. | No cumplido | El historial Git muestra cambios, pero no estado de hitos, indicadores ni acciones ante desviaciones. |
| PM-07 | Gestionar solicitudes y cambios de requisitos. | Parcial | Git conserva cambios de código; no hay registro de cambios del SRS, decisión, impacto, prioridad ni aprobación. |
| PM-08 | Comunicar estado, decisiones y problemas a las partes interesadas. | No cumplido | No hay minutas, informe de estado, registro de decisiones o incidencias. |
| PM-09 | Cerrar la entrega con aceptación y resguardo de productos de trabajo. | No cumplido | No existe versión entregable, acta de aceptación ni registro de configuración liberada. |

### Productos de trabajo de PM

| Producto | Estado | Acción mínima |
|---|---|---|
| Acuerdo de trabajo / declaración de trabajo | Ausente | Crear documento breve con cliente, objetivo, alcance, entregables, aceptación y responsables. |
| Plan de proyecto | Ausente | Crear un plan de una página con tareas, fechas, responsable, estimación y estado. |
| Registro de riesgos | Ausente | Mantener tabla con riesgo, impacto, probabilidad, mitigación, responsable y fecha de revisión. |
| Registro de cambios | Ausente | Versionar el SRS y registrar solicitud, análisis de impacto, decisión y aprobación. |
| Registro de seguimiento | Ausente | Revisar semanalmente avance, bloqueos, riesgos y próximas acciones. |
| Acta de aceptación | Ausente | Asociar CA-01 a CA-09 a resultados y firma/aprobación de cliente o responsable. |

## 5. Evaluación de implementación de software (SI)

| ID | Práctica esperada | Estado | Evidencia o brecha |
|---|---|---|---|
| SI-01 | Analizar requisitos y establecer una especificación base. | Cumplido | `SRS.md` v1.0 define alcance, RF-01 a RF-13, RNF y criterios de aceptación. Falta aprobación formal, tratada en PM-02. |
| SI-02 | Mantener trazabilidad entre requisitos, componentes y pruebas. | Parcial | El SRS enlaza grupos de requisitos con componentes; no hay matriz requisito–caso de prueba–resultado. |
| SI-03 | Definir arquitectura y diseño suficientes para construir el producto. | Parcial | Docker Compose, Nginx, capas Go y componentes React evidencian arquitectura; no existe documento de diseño, diagramas de datos ni decisiones técnicas versionadas. |
| SI-04 | Construir el software conforme a requisitos y estándares definidos. | Parcial | Implementados convertidores, API CRUD y UI; el cliente no compila y existen validaciones inconsistentes, por lo que no se demuestra conformidad con el SRS. |
| SI-05 | Revisar cambios y controlar versiones/configuración. | Parcial | Hay repositorio Git y Dockerfiles; no se evidencia revisión por pares, estrategia de ramas, etiquetas de versión, línea base o control de dependencias reproducible. |
| SI-06 | Ejecutar pruebas unitarias e integración contra los requisitos. | No cumplido | No hay archivos de prueba ni pipeline; `go test ./...` falla por dependencia no declarada. |
| SI-07 | Gestionar defectos hasta su resolución y verificación. | No cumplido | No hay registro de incidencias, severidad, responsable, estado o prueba de corrección. |
| SI-08 | Integrar componentes y verificar la construcción. | No cumplido | `npm run build` falla por props `onUpdate` y `onDelete` no proporcionados desde `App.tsx`; no hay CI. |
| SI-09 | Preparar paquete/release instalable y reproducible. | Parcial | Docker Compose integra servicios; imágenes `latest`, dependencia WebP fuera de `go.mod` y ausencia de versión/release impiden reproducibilidad comprobada. |
| SI-10 | Validar el producto contra criterios de aceptación y obtener aceptación. | No cumplido | El SRS define CA-01 a CA-09, pero no hay casos ejecutados, resultados ni aceptación. |
| SI-11 | Entregar documentación de uso, instalación y operación. | No cumplido | El README del frontend conserva la plantilla Vite; no hay guía del producto, variables, despliegue, respaldo ni operación. |
| SI-12 | Conservar productos de trabajo y configuración de la entrega. | Cumplido | Código, configuración de contenedores, SRS y documentos de evaluación se conservan en el repositorio; falta identificar una línea base de release. |

### Estado de productos de trabajo de SI

| Producto | Estado | Evidencia actual | Acción mínima |
|---|---|---|---|
| Especificación de requisitos | Disponible | `SRS.md` v1.0 | Aprobarla, versionarla y registrar cambios. |
| Arquitectura/diseño | Parcial | Código, Compose y Nginx | Crear diagrama, modelo de datos, contrato API y decisiones. |
| Código fuente | Disponible | Frontend React, backend Go | Corregir bloqueos y aplicar revisiones. |
| Casos y resultados de prueba | Ausente | No hay `*_test.go`, `*.test.*` ni CI | Crear pruebas desde CA-01 a CA-09 y guardar resultados. |
| Registro de defectos | Ausente | — | Registrar cada defecto hasta su verificación. |
| Paquete de entrega | Parcial | Dockerfiles y Compose | Fijar versiones, variables y procedimiento de despliegue. |
| Manuales | Ausente | README plantilla | Documentar usuario, instalación, operación y recuperación. |

## 6. Trazabilidad SRS → implementación → verificación

Esta tabla es el punto de partida para la práctica de trazabilidad del perfil Entry.

| Requisito SRS | Implementación identificada | Verificación requerida | Estado actual |
|---|---|---|---|
| RF-01 a RF-04 | `UploadForm.tsx`, `conversion_handler.go`, `converter.go` | Pruebas de carga válida, inválida, límite y tipos. | Sin evidencia de prueba. |
| RF-05 | `converters/ascii.go` | Pruebas unitarias de dimensiones, caracteres y proporción. | Sin evidencia de prueba. |
| RF-06 | `converters/pixelart.go` | Pruebas unitarias de bloques, PNG y Base64. | Sin evidencia de prueba. |
| RF-07 a RF-08 | `conversion_service.go`, `conversion_repository.go` | Pruebas de integración MongoDB y listado vacío/con registros. | Sin evidencia de prueba. |
| RF-09 | `ConversionList.tsx`, `ResultDisplay.tsx` | Prueba E2E de representación ASCII/PNG y accesibilidad básica. | Bloqueada por build. |
| RF-10 a RF-12 | `conversionApi.ts`, `ResultDisplay.tsx`, handler/service/repository | Pruebas API y E2E de editar/eliminar; validar sincronización del estado. | No conforme: callbacks faltantes en `App.tsx`. |
| RF-13 | `writeError`, cliente API | Pruebas de esquema de error y visualización. | Sin evidencia de prueba. |
| RNF-01 a RNF-06 | Infraestructura y toda la solución | Medición, pruebas de seguridad/accesibilidad, CI y guía de despliegue. | No conforme o parcial según `EVALUACION_ISO_25010.md`. |

## 7. Plan de adopción mínimo

La adopción puede hacerse de forma ligera, sin burocracia ajena al tamaño de la VSE. El orden reduce primero los bloqueos que impiden verificar el producto.

| Fase | Actividades | Evidencias de salida | Procesos |
|---|---|---|---|
| 1. Establecer control | Nombrar responsable/cliente; aprobar SRS; crear acuerdo, plan, riesgos y registro de cambios. | SRS aprobado, plan de una página, riesgos y cambios versionados. | PM-01 a PM-07 |
| 2. Recuperar construcción | Implementar callbacks en `App.tsx`; declarar dependencia WebP; fijar el flujo de build/pruebas en entorno limpio. | `npm run build` y `go test ./...` exitosos. | SI-04, SI-08 |
| 3. Verificar requisitos | Derivar casos de prueba de CA-01 a CA-09; añadir pruebas de convertidores, API y UI; registrar defectos. | Matriz SRS–prueba–resultado y registro de defectos cerrado. | SI-02, SI-06, SI-07, SI-10 |
| 4. Preparar entrega | Documentar arquitectura, API, instalación, operación y respaldo; fijar imágenes/dependencias; etiquetar release. | Paquete reproducible, manuales y línea base de release. | SI-03, SI-09, SI-11, SI-12 |
| 5. Cerrar y mejorar | Ejecutar criterios de aceptación con responsable/cliente; registrar aceptación y lecciones/riesgos pendientes. | Acta de aceptación y retrospectiva breve. | PM-08, PM-09, SI-10 |

## 8. Indicadores ligeros de seguimiento

| Indicador | Objetivo inicial | Fuente |
|---|---|---|
| Requisitos con prueba y resultado | 100 % de RF y CA aplicables | Matriz de trazabilidad. |
| Construcciones exitosas | 100 % en la rama principal | CI o registro de build. |
| Defectos críticos abiertos | 0 antes de aceptación | Registro de defectos. |
| Cambios aprobados antes de implementación | 100 % | Registro de cambios. |
| Riesgos revisados | Semanalmente durante desarrollo | Registro de riesgos/estado. |
| Criterios de aceptación aprobados | 100 % de CA aplicables | Acta de aceptación. |

## 9. Conclusión

El proyecto ya posee dos activos importantes para iniciar una adopción ISO/IEC 29110: un repositorio versionado y un SRS con criterios de aceptación. Aun así, no cuenta con la evidencia mínima para declarar implementación del perfil Entry, especialmente en planificación, trazabilidad de pruebas, gestión de defectos y aceptación.

La prioridad es convertir el SRS en un acuerdo controlado y en pruebas ejecutables, corregir los bloqueos de construcción y registrar el ciclo de entrega. Tras completar las fases 1 a 3, debe realizarse una reevaluación de proceso con evidencia de ejecuciones y aprobaciones.
