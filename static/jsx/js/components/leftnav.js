var TopMenu = React.createClass({
    menuClick: function(e) {
        var $this = $(e.target);
    },
    render: function() {
        var submenu = < SubMenu menus = {this.props.menu.Children}/>;
	    var none=<div></div >
        return ( 
            <li className = "" >
                <a href = "#" className = {this.props.menu.Children == null ? "" : "dropdown-toggle"} >
                    <i className = {this.props.menu.IconClass} > < /i>
				    <span className="menu-text"> {this.props.menu.Name} </span >
                    <b className = {this.props.menu.Children == null ? "" : "arrow fa fa-angle-down"} > < /b>
			     </a >
                <b className = "arrow" > < /b>
			{this.props.menu.Children==null?none:submenu}
		  </li >
        )
    }
});
var MenuItem = React.createClass({
    menuClick: function(e) {
        this.gopage(this.props.menu.Action);
    },
    //发送异步请求，将结果输出到<div id="result"></div>中
    //最后一个参数可以是"html"也可以是"text"
    gopage: function(query, obj) {
        $(".nav").find("li").removeClass("active");

        if (obj != null) {
            $(obj).parent().addClass("active");
        }
        $.post(query, function(data) {
            $("#MainPageContainer").html(data);
        }, "text");
        //屏蔽超级链接跳转
        return false;
    },
    render: function() {
        return ( 
            <li className = "" >
                <a href = "javascript:void(0)" onClick = {this.menuClick} >
                    <i className = "menu-icon fa fa-caret-right" > < /i>
    				{this.props.menu.Name}
    			</a >
                < b className = "arrow" > < /b>
			</li > 
            )
    }
});
var SubMenu = React.createClass({
    render: function() {
        var rows = [];
        if (this.props.menus != undefined) {
            this.props.menus.forEach(function(item) {
                rows.push(< MenuItem menu = {item} key = {item.ID}/>);
            });           
        }
        return (  
    	  	<ul className="submenu nav-show" style={{"display":"block"}}>
    			{rows}
    	    </ul > 
        )
    }
});
var LeftNavMenu = React.createClass({
    loadCommentsFromServer: function() {
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            cache: false,
            success: function(data) {
                //console.log(data);
                this.setState({
                    data: data
                });
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    getInitialState: function() {
        return {
            data: []
        };
    },
    componentDidMount: function() {
        this.loadCommentsFromServer();
    },
    render: function() {
            var rows = [];
            if (this.state.data.Menus != undefined) {
                this.state.data.Menus.forEach(function(item) {
                    rows.push(<TopMenu menu = {item} key = {item.ID} />);
        });           
  	}
    return (
        <ul className="nav nav-list" style={{"top": "0px"}}>
			<li className="active">
				<a href="index">
					<i className="menu-icon fa fa-tachometer"></i >
                    < span className = "menu-text" > Dashboard </span>
				</a >
                <b className = "arrow" > </b>
			</li > 
            {rows} 
        </ul>
    );
  }
});
React.render(<LeftNavMenu name="leftNavMenu" url="/getuserinfo" />, document.getElementById('leftnav'));
