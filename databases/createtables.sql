drop table pa_actionlogformat CASCADE;

drop table pa_actionlogs CASCADE;

drop table pa_actions CASCADE;

drop table pa_employees CASCADE;

drop table pa_r_actionlogs CASCADE;

drop table pa_r_employees_orgunits CASCADE;

drop table pa_systemoptionitems CASCADE;

drop table pa_systemoptions CASCADE;

drop table pa_orgunits CASCADE;

drop table pa_menus CASCADE;

drop table pa_modules CASCADE;

drop table pa_permissions CASCADE;

drop table pa_r_roles_permissions CASCADE;

drop table pa_r_users_roles CASCADE;

drop table pa_roles CASCADE;

drop table pa_users CASCADE;


/*==============================================================*/
/* Table: pa_actionlogformat                                    */
/*==============================================================*/
create table pa_actionlogformat (
   actionlogformatid    serial not null,
   actionid             int4                 null,
   associateitemdefinitionid int4                 not null,
   actionlogformattext  varchar(128)         not null,
   actionlogformatdesc  varchar(256)         not null,
   recordstatus         int2                 not null,
   constraint pk_pa_actionlogformat primary key (actionlogformatid)
);

/*==============================================================*/
/* Table: pa_actionlogs                                         */
/*==============================================================*/
create table pa_actionlogs (
   actionlogid          serial not null,
   actionid             int4                 null,
   actionlognotes       varchar(128)         not null,
   actionlogdesc        varchar(256)         not null,
   createddate          date                 not null,
   createdbyuserid      int4                 not null,
   constraint pk_pa_actionlogs primary key (actionlogid)
);

/*==============================================================*/
/* Table: pa_actions                                            */
/*==============================================================*/
create table pa_actions (
   actionid             int4                 not null,
   actionname           varchar(128)         not null,
   actiondesc           varchar(256)         not null,
   recordstatus         int2                 not null,
   constraint pk_pa_actions primary key (actionid)
);

/*==============================================================*/
/* Table: pa_employees                                          */
/*==============================================================*/
create table pa_employees (
   employeeid           serial not null,
   employeename         varchar(16)          not null,
   employeecode         varchar(16)          not null,
   countryid            int4                 null,
   gender               int4                 null,
   dateofbirth          date                 null,
   placeofbirth         date                 null,
   hometown             varchar(64)          null,
   hukou                varchar(64)          null,
   maritalstatus        int4                 null,
   bloodtype            int4                 null,
   race                 varchar(16)          null,
   party                varchar(128)         null,
   partyjoineddate      date                 null,
   height               varchar(32)          null,
   interests            varchar(128)         null,
   recordstatus         int4                 null,
   createddate          date                 not null,
   createdbyuserid      int4                 not null,
   lastupdated          date                 not null,
   lastupdatedbyuserid  int4                 not null,
   constraint pk_pa_employees primary key (employeeid)
);

/*==============================================================*/
/* Table: pa_r_actionlogs                                       */
/*==============================================================*/
create table pa_r_actionlogs (
   r_actionlogid        serial not null,
   actionlogid          int4                 null,
   associateitemdefinitionid int4                 not null,
   associateitemid      int4                 not null,
   constraint pk_pa_r_actionlogs primary key (r_actionlogid)
);

/*==============================================================*/
/* Table: pa_r_employees_orgunits                               */
/*==============================================================*/
create table pa_r_employees_orgunits (
   r_employee_orgunitid serial not null,
   orgunitid            int4                 not null,
   employeeid           int4                 not null,
   constraint pk_pa_r_employees_orgunits primary key (r_employee_orgunitid)
);

/*==============================================================*/
/* Table: pa_systemoptionitems                                  */
/*==============================================================*/
create table pa_systemoptionitems (
   systemoptionitemid   serial not null,
   systemoptionid       int4                 not null,
   systemoptionitempid  int4                 not null,
   systemoptionitemvalue int4                 not null,
   systemoptionitemname timestamp with time zone not null,
   systemoptionitemdesc varchar(128)         null,
   ispublic             int2                 null,
   recordpermissions    int2                 null,
   recordstatus         int2                 null default 0,
   createddate          date                 null,
   createdbyuserid      int4                 not null,
   lastupdated          date                 not null,
   lastupdatedbyuserid  int4                 not null,
   constraint pk_pa_systemoptionitems primary key (systemoptionitemid)
);

/*==============================================================*/
/* Table: pa_systemoptions                                      */
/*==============================================================*/
create table pa_systemoptions (
   systemoptionid       int4                 not null,
   systemoptionname     varchar(64)          not null,
   systemoptiondesc     varchar(256)         null,
   isuserdefined        bool                 null,
   constraint pk_pa_systemoptions primary key (systemoptionid)
);

/*==============================================================*/
/* Table: pa_orgunits                                           */
/*==============================================================*/
create table pa_orgunits (
   orgunitid            serial not null,
   orgunittype          int4                 null,
   orgunitpid           int4                 null,
   orgunitcode          varchar(32)          null,
   orgunitname          varchar(128)         null,
   orgunitdesc          varchar(256)         null,
   themeid              int4                 null,
   sortindex            int2                 null,
   constraint pk_pa_orgunits primary key (orgunitid)
);

/*==============================================================*/
/* Table: pa_menus                                              */
/*==============================================================*/
create table pa_menus (
   menuid               serial not null,
   moduleid             int4                 not null,
   menupid              int4                 not null,
   menuname             varchar(128)         not null,
   menuiconurl          varchar(256)         null,
   menunavurl           varchar(256)         null,
   menutarget           varchar(64)          null,
   menudesc             varchar(256)         null,
   sortindex            int4                 not null,
   istreeleaf           bool                 not null,
   actionid             varchar(36)          null,
   treelevel            int4                 not null,
   recordstatus         int4                 null,
   createddate          date                 not null,
   createdbyuserid      int4                 not null,
   lastupdated          date                 not null,
   lastupdatedbyuserid  int4                 not null,
   constraint pk_pa_menus primary key (menuid)
);

/*==============================================================*/
/* Table: pa_modules                                            */
/*==============================================================*/
create table pa_modules (
   moduleid             serial not null,
   modulepid            int4                 not null,
   modulename           varchar(64)          not null,
   moduledesc           varchar(512)         null,
   sortindex            int4                 not null,
   istreeleaf           bool                 not null,
   treelevel            int4                 not null,
   recordstatus         int4                 not null,
   createddate          date                 not null,
   createdbyuserid      int4                 not null,
   lastupdated          date                 not null,
   lastupdatedbyuserid  int4                 not null,
   constraint pk_pa_modules primary key (moduleid)
);

/*==============================================================*/
/* Table: pa_permissions                                        */
/*==============================================================*/
create table pa_permissions (
   permissionid         serial not null,
   menuid               int4                 not null,
   actionid             varchar(36)          not null default null,
   permissionname       varchar(128)         not null,
   permissiondesc       varchar(128)         not null,
   permissiongroup      varchar(64)          not null,
   permissionmemo       varchar(512)         null,
   recordstatus         int4                 not null,
   createddate          date                 not null,
   createdbyuserid      int4                 not null,
   lastupdated          date                 not null,
   lastupdatedbyuserid  int4                 not null,
   constraint pk_pa_permissions primary key (permissionid)
);

/*==============================================================*/
/* Table: pa_r_roles_permissions                                */
/*==============================================================*/
create table pa_r_roles_permissions (
   r_role_permissionid  serial not null,
   roleid               int4                 not null,
   permissionid         int4                 not null,
   constraint pk_pa_r_roles_permissions primary key (r_role_permissionid)
);

/*==============================================================*/
/* Table: pa_r_users_roles                                      */
/*==============================================================*/
create table pa_r_users_roles (
   r_user_roleid        serial not null,
   roleid               int4                 not null,
   userid               int4                 not null,
   constraint pk_pa_r_users_roles primary key (r_user_roleid)
);

/*==============================================================*/
/* Table: pa_roles                                              */
/*==============================================================*/
create table pa_roles (
   roleid               serial not null,
   rolename             varchar(128)         not null,
   roledesc             varchar(256)         null,
   recordstatus         int4                 not null,
   createddate          date                 not null,
   createdbyuserid      int4                 not null,
   lastupdated          date                 not null,
   lastupdatedbyuserid  int4                 not null,
   constraint pk_pa_roles primary key (roleid)
);

/*==============================================================*/
/* Table: pa_users                                              */
/*==============================================================*/
create table pa_users (
   userid               serial not null,
   userloginname        varchar(64)          null,
   userpassword         varchar(128)         null,
   userdisplayname      varchar(64)          null,
   userphoto            varchar(256)         null,
   usersignature        varchar(256)         null,
   usermustchangepwd    bool                 null,
   usercannotchangepwd  bool                 null,
   useraccountexpirationdate date                 null,
   recordstatus         int4                 null,
   createddate          date                 not null,
   createdbyuserid      int4                 not null,
   lastupdated          date                 not null,
   lastupdatedbyuserid  int4                 not null,
   constraint pk_pa_users primary key (userid),
   constraint uq_userloginname unique (userloginname)
);

alter table pa_actionlogformat
   add constraint fk_actionlogs_actionid foreign key (actionid)
      references pa_actions (actionid);

alter table pa_actionlogs
   add constraint fk_actionlogs_actionid foreign key (actionid)
      references pa_actions (actionid);

alter table pa_r_actionlogs
   add constraint fk_r_actionlogs_actionlogid foreign key (actionlogid)
      references pa_actionlogs (actionlogid);

alter table pa_r_employees_orgunits
   add constraint fk_r_users_orgunits_empid foreign key (employeeid)
      references pa_employees (employeeid);

alter table pa_r_employees_orgunits
   add constraint fk_r_users_orgunits_orgunitid foreign key (orgunitid)
      references pa_orgunits (orgunitid);

alter table pa_systemoptionitems
   add constraint fk_sysoptitems_systemoptionid foreign key (systemoptionid)
      references pa_systemoptions (systemoptionid);

alter table pa_menus
   add constraint fk_pa_menus_reference_pa_modul foreign key (moduleid)
      references pa_modules (moduleid);

alter table pa_permissions
   add constraint fk_pa_permi_reference_pa_menus foreign key (menuid)
      references pa_menus (menuid);

alter table pa_r_roles_permissions
   add constraint fk_r_roles_permissions_roleid foreign key (roleid)
      references pa_roles (roleid);

alter table pa_r_roles_permissions
   add constraint fk_r_roles_perms_permissionid foreign key (permissionid)
      references pa_permissions (permissionid);

alter table pa_r_users_roles
   add constraint fk_r_users_roles_roleid foreign key (roleid)
      references pa_roles (roleid);

alter table pa_r_users_roles
   add constraint fk_r_users_roles_userid foreign key (userid)
      references pa_users (userid);
