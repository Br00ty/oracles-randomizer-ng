package logic

// these nodes do not define items, only which items can kill which enemies
// under what circumstances, assuming that you've arrived in the room
// containing the enemy.
//
// anything that can be destroyed in more than one way is also included in
// here. bushes, flowers, mushrooms, etc.
//
// technically mystery seeds can be used to kill many enemies that can be
// killed by ember, scent, or gale seeds. mystery seeds are only included as a
// kill option if at all three of these seed types work.
//
// if an enemy is in the same room as a throwable object and is vulnerable to
// thrown objects, than just adding "bracelet" as an OR is sufficient.
//
// animal companions are not included in this logic, since they're only
// available in certain areas.

// when testing how to kill enemies, remember to try:
// - sword
// - beams
// - boomerang L-1
// - boomerang L-2
// - rod
// - seeds (satchel first, slingshot if satchel doesn't work)
// - bombs
// - thrown objects (if applicable)
// - magnet ball (if applicable)
// - fool's ore
// - punch
// - what pushes them into pits (if applicable)
//   - sword
//   - beams
//   - shield
//   - boomerangs (they work on hardhats!)
//   - seeds (satchel first, slingshot if satchel doesn't work)
//   - rod
//   - bombs
//   - shovel
//   - thrown objects (if applicable)
//   - NOT magnet ball; it kills anything pittable
//   - fool's ore
//   - punch

var killNodes = map[string]*Node{
	"gale seed weapon":      And("gale seeds", Or("slingshot", HardAnd("satchel", "jump"))),
	"gale boomerang":        And("gale satchel", "boomerang"), // stun, then drop from satchel
	"slingshot kill normal": And("slingshot", "seed kill normal"),
	"jump kill normal":      And("jump", "kill normal"),
	"jump pit normal":       And("jump", "pit kill normal"),

	// required enemies in normal route-ish order, but with prereqs first
	"seed kill normal": Or("ember seeds", "scent seeds", "gale seed weapon", "gale boomerang", "mystery seeds"),
	"pop maku bubble":  Or("sword", "rod", "seed kill normal", "pegasus slingshot", "bombs", "fool's ore"),

	// the "safe" version is for areas where you can't possibly get stuck from
	// being on the wrong side of a bush.
	"remove bush safe": Or("sword", "boomerang L-2", "bracelet",
		"ember seeds", "gale slingshot", "bombs"),
	"remove bush": Or("sword", "boomerang L-2", "bracelet",
		HardOr("ember seeds", "gale slingshot", "bombs")),

	"kill normal":                     Or("sword", "bombs", "beams", "seed kill normal", "fool's ore", "punch"),
	"pit kill normal":                 Or("sword", "beams", "shield", "scent seeds", "rod", "bombs", Hard("shovel"), "fool's ore", "punch"),
	"kill stalfos":                    Or("kill normal", "rod"),
	"kill stalfos (throw)":            Or("kill stalfos", "bracelet"),
	"hit lever":                       Or("sword", "boomerang", "rod", "ember seeds", "scent seeds", "any slingshot", "fool's ore", "punch", "shovel"),
	"kill goriya bros":                Or("sword", "bombs", "fool's ore", "punch"),
	"kill goriya":                     Or("kill normal"),
	"kill goriya (pit)":               Or("kill goriya", "pit kill normal"),
	"kill aquamentus":                 Or("sword", "beams", "scent seeds", "bombs", "fool's ore", "punch"),
	"hit far switch":                  Or("beams", "boomerang", "bombs", "any slingshot"),
	"toss bombs":                      And("bombs", "toss ring"),
	"kill rope":                       Or("kill normal"),
	"kill hardhat (pit, throw)":       Or("gale seed weapon", "sword", "beams", "boomerang", "shield", "scent seeds", "rod", "bombs", Hard("shovel"), "fool's ore", "bracelet"),
	"kill moblin (gap, throw)":        Or("sword", "beams", "scent seeds", "slingshot kill normal", "bombs", "fool's ore", "punch", "jump kill normal", "jump pit normal"),
	"kill zol":                        Or("kill normal"),
	"remove pot":                      Or("sword L-2", "bracelet"),
	"kill facade":                     Or("bombs"),
	"flip spiked beetle":              Or("shield", "shovel"),
	"damage spiked beetle (throw)":    Or("sword", "bombs", "beams", "seed kill normal", "bracelet", "fool's ore"),
	"flip kill spiked beetle (throw)": And("flip spiked beetle", "damage spiked beetle (throw)"),
	"gale kill spiked beetle":         And("gale seed weapon"),
	"kill spiked beetle (throw)":      Or("flip kill spiked beetle (throw)", "gale kill spiked beetle"),
	"kill mimic":                      Or("kill normal"),
	"damage omuai":                    Or("sword", "bombs", "scent seeds", "fool's ore", "punch"),
	"kill omuai":                      And("damage omuai", "bracelet"),
	"damage mothula":                  Or("sword", "bombs", "scent seeds", "fool's ore", "punch"),
	"kill mothula":                    And("damage mothula", "jump"), // you will basically die without feather
	"remove flower": Or("sword", "boomerang L-2",
		HardOr("ember seeds", "gale slingshot", "bombs")),
	"damage agunima":        Or("sword", "scent seeds", "bombs", "fool's ore", "punch"),
	"kill agunima":          And("ember seeds", "damage agunima"),
	"hit very far lever":    Or("boomerang L-2", "any slingshot"),
	"hit lever gap":         Or("sword", "boomerang", "rod", "any slingshot", "fool's ore"),
	"jump hit lever":        And("jump", "hit lever gap"),
	"long jump hit lever":   And("long jump", "hit lever"),
	"hit far lever":         Or("jump hit lever", "long jump hit lever", "boomerang", "any slingshot"),
	"kill gohma":            And(Or("scent seeds", "ember seeds"), Or("slingshot", Hard("start"))),
	"remove mushroom":       Or("boomerang L-2", "bracelet"),
	"kill moldorm":          Or("sword", "bombs", "punch", "scent seeds", "fool's ore"),
	"kill iron mask":        Or("kill normal"),
	"kill armos":            Or("sword", "bombs", "beams", "boomerang L-2", "scent seeds", "fool's ore"),
	"kill gibdo":            Or("kill normal", "boomerang L-2", "rod"),
	"kill darknut":          Or("sword", "bombs", "beams", "scent seeds", "fool's ore", "punch"),
	"kill darknut (pit)":    Or("sword", "bombs", "beams", "scent seeds", "fool's ore", "punch", "shield", "rod", Hard("shovel")),
	"kill syger":            Or("sword", "bombs", "scent seeds", "fool's ore", "punch"),
	"break crystal":         Or("sword", "bombs", "punch", "bracelet"),
	"kill hardhat (magnet)": Or("magnet gloves", "gale seed weapon"),
	"kill vire":             Or("sword", "bombs", "fool's ore", "punch"),
	"finish manhandla":      Or("sword", "bombs", "any slingshot", "fool's ore"),
	"kill manhandla":        And("boomerang L-2", "finish manhandla"),
	"kill wizzrobe":         Or("kill normal"),
	"kill magunesu":         Or("sword", "fool's ore", "punch"), // even bombs don't work!
	"kill poe sister":       Or("sword", "beams", "ember seeds", "scent seeds", "bombs", "fool's ore", "punch"),
	"kill darknut (across pit)": Or(
		Or("beams", "toss bombs", "scent slingshot", "magnet gloves"),
		And("feather L-2", "kill darknut (pit)")),
	"kill gleeok":      Or("sword", "beams", "bombs", "fool's ore", "punch"),
	"kill frypolar":    Or(And("bracelet", "mystery seeds"), "ember seeds"),
	"kill medusa head": Or("sword", "fool's ore"),
	"kill floormaster": Or("kill normal"),
	"kill onox":        And("sword", "jump"),
}
