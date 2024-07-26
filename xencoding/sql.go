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
	// 此处可自定义 cfg
	sql = cfg.Pretty(stmt.AST)
	// 此处会根据 LineWidth 作 LEFT JOIN ON 截断, 单独处理
	sql = regexp.MustCompile(`ON\s+`).ReplaceAllString(sql, "ON ")
	return sql
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
