## 概要

[ドメイン駆動設計入門 ボトムアップでわかる! ドメイン駆動設計の基本](https://www.amazon.co.jp/%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E9%A7%86%E5%8B%95%E8%A8%AD%E8%A8%88%E5%85%A5%E9%96%80-%E3%83%9C%E3%83%88%E3%83%A0%E3%82%A2%E3%83%83%E3%83%97%E3%81%A7%E3%82%8F%E3%81%8B%E3%82%8B-%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E9%A7%86%E5%8B%95%E8%A8%AD%E8%A8%88%E3%81%AE%E5%9F%BA%E6%9C%AC-%E6%88%90%E7%80%AC-%E5%85%81%E5%AE%A3/dp/479815072X)を読んで。

## 基本概念

### ドメインの概念を表現するオブジェクト

![Untitled](Domain%E9%A7%86%E5%8B%95%E3%83%9B%E3%82%99%E3%83%88%E3%83%A0%E3%82%A2%E3%83%83%E3%83%95%E3%82%9A%2080fa13a71ca642b98c4ba5791cae443b/Untitled.png)

[DDD基礎解説：Entity、ValueObjectってなんなんだ - little hands' lab](https://little-hands.hatenablog.com/entry/2018/12/09/entity-value-object)

### ドメインサービス

- Domain Service(サービス): 値オブジェクトやエンティティではうまく表現できない知識を取り扱う
- Repository(リポジトリ): データの保存・復元といった永続化や再構築を担当する(RDBなどの具体的なデータ永続化を抽象化する役割)

## Value Object

システムに必要とされる処理に従って、値を表現する不変性を持つオブジェクト。

- 表現力が増す
- 不正な値を存在させない
- 誤った代入を防ぐ
- 「色」や「量」のように、その属性だけが重要で、アイデンティティを考えることに意味のないオブジェクト
- ロジックの散在を防ぐ

### 基本例

名前をStringとして扱うと、順序が異なり意図しない値を取得するケースが生まれる

```jsx
const fullName = 'Sato Taro';
const fullName2 = 'John Wick';

const getNames = (fullName: string) => {
    const firstName = fullName.split(' ')[0]
    const lastName = fullName.split(' ')[1]
    return { firstName, lastName }
}
```

クラスで下記のようにすることで柔軟に対応可能

```jsx
class FullName {
  firstName: string;
  lastName: string;

  constructor(firstName: string, lastName: string) {
    this.firstName = firstName
    this.lastName = lastName
  }
}
 
let greeter = new FullName("John", "Wick");
```

### 振る舞いを持った例

- 普遍である必要あり普遍であるべき

```jsx
type CURRENCY = 'JP-YEN' | "US-Dollars";

class Money {
  currency: CURRENCY;
  amount: number;

  constructor(currency: CURRENCY, amount: number) {
    this.currency = currency
    this.amount = amount
  }

  add(currency: string, amount: number) {
    console.log('add', this, amount)
    if(currency !== this.currency) {
        throw new Error("異なった通貨です")
    }

    // Objectは不変として保つ
    return new Money(currency, this.amount + amount)
  }
}
 
const myMoney = new Money("JP-YEN", 10);
const changed = myMoney.add('JP-YEN', 10);
console.log(changed)
```

## Entity

値オブジェクトと対をなす。同一性により識別されるため属性が同じでも別オブジェクトとして扱う(=同姓同名でも異なるオブジェクトとして扱う)

- 可変である(=同じ人の身長は伸びるし・年齢も変わる)
- 同じ属性であっても区別される(同姓同名でも異なるオブジェクトとして扱う)
- 同一性により区別される

### 基本例

Value Objectとは異なり、エンティティの属性を変化させたいときには、そのふるまいを通じて代入を行うことで属性を変更する。

```jsx
class User {
  id: string
  firstName: string;
  lastName: string;

  constructor(firstName: string, lastName: string) {
    this.id = // generate random ID
    this.firstName = firstName
    this.lastName = lastName
  }

  changeLastName(newLastName: string) {
    return this.lastName = newLastName
  }
}
```

## ドメインサービス

値オブジェクトやエンティティに記述すると不自然になる振る舞いを表現するもの。(該当インスタンス自身のふるまいとして存在すると自身の値を調べたり違和感がでる。)

つまりドメインサービスは自身の振る舞いを変更するようなインスタンス特有の状態を持たないオブジェクト。

全ての振る舞いはドメインサービスに記述可能なので、可能な限り使用は避けて、不自然なもののみに限定する。(全てをサービスに書くとドメインモデル貧血症を起こすきっかけになる)

```jsx
class UserService {
    exists(user: User) {
        return !!db.findById(user.id)
    }
}
```

## リポジトリ(Repository)

データストアの操作(永続化・再構築)を抽象的に扱えるようにする(=関連の処理を切り離す)

```jsx
interface IUserRepository {
    find: (userId: string) => User
    create: (user: User) => User
}

class UserRepository implements IUserRepository {
  find(userId: string) {
    // データストアから取得処理
		const user = orm.findById(userId)
    return user
  }

  create(user: User) {
    // データストアへの作成処理
    const user = orm.create(user)
    return user
  }
}

// 🌟🌟ServiceはInterfaceに依存させることで、具体的な実装から切り離す！！
class UserService {
    repository: IUserRepository
    
    constructor(repository: IUserRepository) {
        this.repository = repository
    }

    exists(user: User) {
        return !!this.repository.find(user.id)
    }

    createUser(user: User) {
        return this.repository.create(user)
    }
}
```

## アプリケーションサービス

複数のサービスを跨った実際のアプリケーションのユースケースを表現するサービス層。

ドメインオブジェクトのタスク調整に徹するべきで、ドメインルール等は記述してはならない！！

## その他参考

- [ドメイン駆動設計入門 ボトムアップでわかる! ドメイン駆動設計の基本](https://www.amazon.co.jp/%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E9%A7%86%E5%8B%95%E8%A8%AD%E8%A8%88%E5%85%A5%E9%96%80-%E3%83%9C%E3%83%88%E3%83%A0%E3%82%A2%E3%83%83%E3%83%97%E3%81%A7%E3%82%8F%E3%81%8B%E3%82%8B-%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E9%A7%86%E5%8B%95%E8%A8%AD%E8%A8%88%E3%81%AE%E5%9F%BA%E6%9C%AC-%E6%88%90%E7%80%AC-%E5%85%81%E5%AE%A3/dp/479815072X)
- [プログラミング言語Go完全入門](https://docs.google.com/presentation/d/1RVx8oeIMAWxbB7ZP2IcgZXnbZokjCmTUca-AbIpORGk/edit#slide=id.g4f417182ce_0_0)
- [【Go + レイヤードアーキテクチャー】DDDを意識してWeb APIを実装してみる](https://tech.yyh-gl.dev/blog/go_web_api/)
- [https://github.com/jojoarianto/go-ddd-api](https://github.com/jojoarianto/go-ddd-api)
- [https://github.com/eyazici90/go-ddd](https://github.com/eyazici90/go-ddd)

