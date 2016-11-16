/**
 * Copyright (c) 2015-2016 David Yan
 * Licensed under the MIT license
 */
(function( factory ) {
	if ( typeof define === "function" && define.amd ) {

		// AMD. Register as an anonymous module.
		define([ "jquery" ], factory );
	} else {

		// Browser globals
		factory( jQuery );
	}
}(function( $ ) {


    var easydialog = $.widget( "ui.easydialog", {
    	version: "1.11.2",
    	options: {
    	    dialogId    :   'dialogModal',
    	    width       :   850,
    	    height      :   600,
    	    resizable   :   false,
            title		:	'',
			modal		:	true,
            btns        :   []
            },
        showDialog:function(){
			_showDialog(this.options);
            
        }
    })
	
	function _showDialog(opts) {
        opts.dialogId="#"+opts.dialogId;

        var dialogParent = $(opts.dialogId).parent();
        var dialogOwn = $(opts.dialogId).clone();
        dialogOwn.hide();

        var btn = [];

		if(undefined!=opts.btns){
			$.each(opts.btns,function(i,item){btn.push(item);});
		}
		var btnClose= {text: "关闭","class": "btn",click: function () {$(this).dialog("close");}}
		btn.push(btnClose);
        $.get(opts.url, function (data) {
            $(opts.dialogId).html(data);
            $(opts.dialogId).removeClass('hide').dialog({
                modal: opts.modal,
                resizable: opts.resizable,
                closeOnEscape: false,
                width: opts.width,
                height: opts.height,
                title: "<div class='widget-header widget-header-small'><h4 class='smaller'><i class='ace-icon fa fa-check'></i> " + opts.title + "</h4></div>",
                title_html: true,
                buttons: btn,
                open: function (event, ui) {
                   
                },
                close: function () {
 			dialogOwn.appendTo(dialogParent);
                   	$(this).dialog("destroy").remove();
                }
            });
        }, "text");
    }
	
}))
