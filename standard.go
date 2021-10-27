package asciitxt

const StandardLength = 6

func getStandardLetter(letter string) []string {
	switch letter {
	case "A":
		return []string{
			`    _    `,
			`   / \   `,
			`  / _ \  `,
			` / ___ \ `,
			`/_/   \_\`,
			`         `,
		}
	case "B":
		return []string{
			` ____  `,
			`| __ ) `,
			`|  _ \ `,
			`| |_) |`,
			`|____/ `,
			`       `,
		}
	case "C":
		return []string{
			`  ____ `,
			` / ___|`,
			`| |    `,
			`| |___ `,
			` \____|`,
			`       `,
		}
	case "D":
		return []string{
			` ____  `,
			`|  _ \ `,
			`| | | |`,
			`| |_| |`,
			`|____/ `,
			`       `,
		}
	case "E":
		return []string{
			`  _____ `,
			` | ____|`,
			` |  _|  `,
			` | |___ `,
			` |_____|`,
			`        `,
		}
	case "F":
		return []string{
			` _____ `,
			`|  ___|`,
			`| |_   `,
			`|  _|  `,
			`|_|    `,
			`       `,
		}
	case "G":
		return []string{
			`  ____ `,
			` / ___|`,
			`| |  _ `,
			`| |_| |`,
			` \____|`,
			`      `,
		}
	case "H":
		return []string{
			` _   _ `,
			`| | | |`,
			`| |_| |`,
			`|  _  |`,
			`|_| |_|`,
			`       `,
		}
	case "I":
		return []string{
			` ___ `,
			`|_ _|`,
			` | | `,
			` | | `,
			`|___|`,
			`     `,
		}
	case "J":
		return []string{
			`     _ `,
			`    | |`,
			` _  | |`,
			`| |_| |`,
			` \___/ `,
			`       `,
		}
	case "K":
		return []string{
			` _  __`,
			`| |/ /`,
			`| ' / `,
			`| . \ `,
			`|_|\_\`,
			`      `,
		}
	case "L":
		return []string{
			` _     `,
			`| |    `,
			`| |    `,
			`| |___ `,
			`|_____|`,
			`       `,
		}
	case "M":
		return []string{
			` __  __ `,
			`|  \/  |`,
			`| |\/| |`,
			`| |  | |`,
			`|_|  |_|`,
			`        `,
		}
	case "N":
		return []string{
			` _   _ `,
			`| \ | |`,
			`|  \| |`,
			`| |\  |`,
			`|_| \_|`,
			`       `,
		}
	case "O":
		return []string{
			`  ___  `,
			` / _ \ `,
			`| | | |`,
			`| |_| |`,
			` \___/ `,
			`       `,
		}
	case "P":
		return []string{
			` ____  `,
			`|  _ \ `,
			`| |_) |`,
			`|  __/ `,
			`|_|    `,
			`       `,
		}
	case "Q":
		return []string{
			`  ___  `,
			` / _ \ `,
			`| | | |`,
			`| |_| |`,
			` \__\_\`,
			`       `,
		}
	case "R":
		return []string{
			` ____  `,
			`|  _ \ `,
			`| |_) |`,
			`|  _ < `,
			`|_| \_\`,
			`       `,
		}
	case "S":
		return []string{
			` ____  `,
			`/ ___| `,
			`\___ \ `,
			` ___) |`,
			`|____/ `,
			`       `,
		}
	case "T":
		return []string{
			` _____ `,
			`|_   _|`,
			`  | |  `,
			`  | |  `,
			`  |_|  `,
			`       `,
		}
	case "U":
		return []string{
			` _   _ `,
			`| | | |`,
			`| | | |`,
			`| |_| |`,
			` \___/ `,
			`       `,
		}
	case "V":
		return []string{
			`__     __`,
			`\ \   / /`,
			` \ \ / / `,
			`  \ V /  `,
			`   \_/   `,
			`         `,
		}
	case "W":
		return []string{
			`__        __`,
			`\ \      / /`,
			` \ \ /\ / / `,
			`  \ V  V /  `,
			`   \_/\_/   `,
			`            `,
		}
	case "X":
		return []string{
			`__  __`,
			`\ \/ /`,
			` \  / `,
			` /  \ `,
			`/_/\_\`,
			`      `,
		}
	case "Y":
		return []string{
			`__   __`,
			`\ \ / /`,
			` \ V / `,
			`  | |  `,
			`  |_|  `,
			`       `,
		}
	case "Z":
		return []string{
			` _____`,
			`|__  /`,
			`  / / `,
			` / /_ `,
			`/____|`,
			`      `,
		}
	case "a":
		return []string{
			`       `,
			`  __ _ `,
			" / _` |",
			`| (_| |`,
			` \__,_|`,
			`       `,
		}
	case "b":
		return []string{
			` _     `,
			`| |__  `,
			`| '_ \ `,
			`| |_) |`,
			`|_.__/ `,
			`       `,
		}
	case "c":
		return []string{
			`      `,
			`  ___ `,
			` / __|`,
			`| (__ `,
			` \___|`,
			`      `,
		}
	case "d":
		return []string{
			`     _ `,
			`  __| |`,
			` / _  |`,
			`| (_| |`,
			` \__,_|`,
			`       `,
		}
	case "e":
		return []string{
			`      `,
			`  ___ `,
			` / _ \`,
			`|  __/`,
			` \___|`,
			`      `,
		}
	case "f":
		return []string{
			`  __ `,
			` / _|`,
			`| |_ `,
			`|  _|`,
			`|_|  `,
			`     `,
		}
	case "g":
		return []string{
			`       `,
			`  __ _ `,
			" / _` |",
			`| (_| |`,
			` \__, |`,
			` |___/ `,
		}
	case "h":
		return []string{
			` _     `,
			`| |__  `,
			`| '_ \ `,
			`| | | |`,
			`|_| |_|`,
			`       `,
		}
	case "i":
		return []string{
			` _ `,
			`(_)`,
			`| |`,
			`| |`,
			`|_|`,
			`   `,
		}
	case "j":
		return []string{
			`   _ `,
			`  (_)`,
			`  | |`,
			`  | |`,
			` _/ |`,
			`|__/ `,
		}
	case "k":
		return []string{
			` _    `,
			`| | __`,
			`| |/ /`,
			`|   < `,
			`|_|\_\`,
			`      `,
		}
	case "l":
		return []string{
			` _ `,
			`| |`,
			`| |`,
			`| |`,
			`|_|`,
			`   `,
		}
	case "m":
		return []string{
			`           `,
			` _ __ ___  `,
			"| '_ ` _ \\ ",
			`| | | | | |`,
			`|_| |_| |_|`,
			`           `,
		}
	case "n":
		return []string{
			`       `,
			` _ __  `,
			`| '_  \`,
			`| | | |`,
			`|_| |_|`,
			`       `,
		}
	case "o":
		return []string{
			`       `,
			`  ___  `,
			` / _ \ `,
			`| (_) |`,
			` \___/ `,
			`       `,
		}
	case "p":
		return []string{
			`       `,
			` _ __  `,
			`| '_ \ `,
			`| |_) |`,
			`| .__/ `,
			`|_|    `,
		}
	case "q":
		return []string{
			`       `,
			`  __ _ `,
			" / _` |",
			`| (_| |`,
			` \__, |`,
			`    |_|`,
		}
	case "r":
		return []string{
			`      `,
			` _ __ `,
			`| '__|`,
			`| |   `,
			`|_|   `,
			`      `,
		}
	case "s":
		return []string{
			`     `,
			` ___ `,
			`/ __|`,
			`\__ \`,
			`|___/`,
			`     `,
		}
	case "t":
		return []string{
			` _   `,
			`| |_ `,
			`| __|`,
			`| |_ `,
			` \__|`,
			`     `,
		}
	case "u":
		return []string{
			`       `,
			` _   _ `,
			`| | | |`,
			`| |_| |`,
			` \__,_|`,
			`       `,
		}
	case "v":
		return []string{
			`       `,
			`__   __`,
			`\ \ / /`,
			` \ V / `,
			`  \_/  `,
			`       `,
		}
	case "w":
		return []string{
			`          `,
			`__      __`,
			`\ \ /\ / /`,
			` \ V  V / `,
			`  \_/\_/  `,
			`          `,
		}
	case "x":
		return []string{
			`      `,
			`__  __`,
			`\ \/ /`,
			` >  < `,
			`/_/\_\`,
			`      `,
		}
	case "y":
		return []string{
			`       `,
			` _   _ `,
			`| | | |`,
			`| |_| |`,
			` \__, |`,
			` |___/ `,
		}
	case "z":
		return []string{
			`     `,
			` ____`,
			`|_  /`,
			` / / `,
			`/___|`,
			`     `,
		}

	case "0":
		return []string{
			`  ___  `,
			` / _ \ `,
			`| | | |`,
			`| |_| |`,
			` \___/ `,
			`       `,
		}
	case "1":
		return []string{
			` _ `,
			`/ |`,
			`| |`,
			`| |`,
			`|_|`,
			`   `,
		}
	case "2":
		return []string{
			` ____  `,
			`|___ \ `,
			`  __) |`,
			` / __/ `,
			`|_____|`,
			`       `,
		}
	case "3":
		return []string{
			` _____ `,
			`|___ / `,
			`  |_ \ `,
			` ___) |`,
			`|____/ `,
			`       `,
		}
	case "4":
		return []string{
			` _  _   `,
			`| || |  `,
			`| || |_ `,
			`|__   _|`,
			`   |_|  `,
			`        `,
		}
	case "5":
		return []string{
			` ____  `,
			`| ___| `,
			`|___ \ `,
			` ___) |`,
			`|____/ `,
			`       `,
		}
	case "6":
		return []string{
			`  __   `,
			` / /_  `,
			`| '_ \ `,
			`| (_) |`,
			` \___/ `,
			`       `,
		}
	case "7":
		return []string{
			` _____ `,
			`|___  |`,
			`   / / `,
			`  / /  `,
			` /_/   `,
			`       `,
		}
	case "8":
		return []string{
			`  ___  `,
			` ( _ ) `,
			` / _ \ `,
			`| (_) |`,
			` \___/ `,
			`       `,
		}
	case "9":
		return []string{
			`  ___  `,
			` / _ \ `,
			`| (_) |`,
			` \__, |`,
			`   /_/ `,
			`       `,
		}
	case " ":
		return []string{
			` `,
			` `,
			` `,
			` `,
			` `,
			` `,
		}
	default:
		return []string{
			``,
			``,
			``,
			``,
			``,
			``,
		}
	}
}
