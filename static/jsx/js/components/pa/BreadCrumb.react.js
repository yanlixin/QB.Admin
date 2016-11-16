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
var Reflux = require('reflux');
var AppStore = require('../../stores/AppStore');
var AppActions = require('../../actions/AppActions');
/**
 * Retrieve the current BreadCrumb data from the AppStroe
 */


var BreadCrumb = React.createClass({
  	mixins: [Reflux.connect(AppStore,"currentState")],
        handleEditStart: function(evt) {
		   evt.preventDefault();
	       AppActions.redirect( "redirect");
	  },
  /**
   * @return {object}
   */
  render: function() {
    var items=[];
	  for (var i=0;i<this.state.currentState.breadCrumb.length;i++){
        var item=this.state.currentState.breadCrumb[i];
        if(0==i)
        {
          items.push(  
            <li>
              <i className="ace-icon fa fa-home home-icon"></i>
              <a href="/home">{item.text}</a>
            </li>
          ) ;
        } else
        {
          items.push(  
            <li><a href="/home">{item.text}</a>
            </li>
          ) ;
        }  
      };
 
    return (
      <ul className="breadcrumb">
        {items}
      </ul>
    );
  },

});

module.exports = BreadCrumb;
