package visor

/*
* CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
* AVOID EDITING THIS MANUALLY
 */

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 500000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 0
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
	// MaxDropletPrecision represents the decimal precision of droplets
	MaxDropletPrecision uint64 = 3
	//DefaultMaxBlockSize is max block size
	DefaultMaxBlockSize int = 32768 // in bytes
)

var distributionAddresses = [DistributionAddressesTotal]string{
	"2eBjd4iGebyrWhnv76q2xpVb9fEiv8m2KcM",
"2JUcPgocF5uu4FoFuwnTsZojAtYbN3mob2A",
"8yCgVSFr85zEGky2viimgCxfCE36iw9yJP",
"2bh9YpeAQYwQ7PhMTFSGfrUXiAjvHtCXqDU",
"2jx2YVw2pn1S3DjzNXLANNvGG7kcze4Y3rX",
"maMJoPM3p1TRDzn5F7biuboXnym9XwQiwS",
"fKiCdZFmGXXd1bHYVevD6kBr2ZMBYV6Tdh",
"UWFuxNEKfDabAQ4KUKLuTWQ55t5MVHzw6Z",
"fLcc165FjvP6RmSDyPNt2jqkRBUSDn48d",
"2Mq5TVotZSFB9sucdb4x7ztvaKZhhGijSyN",
"2AWNPQucfV2nDbhsg7dHWxdPuA1zWi2bkvD",
"2ZkQjzjYzBchY7QRJiXJRbavrCZXpgiyQcE",
"2dEsmfZneT5FFDSLT2hQZXHatUbENXGT6Q3",
"45FJNUiHW9nkXi2QgFE74QWeVtm8UMQKHX",
"2V7Xq12GuyVbNgtPAjQCxfrAKYxju1Dtaxy",
"2waYFwFGfF9zojgZEDNrieDH2VsHUdF9mS",
"2jdccBFj9yhtDEg6WoFhceGfdF7VsYsz7zU",
"2aCZuxCg2SMnpk22iGzuhA3XAFPngZ3QW14",
"uik9P1DMwoGkfLpRK4c7Cw7xy3mXAvXgbF",
"2LJXnCRpX2jcEY6cKehkFecbyTs6uVrMLN5",
"dZg5LYkCx8bwcSRFvGuTcPqKRBV7J7dWjg",
"UiUtdJAcjFqqK283un4ofwDbrMKcPrTboz",
"28sGTGaf8aq2LDwpSMuop41dQbuPVQGTbbK",
"R6ULMXNBNv9qX8jPX1J5P9nz1wcwnY2a8z",
"or62ERu1LFWDk1iw5FRFYaY8fUSmZS1qU",
"2Z4v9JbB1NgAheTRXJy8ifHAfxKDHvNEQH7",
"2DBtjWjvEjwCgroZjpwGyihjH2srWAnZJ4M",
"2DgdDVBr82NXLsKMao9WknjUfPpWjybNRA8",
"GHrLnv7CC1PQ8iGaYaU9WCHguHAazmn6PG",
"RaJR455x7vidLc496hgGbhbx6JYrCTTokj",
"Jirdud4rosN9MXDh7SvFihM6RypdmJUsfv",
"2A71qQ3tHDhG7TdnKKFBQGir5cj1ZYkoy2M",
"28UL3kFJU4fvpHGBkyDkW1kEbQrLFwGevHm",
"9bz5mxu1y6Pytj2otSQZT3GZSunw4K3h5F",
"2eJxPafjFZG9V9qBaZFLt4wvyQEZMqtccAP",
"24HCvr8CW1vvCoNv8BUFBzCcZfBYNACkKLQ",
"2821DfKdZpoDkjSvPoNV3c9dQLJrW1uuEqk",
"MnDfotQmjraFQMG9LkbhBGQyPJ5Xkjr4dQ",
"2aBnnHs1zRzfSzxB5vcwaG2qUJvrGpHsWL7",
"NHNqcwMahukiosSutukTGeAwTafP7YaogV",
"22nx29rmDQLgY1TYdxhjQsVuoDWScTqHdEt",
"2aDU7PfJtaBV1hw7jwSskaQhVt5dLq3d6au",
"VxB5z8yWw74CvxS4hWp5ZridGAfKWWWT6U",
"2VYcBRKjM3u7btcjm4TCSJysmDh9ovc8qa9",
"2LzXcvi2Xu6qS5i8qT6bdoXRFaP8bCF6HCz",
"Tinoi2XubcuLcZyr4qRyPPkVkEXTBXcnGe",
"nLQzEut7xJomkLwbjKZjZJzm3Mu4D3pP2Q",
"2iwCLTjCpwSHDtDaGi6kaV9aTZBoxtbyt3y",
"25obz5WkYELRwuuJyYt9s29YkkWr1NNs9tm",
"V48ACygY4SVxQz6P7bEQyRQdgfeGEcZg4A",
"2MjFp1tx1MK49exhYdnocBZw7gWr3Ssx9pX",
"DNVu9Qd4RxsUPqqjCJ9FPgMvQyBEdfkFx5",
"28CpsNBpwC5PRBWnBNMtY6AsTEbam6wrXsb",
"B41qEYhda5EPUbQdat6js1H7WUNzTvHKrb",
"WQM4kd4DD6pEBu8QPwM77k8X9ai52TGe7",
"2VxyjfnFadLNwP726A5BYYNJQDZjyUguqi2",
"2WQb7Y8VjUuTP6UUGDE4AUtMizGXDa4ho4v",
"vAGJiQ7uXH5xCt4cwuXM97CQV42yWYVocE",
"2ayGa9ogD1XiQUJ11pjin9ekcDCzBLxCMPo",
"2GbCraEx9JFaGSECsx1sEfsgXuEMAqraCkD",
"5612oK2taV41Ex77yqne4nAk4M31uq9b7F",
"VTFU84W434B6kLw9R7n2ji6rwGK96ieAGb",
"h4twFnGLygLgTTqWfo4tqiZgQqyv6uZvtn",
"BQKWGmB7gf6FMTvGgZZhEcs6X1y1qoKEXu",
"ASvqRv1fvYK7WdjeyjY36RrKrAD9dzVayB",
"MNoVs7WxN2rRtMndtGZg9zBtVzXnfUVDGd",
"2UJUfHfDtLwSHB4gHu9PUt4kSQTN7PuTDVf",
"23TepiL3s7FmdAyNqohBNSrusWHH5pAruj4",
"PBXBf8orRHqXcHV1pADA6miQZ1MhKqvBqQ",
"21bL9L3Steq7WD6iey4w74pwEvyJv7YDc4e",
"BXQKaGZbM6sZKHnYaHeBJP1LdQFvJk7zGW",
"5EmLkXAoieov3Hy8sCikj5MusSjgBENANL",
"HyoAHtZa2pdVvpy2s5uowXoS2xRCL9Tzcw",
"EF5A136i1YUaLZWP1vHggZXykBgw3YePur",
"27SZE8fK4vMKM7WgSrvCxf4TegzATfZjdvx",
"2g3ZpvETjupRbRPAdNNHNX9Rbo6ihsaUtdc",
"9zESqfijHWS8exZqcA49tU6oBczWaVWsV8",
"UmGXy2sxQTdWNpNShSn4Es8gjcVWUQwmvi",
"2MYiwDKdwGQNxkELTEeaEmoQnCyLd96hQ6k",
"2ju8tWStuuAZzChwCAF4QUcnxaKKVudqget",
"f1upfuWhdkzdQX1v9vg3YF6z3DRpucXedz",
"2hRnhmDiYaUXFncs77XCfcbYzvepw2s7ui4",
"8NCeioJmZtW2TkKu9aExCi7jZEjEZt3BZC",
"2NkRfC9MZ5hZJWLETVVUpJEkfG9zwNVvNrY",
"2E2jMuaoX79YaxzFkCdS3mwWSasYdHzpGyd",
"9Pao2AgKYoNJF3frjKjrmuqoREm7feoFqk",
"2j7p8X68oyvwVSvNbohwCnmPyZbybYpofAV",
"GKfMAnJBNeAsvv9kKSBdQ7Yo3BbzpVQUXp",
"Sv7fS5tSTUX4vR3Bek4bzUYjgU6fi2594m",
"xJnHAEts21x9kBtyDpyvbLtj1QbnvpofiD",
"XzrYgFMA1rksxc8Yihd9qB4JSzUnmUfKqS",
"irXvxF5fs1bJEuXoM8x52H5NRYVK7erdx2",
"esgvDh5i53gFUpHwMmMQWtytfyxaPXiWp6",
"58DMLzth6HQtEgxiBNEn5akEMckQ2j9pET",
"UWueaqTYBnTmnceKdWBvtcuSRhye6WnH6t",
"2KAcd54BBKs2Fs93UqKbcUyapfdBchQhokR",
"2PE8LVACjSKSfBycBpFDixzQ4WKxtFkkFeW",
"2QFsGajPdLdPbfu1DcjizR7DgLCDodnQuU3",
"SCTwpgeDdUz63zGNbM4ZJDhzxp7VYqTzjF",
"6w38gFV6afLoy6uhDd8gEMAZnLho4M6z26",}
