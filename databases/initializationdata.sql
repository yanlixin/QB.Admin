/*source /var/ftp/pub/sogoodsoft.sql;*/
--use qubei;
delete from pa_r_users_roles;
delete from pa_r_roles_permissions;
delete from pa_roles;
delete from pa_permissions;
delete from pa_menus;
delete from pa_modules;
delete from pa_users;
insert into pa_users(userid,userloginname,userpassword,userdisplayname,recordstatus,createddate,createdbyuserid,lastupdated,lastupdatedbyuserid)
values
(1,'master','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','master',0,'2015-01-01',1,'2015-01-01',1),
(2,'david.yan','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','David Yan',0,'2015-01-01',1,'2015-01-01',1),
(3,'jack.wang','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','Jack Wang',0,'2015-01-01',1,'2015-01-01',1),
(4,'张强','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','张小强',0,'2015-01-01',1,'2015-01-01',1),
(5,'yanlx@qiao.top','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','Yanlx@qiao.top',0,'2015-01-01',1,'2015-01-01',1),
(6,'zym_1023','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','Zhangym',0,'2015-01-01',1,'2015-01-01',1),
(7,'wanghong','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','王红',0,'2015-01-01',1,'2015-01-01',1),
(8,'fuyuqiang','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','付宇强',0,'2015-01-01',1,'2015-01-01',1),
(9,'ligang','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','李刚',0,'2015-01-01',1,'2015-01-01',1),
(10,'wangjun','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','王军',0,'2015-01-01',1,'2015-01-01',1),
(11,'zhangliang','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','张亮',0,'2015-01-01',1,'2015-01-01',1),
(12,'admin','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','管理员',0,'2015-01-01',1,'2015-01-01',1),
(13,'zhangqiang','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','张强',0,'2015-01-01',1,'2015-01-01',1),
(14,'lixiaogang','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','李晓刚',0,'2015-01-01',1,'2015-01-01',1),
(15,'wanjun','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','万军',0,'2015-01-01',1,'2015-01-01',1),
(16,'zhangli','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','张丽',0,'2015-01-01',1,'2015-01-01',1),
(17,'毛宇','A0dSqYOeLlJl82BcvexVME4dVzLLGrUP870ARXxBPRMSjpKxYd4GJjoVTG7OHQQIC69H3K4hT3CKMeJYlbjYHg==','毛宇',0,'2015-01-01',1,'2015-01-01',1)
;

insert into pa_modules(moduleid,modulepid,modulename,moduledesc,sortindex,istreeleaf,treelevel,recordstatus,createddate,createdbyuserid,lastupdated,lastupdatedbyuserid)
values
(1,0,'基础管理平台','基础管理平台',1000,FALSE,0,0,'2015-01-01',1,'2015-01-01',1),
(2,0,'客户管理','客户管理系统',1040,FALSE,0,0,'2015-01-01',1,'2015-01-01',1),
(3,0,'合同管理','合同管理系统',1080,FALSE,0,0,'2015-01-01',1,'2015-01-01',1),
(4,0,'债权管理','债权管理系统',1120,FALSE,0,0,'2015-01-01',1,'2015-01-01',1),
(5,0,'人力资源管理','人力资源管理系统',1140,FALSE,0,0,'2015-01-01',1,'2015-01-01',1)
;

insert into pa_menus(menuid ,moduleid,menupid,menuname,menuiconurl,menunavurl,menutarget,menudesc
,sortindex,istreeleaf,actionid,treelevel,recordstatus,createddate,createdbyuserid,lastupdated,lastupdatedbyuserid)
values
(1,1,0,'系统管理','menu-icon fa fa-wrench',null,null,'平台系统管理菜单',10010000,false,null,1,0,'2015-01-01',1,'2015-01-01',1),
(2,1,1,'用户管理',null,'/system/users/list',null,'平台系统用户管理菜单',10010010,true,'85b937084cd483186c79e115f010cd2d',2,0,'2015-01-01',1,'2015-01-01',1),
(3,1,1,'角色管理',null,'/system/roles/list',null,'平台系统角色管理菜单',10010020,true,'e34bb9b5bd3af64e7aab0ce98aea6580',2,0,'2015-01-01',1,'2015-01-01',1),
(4,1,1,'权限管理',null,'/system/permission/list',null,'平台系统权限管理菜单',10010030,true,'2cf1d3ba19982c78a0022481ab799125',2,0,'2015-01-01',1,'2015-01-01',1),
(12,1,1,'菜单管理',null,'/system/menu/list',null,'平台系统菜单管理菜单',10010040,true,'85b937084cd483186c79e115f010cd2c',2,0,'2015-01-01',1,'2015-01-01',1),
(5,1,0,'客户管理','menu-icon fa fa-users',null,null,'客户管理菜单',10040000,false,null,1,0,'2015-01-01',1,'2015-01-01',1),
(6,1,5,'线索管理',null,'#',null,'线索管理菜单',10040010,true,'aebf08b61a7683b16a5a44d950b37082',2,0,'2015-01-01',1,'2015-01-01',1),
(7,1,5,'商机管理',null,'#',null,'商机管理菜单',10040020,true,'a58d644e9e58044b6c9ee052c1b3b217',2,0,'2015-01-01',1,'2015-01-01',1),
(8,1,5,'我的客户',null,'#',null,'我的客户管理菜单',10040030,true,'89a15a3990a4d2c3be8152d8a9e8bf51',2,0,'2015-01-01',1,'2015-01-01',1),
(9,1,0,'合同管理','menu-icon fa fa-book',null,null,'合同管理管理菜单',10080000,false,null,1,0,'2015-01-01',1,'2015-01-01',1),
(10,1,0,'债权管理','menu-icon fa fa-legal',null,null,'债权管理菜单',10120000,false,null,1,0,'2015-01-01',1,'2015-01-01',1),
(11,1,1,'ORM Tools',null,'/model/view',null,'ORM工具',10010030,true,'85b937084cd483186c79e115f010cd2a',2,0,'2015-01-01',1,'2015-01-01',1)
;

insert into pa_permissions(permissionid,menuid,actionid,permissionname,permissiondesc,
permissiongroup,permissionmemo,recordstatus,createddate,createdbyuserid,lastupdated,lastupdatedbyuserid)
values
(1,2,'85b937084cd483186c79e115f010cd2d','用户列表浏览的权限','[系统管理]用户列表浏览权限','用户管理',null,0,'2015-01-01',1,'2015-01-01',1),
(2,2,'807a88e3c90187f39d88976657c713e2','用户浏览的权限','[系统管理]用户浏览权限','用户管理',null,0,'2015-01-01',1,'2015-01-01',1),
(3,2,'bbb9ea9954453d1691b76db96b89a057','用户新增的权限','[系统管理]用户新增权限','用户管理',null,0,'2015-01-01',1,'2015-01-01',1),
(4,2,'7bde5239b6772deed727bff9f7bbeb9c','用户编辑的权限','[系统管理]用户编辑权限','用户管理',null,0,'2015-01-01',1,'2015-01-01',1),
(5,2,'23f7fcb4052b4b19c6c3b8cd95dd40c8','用户删除的权限','[系统管理]用户删除权限','用户管理',null,0,'2015-01-01',1,'2015-01-01',1),
(6,2,'df9981921130a4c91ccf920c074b8f06','用户重置密码的权限' ,'[系统管理]用户重置密码权限','用户管理',null,0,'2015-01-01',1,'2015-01-01',1),
(7,3,'e34bb9b5bd3af64e7aab0ce98aea6580','角色列表浏览的权限','[系统管理]角色列表浏览权限','角色管理',null,0,'2015-01-01',1,'2015-01-01',1),
(8,3,'84983713b0e75dabdf909512f13d6c88','角色浏览的权限','[系统管理]角色浏览权限','角色管理',null,0,'2015-01-01',1,'2015-01-01',1),
(9,3,'9f75bc140204df449828bcaf8a6616ee','角色新增的权限','[系统管理]角色新增权限','角色管理',null,0,'2015-01-01',1,'2015-01-01',1),
(10,3,'9c78c29b87d2c96d276898ed6151779b','角色编辑的权限','[系统管理]角色编辑权限','角色管理',null,0,'2015-01-01',1,'2015-01-01',1),
(11,3,'ee577f771a2472e127dbf3c9b770fdf8','角色删除的权限','[系统管理]角色删除权限','角色管理',null,0,'2015-01-01',1,'2015-01-01',1),
(12,4,'e34bb9b5bd3af64e7aab0ce98aea6580','权限列表浏览的权限','[系统管理]权限列表浏览权限','权限管理',null,0,'2015-01-01',1,'2015-01-01',1),
(13,4,'84983713b0e75dabdf909512f13d6c88','权限浏览的权限','[系统管理]权限浏览权限','权限管理',null,0,'2015-01-01',1,'2015-01-01',1),
(14,4,'9f75bc140204df449828bcaf8a6616ee','权限新增的权限','[系统管理]权限新增权限','权限管理',null,0,'2015-01-01',1,'2015-01-01',1),
(15,4,'9c78c29b87d2c96d276898ed6151779b','权限编辑的权限','[系统管理]权限编辑权限','权限管理',null,0,'2015-01-01',1,'2015-01-01',1),
(16,4,'ee577f771a2472e127dbf3c9b770fdf8','权限删除的权限','[系统管理]权限删除权限','权限管理',null,0,'2015-01-01',1,'2015-01-01',1),
(17,9,'2cf1d3ba19982c78a0022481ab799125','合同管理浏览','[合同管理]浏览权限','合同管理',null,0,'2015-01-01',1,'2015-01-01',1)
;

insert into pa_roles(roleid,rolename,roledesc,recordstatus,createddate,createdbyuserid,lastupdated,lastupdatedbyuserid)
values(1,'系统用户','最基础权限',0,'2015-01-01',1,'2015-01-01',1),
(2,'系统管理员','系统管理员,拥有系统的超级权限',0,'2015-01-01',1,'2015-01-01',1),
(3,'客户管理用户','访问客户管理系统的基础权限',0,'2015-01-01',1,'2015-01-01',1),
(4,'客户管理管理员','拥有客户管理系统的超级权限，拥有此权限可以为其他用户分配系统管理系统的权限',0,'2015-01-01',1,'2015-01-01',1),
(5,'合同管理用户','访问合同管理系统的基础权限',0,'2015-01-01',1,'2015-01-01',1),
(6,'合同管理管理员','拥有合同管理系统的超级权限，拥有此权限可以为其他用户分配合同管理系统的权限',0,'2015-01-01',1,'2015-01-01',1)
;

insert into pa_r_roles_permissions(roleid,permissionid)
values(2,1),(2,2),(2,3),(2,4),(2,5),(2,6),(2,7),(2,8),(2,9),(2,10),(2,11),(2,12),(2,13),(2,14),(2,15),(2,16),(2,17)
,(5,17)
;

insert into pa_r_users_roles(roleid,userid)
values(1,1),(1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8),(1,9),(1,10),
(2,1),(2,2),(2,3),
(3,1),(3,3)
;

