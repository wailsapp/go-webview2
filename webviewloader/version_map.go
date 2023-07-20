//go:build windows

package webviewloader

type Version struct {
	Number         string
	Link           string
	RuntimeVersion string
	Notes          []string
}

var versionMapping = map[string]Version{
	"1.0.1823.32": {
		Number:         "1.0.1823.32",
		Link:           "",
		RuntimeVersion: "114.0.1823.32",
		Notes: []string{
			"Release Date: June 5, 2023",
		},
	},
	"1.0.1905.prerelease": {
		Number:         "1.0.1905.prerelease",
		Link:           "",
		RuntimeVersion: "116.0.1905.0",
		Notes: []string{
			"Release Date: June 12, 2023",
		},
	},
	"1.0.1774.30": {
		Number:         "1.0.1774.30",
		Link:           "",
		RuntimeVersion: "113.0.1774.30",
		Notes: []string{
			"Release Date: May 8, 2023",
		},
	},
	"1.0.1829.prerelease": {
		Number:         "1.0.1829.prerelease",
		Link:           "",
		RuntimeVersion: "115.0.1829.0",
		Notes: []string{
			"Release Date: May 8, 2023",
		},
	},
	"1.0.1722.45": {
		Number:         "1.0.1722.45",
		Link:           "",
		RuntimeVersion: "112.0.1722.45",
		Notes: []string{
			"Release Date: April 13, 2023",
		},
	},
	"1.0.1777.prerelease": {
		Number:         "1.0.1777.prerelease",
		Link:           "",
		RuntimeVersion: "114.0.1777.0",
		Notes: []string{
			"Release Date: April 10, 2023",
		},
	},
	"1.0.1661.34": {
		Number:         "1.0.1661.34",
		Link:           "",
		RuntimeVersion: "111.0.1661.34",
		Notes: []string{
			"Release Date: March 20, 2023",
		},
	},
	"1.0.1724.prerelease": {
		Number:         "1.0.1724.prerelease",
		Link:           "",
		RuntimeVersion: "113.0.1724.0",
		Notes: []string{
			"Release Date: March 20, 2023",
		},
	},
	"1.0.1587.40": {
		Number:         "1.0.1587.40",
		Link:           "",
		RuntimeVersion: "110.0.1587.40",
		Notes: []string{
			"Release Date: February 15, 2023",
		},
	},
	"1.0.1671.prerelease": {
		Number:         "1.0.1671.prerelease",
		Link:           "",
		RuntimeVersion: "112.0.1671.0",
		Notes: []string{
			"Release Date: February 15, 2023",
		},
	},
	"1.0.1518.46": {
		Number:         "1.0.1518.46",
		Link:           "",
		RuntimeVersion: "109.0.1518.46",
		Notes: []string{
			"Release Date: January 17, 2023",
		},
	},
	"1.0.1619.prerelease": {
		Number:         "1.0.1619.prerelease",
		Link:           "",
		RuntimeVersion: "111.0.1619.0",
		Notes: []string{
			"Release Date: January 19, 2023",
		},
	},
	"1.0.1462.37": {
		Number:         "1.0.1462.37",
		Link:           "",
		RuntimeVersion: "108.0.1462.37",
		Notes: []string{
			"Release Date: December 12, 2022",
		},
	},
	"1.0.1549.prerelease": {
		Number:         "1.0.1549.prerelease",
		Link:           "",
		RuntimeVersion: "110.0.1549.0",
		Notes: []string{
			"Release Date: December 12, 2022",
		},
	},
	"1.0.1418.22": {
		Number:         "1.0.1418.22",
		Link:           "",
		RuntimeVersion: "107.0.1418.22",
		Notes: []string{
			"Release Date: October 31, 2022",
		},
	},
	"1.0.1466.prerelease": {
		Number:         "1.0.1466.prerelease",
		Link:           "",
		RuntimeVersion: "109.0.1466.0",
		Notes: []string{
			"Release Date: October 31, 2022",
		},
	},
	"1.0.1370.28": {
		Number:         "1.0.1370.28",
		Link:           "",
		RuntimeVersion: "106.0.1370.28",
		Notes: []string{
			"Release Date: October 11, 2022",
		},
	},
	"1.0.1414.prerelease": {
		Number:         "1.0.1414.prerelease",
		Link:           "",
		RuntimeVersion: "107.0.1414.0",
		Notes: []string{
			"Release Date: October 11, 2022",
		},
	},
	"1.0.1343.22": {
		Number:         "1.0.1343.22",
		Link:           "",
		RuntimeVersion: "105.0.1343.22",
		Notes: []string{
			"Release Date: September 6, 2022",
		},
	},
	"1.0.1369.prerelease": {
		Number:         "1.0.1369.prerelease",
		Link:           "",
		RuntimeVersion: "106.0.1369.0",
		Notes: []string{
			"Release Date: September 6, 2022",
		},
	},
	"1.0.1293.44": {
		Number:         "1.0.1293.44",
		Link:           "",
		RuntimeVersion: "104.0.1293.44",
		Notes: []string{
			"Release Date: August 8, 2022",
		},
	},
	"1.0.1340.prerelease": {
		Number:         "1.0.1340.prerelease",
		Link:           "",
		RuntimeVersion: "105.0.1340.0",
		Notes: []string{
			"Release Date: August 8, 2022",
		},
	},
	"1.0.1264.42": {
		Number:         "1.0.1264.42",
		Link:           "",
		RuntimeVersion: "103.0.1264.42",
		Notes: []string{
			"Release Date: July 4, 2022",
		},
	},
	"1.0.1305.prerelease": {
		Number:         "1.0.1305.prerelease",
		Link:           "",
		RuntimeVersion: "105.0.1305.0",
		Notes: []string{
			"Release Date: July 4, 2022",
		},
	},
	"1.0.1245.22": {
		Number:         "1.0.1245.22",
		Link:           "",
		RuntimeVersion: "102.0.1245.22",
		Notes: []string{
			"Release Date: June 14, 2022",
		},
	},
	"1.0.1210.39": {
		Number:         "1.0.1210.39",
		Link:           "",
		RuntimeVersion: "101.0.1210.39",
		Notes: []string{
			"Release Date: May 9, 2022",
		},
	},
	"1.0.1248.prerelease": {
		Number:         "1.0.1248.prerelease",
		Link:           "",
		RuntimeVersion: "102.0.1248.0",
		Notes: []string{
			"Release Date: May 9, 2022",
		},
	},
	"1.0.1185.39": {
		Number:         "1.0.1185.39",
		Link:           "",
		RuntimeVersion: "100.0.1185.39",
		Notes: []string{
			"Release Date: April 12, 2022",
		},
	},
	"1.0.1222.prerelease": {
		Number:         "1.0.1222.prerelease",
		Link:           "",
		RuntimeVersion: "102.0.1222.0",
		Notes: []string{
			"Release Date: April 12, 2022",
		},
	},
	"1.0.1150.38": {
		Number:         "1.0.1150.38",
		Link:           "",
		RuntimeVersion: "99.0.1150.38",
		Notes: []string{
			"Release Date: March 10, 2022",
		},
	},
	"1.0.1189.prerelease": {
		Number:         "1.0.1189.prerelease",
		Link:           "",
		RuntimeVersion: "100.0.1189.0",
		Notes: []string{
			"Release Date: March 10, 2022",
		},
	},
	"1.0.1108.44": {
		Number:         "1.0.1108.44",
		Link:           "",
		RuntimeVersion: "98.0.1108.44",
		Notes: []string{
			"Release Date: February 6, 2022",
		},
	},
	"1.0.1158.prerelease": {
		Number:         "1.0.1158.prerelease",
		Link:           "",
		RuntimeVersion: "100.0.1158.0",
		Notes: []string{
			"Release Date: February 6, 2022",
		},
	},
	"1.0.1072.54": {
		Number:         "1.0.1072.54",
		Link:           "",
		RuntimeVersion: "97.0.1072.54",
		Notes: []string{
			"Release Date: January 13, 2022",
		},
	},
	"1.0.1133.prerelease": {
		Number:         "1.0.1133.prerelease",
		Link:           "",
		RuntimeVersion: "99.0.1133.0",
		Notes: []string{
			"Release Date: January 13, 2022",
		},
	},
	"1.0.1083.prerelease": {
		Number:         "1.0.1083.prerelease",
		Link:           "",
		RuntimeVersion: "97.0.1083.0",
		Notes: []string{
			"Release Date: November 29, 2021",
		},
	},
	"1.0.1054.31": {
		Number:         "1.0.1054.31",
		Link:           "",
		RuntimeVersion: "96.0.1054.31",
		Notes: []string{
			"Release Date: November 29, 2021",
		},
	},
	"1.0.1056.prerelease": {
		Number:         "1.0.1056.prerelease",
		Link:           "",
		RuntimeVersion: "97.0.1056.0",
		Notes: []string{
			"Release Date: October 29, 2021",
		},
	},
	"1.0.1020.30": {
		Number:         "1.0.1020.30",
		Link:           "",
		RuntimeVersion: "95.0.1020.30",
		Notes: []string{
			"Release Date: October 25, 2021",
		},
	},
	"1.0.992.28": {
		Number:         "1.0.992.28",
		Link:           "",
		RuntimeVersion: "94.0.992.31",
		Notes: []string{
			"Release Date: September 27, 2021",
		},
	},
	"1.0.1018.prerelease": {
		Number:         "1.0.1018.prerelease",
		Link:           "",
		RuntimeVersion: "95.0.1018.0",
		Notes: []string{
			"Release Date: September 20, 2021",
		},
	},
	"1.0.1010.prerelease": {
		Number:         "1.0.1010.prerelease",
		Link:           "",
		RuntimeVersion: "95.0.1010.0",
		Notes: []string{
			"Release Date: September 14, 2021",
		},
	},
	"1.0.961.33": {
		Number:         "1.0.961.33",
		Link:           "",
		RuntimeVersion: "93.0.961.44",
		Notes: []string{
			"Release Date: September 8, 2021",
		},
	},
	"1.0.955.prerelease": {
		Number:         "1.0.955.prerelease",
		Link:           "",
		RuntimeVersion: "93.0.967.0",
		Notes: []string{
			"Release Date: July 26, 2021",
		},
	},
	"1.0.902.49": {
		Number:         "1.0.902.49",
		Link:           "",
		RuntimeVersion: "92.0.902.49",
		Notes: []string{
			"Release Date: July 26, 2021",
		},
	},
	"1.0.902.prerelease": {
		Number:         "1.0.902.prerelease",
		Link:           "",
		RuntimeVersion: "92.0.902.0",
		Notes: []string{
			"Release Date: June 1, 2021",
		},
	},
	"1.0.864.35": {
		Number:         "1.0.864.35",
		Link:           "",
		RuntimeVersion: "91.0.864.35",
		Notes: []string{
			"Release Date: May 31, 2021",
		},
	},
	"1.0.865.prerelease": {
		Number:         "1.0.865.prerelease",
		Link:           "",
		RuntimeVersion: "91.0.865.0",
		Notes: []string{
			"Release Date: April 26, 2021",
		},
	},
	"1.0.818.41": {
		Number:         "1.0.818.41",
		Link:           "",
		RuntimeVersion: "90.0.818.41",
		Notes: []string{
			"Release Date: April 21, 2021",
		},
	},
	"1.0.824.prerelease": {
		Number:         "1.0.824.prerelease",
		Link:           "",
		RuntimeVersion: "91.0.824.0",
		Notes: []string{
			"Release Date: March 8, 2021",
		},
	},
	"1.0.774.44": {
		Number:         "1.0.774.44",
		Link:           "",
		RuntimeVersion: "89.0.774.44",
		Notes: []string{
			"Release Date: March 8, 2021",
		},
	},
	"1.0.790.prerelease": {
		Number:         "1.0.790.prerelease",
		Link:           "",
		RuntimeVersion: "86.0.616.0",
		Notes: []string{
			"Release Date: February 10, 2021",
		},
	},
	"1.0.705.50": {
		Number:         "1.0.705.50",
		Link:           "",
		RuntimeVersion: "86.0.616.0",
		Notes: []string{
			"Release Date: January 25, 2021",
		},
	},
	"1.0.721.prerelease": {
		Number:         "1.0.721.prerelease",
		Link:           "",
		RuntimeVersion: "86.0.616.0",
		Notes: []string{
			"Release Date: December 8, 2020",
		},
	},
	"1.0.664.37": {
		Number:         "1.0.664.37",
		Link:           "",
		RuntimeVersion: "86.0.616.0",
		Notes: []string{
			"Release Date: November 20, 2020",
		},
	},
	"1.0.674.prerelease": {
		Number:         "1.0.674.prerelease",
		Link:           "",
		RuntimeVersion: "86.0.616.0",
		Notes: []string{
			"Release Date: October 19, 2020",
		},
	},
	"1.0.622.22": {
		Number:         "1.0.622.22",
		Link:           "",
		RuntimeVersion: "86.0.616.0",
		Notes: []string{
			"Release Date: October 19, 2020",
		},
	},
	"0.9.622.11": {
		Number:         "0.9.622.11",
		Link:           "",
		RuntimeVersion: "86.0.616.0",
		Notes: []string{
			"Release Date: September 10, 2020",
		},
	},
	"0.9.579": {
		Number:         "0.9.579",
		Link:           "",
		RuntimeVersion: "86.0.579.0",
		Notes: []string{
			"Release Date: July 20, 2020",
		},
	},
	"0.9.538": {
		Number:         "0.9.538",
		Link:           "",
		RuntimeVersion: "85.0.538.0",
		Notes:          []string{},
	},
	"0.9.515.prerelease": {
		Number:         "0.9.515.prerelease",
		Link:           "",
		RuntimeVersion: "84.0.515.0",
		Notes:          []string{},
	},
	"0.9.488": {
		Number:         "0.9.488",
		Link:           "",
		RuntimeVersion: "84.0.488.0",
		Notes:          []string{},
	},
	"0.9.430": {
		Number:         "0.9.430",
		Link:           "",
		RuntimeVersion: "82.0.430.0",
		Notes:          []string{},
	},
	"0.8.355": {
		Number:         "0.8.355",
		Link:           "",
		RuntimeVersion: "80.0.355.0",
		Notes:          []string{},
	},
	"0.8.314": {
		Number:         "0.8.314",
		Link:           "",
		RuntimeVersion: "80.0.314.0",
		Notes:          []string{},
	},
	"0.8.270": {
		Number:         "0.8.270",
		Link:           "",
		RuntimeVersion: "78.0.270.0",
		Notes:          []string{},
	},
	"0.8.230": {
		Number:         "0.8.230",
		Link:           "",
		RuntimeVersion: "77.0.230.0",
		Notes:          []string{},
	},
	"0.8.190": {
		Number:         "0.8.190",
		Link:           "",
		RuntimeVersion: "77.0.190.0",
		Notes:          []string{},
	},
}
