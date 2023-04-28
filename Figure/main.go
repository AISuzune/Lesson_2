package main

import (
	"fmt"
	"math"
)

// Person 结构体表示一个游戏人物，包含人物的名字、等级、经验值、血量和攻击力等信息。
type Person struct {
	Name      string // 名字
	Level     int    // 等级
	Exp       int    // 经验值
	Health    int    // 血量
	AttackPow int    // 攻击力
}

// Attack 方法表示该角色攻击另一个角色。该方法接受一个 Person 类型的参数，表示被攻击的角色。
// 实现该接口的结构体需要具备攻击能力，即能够将被攻击角色的血量减少相应的攻击力值。
func (p *Person) Attack(target *Person) {
	target.Health -= p.AttackPow // 被攻击角色的血量减少相应的攻击力值
}

func main() {
	// 创建两个角色并初始化名字、等级、经验值、血量和攻击力等信息
	player1 := &Person{
		Name:      "yuan神",
		Level:     1,
		Exp:       0,
		Health:    math.MaxInt64,
		AttackPow: math.MaxInt64,
	}

	player2 := &Person{
		Name:      "zty",
		Level:     1,
		Exp:       0,
		Health:    100,
		AttackPow: 10,
	}

	// 打印被攻击角色的血量变化情况
	fmt.Printf("%s's health: %d\n", player2.Name, player2.Health)
	player1.Attack(player2)
	fmt.Printf("%s's health after attack: %d\n", player2.Name, player2.Health)
}
