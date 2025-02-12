# OZON тестовое задание
## Система добавления постов и комментариев с использованием GraphQL
__Для запуска приложения с использованием базы данных Postgres измените docker-compose.yml:__  
`USE_POSTGRES: true`  
___Также необходимо создать таблицы в базе данных с помощью файла init.sql___
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
После подписки комментарии данного поста будут приходить в режиме реального времени
### Пример запроса GraphQL:
__Данный запрос вернет список всех постов__
```
query {  
  lastest(page: 1) {  
    id  
    title  
    text  
    author {  
      id  
      username  
    }  
    comments {  
      author {  
        id  
        username  
      }  
      text  
      comments {  
        text  
        comments {  
          text  
        }  
      }  
    }  
  }  
}  
```