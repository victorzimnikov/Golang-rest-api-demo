# Демо проект на Golang

Простой REST API на Golang без авторизации с генерацией SWAGGER

### Стек

- Golang
- Fiber
- Postgres
- swaggo
- sqlc

### Роуты

- GET /projects - получить список проектов
- POST /projects - создать проект
- GET /projects/{projectID} - получить проект
- PATCH /projects/{projectID} - обновить проект
- DELETE /projects/{projectID} - удалить проект
- GET /projects/{projectID}/issues - получить список задач проекта
- POST /projects/{projectID}/issues - создать задачу проекта
- GET /issues/{issueID} - получить задачу
- PATCH /issues/{issueID} - обновить задачу
- DELETE /issues/{issueID} - удалить задачу
- GET /issues/{issueID}/comments - получить список комментариев задачи
- POST /issues/{issueID}/comments - создать комментарий задачи
- GET /comments/{issueID} - получить комментарий
- PATCH /comments/{issueID} - обновить комментарий
- DELETE /comments/{issueID} - удалить комментарий
