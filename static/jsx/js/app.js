var React = require('react/addons');
var ReactRouter = require('react-router');
var Reflux = require('reflux');
var AppStore = require('./stores/AppStore');
var AppActions = require('./actions/AppActions');
var BreadCrumb = require('./components/pa/BreadCrumb.react');
var LeftNavMenu = require('./components/pa/LeftMenu.react.js');
var TopMenu = require('./components/pa/TopMenu.react.js');
var SelectInteraction = require('./components/pa/SelectInteraction.react.js');

React.render(<TopMenu /> , document.getElementById('topMenu'));
React.render(<BreadCrumb id = "breadCrumb" /> , document.getElementById('breadcrumbs'));
React.render(<LeftNavMenu name = "leftNavMenu" url = "/getuserinfo" /> , document.getElementById('leftnav'));