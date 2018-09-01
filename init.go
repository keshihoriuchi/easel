package main

import (
	"fmt"
	"os"
)

var cmdInit = &Command{
	Run:       runInit,
	UsageLine: "init [OPTIONS]",
	Short:     "Generate markdown template",
	Long: `
Options:
	-h, --help     Print usage
`,
}

func init() {
}

func runInit(args []string) int {

	if len(args) > 0 {
		fmt.Fprintln(os.Stderr, "Too many arguments given.")
		return 1
	}

	base := `# 課題

上位3つの課題

## 別の解決手段



# 顧客セグメント


## アーリーアダプター


# 独自の価値提案



## ハイレベルコンセプト



# ソリューション

上位3つの機能

# チャネル



# 収益の流れ



# コスト構造



# 主要指標



# 圧倒的な優位性




`
	fmt.Printf(base)
	return 0
}
