# Evaluación del sistema de gestión de la calidad — ISO 9001:2026

**Proyecto evaluado:** Conversor de imágenes ASCII / Pixel Art  
**Fecha:** 21 de septiembre de 2026  
**Evidencia revisada:** [SRS.md](SRS.md), [EVALUACION_ISO_25010.md](EVALUACION_ISO_25010.md), [EVALUACION_ISO_29110.md](EVALUACION_ISO_29110.md), código fuente, configuración Docker/Nginx e historial Git  
**Alcance de la evaluación:** evidencia disponible en el repositorio. No se entrevistó a la dirección, clientes, proveedores o equipo; tampoco se revisaron registros externos.

## 1. Referencia y propósito

ISO 9001 define requisitos para un sistema de gestión de la calidad (SGC) que permita a una organización entregar productos y servicios que satisfagan requisitos del cliente y los requisitos aplicables, y mejorar la satisfacción del cliente. La edición vigente es [ISO 9001:2026](https://www.iso.org/standard/9001); ISO indica que ISO 9001:2015 fue retirada y sustituida por esa edición el 16 de septiembre de 2026.

Este documento aplica los temas de un SGC al proyecto de software: contexto, liderazgo, planificación, apoyo, operación, evaluación del desempeño y mejora. Su fin es establecer el punto de partida para implantar un SGC proporcional al proyecto. **No es una auditoría de certificación ni una declaración de conformidad con ISO 9001.** La certificación exige evaluar a la organización y su evidencia operativa mediante una entidad competente.

## 2. Método y escala

La evaluación busca evidencia objetiva en el repositorio y la contrasta con cada bloque del SGC. La norma se aplica a la organización, no solo a un repositorio; por eso, la ausencia de documentos organizacionales se registra como falta de evidencia y no como una afirmación sobre prácticas que pudieran existir fuera de este alcance.

| Estado | Criterio |
|---|---|
| Cumplido | Existe evidencia actual, controlada y suficiente dentro del alcance. |
| Parcial | Existe una práctica o documento, pero faltan control, responsables, medición, aprobación o evidencia de uso. |
| Sin evidencia | No se encontró evidencia en el repositorio. |
| No aplicable | No corresponde al alcance; requiere justificación aprobada. |

## 3. Resultado ejecutivo

El proyecto posee activos útiles para un SGC inicial: un SRS con requisitos y criterios de aceptación, código versionado, configuración de despliegue y evaluaciones de calidad. Sin embargo, estos activos no forman todavía un sistema de gestión: no se evidencia una política u objetivos de calidad, partes interesadas, alcance formal del SGC, responsables, gestión de riesgos, control de cambios documentales, medición de satisfacción, auditorías internas, revisión por la dirección, acciones correctivas o mejora continua.

El dictamen es **preparación baja para ISO 9001:2026**. La operación de desarrollo tiene documentación parcial, pero el SGC organizacional no está establecido con la evidencia evaluada.

| Bloque del SGC | Estado | Conclusión |
|---|---|---|
| 4. Contexto de la organización | Sin evidencia | No se han definido partes interesadas, alcance formal ni procesos del SGC. |
| 5. Liderazgo | Sin evidencia | No hay política, roles, responsabilidades ni compromiso de dirección documentados. |
| 6. Planificación | Sin evidencia | No hay riesgos, oportunidades, objetivos medibles ni planificación de cambios. |
| 7. Apoyo | Parcial | Hay documentación técnica y Git, pero faltan competencias, comunicaciones y control documental. |
| 8. Operación | Parcial | El SRS y la implementación proporcionan una base; faltan control de requisitos, validación, proveedores y liberación. |
| 9. Evaluación del desempeño | Sin evidencia | No hay indicadores, satisfacción de cliente, auditoría interna ni revisión de dirección. |
| 10. Mejora | Sin evidencia | No hay no conformidades, acciones correctivas ni ciclo de mejora controlado. |

## 4. Evaluación detallada

### 4.1 Contexto de la organización

| Práctica del SGC | Estado | Evidencia / brecha | Acción requerida |
|---|---|---|---|
| Analizar cuestiones internas y externas relevantes. | Sin evidencia | El SRS describe alcance técnico, no contexto organizacional, mercado, obligaciones, riesgos externos o entorno de operación. | Crear análisis breve de contexto y revisar periódicamente. |
| Determinar necesidades de partes interesadas. | Parcial | El SRS menciona usuario, frontend, API y base de datos, pero no identifica cliente, propietario, equipo, proveedor, regulador ni sus necesidades. | Registrar partes interesadas, requisitos y responsable de seguimiento. |
| Definir alcance del SGC. | Sin evidencia | No se establece qué productos, servicios, ubicaciones o procesos cubre el SGC. | Publicar alcance: desarrollo, despliegue y soporte del conversor, con exclusiones justificadas. |
| Determinar procesos, entradas, salidas, responsables e interacción. | Parcial | Arquitectura técnica y SRS describen flujo de software, no procesos del SGC. | Mapear gestión de requisitos, desarrollo, pruebas, entrega, soporte y mejora. |

### 4.2 Liderazgo

| Práctica del SGC | Estado | Evidencia / brecha | Acción requerida |
|---|---|---|---|
| Demostrar compromiso de la dirección con el SGC y enfoque al cliente. | Sin evidencia | No hay evidencia de responsable de calidad, decisiones, recursos o revisión de requisitos del cliente. | Nombrar dueño del SGC y registrar compromiso/revisiones. |
| Establecer una política de calidad comunicada y disponible. | Sin evidencia | No existe política de calidad en el repositorio. | Aprobar una política corta centrada en requisitos, calidad y mejora. |
| Asignar roles, responsabilidades y autoridades. | Sin evidencia | Los documentos describen componentes, no quién aprueba SRS, cambios, pruebas o releases. | Crear matriz RACI ligera. |

### 4.3 Planificación

| Práctica del SGC | Estado | Evidencia / brecha | Acción requerida |
|---|---|---|---|
| Tratar riesgos y oportunidades del SGC. | Parcial | Las evaluaciones ISO 25010/29110 identifican riesgos técnicos, pero no hay matriz controlada con prioridad, dueño, tratamiento y revisión. | Mantener registro de riesgos y oportunidades. |
| Definir objetivos de calidad medibles y planes para lograrlos. | Parcial | ISO 25010 propone métricas, pero no son objetivos aprobados ni tienen responsable, plazo o línea base. | Aprobar objetivos, por ejemplo build exitoso, satisfacción, defectos y tiempo de resolución. |
| Planificar cambios al SGC y al producto. | Sin evidencia | Git guarda cambios, pero no hay evaluación de impacto, aprobación ni actualización del SRS. | Establecer solicitud de cambio y plan de transición. |
| Considerar pertinencia del cambio climático cuando aplique. | Sin evidencia | No se evaluó si el clima es una cuestión relevante o un requisito de parte interesada. | Registrar una decisión de aplicabilidad; para este servicio digital puede ser no material, pero debe quedar justificado. |

### 4.4 Apoyo

| Práctica del SGC | Estado | Evidencia / brecha | Acción requerida |
|---|---|---|---|
| Proveer recursos, personas e infraestructura. | Parcial | Docker, Nginx, MongoDB, Go y React describen infraestructura; no hay capacidad asignada, responsable ni plan de continuidad. | Definir recursos, responsables, presupuesto y respaldo operativo. |
| Asegurar competencia y toma de conciencia. | Sin evidencia | No hay perfiles, formación, evaluación de competencias ni inducción al proceso. | Mantener matriz simple de competencias y formación necesaria. |
| Establecer comunicación interna/externa. | Sin evidencia | No hay canales, frecuencia, responsables o comunicación con cliente. | Definir cómo se informan estado, cambios, incidentes y aceptación. |
| Controlar información documentada. | Parcial | Git versiona el código y existen SRS/evaluaciones; no hay propietario, estado de aprobación, versión de release, retención ni control de documentos externos. | Crear lista maestra de documentos con versión, dueño, aprobación y retención. |

### 4.5 Operación

| Práctica del SGC | Estado | Evidencia / brecha | Acción requerida |
|---|---|---|---|
| Planificar y controlar la prestación/desarrollo. | Parcial | SRS, Compose y código describen el producto; no hay plan de trabajo, criterios de liberación o registros de ejecución. | Definir flujo: requisito → cambio → revisión → prueba → release → aceptación. |
| Determinar, revisar y comunicar requisitos de cliente/producto. | Parcial | `SRS.md` contiene requisitos y CA-01 a CA-09, pero no evidencia origen, revisión, viabilidad o aprobación del cliente. | Aprobar SRS y registrar revisiones/aceptaciones. |
| Controlar cambios de requisitos y producción. | Sin evidencia | Cambios en Git no incluyen impacto, autorización, actualización del SRS ni trazabilidad a pruebas. | Implantar registro de cambios con impacto, decisión y versión liberada. |
| Controlar proveedores o servicios externos. | Sin evidencia | Se usan imágenes Docker, MongoDB, ngrok, npm y módulos Go, sin evaluación de dependencias, condiciones o vulnerabilidades. | Inventario de proveedores/dependencias, criterios de selección y revisión periódica. |
| Verificar y validar antes de liberar. | Sin evidencia | El SRS define criterios, pero no hay pruebas; `npm run build` y `go test ./...` fallan en el estado evaluado. | Corregir build y ejecutar pruebas contra cada criterio de aceptación. |
| Identificar, preservar y dar trazabilidad a la salida. | Parcial | Git y registros Mongo dan identificadores técnicos; no hay versión de release, notas, hash de imagen ni retención de evidencia. | Etiquetar releases y conservar resultados de construcción/pruebas. |
| Controlar salidas no conformes. | Sin evidencia | Se han identificado fallos en evaluaciones, sin registro de defectos, contención, decisión o verificación de corrección. | Mantener registro de no conformidades/defectos. |

### 4.6 Evaluación del desempeño

| Práctica del SGC | Estado | Evidencia / brecha | Acción requerida |
|---|---|---|---|
| Medir, analizar y evaluar el desempeño del SGC. | Sin evidencia | No hay tablero, línea base, objetivos ni resultados de indicadores. | Medir builds, cobertura, defectos, tiempo de ciclo, disponibilidad y cumplimiento de CA. |
| Obtener y analizar satisfacción del cliente. | Sin evidencia | No hay mecanismo ni registro de retroalimentación del usuario/cliente. | Definir encuesta, retrospectiva de entrega o canal de feedback y analizar resultados. |
| Realizar auditorías internas. | Sin evidencia | No hay programa, criterios, hallazgos o acciones de auditoría. | Planificar auditoría interna proporcional antes de una evaluación externa. |
| Revisar el SGC por la dirección. | Sin evidencia | No hay actas que revisen objetivos, riesgos, desempeño, recursos o mejora. | Realizar revisión periódica y registrar entradas, decisiones y acciones. |

### 4.7 Mejora

| Práctica del SGC | Estado | Evidencia / brecha | Acción requerida |
|---|---|---|---|
| Gestionar no conformidades y acciones correctivas. | Sin evidencia | Los fallos de build/dependencia no tienen análisis causal, responsable, plazo ni verificación de eficacia. | Registrar no conformidad, causa raíz, acción, responsable, fecha y evidencia de eficacia. |
| Mejorar continuamente la idoneidad, adecuación y eficacia del SGC. | Sin evidencia | No existe ciclo PDCA documentado ni historial de revisiones de proceso. | Planificar ciclos trimestrales de objetivos, medición, acciones y revisión. |

## 5. Trazabilidad con los documentos existentes

| Activo existente | Aporta a ISO 9001 | Limitación actual |
|---|---|---|
| `SRS.md` | Requisitos, alcance del producto y criterios de aceptación para operación. | No tiene aprobación, control de cambios ni vínculo a resultados de prueba. |
| `EVALUACION_ISO_25010.md` | Riesgos de calidad, métricas y acciones técnicas. | No es un registro de riesgos/objetivos del SGC ni evidencia de ejecución. |
| `EVALUACION_ISO_29110.md` | Diagnóstico de proceso y plan de adopción para VSE. | No sustituye política, revisión de dirección, auditoría o registros del SGC. |
| Git | Control de versiones de código y documentos. | No sustituye revisión, aprobación, versión de release o gestión de cambios. |
| Docker/Compose | Evidencia de infraestructura y entrega potencial. | No garantiza release reproducible: existen etiquetas flotantes y verificaciones fallidas. |

## 6. Plan de implantación proporcional

| Prioridad | Acción | Evidencia de salida | Bloques ISO 9001 |
|---|---|---|---|
| P0 | Definir alcance del SGC, partes interesadas, dueño del SGC, política y objetivos medibles. | `SGC.md` aprobado con alcance, política, roles, objetivos e indicadores. | 4, 5, 6 |
| P0 | Aprobar el SRS y establecer control de cambios. | SRS con versión, aprobador, fecha; registro de cambios e impacto. | 7, 8 |
| P0 | Corregir los bloqueos de build y registrar los defectos/acciones correctivas. | Builds verdes, análisis causal y verificación de eficacia. | 8, 10 |
| P1 | Crear matriz requisito–prueba–resultado y ejecutar CA-01 a CA-09. | Evidencia de pruebas, defectos y aceptación/rechazo por criterio. | 8, 9 |
| P1 | Establecer control de proveedores/dependencias, releases y documentación operativa. | Inventario, criterios de revisión, release etiquetado, guía de despliegue/recuperación. | 7, 8 |
| P1 | Medir objetivos y satisfacción; hacer revisión de dirección. | Tablero periódico, feedback analizado y acta de revisión con decisiones. | 9 |
| P2 | Auditar internamente el SGC y cerrar hallazgos. | Programa, informe de auditoría, acciones correctivas y seguimiento. | 9, 10 |

## 7. Indicadores mínimos sugeridos

| Objetivo de calidad | Indicador | Meta inicial | Frecuencia | Responsable |
|---|---|---:|---|---|
| Entregar software verificable | Builds exitosos en rama principal | 100 % | Cada cambio | Responsable técnico |
| Cumplir requisitos | Criterios de aceptación con resultado aprobado | 100 % aplicables | Cada release | Responsable de calidad |
| Reducir defectos críticos | Defectos críticos abiertos antes de release | 0 | Cada release | Responsable técnico |
| Responder a calidad percibida | Satisfacción/aceptación de cliente | Objetivo a acordar | Cada entrega | Responsable del producto |
| Mejorar el proceso | Acciones correctivas cerradas en plazo | ≥ 90 % | Mensual | Dueño del SGC |

Los responsables y metas definitivas deben ser nombrados y aprobados por la organización; los de la tabla son propuestas, no asignaciones vigentes.

## 8. Conclusión

El proyecto cuenta con documentación de requisitos y una base de implementación que pueden alimentar un SGC ligero. Para alinearse con ISO 9001:2026 debe evolucionar de documentación técnica aislada a un sistema repetible y evidenciable: liderazgo, alcance, riesgos, objetivos, control de documentos/cambios, validación de entregas, medición, revisión y acciones correctivas.

La primera revisión del SGC debe realizarse después de completar P0. En ella se debe verificar con evidencia que el SRS fue aprobado, que la construcción es reproducible, que las no conformidades actuales tienen acción correctiva eficaz y que los objetivos de calidad cuentan con responsables y medición.
