# API REST con Go y DDD (Domain-Driven Design)

Este proyecto implementa una API REST siguiendo los principios de Domain-Driven Design (DDD) en Go, demostrando una arquitectura limpia y mantenible para gestionar usuarios.

## Introducción a DDD

Domain-Driven Design es un enfoque de desarrollo de software que:
- Prioriza el dominio del negocio y su lógica
- Basa el diseño en un modelo del dominio
- Establece límites claros entre diferentes partes del software

## Estructura del Proyecto

```bash
├── cmd
│   └── api
│       └── main.go           # Punto de entrada de la aplicación
├── internal
│   ├── domain               # Capa de dominio
│   │   ├── user
│   │   │   ├── entity.go    # Entidades y reglas de negocio
│   │   │   ├── repository.go # Interfaces del repositorio
│   │   │   └── service.go    # Lógica de negocio
│   ├── infrastructure       # Capa de infraestructura
│   │   ├── http
│   │   │   ├── handler      # Manejadores HTTP
│   │   │   │   └── user_handler.go
│   │   │   └── router       # Configuración de rutas
│   │   │       └── router.go
│   │   └── persistence      # Implementaciones de persistencia
│   │       └── memory
│   │           └── user_repository.go
├── pkg                      # Código reutilizable
├── config/                  # Configuración del proyecto
├── utils/                   # Funciones auxiliares
```

### Explicación de las Capas

1. **Domain Layer** (`internal/domain`)
   - Contiene las entidades centrales del negocio
   - Define interfaces y contratos
   - Implementa la lógica de negocio core
   - Es independiente de frameworks y tecnologías externas

2. **Infrastructure Layer** (`internal/infrastructure`)
   - Implementa las interfaces definidas en el dominio
   - Maneja la comunicación HTTP
   - Gestiona la persistencia de datos
   - Integra con servicios externos

3. **Entry Point** (`cmd/api`)
   - Inicializa y configura la aplicación
   - Maneja la inyección de dependencias
   - Arranca el servidor HTTP

## Endpoints de la API

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/users` | Obtiene todos los usuarios |
| GET | `/users/{id}` | Obtiene un usuario por ID |
| POST | `/users` | Crea un nuevo usuario |
| PUT | `/users/{id}` | Actualiza un usuario existente |
| DELETE | `/users/{id}` | Elimina un usuario |

## Beneficios de esta Arquitectura

1. **Separación de Responsabilidades**
   - Cada capa tiene un propósito específico
   - Facilita el mantenimiento y las pruebas
   - Reduce el acoplamiento entre componentes

2. **Escalabilidad**
   - Fácil de agregar nuevas características
   - Simple de modificar implementaciones existentes
   - Preparado para crecer con el negocio

3. **Testabilidad**
   - Arquitectura orientada a pruebas
   - Fácil mock de dependencias
   - Pruebas unitarias más limpias

4. **Mantenibilidad**
   - Código organizado y predecible
   - Fácil de entender y modificar
   - Documentación implícita en la estructura


### Modulos en Go

```bash
go mod init api-rest-postgresql
go mod tidy
```

- Que es go mod?

    - Es un sistema de gestión de dependencias para Go
    - Permite a los desarrolladores especificar y gestionar las dependencias de sus proyectos

- Que es go mod init?

    - Inicializa un nuevo módulo Go
    - Crea un archivo go.mod en el directorio actual
    - Agrega el módulo al sistema de módulos de Go

- Que hace go mod tidy?

    - Agrega los modulos necesarios para el proyecto
    - Elimina los modulos que no se usan
    - Actualiza el archivo go.mod con las dependencias necesarias


### Instalar PostgreSQL

```bash
brew install postgresql
```

### Instalar pq en go

```bash
go get github.com/lib/pq
```

- Que es pq?

    - Es una librería para interactuar con PostgreSQL desde Go
    - Proporciona funciones para ejecutar consultas SQL y manejar errores
    - Facilita la manipulación de datos en la base de datos


- Linter

Un linter es una herramienta que analiza el código fuente para identificar errores, problemas de estilo, y posibles bugs. Su objetivo principal es mejorar la calidad del código y asegurar que sigue ciertas convenciones y estándares.

Instalación de golangci-lint

Usando Homebrew (macOS):

   brew install golangci-lint

Ejecutar golangci-lint

   golangci-lint run

### Despliegue

El archivo `deploy.yml` en el contexto de GitHub Actions es un archivo de configuración que define un flujo de trabajo automatizado para tu proyecto. Este flujo de trabajo puede incluir tareas como la ejecución de pruebas, el linting del código, la construcción de artefactos, y el despliegue de tu aplicación. Aquí te explico cómo funciona y qué son los componentes clave como los "jobs".

### ¿Qué es un `deploy.yml`?

- **Flujo de Trabajo Automatizado:** Define una serie de pasos que se ejecutan automáticamente en respuesta a eventos específicos en tu repositorio, como un push o una pull request.
- **Integración Continua (CI):** Permite verificar automáticamente que el código nuevo no rompa la aplicación existente.
- **Despliegue Continuo (CD):** Automatiza el proceso de despliegue de la aplicación a un entorno de producción o de prueba.

### Componentes Clave

1. **Eventos (`on`):**
   - Especifica los eventos que disparan el flujo de trabajo. Por ejemplo, `push` o `pull_request` en la rama `main`.

2. **Jobs:**
   - Un "job" es una unidad de trabajo que se ejecuta en un entorno de ejecución. Cada job puede contener múltiples pasos.
   - Los jobs se ejecutan en paralelo de forma predeterminada, pero puedes configurarlos para que se ejecuten secuencialmente si uno depende de otro.

3. **Steps:**
   - Cada job contiene una serie de pasos (`steps`) que se ejecutan secuencialmente.
   - Los pasos pueden incluir acciones predefinidas de GitHub, comandos de shell, o acciones personalizadas.

4. **Runners:**
   - Los jobs se ejecutan en "runners", que son entornos de ejecución proporcionados por GitHub o autohospedados.

### Ejemplo de `deploy.yml`

```yaml
name: Lint and Deploy

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  lint_and_deploy:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: '1.23'

    - name: Install golangci-lint
      run: |
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

    - name: Run golangci-lint
      run: |
        cd api-rest-postgresql
        golangci-lint run

    - name: Deploy to Render
      if: success()
      env:
        RENDER_DEPLOY_HOOK_URL: ${{ secrets.RENDER_DEPLOY_HOOK_URL }}
      run: |
        curl "$RENDER_DEPLOY_HOOK_URL"
```