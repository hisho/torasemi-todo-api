# Todo GraphQL Server

## 概要
とらゼミのハンズオンで使用するGraphQLサーバです

## 技術仕様
|Name|Description|
|---|---|
|golang|基本言語|
|sqlboiler|ORマッパー|
|golang-migrate|マイグレーション|
|go-chi|ルーター|
|gqlgen|GraphQLサーバ|

## ハンズオン

### 1. クエリ

#### まずは全部のユーザを取ってきます
```graphql
query AllUsers {
    allUsers {
        id,
        name,
    }
}
```

#### IDを指定して、ユーザを1個取ってきます
```graphql
query User {
    user(id: 1) {
        id,
        name,
    }
}
```

#### フラグメント
```graphql
query AllUsers {
    allUsers {
        ...UserFragment,
    }
}

fragment UserFragment on User {
    id,
    name,
    created_at,
    updated_at,
}
```

#### フィルターで絞り込み
```graphql
query Users {
    users(filter: {
        name: "空",
    }) {
        ...UserFragment,
    }
}

fragment UserFragment on User {
    id,
    name,
    created_at,
    updated_at,
}
```

### 2. ミューテーション

#### ユーザ作成
```graphql
mutation CreateUser {
    createUser(input: {
        name: "ジョルノ・ジョバーナ",
    }) {
        ...UserFragment,
    }
}

fragment UserFragment on User {
    id,
    name,
    created_at,
    updated_at,
}
```

#### ユーザ更新
```graphql
mutation UpdateUser {
    updateUser(input: {
        id: 4,
        name: "DIO",
    }) {
        ...UserFragment,
    }
}

fragment UserFragment on User {
    id,
    name,
    created_at,
    updated_at,
}
```

#### ユーザ削除
```graphql
mutation DeleteUser {
    deleteUser(input: {
        id: 6,
    }) {
        ...UserFragment,
    }
}

fragment UserFragment on User {
    id,
    name,
    created_at,
    updated_at,
}
```

#### Todo作成
```graphql
mutation CreateTodo {
    createTodo(input: {
        todo: "todo 3",
        userId: 3,
    }) {
        ...TodoFragment,
    }
}

fragment TodoFragment on Todo {
    id,
    todo,
    finished,
    userId,
    createdAt,
    updatedAt,
}
```

#### Todo更新
```graphql
mutation UpdateTodo {
    updateTodo(input: {
        id: 4,
        todo: "updated todo 3",
        finished: true,
        userId: 3,
    }) {
        ...TodoFragment,
    }
}

fragment TodoFragment on Todo {
    id,
    todo,
    finished,
    userId,
    createdAt,
    updatedAt,
}
```

#### Todo削除
```graphql
mutation DeleteTodo {
    deleteTodo(input: {
        id: 4,
    }) {
        ...TodoFragment,
    }
}

fragment TodoFragment on Todo {
    id,
    todo,
    finished,
    userId,
    createdAt,
    updatedAt,
}
```
