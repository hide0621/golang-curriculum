# golang-curriculum

こちらの教材を使用しています

https://www.udemy.com/course/go-crash-course/

# workspaceモード（複数モジュールの使用）について

こちらの記事を参照しています

https://selfnote.work/20220812/programming/golang-go-work/

https://yossi-note.com/golang_workspaces/

https://qiita.com/Rqixy/items/6bdead71dc02eb233376

なお、`go work init`については`go.work`ファイルと`go.work.sum`ファイルを作りたい階層（レポジトリ内の一番上の階層が多い？）で使用し、

workspaceモードで管理する為にも、管理下に置きたいモジュールを作成してから上記コマンドを実行すること

これは新しく作成したモジュールや、新たにworkspaceモードで管理したいモジュールを当該モードで管理する（`go work use モジュール（go.modファイル）が含まれているディレクトリ名（モジュール名ではない）`にて）場合でも同様

また、`go work init`などで指定するモジュールは、そのモジュール（go.modファイル）が含まれているディレクトリ名（モジュール名ではない）を指定すること

# 同一モジュール内にmainパッケージが複数ある場合の実行の方法

`go run main.go` だと 

command-line-arguments

./main.go:6:12: undefined: Add

このようなエラーが出るので...

`go run *.go` とする

# スタックとヒープの割り出しについて

以下のコマンドを実行すると、どの変数や引数、値がスタックにあるのか、それともヒープにあるのかがわかる

`go build -gcflags '-m -l' main.go`

上記コマンドを実行して、

`moved to heap`部分がヒープに割り当てられたもの、

`... argument does not escape`部分がスタックに割り当てられたもの（ex. 関数の引数や値、関数内のローカル変数など）

とのように読み解く

また上記のコマンドにより、コマンドを実行したディレクトリ内に`main`というファイルが作成される

# gitの管理下に置きたくないファイルなどの扱い方について

既にgithubに表題のようなファイルなどを上げている場合は以下のようなコマンドを実行する

`git rm --cached 4/main`（golang-curriculumがカレントディレクトリであり、その中の4/mainをローカルのgitの管理下から外す場合）

上記コマンドにより対象ファイルなどをローカルのgitの管理下から外すことができたので、

この変更分をcommitして、pushすれば表題のファイルなどをローカル/リモートのgitの管理下から外すことができ、github上のリポジトリからもそのファイルなどが削除される（ローカルのディレクトリにはそのファイルは残る）

また、まだ対象のファイルなどをgithubに上げていないのであれば、プロジェクトのルートディレクトリ（今回の場合はgolang-curriculum）に`.gitignore`ファイルを作成して、

対象のファイルなどのパスをそこに記載し、新たに追跡を開始するファイルなどをgitの管理下から外すことができる

書き方についてはこのディレクトリ内の`.gitignore`ファイルを参照

また環境毎に`gitignore`ファイルを生成してくれるツールもあるので、それを活用しても良い

詳しくは以下の記事と公式サイトを参照

[Githubのgitignoreファイルの使い方](https://qiita.com/gk12/items/944b0ef5932adb9a6bd6)

[gitignore.io](https://www.toptal.com/developers/gitignore)

# MySQLの環境構築（パスワード無しでログインできないようにする）について

[前提] **homebrew**で入れたものとする

以下のコマンドを実行する

`mysql_secure_installation`

上記のコマンドを実行することで、

・初期パスワードからユーザーが指定したパスワード（rootユーザーのパスワード）に変更できる

・パスワードにバリデーション（パスワードの強度）を設定できる

・ローカル以外のホストが遠隔でrootユーザーとしてログインできないようになる

・匿名ユーザーの削除

・testデータベースの削除

・「test_」から始まるデータベースへの接続権限の削除

・特権テーブルのリロード（更新内容の反映）

上記のようなメリットがあり、セキュリティー対策としてやっておいた方がいい

詳しくは以下のサイトを参照

https://dev.mysql.com/doc/refman/8.0/en/mysql-secure-installation.html

https://qiita.com/hato_poppo/items/d3c2bb3aa97e09b123c3

# envファイルとgitignoreファイルで繊細な情報を管理する

MySQLのユーザーパスワード等、gitの管理下に置きたくない情報については以下の手順を踏むこと

・.envファイルに上記の繊細な情報を書く

・.gitignoreファイルにルートディレクトリ以下の全ての.envファイルがgitの管理下から外れるように書く（詳しくはこのレポジトリー内の.gitignoreファイルや参考ドキュメントを参照）

詳しくは以下の記事を参照

[Go + MySQLに.envを使う](https://zenn.dev/joo_hashi/scraps/3f5eef2103c6f9)

[Ignoring .env files in my project with .gitignore duplicate](https://stackoverflow.com/questions/59055847/ignoring-env-files-in-my-project-with-gitignore)

# DBeaverの設定について

**オートコミット（自動コミット）はオフ**にして、明示的にトランザクションとコミットのステートメントを使用して、DBに反映させるようにしている

オートコミットがオンだと、反映させるつもりのない更新系のクエリを誤って実行してしまうことも考えられるため...

ex.
[DBeaverでやらかした話](https://qiita.com/tosiooooooo/items/46b2ae7c0687dc061775)

参考：

[業務でDBeaverを利用する上でのオススメ設定](https://qiita.com/hiroga/items/34ac29b2a81b7f64eb1f)










