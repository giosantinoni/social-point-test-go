# Social Point Test Project

## Para Ejecutar el Proyecto

### Build del Contenedor

Para construir la imagen del contenedor, ejecute el siguiente comando en la terminal:

```sh
docker-compose build
```
Para levantar el contenedor
```sh
docker-compose up -d
```

### Endpoints

* Base URL: `localhost:8080`
* [POST] `/user/{userId}/score`
```
{
  "score": 500,
  "operator": "+"
}
```

* [GET] `/ranking?type=Top100`

Nota: El parámetro type debe seguir uno de los siguientes patrones:

`^[Tt]op[1-9]\d*$` (Ejemplo: Top10, top15, etc.)

`^At[1-9]\d*/[1-9]$` (Ejemplo: At10/2, At20/5, etc.)