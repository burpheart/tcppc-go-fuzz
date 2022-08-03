package tcppc

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

const data1 = `HTTP/1.1 200 OK
content-type: text/html; charset=utf-8
server: Jexus/5.8.1.3 Linux
p3p: CP="CAO PSA OUR"
cache-control: private
set-cookie: ASP.NET_SessionId=123456789; path=/`

const data2 = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<head>
<meta charset="UTF-8" />
<meta name="renderer" content="webkit" />
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
<meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
<meta name="format-detection" content="telephone=no,email=no" />
<meta name="apple-mobile-web-app-capable" content="yes" />
<meta name="apple-mobile-web-app-status-bar-style" content="black" />
<meta name="author" content="Tencent-CDC" />
<meta name="copyright" content="Tencent" />
<meta name="robots" content="index,follow" />
<title>Web统一登录</title>
<link rel="stylesheet" type="text/css" href="../Styles/login.16bd.css" />

<link rel="stylesheet" type="text/css" href="../Styles/normalize.css" />
<script type="text/javascript" src="../Scripts/jquery-1.9.0.min.js"></script>
<script type="text/javascript" src="../Scripts/jquery.cookie.js"></script>
<script type="text/javascript" src="../Scripts/login.js"></script>
</head>

<body id="thebody" >
<div class="login_wr zoom">
<div class="login_con">
<h1 class="logo">
<img class="logo_img" src="../Images/logo@1x.d311.png" srcset="../Images/logo@2x.9cab.png 2x" alt="腾讯">
<span class="sub_title">企业应用</span>
</h1>
<div class="login_normal">
<form method="post" action="/Modules/SignIn.aspx?url=https%3a%2f%2fhrc.tencent.com%2f_login%2f%3furl%3d%252F%26app%3d&amp;appkey=a6c2e79693674c2e8180979f41d1dad7" onsubmit="javascript:return window.WebForm_OnSubmit();" id="form1">
<script type="text/javascript">//<![CDATA[
var theForm;
if (document.getElementById) { theForm = document.getElementById ('form1'); }
else { theForm = document.form1; }
//]]></script>

<div class="aspNetHidden">


<input type="hidden" name="__SCROLLPOS" id="__SCROLLPOS" value="0" /><input type="hidden" name="__COMPRESSED_VIEWSTATE" id="__COMPRESSED_VIEWSTATE" value="QlpoMTFBWSZTWbYCXBoAAJV/+//v6ABECj/QBCFVAD5v/6AgjAEAQBCgAAEAAeAAALAA2IKp6PSTNRoMhoMjI0DTQB6R6m0gND0hzAmJoMJkyZMjCYJppkYmAIYBKmmiaI9IT1DAE0ZGmg0YEZANDQTNlFOpBz2pAmIjmFItQVwJLWYRjgRz9KGvhP6eSHlXojPlXT51CYlAgSC8AyKDRSog5QgTtFRwxDiewbBr0WvZfcuPkCYEY17AlMa5KMXP04saufR3ZEgRYYl9I1AIAyKuxTRwKwvanRG1ohODTcTUnIRow4NF3soOAilO/IOnAg4OyuCG/jZExzIeqHWW1M2gza4qjwqyu4hRjIpdmXWC2wPz79Jn/JoD6hM7iTUfxdyRThQkLYCXBoA=" /><input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="" />
</div>

<script src="/WebResource.axd?d=QNuAqBjD%2bG53zN9aD85xJKiBmozdxfjFxpyPCSOiJKE%3d_F6U03CPPBZWsndHYvZnjEgQUgF5MEW%2fhWczbUre%2fpGo%3d_f&t=637902724330000000" type="text/javascript"></script>
<script type="text/javascript">//<![CDATA[
WebFormValidation_Initialize(window);
window.WebForm_OnSubmit = function () {
if (!window.ValidatorOnSubmit()) return false;
return true;
}
//]]></script>


<div class="ipt_item ipt_name ipt_item_focus">
<input type="text" maxlength="128" name="txtLoginName" id="txtLoginName" tabindex="1" class="ipt" value="" name="account" placeholder="英文ID" />
<span id="LoginNameRequiredFieldValidator" style="visibility:hidden;display:inline-block;color:Red;height:0px;">&nbsp</span>
<i class="icon_account"></i>
</div>
<div class="ipt_item ipt_password ipt_item_focus">
<input type="text" value="1" name="txtPasswordType" id="txtPasswordType" style="display: none" />
<input type="password" maxlength="32" name="txtPassword" id="txtPassword" tabindex="2" class="ipt" value="" name="password" placeholder="PIN+Token" />
<span id="PasswordRequiredFieldValidator" style="visibility:hidden;display:inline-block;color:Red;height:0px;">&nbsp</span>
<i class="icon_password"></i>
</div>

<div class="btn_wrap">
<input type="submit" name="ibnLogin" value="登  录" onclick="setCookieLoginName();WebForm_DoPostback(&quot;ibnLogin&quot;,&quot;&quot;,null,false,true,false,false,&quot;&quot;)" id="ibnLogin" class="btn_login" />
<div id="ValidationSummary1" style="color:Red;display:none;">

</div>
</div>
<div class="account_warning">
<!--<!--<i class="icon_warning">-->


</div>

<script type="text/javascript">//<![CDATA[
var LoginNameRequiredFieldValidator = document.all ? document.all ["LoginNameRequiredFieldValidator"] : document.getElementById ("LoginNameRequiredFieldValidator");
LoginNameRequiredFieldValidator.evaluationfunction = "RequiredFieldValidatorEvaluateIsValid";
LoginNameRequiredFieldValidator.initialvalue = "";
LoginNameRequiredFieldValidator.controltovalidate = "txtLoginName";
LoginNameRequiredFieldValidator.errormessage = "请输入用户名!";
var ValidationSummary1 = document.all ? document.all ["ValidationSummary1"] : document.getElementById ("ValidationSummary1");
ValidationSummary1.showmessagebox = "True";
ValidationSummary1.showsummary = "False";
var PasswordRequiredFieldValidator = document.all ? document.all ["PasswordRequiredFieldValidator"] : document.getElementById ("PasswordRequiredFieldValidator");
PasswordRequiredFieldValidator.evaluationfunction = "RequiredFieldValidatorEvaluateIsValid";
PasswordRequiredFieldValidator.initialvalue = "";
PasswordRequiredFieldValidator.controltovalidate = "txtPassword";
PasswordRequiredFieldValidator.errormessage = "请输入密码!";
//]]></script>


<div class="aspNetHidden">



<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEdAAEAAAD/////AQAAAAAAAAAPAQAAAAQAAAAIrHy4eLv3ptWDlukUTsyIawsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJw/SLKWrdR0XparDtO2eZGXxpXAhp/NwnpO0c8BMkJF" /><input type="hidden" name="__PREVIOUSPAGE" id="__PREVIOUSPAGE" value="/Modules/SignIn.aspx" /><input type="hidden" name="__EVENTARGUMENT" id="__EVENTARGUMENT" value="" /><input type="hidden" name="__EVENTTARGET" id="__EVENTTARGET" value="" />
</div>

<script type="text/javascript">//<![CDATA[
window._form = theForm;
window.__doPostBack = function (eventTarget, eventArgument) {
if(theForm.onsubmit && theForm.onsubmit() == false) return;
theForm.__EVENTTARGET.value = eventTarget;
theForm.__EVENTARGUMENT.value = eventArgument;
theForm.submit();
}
//]]></script>

<script src="/WebResource.axd?d=QNuAqBjD%2bG53zN9aD85xJKiBmozdxfjFxpyPCSOiJKE%3d_MH19Ul1tHTHq6yNxXltGK4FStaY6nLhiMrWTaLf9SfY%3d_f&t=637902724330000000" type="text/javascript"></script>
<script type="text/javascript">//<![CDATA[
WebForm_Initialize(window);
//]]></script>

<script type="text/javascript">//<![CDATA[
var Page_Validators =  new Array(document.getElementById ('LoginNameRequiredFieldValidator'), document.getElementById ('PasswordRequiredFieldValidator'));
var Page_ValidationSummaries =  new Array(document.getElementById ('ValidationSummary1'));
//]]></script>


<script language='javascript'>function saveScrollPosition() {document.forms[0].__SCROLLPOS.value = thebody.scrollTop;}thebody.onscroll=saveScrollPosition;</script>
<script type="text/javascript">//<![CDATA[

window.Page_ValidationActive = false;
window.ValidatorOnLoad();
window.ValidatorOnSubmit = function () {
if (this.Page_ValidationActive) {
return this.ValidatorCommonOnSubmit();
}
return true;
};
//]]></script>
</form>
</div>
</div>

<div class="login_footer">
<p>Power By IT Login</p>
<p>Copyright © 1998 - 2016 Tencent. All Rights Reserved.</p>
<p>腾讯公司 版权所有</p>
</div>
</div>

<script type="text/javascript">

$(document).ready(initialize);

//刷新验证码
function reflushCaptcha() {
$("#imgCaptcha").attr("src", "https://idcw.rio.tencent.com/tof4/captcha/getimage?aid=1600000411" + "&" + Math.random());
}
</script>
</body>
</html>
<!-
root
uid=0(root) gid=0(root) groups=0(root)

root:x:0:0:root:/root:/bin/bash
bin:x:1:1:bin:/bin:/sbin/nologin
daemon:x:2:2:daemon:/sbin:/sbin/nologin
adm:x:3:4:adm:/var/adm:/sbin/nologin
lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin
sync:x:5:0:sync:/sbin:/bin/sync
shutdown:x:6:0:shutdown:/sbin:/sbin/shutdown
halt:x:7:0:halt:/sbin:/sbin/halt
mail:x:8:12:mail:/var/spool/mail:/sbin/nologin
operator:x:11:0:operator:/root:/sbin/nologin
games:x:12:100:games:/usr/games:/sbin/nologin
ftp:x:14:50:FTP User:/var/ftp:/sbin/nologin
nobody:x:99:99:Nobody:/:/sbin/nologin

success
SUCCCESS
alibaba-inc
alibaba-inc.com
tencent
tencent.com
MikroTik
RouterOS
->`

type NewConn struct {
	net.Conn
	r io.Reader
}

func (c NewConn) Read(p []byte) (int, error) {
	return c.r.Read(p)
}

func HandleTLSSession(conn *tls.Conn, writer *RotWriter, timeout int) {
	defer conn.Close()
	defer counter.dec()
	counter.inc()

	var src, dst *net.TCPAddr
	src = conn.RemoteAddr().(*net.TCPAddr)
	dst = conn.LocalAddr().(*net.TCPAddr)

	flow := NewTLSFlow(src, dst)
	session := NewSession(flow)

	log.Printf("TLS: Established: %s (#Sessions: %d)\n", session, counter.count())

	var length uint
	var err error

	buf := make([]byte, 4096)

	for {
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))

		length, err := conn.Read(buf)
		if err != nil {
			break
		}

		data := make([]byte, length)
		copy(data, buf[:length])
		if length > 10 {
			if string(data[0:5]) == "POST " || string(data[0:4]) == "GET " || string(data[0:5]) == "HEAD " || string(data[0:8]) == "OPTIONS " || string(data[0:7]) == "DELETE " || string(data[0:4]) == "PUT " || string(data[0:6]) == "TRACE " || string(data[0:8]) == "CONNECT " {
				conn.Write([]byte(data1))
				a := session.String()
				conn.Write([]byte("\n" + a))

				conn.Write([]byte("\ncontent-length: " + strconv.Itoa(len(data2))))
				conn.Write([]byte("\n\n"))
				conn.Write([]byte(data2))
				log.Printf("TLS: Send: %s\n", session)
			} else {
				//log.Printf("TLS: noSend: %s : %q\n", session, data[0:5])
			}

		} else {
			//log.Printf("TLS: short: %s : %q\n", session, data)
		}

		session.AddPayload(data)

		log.Printf("TLS: Received: %s: %q (%d bytes)\n", session, buf[:length], length)
	}

	if writer != nil {
		outputJson, err := json.Marshal(session)
		if err == nil {
			log.Printf("Wrote data: %s\n", session)
			writer.Write(outputJson)
		} else {
			log.Printf("Failed to encode data as json: %s\n", err)
		}
	}

	if length == 0 {
		log.Printf("Closed: %s (#Sessions: %d)\n", session, counter.count())
	} else {
		log.Printf("Aborted: %s %s (#Sessions: %d)\n", session, err, counter.count())
	}
}
