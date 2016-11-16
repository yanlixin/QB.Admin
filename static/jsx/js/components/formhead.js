var FormHead = React.createClass({
    render: function() {
        return (
                <h4>
                    {this.props.title}
                    <small className="red">
                        <i className="ace-icon fa fa-angle-double-right"></i>
                        {this.props.message}
                    </small>
                </h4>

        );
    }
});
React.render(<FormHead title="form1" message="消息" />, document.getElementById('formuser'));