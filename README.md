# 💬 Blog

Simple Go blog

## 📄 Routes

- /auth
  - /sign-up `POST` SignUp
  - /sign-in `POST` SignIn
- /api
  - /users
    - /:user_id `GET` GetUserById
    - / `GET` GetAllUsers
    - / `PUT` UpdateUserById 🧑 `auth`
    - / `DELETE` DeleteUserById 🧑 `auth`
  - /posts
    - / `POST` CreatePost 👑 `auth / admin`
    - / `GET` GetAllPosts
    - /:post_id `GET` GetPostById
    - /:post_id `PUT` UpdatePostById 👑 `auth / admin`
    - /:post_id `DELETE` DeleteUserById 👑 `auth / admin`
    - /:post_id/comments
      - / `POST` CreateCommentByPostId 🧑 `auth`
      - / `GET` GetCommentsByPostId
  - /comments
    - /:comment_id `PUT` UpdateCommentById 🧑 `auth`
    - /:comment_id `DELETE` DeleteCommentById 🧑 `auth`

## 🛠️ Built with

- Go `1.21.4`
- MySQL `8.0.34`

## 📦 Libraries

| Name   | Module                                                                   | Version  |
| :----- | :----------------------------------------------------------------------- | :------- |
| gin    | [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)             | `1.9.1`  |
| jwt    | [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)        | `5.1.0`  |
| viper  | [github.com/spf13/viper](https://github.com/spf13/viper)                 | `1.17.0` |
| logrus | [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)         | `1.9.3`  |
| sqlx   | [github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx)               | `1.3.5`  |
| mysql  | [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) | `1.7.1`  |
| crypto | [golang.org/x/crypto](https://cs.opensource.google/go/x/crypto)          | `0.15.0` |
