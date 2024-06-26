package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defence int
}

func NewCreature(name string, attack int, defence int) *Creature {
	return &Creature{
		Name:    name,
		Attack:  attack,
		Defence: defence}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defence)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type IncreasedDefenceModifier struct {
	CreatureModifier
}

func NewIncreasedDefenceModifier(c *Creature) *IncreasedDefenceModifier {
	return &IncreasedDefenceModifier{CreatureModifier{creature: c}}
}

func (i *IncreasedDefenceModifier) Handle() {
	fmt.Println("Increasing", i.creature.Name, "\b's defence")
	i.creature.Defence++
	i.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{creature: c}}
}

func (n *NoBonusesModifier) Handle() {
	// empty
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())
	root := NewCreatureModifier(goblin)
	// root.Add(NewNoBonusesModifier(goblin)) // Goblin (1/1)
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreasedDefenceModifier(goblin))
	root.Add(NewIncreasedDefenceModifier(goblin)) //Goblin (4/3)
	root.Handle()
	fmt.Println(goblin.String())
}
