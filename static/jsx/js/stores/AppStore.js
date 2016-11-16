/*
 * Copyright (c) 2014, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 *
 * AppStore
 */

var Reflux = require('reflux');
var EventEmitter = require('events').EventEmitter;
var assign = require('object-assign');
var AppActions = require('../actions/AppActions');
var common = require('../utils/common')
var todoCounter = 0,
        localStorageKey = "todos";


var AppStore = Reflux.createStore({
        // this will set up listeners to all publishers in TodoActions, using onKeyname (or keyname) as callbacks
        listenables: [AppActions],
        onRedirect: function(data) {
    		common.redirect(data.Url);
            items=[];
            var item=data;
            while(item !=undefined){
                items.push({id:item.ID,url:"",text:item.Name});
                var item=item.Parent;
            }
            bcData=[];
            for(var i=items.length-1;i>=0;i--){
                bcData.push(items[i]);
            }
            this.updateState({breadCrumb:bcData});
	    },
        updateState: function(appState){
            appState= assign({}, this.currentState, appState); 
            //localStorage.setItem(localStorageKey, JSON.stringify(appState));
            // if we used a real database, we would likely do the below in a callback
            this.currentState = appState;
            this.trigger(appState); // sends the updated list to all listening components (TodoApp)
        },            
        // this will be called by all listening components as they register their listeners
        getInitialState: function() {
        var currentStorage = localStorage.getItem(localStorageKey);
         if (!currentStorage) {
                // If no list is in localstorage, start out with a default one
                this.currentState = {
                    key: todoCounter++,
                    created: new Date(),
                    isComplete: false,
                    label: 'Rule the web',
                    userInfo:{userName:""},
                    breadCrumb:[{id:1,url:"",text:"Dashboard"}]
                };
            } else {
                this.currentState = JSON.parse(currentStorage)
            }
            return this.currentState;
        }
        
    });
module.exports = AppStore;
