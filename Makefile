run:
	go run main.go

# 测试 ebiten 是否安装成功
test-ebiten:
	go run -tags=example github.com/hajimehoshi/ebiten/examples/rotate