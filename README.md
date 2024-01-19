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


