var React = require('react/addons');
var Reflux = require('reflux');
var AppStore = require('../../stores/AppStore');
var AppActions = require('../../actions/AppActions');
/**
 * Retrieve the current SelectInteraction data from the AppStroe
 */
var UserInfo=React.createClass({
    mixins: [Reflux.connect(AppStore,"currentState")],
        handleEditStart: function(evt) {
           evt.preventDefault();
           AppActions.redirect( "redirect");
      },
    handleProfile:function(e){
          $(".nav").find("li").removeClass("active");
          var obj=undefined; 
          if (obj != undefined) {
              $(obj).parent().addClass("active");
          }
          var bc={ID:0,Url:"/user/profile",Name:"个人信息",Parent:undefined};
          AppActions.redirect( bc);
      },    
    render: function() {
        
        return (
                <li className="light-blue">
                    <a data-toggle="dropdown" href="#" className="dropdown-toggle">
                                <img className="nav-user-photo" src="<%.cdn%>/avatars/user.jpg" alt="Jason's Photo" />
                                <span className="user-info">
                                    <small>Welcome,</small>{this.state.currentState.userInfo.userName}
                                </span>

                                <i className="ace-icon fa fa-caret-down"></i>
                            </a>

                    <ul className="user-menu dropdown-menu-right dropdown-menu dropdown-yellow dropdown-caret dropdown-close">
                        <li>
                            <a href="javascript:void(0)" onClick={this.handleProfile} >
                                <i className="ace-icon fa fa-user"></i>
                                Profile
                            </a>
                        </li>

                        <li className="divider"></li>

                        <li>
                            <a href="/logout">
                                <i className="ace-icon fa fa-power-off"></i>
                                Logout
                            </a>
                        </li>
                    </ul>
                </li>
        );
    }
});
var TopMenuGrey=React.createClass({
    render: function() {
        return (
                <li className="grey">
                    <a data-toggle="dropdown" className="dropdown-toggle" href="#">
                        <i className="ace-icon fa fa-tasks"></i>
                        <span className="badge badge-grey">4</span>
                    </a>

                    <ul className="dropdown-menu-right dropdown-navbar dropdown-menu dropdown-caret dropdown-close">
                        <li className="dropdown-header">
                            <i className="ace-icon fa fa-check"></i>
                            4 Tasks to complete
                        </li>

                        <li className="dropdown-content">
                            <ul className="dropdown-menu dropdown-navbar">
                                <li>
                                    <a href="#">
                                        <div className="clearfix">
                                            <span className="pull-left">Software Update</span>
                                            <span className="pull-right">65%</span>
                                        </div>

                                        <div className="progress progress-mini">
                                            <div style={{"width":"65%"}} className="progress-bar"></div>
                                        </div>
                                    </a>
                                </li>

                                <li>
                                    <a href="#">
                                        <div className="clearfix">
                                            <span className="pull-left">Hardware Upgrade</span>
                                            <span className="pull-right">35%</span>
                                        </div>

                                        <div className="progress progress-mini">
                                            <div style={{"width":"35%"}} className="progress-bar progress-bar-danger"></div>
                                        </div>
                                    </a>
                                </li>

                                <li>
                                    <a href="#">
                                        <div className="clearfix">
                                            <span className="pull-left">Unit Testing</span>
                                            <span className="pull-right">15%</span>
                                        </div>

                                        <div className="progress progress-mini">
                                            <div style={{"width":"15%"}} className="progress-bar progress-bar-warning"></div>
                                        </div>
                                    </a>
                                </li>

                                <li>
                                    <a href="#">
                                        <div className="clearfix">
                                            <span className="pull-left">Bug Fixes</span>
                                            <span className="pull-right">90%</span>
                                        </div>

                                        <div className="progress progress-mini progress-striped active">
                                            <div style={{"width":"90%"}} className="progress-bar progress-bar-success"></div>
                                        </div>
                                    </a>
                                </li>
                            </ul>
                        </li>

                        <li className="dropdown-footer">
                            <a href="#">
                                        See tasks with details
                                        <i className="ace-icon fa fa-arrow-right"></i>
                                    </a>
                        </li>
                    </ul>
                </li>
        );
    }
});
var TopMenuPurple=React.createClass({
    render: function() {
        return (
                <li className="purple">
                    <a data-toggle="dropdown" className="dropdown-toggle" href="#">
                        <i className="ace-icon fa fa-bell icon-animated-bell"></i>
                        <span className="badge badge-important">8</span>
                    </a>

                    <ul className="dropdown-menu-right dropdown-navbar navbar-pink dropdown-menu dropdown-caret dropdown-close">
                        <li className="dropdown-header">
                            <i className="ace-icon fa fa-exclamation-triangle"></i>
                            8 Notifications
                        </li>

                        <li className="dropdown-content">
                            <ul className="dropdown-menu dropdown-navbar navbar-pink">
                                <li>
                                    <a href="#">
                                        <div className="clearfix">
                                            <span className="pull-left">
                                                        <i className="btn btn-xs no-hover btn-pink fa fa-comment"></i>
                                                        New Comments
                                                    </span>
                                            <span className="pull-right badge badge-info">+12</span>
                                        </div>
                                    </a>
                                </li>

                                <li>
                                    <a href="#">
                                        <i className="btn btn-xs btn-primary fa fa-user"></i>
                                        Bob just signed up as an editor ...
                                    </a>
                                </li>

                                <li>
                                    <a href="#">
                                        <div className="clearfix">
                                            <span className="pull-left">
                                                        <i className="btn btn-xs no-hover btn-success fa fa-shopping-cart"></i>
                                                        New Orders
                                                    </span>
                                            <span className="pull-right badge badge-success">+8</span>
                                        </div>
                                    </a>
                                </li>

                                <li>
                                    <a href="#">
                                        <div className="clearfix">
                                            <span className="pull-left">
                                                        <i className="btn btn-xs no-hover btn-info fa fa-twitter"></i>
                                                        Followers
                                                    </span>
                                            <span className="pull-right badge badge-info">+11</span>
                                        </div>
                                    </a>
                                </li>
                            </ul>
                        </li>

                        <li className="dropdown-footer">
                            <a href="#">
                                        See all notifications
                                        <i className="ace-icon fa fa-arrow-right"></i>
                                    </a>
                        </li>
                    </ul>
                </li>
        );
    }
});
var TopMenuGreen=React.createClass({
    render: function() {
        return (
                <li className="green">
                    <a data-toggle="dropdown" className="dropdown-toggle" href="#">
                        <i className="ace-icon fa fa-envelope icon-animated-vertical"></i>
                        <span className="badge badge-success">5</span>
                    </a>

                    <ul className="dropdown-menu-right dropdown-navbar dropdown-menu dropdown-caret dropdown-close">
                        <li className="dropdown-header">
                            <i className="ace-icon fa fa-envelope-o"></i>
                            5 Messages
                        </li>

                        <li className="dropdown-content">
                            <ul className="dropdown-menu dropdown-navbar">
                                <li>
                                    <a href="#" className="clearfix">
                                                <img src="<%.cdn%>/avatars/avatar.png" className="msg-photo" alt="Alex's Avatar" />
                                                <span className="msg-body">
                                                    <span className="msg-title">
                                                        <span className="blue">Alex:</span>
                                                        Ciao sociis natoque penatibus et auctor ...
                                                    </span>

                                                    <span className="msg-time">
                                                        <i className="ace-icon fa fa-clock-o"></i>
                                                        <span>a moment ago</span>
                                                    </span>
                                                </span>
                                            </a>
                                </li>

                                <li>
                                    <a href="#" className="clearfix">
                                                <img src="<%.cdn%>/avatars/avatar3.png" className="msg-photo" alt="Susan's Avatar" />
                                                <span className="msg-body">
                                                    <span className="msg-title">
                                                        <span className="blue">Susan:</span>
                                                        Vestibulum id ligula porta felis euismod ...
                                                    </span>

                                                    <span className="msg-time">
                                                        <i className="ace-icon fa fa-clock-o"></i>
                                                        <span>20 minutes ago</span>
                                                    </span>
                                                </span>
                                            </a>
                                </li>

                                <li>
                                    <a href="#" className="clearfix">
                                                <img src="<%.cdn%>/avatars/avatar4.png" className="msg-photo" alt="Bob's Avatar" />
                                                <span className="msg-body">
                                                    <span className="msg-title">
                                                        <span className="blue">Bob:</span>
                                                        Nullam quis risus eget urna mollis ornare ...
                                                    </span>

                                                    <span className="msg-time">
                                                        <i className="ace-icon fa fa-clock-o"></i>
                                                        <span>3:15 pm</span>
                                                    </span>
                                                </span>
                                            </a>
                                </li>

                                <li>
                                    <a href="#" className="clearfix">
                                                <img src="<%.cdn%>/avatars/avatar2.png" className="msg-photo" alt="Kate's Avatar" />
                                                <span className="msg-body">
                                                    <span className="msg-title">
                                                        <span className="blue">Kate:</span>
                                                        Ciao sociis natoque eget urna mollis ornare ...
                                                    </span>

                                                    <span className="msg-time">
                                                        <i className="ace-icon fa fa-clock-o"></i>
                                                        <span>1:33 pm</span>
                                                    </span>
                                                </span>
                                            </a>
                                </li>

                                <li>
                                    <a href="#" className="clearfix">
                                                <img src="<%.cdn%>/avatars/avatar5.png" className="msg-photo" alt="Fred's Avatar" />
                                                <span className="msg-body">
                                                    <span className="msg-title">
                                                        <span className="blue">Fred:</span>
                                                        Vestibulum id penatibus et auctor  ...
                                                    </span>

                                                    <span className="msg-time">
                                                        <i className="ace-icon fa fa-clock-o"></i>
                                                        <span>10:09 am</span>
                                                    </span>
                                                </span>
                                            </a>
                                </li>
                            </ul>
                        </li>

                        <li className="dropdown-footer">
                            <a href="inbox.html">
                                        See all messages
                                        <i className="ace-icon fa fa-arrow-right"></i>
                                    </a>
                        </li>
                    </ul>
                </li>
        );
    }
});
var TopMenu=React.createClass({
    render: function() {
        return (
            <ul className="nav ace-nav">
                <TopMenuGrey></TopMenuGrey>
                <TopMenuPurple></TopMenuPurple>
                <TopMenuGreen></TopMenuGreen>
                <UserInfo></UserInfo>
            </ul>
    );
}
}); 
module.exports = TopMenu;

          