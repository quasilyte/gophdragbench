package yourock

import . "github.com/quasilyte/gophers-and-dragons/game"

const MagicArrowMP = 1
const FireboltMP = 3
const HealMP = 4

const HealAvg = 12.5

const DragonDmgAvg = 5.5
const PowerAttackDmgAvg = 4.5
const AttackDmgAvg = 3

func canSurviveDragon(s State) bool {
	dragonHP := float64(s.Creep.HP)
	myHP := float64(s.Avatar.HP)

	manaLeft := s.Avatar.MP
	healsLeft := s.Deck[CardHeal].Count
	parriesLeft := s.Deck[CardParry].Count
	powersLeft := s.Deck[CardPowerAttack].Count
	stunsLeft := s.Deck[CardStun].Count

	for dragonHP > 0 && myHP > 0 {
		if myHP <= 25 && healsLeft > 0 && manaLeft >= HealMP {
			myHP += HealAvg - DragonDmgAvg
			healsLeft--
			manaLeft -= HealMP
			continue
		}

		if parriesLeft > 0 {
			dragonHP -= DragonDmgAvg
			continue
		}

		if stunsLeft > 0 {
			stunsLeft--
			if powersLeft > 0 {
				powersLeft--
				dragonHP -= PowerAttackDmgAvg
			} else {
				dragonHP -= AttackDmgAvg
			}
			continue
		}

		myHP -= DragonDmgAvg

		if powersLeft > 0 {
			powersLeft--
			dragonHP -= PowerAttackDmgAvg
			continue
		}

		dragonHP -= AttackDmgAvg
	}

	println("dragon hp: ", dragonHP, ", my hp: ", myHP)

	return myHP > 0 && dragonHP < 0
}

func fightDragon(s State) CardType {
	if s.Creep.IsFull() && !canSurviveDragon(s) {
		if s.Can(CardHeal) {
			return CardHeal
		}

		println("retreated, dragon hp: ", s.Creep.HP, ", my hp: ", s.Avatar.HP)

		return CardRetreat
	}

	if s.Avatar.HP <= 25 && s.Can(CardHeal) {
		println("Before heal:", s.Avatar.HP)
		return CardHeal
	}

	if !s.Creep.IsStunned() && s.Can(CardParry) {
		return CardParry
	}

	if !s.Creep.IsStunned() && s.Can(CardStun) {
		return CardStun
	}

	if s.Can(CardPowerAttack) {
		return CardPowerAttack
	}

	return CardAttack
}

func fightMummy(s State) CardType {
	if s.Creep.IsFull() {
		if s.Avatar.MP >= 2*FireboltMP+HealMP && s.Deck[CardFirebolt].Count >= 2 {
			return CardFirebolt
		}

		return CardRetreat
	}

	if s.Creep.HP <= 2 {
		return CardAttack
	} else if s.Creep.HP <= 3 && s.Avatar.MP >= HealMP+MagicArrowMP {
		return CardMagicArrow
	} else if s.Avatar.MP >= FireboltMP+HealMP && s.Can(CardFirebolt) {
		return CardFirebolt
	}

	return CardRetreat
}

func Hero(s State) CardType {
	println("#", s.Turn, "HP: ", s.Avatar.HP)

	if s.Creep.Type == CreepCheepy {
		if s.Avatar.HP <= 25 && s.Can(CardHeal) {
			println("Before heal:", s.Avatar.HP)
			return CardHeal
		}

		if s.Avatar.HP <= 20 && s.Can(CardRest) {
			println("Before rest:", s.Avatar.HP)
			return CardRest
		}

		return CardAttack
	}

	if s.Creep.Type == CreepDragon {
		return fightDragon(s)
	}

	if s.Avatar.HP <= 20 && s.Can(CardHeal) {
		return CardHeal
	}

	if s.Creep.Type == CreepMummy {
		return fightMummy(s)
	}

	if s.Turn%2 == 0 && s.Creep.Damage.High() >= 4 && s.Creep.Type != CreepFairy && !s.Creep.IsStunned() && s.Can(CardParry) {
		return CardParry
	}

	if s.Turn%2 == 1 && !s.Creep.IsStunned() && s.Can(CardStun) {
		return CardStun
	}

	if s.Round >= 8 && s.Avatar.MP >= FireboltMP+HealMP && s.Can(CardFirebolt) {
		return CardFirebolt
	}

	if s.Creep.Damage.High() >= 4 && s.Can(CardPowerAttack) {
		return CardPowerAttack
	}

	if s.Creep.Type == CreepFairy && s.Creep.HP%3 == 0 && s.Avatar.MP >= MagicArrowMP+HealMP {
		return CardMagicArrow
	}

	return CardAttack
}