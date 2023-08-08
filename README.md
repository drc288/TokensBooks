# Api de authenticacion

Esta api se conecta a una base de datos para logearse y retornar un token JTW

## APIs

### /api/health

Status de la aplicacion

### /api/auth/registry

Crea un usuario y returna su ID

### /api/auth/login

Returna el Token JWT

## Docker
```bash
docker pull drc288/mariadb-contable:0.0.1
docker run --detach --name mariabd-contable -p 3306:3306 drc288/mariadb-contable:0.0.1
```

### Base de datos

Se integra en el proyecto una imagen de docker la cual emula la base de datos con la base de datos creada
