# Estrategia de pruebas de software — ISTQB CTFL v4.0.1

**Proyecto:** Conversor de imágenes ASCII / Pixel Art  
**Identificador:** ISTQB-TP-001  
**Versión:** 1.0  
**Fecha:** 21 de septiembre de 2026  
**Base de prueba:** [SRS.md](SRS.md), [PLAN_PRUEBAS_IEEE_829.md](PLAN_PRUEBAS_IEEE_829.md), código fuente y evaluaciones de calidad  
**Estado:** Plan preparado; ejecución inicial bloqueada por incidencias conocidas.

## 1. Referencia y propósito

ISTQB no es una norma de certificación de producto: es un esquema internacional de cualificaciones y buenas prácticas para pruebas de software. Esta estrategia aplica el programa [Certified Tester Foundation Level (CTFL) v4.0](https://www.istqb.org/certifications/certified-tester-foundation-level-ctfl-v4-0/), cuya versión de temario disponible es v4.0.1. El temario cubre fundamentos, ciclo de vida, pruebas estáticas, técnicas de diseño, gestión de pruebas, defectos y automatización, y es aplicable a enfoques ágiles, DevOps y entrega continua.

El propósito es convertir los requisitos del SRS en pruebas priorizadas por riesgo, definir cómo se analizarán defectos y establecer una ruta de automatización. No acredita una certificación ISTQB ni sustituye la ejecución y aprobación de las pruebas.

## 2. Objetivos de prueba

- Detectar defectos que impidan cumplir RF-01 a RF-13 y CA-01 a CA-09 del SRS.
- Verificar los resultados de conversión ASCII y Pixel Art, la persistencia y los flujos de historial.
- Reducir el riesgo de release validando construcción reproducible, límites de carga, errores de API, accesibilidad básica y acceso a datos.
- Aportar evidencia objetiva para corregir defectos y decidir la liberación.

## 3. Aplicación de los siete principios ISTQB

| Principio | Aplicación al proyecto |
|---|---|
| Las pruebas muestran presencia de defectos, no su ausencia. | Un build exitoso o casos aprobados no demostrarán que no hay fallos; se reportará riesgo residual. |
| Las pruebas exhaustivas son imposibles. | Se priorizan carga, conversión, persistencia, mutaciones, construcción y seguridad por riesgo. |
| Pruebas tempranas ahorran tiempo y dinero. | Revisar SRS y pull requests, y ejecutar unitarias antes de integración/E2E. |
| Agrupamiento de defectos. | Concentrar revisión en handlers, servicio de actualización, estado React, configuración y dependencias: son zonas con defectos conocidos. |
| Paradoja del pesticida. | Variar imágenes, tamaños, formatos, IDs y errores; revisar casos cuando dejen de encontrar defectos. |
| Las pruebas dependen del contexto. | Usar una base aislada, imágenes pequeñas y grandes, y distinguir pruebas locales, integración y aceptación. |
| Ausencia de errores es una falacia. | Además de arreglar builds, verificar que el resultado conserva coherencia: no permitir cambiar `type` sin reconversión o una regla explícita. |

## 4. Enfoque de equipo completo y responsabilidades

| Rol | Responsabilidades |
|---|---|
| Responsable de producto | Aclara SRS, prioriza riesgo, acepta criterios CA y decide sobre defectos de negocio. |
| Desarrollo | Escribe/mantiene unitarias, corrige defectos, revisa cambios y aporta evidencia técnica. |
| QA o responsable de pruebas | Diseña pruebas, mantiene trazabilidad, ejecuta integración/E2E, registra defectos e informa estado. |
| Operaciones | Mantiene entorno de pruebas, configuración, datos aislados y evidencia de despliegue. |

En un equipo pequeño una persona puede asumir varios roles, pero la revisión/aceptación debe ser independiente del cambio cuando sea posible.

## 5. Proceso de prueba ISTQB aplicado

| Actividad | Aplicación concreta | Producto de trabajo |
|---|---|---|
| Planificación | Definir alcance, riesgos, prioridades, ambiente, cronograma y salida. | Este plan y registro de riesgos. |
| Monitoreo y control | Comparar casos ejecutados, fallos y bloqueos contra el plan; ajustar prioridades. | Tablero de ejecución semanal. |
| Análisis | Convertir SRS, cambios y riesgos en condiciones de prueba. | Matriz requisito–condición–caso. |
| Diseño | Seleccionar técnicas de caja negra/blanca y definir datos/resultado esperado. | Casos TC-01 a TC-20 del plan IEEE. |
| Implementación | Crear datos, scripts, fixtures, suites unitarias/API/E2E y ambiente aislado. | Código de prueba, datos D-01 a D-07 y configuración. |
| Ejecución | Ejecutar, comparar esperado/real, registrar evidencia y abrir defectos. | Registro TLG-001 e incidencias INC-XXX. |
| Finalización | Evaluar criterios de salida, riesgo residual, lecciones y artefactos reutilizables. | Informe resumen y decisión de release. |

## 6. Riesgos del producto y prioridad de prueba

La prioridad combina impacto al usuario, probabilidad observada y detectabilidad. P0 bloquea release; P1 debe completarse para aceptación; P2 se programa tras estabilizar P0/P1.

| ID | Riesgo | Prob. | Impacto | Prioridad | Cobertura / tratamiento |
|---|---|---:|---:|---|---|
| R-01 | El frontend no construye por props faltantes. | Alta | Crítico | P0 | TC-18; INC-001; corregir antes de E2E. |
| R-02 | Backend no prueba/compila desde clon limpio por WebP no declarado. | Alta | Alto | P0 | TC-18, TC-03; INC-002. |
| R-03 | Usuarios no autorizados leen o modifican datos. | Alta | Crítico | P0 | TC-19; INC-003; implementar identidad y autorización. |
| R-04 | Cambio de tipo deja resultado incoherente. | Alta | Alto | P0 | TC-13, TC-15; INC-004; definir regla de negocio. |
| R-05 | Carga maliciosa o grande consume recursos. | Media | Alto | P0 | TC-04 a TC-06; añadir límite de píxeles y `MaxBytesReader`. |
| R-06 | Conversión produce salida inválida/corrupta. | Media | Alto | P1 | TC-07 a TC-10, pruebas unitarias y propiedades. |
| R-07 | Historial/persistencia no refleja datos esperados. | Media | Alto | P1 | TC-11, TC-12, TC-14, TC-16, TC-17. |
| R-08 | Usuario no recibe errores ni puede operar con asistencia. | Media | Medio | P1 | TC-20 y prueba de error inicial. |
| R-09 | Configuración no reproducible por etiquetas flotantes. | Media | Medio | P1 | TC-18 y revisión de Compose/Dockerfiles. |

## 7. Niveles y tipos de prueba

| Nivel | Objeto y alcance | Tipos principales | Criterio de entrada |
|---|---|---|---|
| Unitario | Convertidores, validación del servicio y helpers. | Funcional, límites, regresión. | Dependencias Go declaradas y testables. |
| Integración | Handler ↔ servicio ↔ repositorio ↔ MongoDB. | API, persistencia, error y contrato. | MongoDB aislado y backend construido. |
| Sistema | Frontend ↔ Nginx ↔ backend ↔ MongoDB. | Funcional E2E, usabilidad básica, compatibilidad básica. | `npm run build` exitoso y ambiente Compose disponible. |
| Aceptación | SRS desde perspectiva del responsable de producto. | Aceptación funcional y exploratoria. | Casos P0/P1 aprobados y ambiente estable. |

Las pruebas no funcionales prioritarias son seguridad básica, desempeño frente a cargas límite, accesibilidad básica, portabilidad de build y fiabilidad de persistencia. Las pruebas de penetración y carga masiva requieren un plan especializado posterior.

## 8. Técnicas de diseño aplicadas

| Técnica ISTQB | Aplicación | Casos / pruebas a implementar |
|---|---|---|
| Partición de equivalencia | Imagen válida, imagen inválida, sin archivo; tipos `ascii`, `pixelart`, no admitido; ID válido/inválido/inexistente. | TC-01 a TC-05, TC-15, TC-17. |
| Análisis de valores límite | 0/1/10 MiB/10 MiB+1, dimensiones 1×1, bordes no múltiplos de 8, ancho ASCII 100. | TC-06 a TC-10. |
| Tabla de decisión | Actualizar: campos presentes/ausentes y tipo válido/no válido; eliminar: existente/no existente/ID inválido. | TC-13 a TC-17. |
| Transición de estados | Conversión: seleccionada → procesando → éxito/error; registro: visible → edición → guardado/cancelado/eliminado. | TC-01, TC-12 a TC-16. |
| Prueba de sentencias/rutas | Ramas de validación de upload, selección de conversor y actualización. | Unitarias de handler, `NewConverter` y servicios. |
| Prueba exploratoria | Navegación real con combinaciones de formato/tipo/nombre, interrupciones y respuestas lentas. | Sesiones ET-01 y ET-02. |
| Prueba basada en checklist | Accesibilidad, API y configuración de release. | TC-18 a TC-20. |

### 8.1 Tablas de decisión esenciales

#### Actualización parcial (`PATCH`)

| `filename` | `type` | Resultado esperado |
|---|---|---|
| Válido | Ausente | 200; cambiar solo nombre. |
| Ausente | Válido | 200 solo si se define reconversión o cambio de tipo permitido; si no, 400. |
| Válido | Válido | 200 solo si ambos son coherentes según regla aprobada. |
| Ausente | Ausente | 400; no cambiar registro. |
| Vacío/no válido | Cualquiera | 400; no cambiar registro. |

#### Eliminación

| ID | Confirmación UI | Resultado esperado |
|---|---|---|
| Válido existente | Sí | 204, registro eliminado de API y UI. |
| Válido existente | No | No se llama DELETE; UI conserva registro. |
| Válido inexistente | Sí | Error coherente; UI no cambia otro registro. |
| Malformado | Sí | Error 400; UI no cambia. |

## 9. Condiciones y casos de prueba

Los casos detallados, datos y resultados esperados están en [PLAN_PRUEBAS_IEEE_829.md](PLAN_PRUEBAS_IEEE_829.md#6-tcs-001--especificación-de-casos-de-prueba). Esta estrategia los prioriza así:

| Prioridad | Casos obligatorios |
|---|---|
| P0 | TC-01, TC-02, TC-04 a TC-06, TC-13, TC-16, TC-18, TC-19. |
| P1 | TC-03, TC-07, TC-09, TC-11, TC-12, TC-14, TC-15, TC-17, TC-20. |
| P2 | TC-08, TC-10 y sesiones exploratorias. |

### Sesiones exploratorias

| ID | Carta de prueba | Duración sugerida | Evidencia |
|---|---|---:|---|
| ET-01 | Explorar carga y conversión con formatos, tamaños, dimensiones y nombres atípicos para descubrir errores de validación, resultado o memoria. | 60 min | Notas, datos usados, incidencias, capturas. |
| ET-02 | Explorar historial, edición y eliminación con interrupciones, respuestas fallidas, recarga de página y teclado. | 45 min | Notas, rutas recorridas, incidencias, capturas. |

## 10. Pruebas estáticas y revisiones

| Objeto | Lista de comprobación | Salida esperada |
|---|---|---|
| SRS | Cada requisito es claro, verificable, sin ambigüedad, trazable y con criterio de aceptación. | Hallazgos y cambios aprobados; resolver regla de cambio de `type`, paginación y límite de píxeles. |
| Código de cambio | Compila, tiene prueba asociada, maneja error, no expone secreto, respeta contrato API y actualiza documentación. | Revisión de pull request aprobada. |
| Configuración | Versiones fijadas, variables documentadas, CORS restringido, TLS/secretos fuera del repo, health check. | Checklist de release aprobada. |
| Casos de prueba | Precondición, pasos, datos, esperado y requisito trazados; no dependen de orden no declarado. | Casos revisados y listos. |

Toda revisión debe registrar autor, revisor, fecha, hallazgo, decisión y acción. La revisión no sustituye la ejecución dinámica.

## 11. Gestión de defectos

### 11.1 Flujo

```text
Nuevo → Triado → En progreso → Resuelto → Reprueba → Cerrado
                    └──────────────→ Rechazado / Duplicado / Diferido
```

Un defecto solo se cierra tras repetir el caso que lo detectó y una regresión proporcional. La depuración identifica la causa y corrige el código; la prueba confirma el fallo y su corrección.

### 11.2 Severidad y prioridad

| Severidad | Definición | Ejemplo actual |
|---|---|---|
| Crítica | Impide el sistema, pérdida/exposición grave o no hay alternativa. | Frontend no construye (INC-001); datos sin control de acceso (INC-003). |
| Alta | Función esencial no cumple o no puede verificarse. | Dependencia WebP no declarada (INC-002); tipo incoherente (INC-004). |
| Media | Función degradada con alternativa o impacto moderado. | Error inicial solo en consola, sin mensaje al usuario. |
| Baja | Mejora visual, texto o documentación sin impacto funcional directo. | Inconsistencias tipográficas. |

Cada informe debe incluir ID, ambiente, hash Git, pasos, datos, esperado, real, evidencia, severidad, responsable, causa, corrección y resultado de reprueba.

## 12. Automatización, configuración y CI

| Capa | Automatización recomendada | Comando/artefacto objetivo |
|---|---|---|
| Frontend | Typecheck, build, pruebas de componentes/API simulada. | `npm run build`, suite de tests a añadir. |
| Backend | Unitarias de convertidores/servicios y pruebas HTTP. | `go test ./...`. |
| Integración | API contra MongoDB efímero. | Job de CI con contenedor de servicio. |
| E2E | Flujos cargar, listar, editar y eliminar. | Suite de navegador a añadir. |
| Calidad | Lint, dependencia, imagen y configuración. | `npm run lint`, análisis de módulos e imagen. |

Todo resultado debe estar asociado a hash Git, versión de dependencias y ambiente. La rama principal no deberá aceptar cambios si fallan build, pruebas P0 automatizadas o análisis de seguridad definido.

## 13. Métricas y reporte

| Métrica | Fórmula | Objetivo inicial |
|---|---|---:|
| Progreso de pruebas | Casos ejecutados / casos planificados | 100 % P0/P1 antes de aceptación. |
| Éxito | Casos aprobados / casos ejecutados | 100 % P0; objetivo acordado para P1. |
| Cobertura de requisitos | Requisitos con al menos un caso / requisitos aplicables | 100 %. |
| Defectos abiertos | Por severidad y estado | 0 críticos/altos antes de release. |
| Tasa de corrección | Defectos cerrados / defectos resueltos | 100 % con reprueba. |
| Automatización | Casos P0 automatizados / casos P0 | Incremental; 100 % antes de operación estable. |

El informe de cada ciclo debe indicar periodo, versión, casos por estado, defectos por severidad, riesgos residuales, cambios de alcance, decisión de salida y acciones siguientes.

## 14. Estado actual y decisión inicial

| Evidencia | Estado | Implicación |
|---|---|---|
| `npm run build` | Fallido: TS2739 en `App.tsx`. | No ejecutar pruebas de sistema/aceptación hasta corregir. |
| `go test ./...` | Fallido: `golang.org/x/image/webp` no está declarado. | No validar backend desde clon limpio. |
| SRS | Disponible, sin aprobación formal. | Puede ser base de prueba, pero requiere revisión estática/aprobación. |
| Casos/pruebas automatizadas | No existen en el proyecto. | Se debe implementar la pirámide de pruebas propuesta. |
| Autenticación/autorización | Ausentes. | CA-09 falla y el riesgo R-03 bloquea release público. |

**Decisión:** no liberar. Primero cerrar INC-001, INC-002 e INC-003; definir la regla del cambio de tipo; después ejecutar la suite P0, continuar con P1 y emitir un informe de finalización con riesgo residual.

## 15. Criterios de finalización

- Todos los casos P0 y P1 se ejecutaron con evidencia y trazabilidad al SRS.
- No existen defectos críticos o altos abiertos; las excepciones están aceptadas explícitamente por responsable de producto y seguridad.
- Los builds y pruebas automatizadas se ejecutan correctamente desde un clon limpio.
- Los requisitos CA-01 a CA-09 aplicables cuentan con resultado y aceptación.
- El informe final registra métricas, defectos diferidos, riesgos residuales, configuración entregada y lecciones aprendidas.
