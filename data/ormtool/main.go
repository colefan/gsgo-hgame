package main

import "github.com/colefan/gsgo/orm"

func main() {
	tool := orm.OrmCmd{}
	tool.Run("../data")

}
