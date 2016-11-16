var common={
	redirect:function(url) {
	    $.post(url, function (data) {
	        $("#MainPageContainer").html(data);
	    }, "text"
	    );
	    //屏蔽超级链接跳转
	    return false;
	},
	helloWorld:function() {
	    console.log("Hello World!");
	},
	dateFormat:function(data ){
		return data;
	}
}
module.exports = common
