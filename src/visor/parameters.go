package visor

/*
* CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
* AVOID EDITING THIS MANUALLY
 */

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 100000000
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
	"23bNjBuUQEfhGdWZNiwb2f7Mu47xyiudYET",
"MYL1eQk3CAJvxQoXjfPmprhSiP73svL1XS",
"2Nm4oJ96Sb629114ERPi7tqEvpBDgt92tyF",
"2SJo7pXfY2w7k4SwmNsjb7Jy2B1vz4BEWHC",
"FWiv9GmDsZR7rtfXYtWPybyVT1iKKFu6GB",
"28cuBTUnfB5MeFGcqtmzb9gTh4zUfKzVVDK",
"5NWMnPw1mMcJRJdAhYkfzUo5uDNUcrgNN1",
"JEACD2rBDGvYcE8qNxYEFDJEE1dexhFhjE",
"2fDUoJzsddK2k8X1Y6QbFisxKFMRjNwErNB",
"uEEo7KrSBYUZjYSvY8Z9PLCLDJBxTFheQv",
"2TPcUbKzfeNWDXAYowJCybDZGiRgw32kgMK",
"WpSSxNVMVfWFimNcTGfgEA3AVCNrchDvfr",
"2N7p2MVvXbr9PbwQWGk7rGG6H2M9riGQcUB",
"yqoecGz1LBz3wQdtPJVHnuZX1EY6hVMjFq",
"2WH6cQa4LirnwUKnJkTmsNpyZLMXEwy4UjS",
"n1nuPRHjtHrmM2yGKm2Uz5AhKmKT8JTX7p",
"yhTPvq73ezReis3dZPA5YpzNnxHzHTkALB",
"2QCdCYmEZZYrhvcoms8qJmW57hfbPBEBwXL",
"2X3L48KCcHi2y5xpvnQC21x1PDPN6QLqzRJ",
"HFrbigTAqNc4n51YFAQ8ta8Ey3XqmpHqPJ",
"2W59hCZeqc1kURS1hKq9Swyjmf535DcbwBE",
"hNdDNtak3JYnpdRapoQZtT1mBJt56uMg6q",
"VPTBXcHsSY6ExfXkEvmFVvkYeLsJp7RdME",
"2LmPHmsKpS83g4QHfMvxvTHqc6GcYumfJvT",
"2cHn2Z3uPmraS8cGRcQeeZv2QbkFhKCmT1C",
"2XzE3ec6WNSPTo8WFkHYHjCp8NZzMT2VPfp",
"BaCCBy6eACE7ZgdT38TmU9n6TwTFW7agWn",
"W93CcpUmAwyyWJfaoz86vJuQbTXHMUXX9s",
"PdTiWAxjNc54D5GZoyF1qqfU71qXUyEeth",
"2WTKLguHpfJxm7Jc7HDdEMFww8Ssf4TC93h",
"23G1k2fYwQYqY9ACkeukWEhbpMx62auamYF",
"D91krziPz4snSRvkdaDX939FVBHA1fiwUa",
"J9hgP4iJ9femPEmrQU1NovMaG4qBQjQ92x",
"i3aRzqgjivBZeX2wFhazXqQAfSwzWXHLZD",
"26VWjg4x3wSEHE1vK1FyTUHfcBbS8gp8dsT",
"8UchvQEBBEytxndRTpXG3T32k3aPykqpfd",
"kPutyqB4usBsiDyqjDYyB8PeKFsJPZBsY4",
"2T64cBRvZrfw2fHXDYhT2QtvJtCBEAw4xbb",
"gASYNqkQwT4b1uJu5nhFqQzaEYiUYMXKhW",
"2HAFcb7dkSSEirBb8S8DBoz2yt8MxPSFyxZ",
"2Vm1K3yvRqXVxg3TGBvryvHHGhL18GmyUd9",
"GJPHA5ELhcqvdvtSs942bfXq5aeaqD31d3",
"TfaH8cLx92yZ8Czq2JVozxvCAzX8ToksMy",
"T2yhCt8NQf8z4PztyXbe3XGjwJjggz8cYD",
"2GGURjuGfwmsUpxtKqBvXaYtqTd91qqLU15",
"BKq56WV5iWjvE4Ch47rKfN6FqhqBa76BS9",
"2DpCcC64JX5hbV7wdz1hxTQjbkcTEi2Ty4V",
"2TroiWS4uB6hy8xW69XY7PRdk1Jubc654SZ",
"2nZNjwyV8gy91V3U4eyiAssdc7Ls7BNjSq",
"5ecmCxF7zWyx97KoSbGPLLXxVsfXFkVQoz",
"2Ve67p5a8VuesR9QJakQfmBKbuB2E3YXfiZ",
"xxYvUHJXZ7huisjkBC3TqKjaZqxJyataKq",
"BcZKm3CMwaJ611RDLQm2sv3EhDKiqbwebX",
"LS8SdBCpMW6dGKYNMZhUvTSFL8KgGFes4a",
"2KwE6a518EN3WyufuJsZj2M6T2nVcW8x1xb",
"21ijQHwFuYuWn3oiP5YeiQDCLyBswphyqsj",
"2YrexqhQbuzy57u2MdQAa4zRM7TcvyWEkdQ",
"ptKy9M67zTWX8uA84UqQ9h1pNgTwHmHejD",
"2kwYYHt5rCEbTCAsJ8D4pEbb8CiyqjXFSri",
"2hF9FsnqPR4H2bT5MasWo7QqWfAuLx1dxb5",
"WMVV8pxc5iswLepL89npQaMnC7aoqryN4F",
"28YYCP7mV1DfMLiGmf1ZAyr1AeiZmehBnsX",
"2JN5JigTv9Mb6kUBbPXydEBhMXEsvf32VJk",
"qP3cGCvVXykvkhEXVjXssrpxCLV5MhTAMJ",
"2MJNzJU17LTdbDmm84gLR2qTNk4San1DJJY",
"92nifkaphsBkZ9VRgktkEqqSGkroDaNRhQ",
"cnC5mTDjQMPnmv6jzSkDqdPjrpNRVVsfNB",
"2SQ2ZYp6eNtxUitRhnHgmsGqg3Ve4YEFYbC",
"2RDu9UT49XuzVyBe67zBqxQR5cAt2t2KkRi",
"d2GEgVr1EkCjC664MqL9Mx1AtWscvk9VXP",
"22znryMC36GVJkoDtruxDvUfP6rFQbF6zob",
"289i7SnAzPqtuVXWkFVBQyWiLZhZygXhmSG",
"2e32aNdnZ63NHMasvxkE3b5cGvre1CACExQ",
"jX3xJjitR8nBtWJXG4Q4wLKctq7SwCqxRy",
"QfySmE5E9NyuYJ5xSJ1NLqjx9EQ6X9UZvK",
"S7fksUk2bgueUTQ1oWEi4S7CtfPE4ALnuL",
"2dLV4qmJEu5BcBNNjy7nxsyxSQTLW8W87uu",
"U7FMBy5D81uWuE5yfFWg5R9VufokDt8zmz",
"2M4KAUsKzPrkoGuSTa9vXTqRKH7h3b25MfB",
"22pGBTfcpNcGgnQsY8wHXDaUCYp1THPHDbd",
"PKhPArP4Yy1s9nKta2LfZVuFBJoYFFgLPq",
"2QmUP7AL3W8eVW9j9AgxKxXe8iqYELV7Lu8",
"pdRYWyjpaSaRswx33wzWCPdwAeWcPvvRq1",
"2f75891emB82nZmrkrjRok7U5XJrATR1LCb",
"wxEKdg6AcTp7Ueo1isnbR4EKi9ufR5JTPa",
"z4Es9X7xvtuB1qfV5CZT5JG1p2GxAgrnko",
"2ZpnnkTSVvaA9ENWFNyXJWycrRiEntroJaJ",
"2W1VPJLs8yXpXuXPcZcAzEtrW5p2pbKFYCD",
"2Hcokk8EjjcSXGvHgRzsQ8SyswCPthWu9Lc",
"GjaWu68GcnsxtnTpsaRwqhBCRP6ApDhEwX",
"2MTmVkCfkbSbf2WFF7wD5wSoBk3TosX6LbN",
"heLVT5819tVrSy4VZ2G9rVkpKLojNLGjT1",
"28bShKHnV4W8qRZti8uE63HaCdb4jKTqKqp",
"3nTkzrxENBvu53qvgxbiLN55yP7Wf5v8b5",
"2HcXd1p9yBtnYSgRJ5KkBNe3nc2MKKb8EJf",
"vmzDMTohNYVJP6PQJhrrcSeaCzRbDELpQs",
"2PVzujZpkZsayMTMLBphSNdUPp3DPG2eKv3",
"2msYoM9CnJwHX7W9th85CYoXbL61jyCCTkF",
"2Hpw5t4MSnNoFiFgHXKjWuszdVJFuMkfM3S",
"2YuyJoUsRXSj2J7LHL52zDTdveWzkj8eywT",}
