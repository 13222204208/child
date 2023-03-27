/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{u as e,r as a,a as l,E as s,b as t,o,c as d,d as p,n as r,e as i,w as m,f as u,g as n,h as c,i as b,j as f}from"./index.61dd0a2d.js";import{E as v,a as g}from"./form-item.a3ced827.js";/* empty css              */import"./tag.3ecd1fed.js";import{E as y,a as h}from"./select.31449efc.js";import"./scrollbar.ef9dc6d6.js";/* empty css               *//* empty css               */import{i as _}from"./initdb.74f5eeea.js";import{_ as j}from"./_plugin-vue_export-helper.cdc0426e.js";import"./_Uint8Array.0caf8bf8.js";import"./_initCloneObject.e29c0e59.js";import"./index.cbcb4c8f.js";import"./index.ba9b69da.js";import"./strings.ed5ba909.js";import"./isEqual.3dfeee36.js";const q={class:"init_page"},V={class:"init_page_panel"},w=n('<div class="hello_title" data-v-f5ae62a5>GIN-VUE-ADMIN</div><p class="in-two a-fadeinT" data-v-f5ae62a5>初始化须知</p><p class="init_p" data-v-f5ae62a5>1.您需有用一定的VUE和GOLANG基础</p><p class="init_p" data-v-f5ae62a5>2.请您确认是否已经阅读过官方文档</p><p class="init_p" data-v-f5ae62a5>3.请您确认是否了解后续的配置流程</p><p class="init_p" data-v-f5ae62a5>4.如果您使用mysql数据库，请确认数据库引擎为<span style="color:red;" data-v-f5ae62a5>innoDB</span></p><p class="init_p" data-v-f5ae62a5>注：开发组不为文档中书写过的内容提供无偿服务</p>',7),N={class:"init_btn"},k={style:{"text-align":"right"}},U={name:"Init"},x=j(Object.assign(U,{setup(n){const j=e(),U=a(0),x=()=>{U.value=U.value+1},T=()=>{window.open("https://www.gin-vue-admin.com/guide/start-quickly/env.html")},O=a(!1),E=l({dbType:"mysql",host:"127.0.0.1",port:"3306",userName:"root",password:"",dbName:"gva"}),C=e=>{switch(e){case"mysql":default:Object.assign(E,{dbType:"mysql",host:"127.0.0.1",port:"3306",userName:"root",password:"",dbName:"gva"});break;case"pgsql":Object.assign(E,{dbType:"pgsql",host:"127.0.0.1",port:"5432",userName:"postgres",password:"",dbName:"gva"});break;case"oracle":Object.assign(E,{dbType:"oracle",host:"127.0.0.1",port:"1521",userName:"oracle",password:"",dbName:"gva"});break;case"mssql":Object.assign(E,{dbType:"mssql",host:"127.0.0.1",port:"1433",userName:"mssql",password:"",dbName:"gva"})}},I=async()=>{const e=b.service({lock:!0,text:"正在初始化数据库，请稍候",spinner:"loading",background:"rgba(0, 0, 0, 0.7)"});try{const a=await _(E);0===a.code&&(O.value=!0,f({type:"success",message:a.msg}),j.push({name:"Login"})),e.close()}catch(a){e.close()}};return(e,a)=>{const l=s,n=y,b=h,f=v,_=t,j=g;return o(),d("div",q,[p("div",V,[U.value<2?(o(),d("div",{key:0,id:"hello",class:r([[U.value<1?"slide-in-fwd-top":"slide-out-right"],"hello"])},[p("div",null,[w,p("p",N,[i(l,{type:"primary",onClick:T},{default:m((()=>[c(" 阅读文档 ")])),_:1}),i(l,{type:"primary",onClick:x},{default:m((()=>[c(" 我已确认 ")])),_:1})])])],2)):u("",!0),U.value>0?(o(),d("div",{key:1,class:r([[U.value>0&&!O.value?"slide-in-left":"",O.value?"slide-out-right":""],"form"])},[i(j,{ref:"formRef",model:E,"label-width":"100px"},{default:m((()=>[i(f,{label:"数据库类型"},{default:m((()=>[i(b,{modelValue:E.dbType,"onUpdate:modelValue":a[0]||(a[0]=e=>E.dbType=e),placeholder:"请选择",onChange:C},{default:m((()=>[i(n,{key:"mysql",label:"mysql",value:"mysql"}),i(n,{key:"pgsql",label:"pgsql",value:"pgsql"}),i(n,{key:"oracle",label:"oracle",value:"oracle"}),i(n,{key:"mssql",label:"mssql",value:"mssql"})])),_:1},8,["modelValue"])])),_:1}),i(f,{label:"host"},{default:m((()=>[i(_,{modelValue:E.host,"onUpdate:modelValue":a[1]||(a[1]=e=>E.host=e),placeholder:"请输入数据库链接"},null,8,["modelValue"])])),_:1}),i(f,{label:"port"},{default:m((()=>[i(_,{modelValue:E.port,"onUpdate:modelValue":a[2]||(a[2]=e=>E.port=e),placeholder:"请输入数据库端口"},null,8,["modelValue"])])),_:1}),i(f,{label:"userName"},{default:m((()=>[i(_,{modelValue:E.userName,"onUpdate:modelValue":a[3]||(a[3]=e=>E.userName=e),placeholder:"请输入数据库用户名"},null,8,["modelValue"])])),_:1}),i(f,{label:"password"},{default:m((()=>[i(_,{modelValue:E.password,"onUpdate:modelValue":a[4]||(a[4]=e=>E.password=e),placeholder:"请输入数据库密码（没有则为空）"},null,8,["modelValue"])])),_:1}),i(f,{label:"dbName"},{default:m((()=>[i(_,{modelValue:E.dbName,"onUpdate:modelValue":a[5]||(a[5]=e=>E.dbName=e),placeholder:"请输入数据库名称"},null,8,["modelValue"])])),_:1}),i(f,null,{default:m((()=>[p("div",k,[i(l,{type:"primary",onClick:I},{default:m((()=>[c("立即初始化")])),_:1})])])),_:1})])),_:1},8,["model"])],2)):u("",!0)])])}}}),[["__scopeId","data-v-f5ae62a5"]]);export{x as default};