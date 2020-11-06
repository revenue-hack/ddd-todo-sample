# ddd-todo-sample

## ユースケース
- ユーザ作成
    - 必須項目
        - ユーザ名: asciiのみ, 20文字以内
        - email: Email形式
        - パスワード: 12文字以内
- TODOの作成
    - 必須項目
        - タイトル: 100文字以内
        - ステータス: icebox, todo, confirm, done
        - 内容: 1000文字以内
        - アサイニー: 必ず一人存在する
        - 参考URL: 任意(一つまで)
    - 一人のユーザにつきtodoは3個まで
    - 一人のユーザにつきconfirmは5個まで
    - doneステータスになると表示されない
    - 新規作成時は必ずiceboxステータスになる
    - todoを作成以降浅いニーはDONEになるまで変更できない
