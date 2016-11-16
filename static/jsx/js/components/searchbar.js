var ENTER_KEY_CODE = 13;
var ReactPropTypes = React.PropTypes;
var SearchTextInput = React.createClass({

  propTypes: {
    className: ReactPropTypes.string,
    id: ReactPropTypes.string,
    placeholder: ReactPropTypes.string,
    doSearch: ReactPropTypes.func.isRequired,
    value: ReactPropTypes.string
  },

  getInitialState: function() {
    return {
      value: this.props.value || ''
    };
  },

  /**
   * @return {object}
   */
  render: function() /*object*/ {
    return (
      <input
        className="nav-search-input width-80"
        id={this.props.id}
        autoComplete="off" 
        placeholder={this.props.placeholder}
        onBlur={this._save}
        onChange={this._onChange}
        onKeyDown={this._onKeyDown}
        value={this.state.value}
        autoFocus={true}
      />
    );
  },

  /**
   * Invokes the callback passed in as onSave, allowing this component to be
   * used in different ways.
   */
  _save: function() {
    this.props.doSearch(this.state.value);
    this.setState({
      value: ''
    });
  },

  /**
   * @param {object} event
   */
  _onChange: function(/*object*/ event) {
    this.setState({
      value: event.target.value
    });
  },

  /**
   * @param {object} event
   */

  _onKeyDown: function(event) {
    if (event.keyCode === ENTER_KEY_CODE) {
      this._save();
    }
  }

});

var SearchBar = React.createClass({
    getInitialState: function() {
      return {
        collapse: this.props.collapse || false,
        msgclass: "ace-icon fa fa-chevron-down"
      };
    },
    _doSearch: function(text) {
        console.log(text)
        //TodoActions.create(text);
    },
    _onCollapse:function(){
    if(this.state.collapse)
    {

        this.setState({
          msgclass: "ace-icon fa fa-chevron-up"
        });
    }
    else
    {
        this.setState({
          msgclass: "ace-icon fa fa-chevron-down"
        });
        //$("#more_filter").addClass("hide");
    }
                    
      this.setState({
        collapse: !this.state.collapse
      });
    },
    render: function() {
        return (
               <div className="nav-search width-40" id={this.props.name}>
                    <span className="input-icon width-100">
                        <SearchTextInput id="searchBarInput" doSearch={this._doSearch}  placeholder={this.props.placeholder} />
                        <i className="ace-icon nav-search-icon  fa fa-search "></i>
                        <span className="nav-search-btn">
                            <div className="widget-toolbar">
                                <a href="#" data-action="collapse" onClick={this._onCollapse}   >
                                    <i className={this.state.msgclass} ></i>
                                    
                                </a>
                            </div>
                        </span>
                    </span>
                </div>
        );
    }
});
