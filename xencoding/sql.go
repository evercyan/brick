package xencoding

import (
	"regexp"
	"strings"

	"github.com/auxten/postgresql-parser/pkg/sql/parser"
	"github.com/auxten/postgresql-parser/pkg/sql/sem/tree"
)

// SQLPretty ...
func SQLPretty(sql string) string {
	stmt, err := parser.ParseOne(sql)
	if err != nil {
		return sql
	}
	cfg := tree.DefaultPrettyCfg()
	// custom config
	return cfg.Pretty(stmt.AST)
}

// SQLMinify ...
func SQLMinify(sql string) string {
	// 过滤首尾空格
	sql = strings.TrimSpace(sql)
	// 换行替换成空格
	sql = strings.ReplaceAll(sql, "\n", " ")
	// 过滤注释
	sql = regexp.MustCompile(`--.*$|/\*.*?\*/`).ReplaceAllString(sql, "")
	// 替换连续空格为单空格
	sql = regexp.MustCompile(`\s+`).ReplaceAllString(sql, " ")
	return sql
}
