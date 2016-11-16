var ENTER_KEY_CODE = 13;
var ReactPropTypes = React.PropTypes;
var DataGridToolbar = React.createClass({

	propTypes: {
		id: ReactPropTypes.string,
		doNew: ReactPropTypes.func.isRequired,
		btnNewTitle: ReactPropTypes.string,
		doReload: ReactPropTypes.func.isRequired,
		value: ReactPropTypes.string
	},

	getInitialState: function() {
		return {
			value: this.props.value || ''
		};
	},

	/**
	 *        *    * @return {object}
	 *               *       */
	render: function() /*object*/ {
		var rows = [];
		if (this.props.btnNewTitle != undefined) {
			rows = < a href = "#"
			data - action = "new"
			onClick = {
				this._showDialog
			} > < i className = "ace-icon fa fa-plus white" > & nbsp; <
			/i><span className="white">{this.props.btnNewTitle}</span > < /a>;

		}
		return ( < div className = "widget-header btn-info" >
			< h5 className = "widget-title" >
			< i className = "ace-icon fa  fa-users" > < /i> {
			this.props.title
		} < /h5>

		< div className = "widget-toolbar" >

			< a href = "#"
		data - action = "reload"
		onClick = {
				this._doReload
			} >
			< i className = "ace-icon fa fa-refresh" > < /i> < /a > < /div>

		< div className = "widget-toolbar no-border" > {
			rows
		} < /div> < /div >
	);
},

/**
 * 	 *    * Invokes the callback passed in as onSave, allowing this component to be
 * 	 	 *       * used in different ways.
 * 	 	 	 *          */
_showDialog: function() {
	this.props.doNew();
},

/**
 * 	   *    * @param {object} event
 * 	   	   *       */
_doReload: function( /*object*/ event) {
	this.props.doReload();
},

});
var DataRow = React.createClass({
			render: function() {
					var rows = [];

					if (this.props.columns == undefined) {
						if (this.props.rowType == "foot")
							rows.push( < td className = "dataTables_empty" > 表格定义信息错误 < /td>)
							} else {
								if (this.props.rowType == "foot") {
									rows.push( < td colSpan = {
											this.props.columns.length
										}
										className = "dataTables_empty" > 数据加载中... < /td>)
									} else {
										this.props.columns.forEach(function(item) {
												if (undefined == item.mTitle) {
													rows.push( < th > {
															item.mData
														} < /th>);}
														else
															rows.push( < th > {
																	item.mTitle
																} < /th>);
															});
												}
											}
											return ( < tr > {
													rows
												} < /tr>);
											}
										});
									var DataGrid = React.createClass({

											_doNew: function( /*object*/ event) {
												this.props.doNew();
											},
											_doReload: function( /*object*/ event) {
												var displayLength = 15;
												var lengthMenu = [
													[15, 20, 25, 50, 100],
													[15, 20, 25, 50, 100]
												];
												if (undefined != this.props.displayLength) {
													displayLength = this.props.displayLength;
													lengthMenu = [
														[displayLength],
														[displayLength]
													];
												}
												$("#" + this.props.tableId).dataTable({
													"iDisplayLength": displayLength,
													"aLengthMenu": lengthMenu,
													"bAutoWidth": true,
													"bServerSide": true,
													"ordering": false,
													"paging": displayLength > 0,
													"sAjaxSource": this.props.ajaxSource,
													"sServerMethod": "GET",
													"fnServerParams": function(aoData) {
														aoData.push({
															"name": "extPara",
															"value": ""
														});
													},
													"aoColumns": this.props.columns,
													"bFilter": false,
													"bInfo": false,
													"bProcessing": false,
													"sDom": "<'top'f>rt<'row'<'col-sm-6'l><'col-sm-6'p>>",
													"oLanguage": {
														"sLengthMenu": "每页 _MENU_ 条",
														"sZeroRecords": "没有数据",
														"sSearch": "查找",
														"oPaginate": {
															"sFirst": "首页",
															"sPrevious": "前一页",
															"sNext": "后一页",
															"sLast": "尾页"
														},
													},

												});

											},
											componentDidMount: function() {
												this._doReload()
											},
											render: function() {
												return ( < div className =
													"widget-box widget-color-blue3 no-border" >
													< DataGridToolbar id = "searchBarInput"
													doReload = {
														this._doReload
													}
													doNew = {
														this._doNew
													}
													title = {
														this.props.title
													}
													btnNewTitle = {
														this.props.btnNewTitle
													} > < /DataGridToolbar> < div className = "widget-body" > <
													table id = {
														this.props.tableId
													}
													className =
													"table table-striped table-bordered table-hover dataTable no-footer" >
													< thead >
													< DataRow columns = {
														this.props.columns
													}
													rowType = "head" / >
													< /thead> < tbody > < DataRow columns = {
													this.props.columns
												}
												rowType = "foot" / >
													< /tbody> < /table > < /div> < /div >
											);
										}
									});
								/*
								 * var columns =[
								 *                 { "mData": "MenuID", "mTitle":"ID","bSortable": false,  "width": "7%" },
								 *                 		                { "mData": "MenuIconUrl", "mTitle":"图标", "bSortable": false,  "width": "7%" ,
								 *                 		                					                "mRender":function(data,type,full){
								 *                 		                					                								        				return '<i class="'+data+'"></i>';
								 *                 		                					                								        																	        			}
								 *                 		                					                								        																	        										                 },
								 *                 		                					                								        																	        										                 									                 { "mData": "MenuName", "mTitle":"名称", "bSortable": false,  "width": "20%",
								 *                 		                					                								        																	        										                 									                 												         			"mRender":function(data,type,full){
								 *                 		                					                								        																	        										                 									                 												         																				        		 		var out=""
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        				for(var i=1;i<full.TreeLevel;i++){
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        					out+="&nbsp;&nbsp;&nbsp;&nbsp;"
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        				}
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        				return out+data;
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        			}
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 },
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 { "mData": "MenuNavUrl", "mTitle":"地址", "bSortable": false, "width": "60%" },
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 { "mData": "SortIndex", "mTitle":"编号", "bSortable": false, "width": "10%"},
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 {
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     "mData": "MenuID", "mTitle":"", "bSortable": false,"width":"10%",
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         "mRender": function (data, type, full) {
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         var fun = "showDialog('/system/menu/view/" + data + "','查看菜单信息');";
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         																																			                         var result= '<button class="btn btn-success btn-minier" onclick=' + fun + '><i class="icon-plus " />查看</button>';
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         																																			                         																																						                         return result;
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         																																			                         																																						                         																																									                     }
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         																																			                         																																						                         																																									                     																												                     }
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         																																			                         																																						                         																																									                     																												                                 ];
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         																																			                         																																						                         																																									                     																												                                 	    React.render(<DataGrid name="grid1" tableId="table_menu" title="user" btnNewTitle="New" columns={columns} ajaxSource="/system/menu/flat/json" />
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         																																			                         																																						                         																																									                     																												                                 	    			    , document.getElementById('grid1'));
								 *                 		                					                								        																	        										                 									                 												         																				        		 																									        																																	        																																									        																										        																															        																			        		 																			                 																					                 																							                 																										                     																												                         																																                         																																			                         																																						                         																																									                     																												                                 	    			    */
