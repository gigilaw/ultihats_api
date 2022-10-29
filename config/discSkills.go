package config

var (
	PRIMARY_ROLE = map[int]string{
		1: "Handler",
		2: "Cutter",
		3: "Hybrid",
	}
	THROWING = map[int]string{
		1: "Can throw accurately to stationary targets",
		2: "Can throw accurately to moving targets",
		3: "Can throw into space for reciever to attack the disc",
		4: "Can throw a reciever open",
		5: "Can throw a disc anytime to anywhere I want",
	}
	CATCHING = map[int]string{
		1: "Can catch consitently while stationary without any defensive pressure",
		2: "Can catch consitently while moving with some defensive pressure",
		3: "Can catch while boxing out my defender",
		4: "Can catch misthrown disc within my proximity",
		5: "Just toss it up, I'll get it",
	}
	OFFENSIVE_STRATEGIES = map[int]string{
		1: "I only know how to move when specifically told against person defense/zone",
		2: "I can follow along and make simple continuation cuts when it is obvious",
		3: "I can make instant decisions to change my cuts based on what is happening",
		4: "I know how to move to get open and create space against all defensive strategies",
		5: "I can anchor any type of offensive",
	}
	DEFENSIVE_STRATEGIES = map[int]string{
		1: "I understand marks and person defense but struggle to hold my own",
		2: "I can hold my own against person defense but I struggle to understand zone positioning",
		3: "It will take very creative cutting to get open against me",
		4: "I can adjust my defensive positioning against my mark as I read the play",
		5: "I can read the field and help out my teammates while maintaining my role",
	}
)
