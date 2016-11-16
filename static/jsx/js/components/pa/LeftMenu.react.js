/**
 * Copyright (c) 2015, David Yan, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

/**
 * This component operates as a "Controller-View".  It listens for changes in
 * the AppStroe and passes the new data to its children.
 */
var React = require('react/addons');
var ReactRouter = require('react-router');
var Reflux = require('reflux');
var AppStore = require('../../stores/AppStore');
var AppActions = require('../../actions/AppActions');

var MenuItem = React.createClass({
  mixins: [Reflux.connect(AppStore,"currentState")],
  menuClick:function(e){
      $(".nav").find("li").removeClass("active");
      var obj=undefined; 
      if (obj != undefined) {
          $(obj).parent().addClass("active");
      }
      var bc={ID:this.props.ID,Url:this.props.menu.Action,Name:this.props.menu.Name,Parent:this.props.menu.Parent};
      AppActions.redirect( bc);
  },   
  render:function(){
  		return(
  			<li className="" >
  				<a href="javascript:void(0)" onClick={this.menuClick}>
  					<i className="menu-icon fa fa-caret-right"></i>
  					{this.props.menu.Name}
  				</a>
  				<b className="arrow"></b>
  			</li>)
  	}
});
var SubMenu = React.createClass({
   render: function(){ 
      var rows = [];
      if(this.props.menus !=undefined){
             this.props.menus.forEach(function(item) {
                rows.push(<MenuItem menu={item} key={item.ID} />);
            });           
      }
      return (  
	  	<ul className="submenu nav-show" style={{"display":"none"}}>
			{rows}
	    </ul>)
    }
});
var TopMenu = React.createClass({
    menuClick: function(e){
      var $this = $(e.target);
    },
    render: function() {
  var submenu= <SubMenu menus={this.props.menu.Children} />;
  var none=<div></div>
  return (
      <li className="">
      <a href="#" className="dropdown-toggle" >
        <i className={this.props.menu.IconClass}></i>
        <span className="menu-text"> {this.props.menu.Name} </span>
        <b className={this.props.menu.Children==null?"":"arrow fa fa-angle-down"}></b>
      </a>
      <b className="arrow"></b>
      {this.props.menu.Children==null?none:submenu}
      </li>  
      )
    }
});
var LeftMenu = React.createClass({
    mixins: [Reflux.connect(AppStore,"currentState")],
    loadCommentsFromServer: function() {
    $.ajax({
      url: this.props.url,
      dataType: 'json',
      cache: false,
      success: function(data) {
        AppActions.updateState({userInfo:{userName:data.User.UserName}});
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },
  getInitialState: function() {
    return {data: []};
  },
  componentDidMount: function() {
    this.loadCommentsFromServer();
  },
  render: function() {
    var rows = [];
  	if(this.state.data.Menus !=undefined){
          	this.state.data.Menus.forEach(function(item) {
             		rows.push(<TopMenu menu={item} key={item.ID} />);
          	});           
  	}
    return (
        <ul className="nav nav-list" style={{"top": "0px"}}>
  			  <li className="active">
  				<a href="index">
  					<i className="menu-icon fa fa-tachometer"></i>
  					<span className="menu-text"> Dashboard </span>
  				</a>
  				<b className="arrow"></b>
  			  </li>
          {rows}
        </ul>
    );
  }
});

module.exports = LeftMenu;
