// 建造者模式
// https://design-patterns.readthedocs.io/zh_CN/latest/creational_patterns/builder.html
// https://github.com/senghoo/golang-design-pattern/blob/master/06_builder/builder.go

package creative

// Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// 建造师 负责智慧建造流程和工艺 用户只需调用构建接口即可
type Director struct {
	builder Builder
}

// NewDirector
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct Product
// 建造者拼装流程
func (this *Director) Construct() {
	this.builder.Part1()
	this.builder.Part2()
	this.builder.Part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}
