# OZON тестовое задание
## Система добавления постов и комментариев с использованием GraphQL
__Для запуска приложения с использованием базы данных Postgres измените docker-compose.yml:__  
`USE_POSTGRES: true`  
### Endpoints:
`/user/add - POST` - Запрос на добавление пользователя  
`/post/add - POST` - Запрос на создание поста  
`/query - POST` - Запросы с использованием GraphQL  
`/` - Веб интерфейс для тестирования GraphQL запросов  
### GrapgQL queries:
#### __query:__  
`lastest(page: int)` - список всех постов  
`getPostById(id: int)` - посмотреть пост по id    

#### __mutation__
`addComment` - проккоментировать пост

#### __subscription__
`commentAdded(postID: int)` - подписаться на пост