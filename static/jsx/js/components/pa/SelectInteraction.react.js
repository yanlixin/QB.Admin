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
 * Retrieve the current SelectInteraction data from the AppStroe
 */

var SecondSelect=React.createClass({

   render: function() {
    
    var items = $.map(this.props.selectData,function (child) {
      return (
        <option value={child.ID}>{child.Name}</option>
      );
    });
    return (
          <select className="input-medium"  name={this.props.secName} id={this.props.secId} >
          <option value=" ">--select--</option>
          {items}
          </select>
    );
  }
});

var SelectInteraction = React.createClass({
  loadCommentsFromServer: function() {
    $.ajax({
      url: this.props.firUrl,
      dataType: 'json',
      cache: false,
      success: function(jsonData) {
        this.setState({data: jsonData.aaData});
      }.bind(this),
      error: function(xhr, status, err) {
        
      }.bind(this)
    });
  },
  getInitialState: function() {
    return {data: [],secData:[]};
  },
  componentDidMount: function() {
    this.loadCommentsFromServer();
  },
  handleUserChange: function(evt){
    var selectId=$(evt.target).val();
    $.ajax({
      url: this.props.secUrl+"/"+selectId,
      dataType: 'json',
      cache: false,
      success: function(jsonData) {
        this.setState({secData: jsonData.aaData});
      }.bind(this),
      error: function(xhr, status, err) {
        
      }.bind(this)
    });
  },
  /**
   * @return {object}
   */
  render: function() {
    var items = $.map(this.state.data,function (child) {
      return (
        <option value={child.ID}>{child.Name}</option>
      );
    });
    return (
      <div className="input-group">
          <select className="input-medium"  onChange={this.handleUserChange} >
          <option value=" ">--select--</option>
            {items}
          </select><br/>
          <SecondSelect selectData={this.state.secData} id={this.props.secId} name={this.props.secName} />
      </div>
    );
  },

});

module.exports = SelectInteraction;
