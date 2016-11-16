var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        var timespan=Math.round(new Date().getTime()/1000)
        url = options.url+"?"+timespan;
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, {
                'X-Xsrftoken': xsrftoken
            });
        }
        options.headers = headers;
        return ajax(url, options);
    }
});

function goBack(){
    url=window.location.hash.substring(1, window.location.hash.length); 
    redirect(url)
}
function redirect(url,refererUrl, objId) {
    if (undefined == objId) {
        objId = "#MainPageContainer";
    }
    
    if(undefined!=refererUrl)
        window.location.hash=refererUrl;
    $.get(url, function(data) {
        $(objId).html(data);
    }, "text");
}

function show_point_on_map(url) {
    window.open(url);
}

//数字格式化
//s要格式化的字符, n保留小数点位数
function fmoney(s, n) {
    n = n > 0 && n <= 20 ? n : 2;
    s = parseFloat((s + "").replace(/[^\d\.-]/g, "")).toFixed(n) + "";
    var l = s.split(".")[0].split("").reverse(),
        r = s.split(".")[1];
    t = "";
    for (i = 0; i < l.length; i++) {
        t += l[i] + ((i + 1) % 3 == 0 && (i + 1) != l.length ? "," : "");
    }
    return t.split("").reverse().join("") + "." + r;
}

//Html标签转换
function HTMLEncodeFull(strVal) {
    if (strVal != undefined && strVal != "") {
        //strVal = ReplaceAll(strVal, "&", "&amp;");
        strVal = ReplaceAll(strVal, ">", "&gt;");
        strVal = ReplaceAll(strVal, "<", "&lt;");
        strVal = ReplaceAll(strVal, "\"", "&quot;");
        strVal = ReplaceAll(strVal, "\r", "");
    }
    return strVal;
}

//文本全部替换
function ReplaceAll(str, sptr, sptr1) {
    while (str.indexOf(sptr) >= 0) {
        str = str.replace(sptr, sptr1);
    }
    return str;
}

//json日期格式转换为正常格式
function jsonDateFormat(jsonDate, formart) {
    try { //出自http://www.cnblogs.com/ahjesus 尊重作者辛苦劳动成果,转载请注明出处,谢谢!

        var date = new Date(jsonDate);

        var month = date.getMonth() + 1 < 10 ? "0" + (date.getMonth() + 1) : date.getMonth() + 1;
        var day = date.getDate() < 10 ? "0" + date.getDate() : date.getDate();
        var hours = date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
        var minutes = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();
        var seconds = date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
        var milliseconds = date.getMilliseconds();

        if (formart == "yyyy-MM-dd") {
            return date.getFullYear() + "-" + month + "-" + day;
        } else if (formart == "yyyy-MM-dd HH:mm:ss") {
            return date.getFullYear() + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds;
        } else if (formart == "yyyy年MM月dd日") {
            return date.getFullYear() + "年" + month + "月" + day + "日";
        } else if (formart == "HH:mm") {
            return hours + ":" + minutes;
        }
        return date.getFullYear() + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds + "." + milliseconds;

    } catch (ex) { //出自http://www.cnblogs.com/ahjesus 尊重作者辛苦劳动成果,转载请注明出处,谢谢!
        return "";
    }
}

function setclick(table, havesub) {
    var dTable;
    dTable = $(table).dataTable();
    $(table + ' tbody tr').click(function(e) {
        if ($(this).hasClass('row_selected')) {
            $(this).removeClass('row_selected');
        } else {
            dTable.$('tr.row_selected').removeClass('row_selected');
            $(this).addClass('row_selected');
        }
        if (havesub == 1) {
            loadPMR(this.id);
        }
    });
}

//发送异步请求，将结果输出到<div id="result"></div>中
//最后一个参数可以是"html"也可以是"text"
function gopage(query, obj) {
    $(".nav").find("li").removeClass("active");

    if (obj != null) {
        $(obj).parent().addClass("active");
    }
    $.post(query, function(data) {
        $("#MainPageContainer").html(data);
    }, "text");
    //屏蔽超级链接跳转
    return false;
}



//弹出confirm,执行操作后刷新列表
//参数说明：
//title：提示标题，
//content：提示内容，
///confirm_url：执行操作的链接，
//dataTableId：要刷新的列表的id(若不需要刷新列表，可不传此参数，或者传"")
//other_close_dialog_id：其他需要关闭的窗体
//vardata：提交的json数据
function confirm_refresh_datatable(title, content, confirm_url, dataTableId, other_close_dialog_id, vardata,callback) {
    var dialogParent = $("#confirm-dialog-message").parent();
    var dialogOwn = $("#confirm-dialog-message").clone();
    dialogOwn.hide();

    $("#confirmMsg").html(content);
    $("#confirm-dialog-message").removeClass('hide').dialog({
        resizable: false,
        modal: true,
        closeOnEscape: false,
        title: "<div class='widget-header widget-header-small'><h4 class='smaller'><i class='ace-icon fa fa-warning'></i> " + title + "</h4></div>",
        title_html: true,
        buttons: [{
            text: "确定",
            "class": "btn btn-primary",
            click: function() {
                $.ajax({
                    url: confirm_url,
                    timeout: 30000,
                    type: "PUT",
                    data: vardata,
                    dataType: 'json',
                    error: function (XMLHttpRequest, textStatus, errorThrown) {
                        alert("An error has occurred making the request: " + errorThrown)
                    },
                    success: function (data) {
                        if (data.code == "0000") {
                            if (other_close_dialog_id != undefined && other_close_dialog_id != null && other_close_dialog_id != "")
                                $(other_close_dialog_id).dialog("close");

                            if (dataTableId != undefined && dataTableId != null && dataTableId != "")
                                $(dataTableId).dataTable().fnDraw(false);
                            if (callback && typeof(callback) === "function") { //对callback做判断是否存在并且是一个函数
                                    callback();
                            }                           
                            return true;
                        } else {
                            ace_alert("错误", data.message);
                            return false;
                        }

                    }
                });
                $(this).dialog("close");
            }
        }, {
            text: "取消",
            "class": "btn",
            click: function() {
                $(this).dialog("close");
            }
        }],
        close: function() {
            dialogOwn.appendTo(dialogParent);
            $(this).dialog("destroy").remove();
        }
    });
}

//弹出提示框
//参数说明：
//title：提示框标题，
///content：提示框内容
function ace_alert(title, content) {
    var dialogParent = $("#alert-dialog-message").parent();
    var dialogOwn = $("#alert-dialog-message").clone();
    dialogOwn.hide();

    $("#alertMsg").html(content);
    $("#alert-dialog-message").removeClass('hide').dialog({
        resizable: false,
        modal: true,
        closeOnEscape: false,
        width: 400,
        title: "<div class='widget-header widget-header-small'><h4 class='smaller'><i class='ace-icon fa fa-warning'></i> " + title + "</h4></div>",
        title_html: true,
        buttons: [{
            text: "关闭",
            "class": "btn btn-minier",
            click: function() {
                $(this).dialog("close");
            }
        }],
        close: function() {
            dialogOwn.appendTo(dialogParent);
            $(this).dialog("destroy").remove();
        }
    });
}

//动态弹出dialog，dialog页面为url加载页面
//参数说明：
//第一个参数：dialog页面加载的url，第二个参数：装载dialog的div的id，第三个参数：宽度，第四个参数：高度，第五个参数：标题
function showDialog(url, dialogId, mWidth, mHeight, title, resizable) {
    if (undefined == dialogId)
        dialogId = "dialogModal";
    if (undefined == resizable)
        resizable = false;
    if (undefined == mWidth)
        mWidth = 850;
    if (undefined == mHeight)
        mHeight = 600;
    $.get(url, function(data) {
        var dialogParent = $(dialogId).parent();
        var dialogOwn = $(dialogId).clone();
        dialogOwn.hide();

        $(dialogId).html(data);
        $(dialogId).removeClass('hide').dialog({
            modal: true,
            resizable: resizable,
            closeOnEscape: false,
            width: mWidth,
            height: mHeight,
            title: "<div class='widget-header widget-header-small'><h4 class='smaller'><i class='ace-icon fa fa-check'></i> " + title + "</h4></div>",
            title_html: true,
            buttons: [{
                text: "保存并关闭",
                "class": "btn btn-primary",
                click: function() {
                    submitdata();
                }
            }, {
                text: "关闭",
                "class": "btn",
                click: function() {
                    $(this).dialog("close");
                }
            }],
            close: function() {
                dialogOwn.appendTo(dialogParent);
                $(this).dialog("destroy").remove();
            }
        });
    }, "text");
}

//动态弹出只有关闭按钮的弹出框
//参数说明：
//第一个参数：dialog页面加载的url，第二个参数：装载dialog的div的id，第三个参数：宽度，第四个参数：高度，第五个参数：标题
function open_view_dialog(url, dialogId, mWidth, mHeight, title, isResizable) {
    $.get(url, function(data) {
        if (isResizable == undefined) {
            isResizable = false;
        }

        var dialogParent = $(dialogId).parent();
        var dialogOwn = $(dialogId).clone();
        dialogOwn.hide();

        $(dialogId).html(data);
        $(dialogId).removeClass('hide').dialog({
            modal: true,
            resizable: isResizable,
            closeOnEscape: false,
            width: mWidth,
            height: mHeight,
            title: "<div class='widget-header widget-header-small'><h4 class='smaller'><i class='ace-icon fa fa-check'></i> " + title + "</h4></div>",
            title_html: true,
            buttons: [{
                text: "关闭",
                "class": "btn",
                click: function() {
                    $(this).dialog("close");
                }
            }],
            close: function() {
                dialogOwn.appendTo(dialogParent);
                $(this).dialog("destroy").remove();
            }
        });
    }, "text");
}

//动态弹出dialog，dialog页面为url加载页面, 确定时调用submitchilddata()
//参数说明：
//第一个参数：dialog页面加载的url，第二个参数：装载dialog的div的id，第三个参数：宽度，第四个参数：高度，第五个参数：标题
function open_new_children_dialog(url, dialogId, mWidth, mHeight, title, resizable) {
    if (resizable == undefined) {
        resizable = false;
    }
    var dialogParent = $(dialogId).parent();
    var dialogOwn = $(dialogId).clone();
    dialogOwn.hide();

    $.get(url, function(data) {
        $(dialogId).html(data);
        $(dialogId).removeClass('hide').dialog({
            modal: true,
            resizable: resizable,
            closeOnEscape: false,
            width: mWidth,
            height: mHeight,
            title: "<div class='widget-header widget-header-small'><h4 class='smaller'><i class='ace-icon fa fa-check'></i> " + title + "</h4></div>",
            title_html: true,
            buttons: [{
                text: "保存并关闭",
                "class": "btn btn-primary",
                click: function() {
                    submitchilddata();
                }
            }, {
                text: "关闭",
                "class": "btn",
                click: function() {
                    $(this).dialog("close");
                }
            }],
            close: function() {
                dialogOwn.appendTo(dialogParent);
                $(this).dialog("destroy").remove();
            }
        });
    }, "text");
}


//弹出dialog，dialog页面静态加载页面, 确定时调用submitchilddata()
//参数说明：
//第一个参数：dialog页面加载的url，第二个参数：装载dialog的div的id，第三个参数：宽度，第四个参数：高度，第五个参数：标题
function open_new_children_dialog_nourl(dialogId, mWidth, mHeight, title) {

    var dialogParent = $(dialogId).parent();
    var dialogOwn = $(dialogId).clone();
    dialogOwn.hide();

    $(dialogId).removeClass('hide').dialog({
        modal: true,
        resizable: false,
        closeOnEscape: false,
        width: mWidth,
        height: mHeight,
        title: "<div class='widget-header widget-header-small'><h4 class='smaller'><i class='ace-icon fa fa-check'></i> " + title + "</h4></div>",
        title_html: true,
        buttons: [{
            text: "保存并关闭",
            "class": "btn btn-primary",
            click: function() {
                submitchilddata();
            }
        }, {
            text: "关闭",
            "class": "btn",
            click: function() {
                $(this).dialog("close");
            }
        }],
        close: function() {
            dialogOwn.appendTo(dialogParent);
            $(this).dialog("destroy").remove();
        }
    });
}


//单选控件弹出框，无按钮，选择
function open_new_single_contorl_dialog(url, dialogId, mWidth, mHeight, title) {
    var dialogParent = $(dialogId).parent();
    var dialogOwn = $(dialogId).clone();
    dialogOwn.hide();

    $.get(url, function(data) {
        $(dialogId).html(data);
        $(dialogId).removeClass('hide').dialog({
            modal: true,
            resizable: false,
            closeOnEscape: false,
            width: mWidth,
            height: mHeight,
            title: "<div class='widget-header widget-header-small'><h4 class='smaller'><i class='ace-icon fa fa-check'></i> " + title + "</h4></div>",
            title_html: true,
            buttons: [],
            close: function() {
                dialogOwn.appendTo(dialogParent);
                $(this).dialog("destroy").remove();
            }
        });
    }, "text");
}

$.widget("ui.dialog", $.extend({}, $.ui.dialog.prototype, {
    _title: function(title) {
        var $title = this.options.title || '&nbsp;';
        if (("title_html" in this.options) && this.options.title_html == true)
            title.html($title);
        else title.text($title);
    }
}));
///关闭弹出dialog，并刷新列表
///参数说明：
///第一个参数：需要关闭的dialog的ID，第二个参数：需要刷新的datatable的id
function closedialog_and_refreshdatatable(dialogId, dataTableId) {
    $(dialogId).dialog('close');
    $(dataTableId).dataTable().fnDraw(false);
    $(dialogId).dialog("destroy").remove();
}

/**
 * json对象转字符串形式
 */
function json2str(o) {
    var arr = [];
    var fmt = function(s) {
        if (typeof s == 'object' && s != null) return json2str(s);
        return /^(string|number)$/.test(typeof s) ? "'" + s + "'" : s;
    }
    for (var i in o) arr.push("'" + i + "':" + fmt(o[i]));
    return '{' + arr.join(',') + '}';
}

///初始化相册
function initColorbox() {
    var $overflow = '';
    var colorbox_params = {
        rel: 'colorbox',
        reposition: true,
        scalePhotos: true,
        scrolling: false,
        previous: '<i class="ace-icon fa fa-arrow-left"></i>',
        next: '<i class="ace-icon fa fa-arrow-right"></i>',
        close: '&times;',
        current: '{current} of {total}',
        maxWidth: '100%',
        maxHeight: '100%',
        onOpen: function() {
            $overflow = document.body.style.overflow;
            document.body.style.overflow = 'hidden';
        },
        onClosed: function() {
            document.body.style.overflow = $overflow;
        },
        onComplete: function() {
            $.colorbox.resize();
        }
    };

    $('.ace-thumbnails [data-rel="colorbox"]').colorbox(colorbox_params);
    $("#cboxLoadingGraphic").html("<i class='ace-icon fa fa-spinner orange fa-spin'></i>"); //let's add a custom loading icon


    $(document).one('ajaxloadstart.page', function(e) {
        $('#colorbox, #cboxOverlay').remove();
    });
}

//messages处，如果某个控件没有message，将调用默认的信息
jQuery.extend(jQuery.validator.messages, {
    required: "必填字段",
    remote: "请修正该字段",
    email: "请输入正确格式的电子邮件",
    url: "请输入合法的网址",
    date: "请输入合法的日期",
    dateISO: "请输入合法的日期 (ISO).",
    number: "请输入合法的数字",
    digits: "只能输入整数",
    creditcard: "请输入合法的信用卡号",
    equalTo: "请再次输入相同的值",
    accept: "请输入拥有合法后缀名的字符串",
    maxlength: jQuery.validator.format("请输入一个长度最多是 {0} 的字符串"),
    minlength: jQuery.validator.format("请输入一个长度最少是 {0} 的字符串"),
    rangelength: jQuery.validator.format("请输入一个长度介于 {0} 和 {1} 之间的字符串"),
    range: jQuery.validator.format("请输入一个介于 {0} 和 {1} 之间的值"),
    max: jQuery.validator.format("请输入一个最大为 {0} 的值"),
    min: jQuery.validator.format("请输入一个最小为 {0} 的值")
});


$.ajaxSetup({

    complete: function(request, status) {

        if (typeof(request) != 'undefined') {

            var responseText = request.getResponseHeader("X-Responded-JSON");

            if (responseText != null) {

                window.tipError('系统提示', '登录超时，请重新登录', null, null, function() {

                    window.location.href = window.location.href;

                });

            }

        }

    },

    error: function(jqXHR, textStatus, errorThrown) {

        var status = 0;

        switch (jqXHR.status) {

            case (500):

                //TODO 服务器系统内部错误

                status = 500;

                break;

            case (401):
                //TODO 未登录
                window.location.href = '/admin/account/login';
                break;

            case (403):

                //TODO 无权限执行此操作

                status = 403;

                break;

            case (408):

                //TODO 请求超时

                status = 408;

                break;

            case (0):

                //TODO cancel

                break;

            default:

                status = 1;

                //TODO 未知错误

        }

        if (status > 0) {

        }

    }

});
