var CheckBox=React.createClass({
 getInitialState: function() {
    return {value: ''};
  },
  handleChange: function(event) {
    this.setState({value: event.target.value});
  },
 
    render: function() {
        var value = this.props.data.value;
        return (
    			<label>
    				<input name="agree" name={this.props.data.id} id={this.props.data.id} type="checkbox" value={value} className="ace" />
    				<span class="lbl"> {this.props.data.title}</span>
    			</label>
        )
    }
});
var DateBox=React.createClass({
 getInitialState: function() {
    return {value: ''};
  },
  handleChange: function(event) {
    this.setState({value: event.target.value});
  },
 
    render: function() {
        var value = this.props.data.value;
        return (
                <div className="input-group">
                    <input className="input-medium date-picker"  type="text"  name={this.props.data.id} id={this.props.data.id} value={value} onChange={this.handleChange}
                     data-date-format="yyyy-mm-dd" placeholder="yyyy-mm-dd"/>
                    <span className="input-group-addon">
                        <i className="ace-icon fa fa-calendar"></i>
                    </span>
                </div>
        )
    }
});

var PasswordBox=React.createClass({
  getInitialState: function() {
    return {value: ''};
  },
  handleChange: function(event) {
    this.setState({value: event.target.value});
  },
  
    render: function() {
        var value = this.props.data.value;
        return (
             <div className="clearfix">
                <input type="password" name={this.props.data.id} id={this.props.data.id} value={value} onChange={this.handleChange} />
             </div>
        )
    }
});

var TextBox=React.createClass({
getInitialState: function() {
    return {value: ''};
  },
  handleChange: function(event) {
    this.setState({value: event.target.value});
  },
  
    render: function() {
        var value = this.props.data.value;
        return (
            <div className="clearfix">
               <input type="text" name={this.props.data.id} id={this.props.data.id} value={value} onChange={this.handleChange} />
            </div>
        )
   } 
});


var DropDownList=React.createClass({
    render: function() {
        var options=[];
               this.props.options.forEach(function(item) {
            	options.push(<option value={item.value}>{item.text}</option>);
        });
 
        return (
            <div className="clearfix">
                <select className="input-medium" name={this.props.data.id} id={this.props.data.id}>
                    {options}
    			</select>
            </div>
        )
    }
});
var FormGroup=React.createClass({
    render: function() {
        var item=[];
        switch(this.props.formItem.type){
           case "text":item.push(<TextBox data={this.props.formItem}/>) ;break;
           case "password":item.push(<PasswordBox data={this.props.formItem.value}/>) ;break;
           case "select":item.push(<DropDownList data={this.props.formItem.value}  options={this.props.formItem.options}/>) ;break;
           case "date":item.push(<DateBox data={this.props.formItem.value} />) ;break;
           case "checkbox":item.push(<CheckBox data={this.props.formItem.value} />) ;break;
           case "rediobox":item.push(<CheckBox data={this.props.formItem.value} />) ;break;
        }
        return (
            <div className="form-group group">
                <label className="control-label col-xs-12 col-sm-2 no-padding-right" for="roles">{this.props.formItem.title}</label>
                <div className="col-xs-12 col-sm-4">
                    <div className="clearfix">
                        {item}
                    </div>
                </div>

            </div>
        )
    }
});

var Form=React.createClass({
    render: function() {
        var groups =[];
        
        this.props.formItems.forEach(function(item) {
            	groups.push(<FormGroup formItem={item} />);
        });
    
        return (
        <form className="form-horizontal" id="validation-form" method="get">
            {groups}
            <div className="space-2"></div>
        </form>)
    }
});
var FormEdit = React.createClass({
    render: function() {
        return (
         <div className="widget-main">

            <Form formItems={this.props.formItems}>	

            </Form>

        </div>

        );
    }
});
/*
var formItems=[
{id:"dispalyName",title:"显示名称",value:"张三",type:"text"}
,{id:"password",title:"密码",value:"123",type:"password"}
,{id:"repassword",title:"密码",value:"123",type:"password"}
,{id:"exdate",title:"失效日期",value:"张三",type:"date"}
,{id:"recordstatus",title:"状态",value:"1",type:"select",options:[{value:0,text:"正常"},{value:-1,text:"已删除"}]}
]
React.render(<FormEdit name="form1" formItems={formItems} />, document.getElementById('formuser'));
*/