/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{r as a,o as e,c as t,e as l,d as r,w as o,h as i,a0 as u,f as d,bm as s,j as n,E as m,b as p}from"./index.61dd0a2d.js";import{E as c}from"./drawer.f8ea4ac7.js";import"./overlay.533fa516.js";import{E as y,a as h}from"./tab-pane.2818a6c6.js";import{E as v}from"./dialog.a2b2142c.js";import{E as f,a as b}from"./form-item.a3ced827.js";/* empty css              */import{E as I}from"./radio.64740f4f.js";/* empty css               */import"./tag.3ecd1fed.js";import"./checkbox.384d52e9.js";import"./scrollbar.ef9dc6d6.js";import{E as g,a as j}from"./table-column.720e0d1a.js";/* empty css                *//* empty css               */import{g as w,d as k,c as _,u as C,a as N}from"./authority.120a8f54.js";import V from"./menus.a455c93c.js";import x from"./apis.abb5eac9.js";import E from"./datas.00424527.js";import{W as A}from"./warningBar.3ea93790.js";import"./strings.ed5ba909.js";import"./refs.1f88eb22.js";import"./_Uint8Array.0caf8bf8.js";import"./_initCloneObject.e29c0e59.js";import"./rand.6d0d1e93.js";import"./isEqual.3dfeee36.js";import"./arrays.6dddb988.js";import"./index.cbcb4c8f.js";import"./index.ba9b69da.js";import"./tree.14d81143.js";import"./index.6a5a2fdc.js";import"./authorityBtn.0a9d8751.js";import"./_plugin-vue_export-helper.cdc0426e.js";import"./api.e309be6a.js";/* empty css             */const U={class:"authority"},q={class:"gva-table-box"},B={class:"gva-btn-list"},S={class:"dialog-footer"},z={name:"Authority"},D=Object.assign(z,{setup(z){const D=a([{authorityId:0,authorityName:"根角色"}]),R=a(!1),F=a("add"),O=a({}),T=a("新增角色"),P=a(!1),W=a(!1),$=a({}),G=a({authorityId:0,authorityName:"",parentId:0}),H=a({authorityId:[{required:!0,message:"请输入角色ID",trigger:"blur"},{validator:(a,e,t)=>/^[0-9]*[1-9][0-9]*$/.test(e)?t():t(new Error("请输入正整数")),trigger:"blur",message:"必须为正整数"}],authorityName:[{required:!0,message:"请输入角色名",trigger:"blur"}],parentId:[{required:!0,message:"请选择父角色",trigger:"blur"}]}),J=a(1),K=a(0),L=a(999),M=a([]),Q=a({}),X=async()=>{const a=await w({page:J.value,pageSize:L.value,...Q.value});0===a.code&&(M.value=a.data.list,K.value=a.data.total,J.value=a.data.page,L.value=a.data.pageSize)};X();const Y=(a,e)=>{O.value[a]=e},Z=a(null),aa=a(null),ea=a(null),ta=(a,e)=>{const t=[Z,aa,ea];e&&t[e].value.needConfirm&&(t[e].value.enterAndNext(),t[e].value.needConfirm=!1)},la=a(null),ra=()=>{la.value&&la.value.resetFields(),G.value={authorityId:0,authorityName:"",parentId:0}},oa=()=>{ra(),P.value=!1,W.value=!1},ia=()=>{if(G.value.authorityId=Number(G.value.authorityId),0===G.value.authorityId)return n({type:"error",message:"角色id不能为0"}),!1;la.value.validate((async a=>{if(a){switch(F.value){case"add":0===(await N(G.value)).code&&(n({type:"success",message:"添加成功!"}),X(),oa());break;case"edit":0===(await C(G.value)).code&&(n({type:"success",message:"添加成功!"}),X(),oa());break;case"copy":{const a={authority:{authorityId:0,authorityName:"",datauthorityId:[],parentId:0},oldAuthorityId:0};a.authority.authorityId=G.value.authorityId,a.authority.authorityName=G.value.authorityName,a.authority.parentId=G.value.parentId,a.authority.dataAuthorityId=$.value.dataAuthorityId,a.oldAuthorityId=$.value.authorityId;0===(await _(a)).code&&(n({type:"success",message:"复制成功！"}),X())}}ra(),P.value=!1}}))},ua=()=>{D.value=[{authorityId:0,authorityName:"根角色"}],da(M.value,D.value,!1)},da=(a,e,t)=>{G.value.authorityId=String(G.value.authorityId),a&&a.forEach((a=>{if(a.children&&a.children.length){const l={authorityId:a.authorityId,authorityName:a.authorityName,disabled:t||a.authorityId===G.value.authorityId,children:[]};da(a.children,l.children,t||a.authorityId===G.value.authorityId),e.push(l)}else{const l={authorityId:a.authorityId,authorityName:a.authorityName,disabled:t||a.authorityId===G.value.authorityId};e.push(l)}}))},sa=a=>{ra(),T.value="新增角色",F.value="add",G.value.parentId=a,ua(),P.value=!0};return(a,w)=>{const _=m,C=g,N=j,z=I,W=f,K=p,L=b,Q=v,ra=y,da=h,na=c;return e(),t("div",U,[l(A,{title:"注：右上角头像下拉可切换角色"}),r("div",q,[r("div",B,[l(_,{type:"primary",icon:"plus",onClick:w[0]||(w[0]=a=>sa(0))},{default:o((()=>[i("新增角色")])),_:1})]),l(N,{data:M.value,"tree-props":{children:"children",hasChildren:"hasChildren"},"row-key":"authorityId",style:{width:"100%"}},{default:o((()=>[l(C,{label:"角色ID","min-width":"180",prop:"authorityId"}),l(C,{align:"left",label:"角色名称","min-width":"180",prop:"authorityName"}),l(C,{align:"left",label:"操作",width:"460"},{default:o((a=>[l(_,{icon:"setting",type:"primary",link:"",onClick:e=>{return t=a.row,R.value=!0,void(O.value=t);var t}},{default:o((()=>[i("设置权限")])),_:2},1032,["onClick"]),l(_,{icon:"plus",type:"primary",link:"",onClick:e=>sa(a.row.authorityId)},{default:o((()=>[i("新增子角色")])),_:2},1032,["onClick"]),l(_,{icon:"copy-document",type:"primary",link:"",onClick:e=>(a=>{ua(),T.value="拷贝角色",F.value="copy";for(const e in G.value)G.value[e]=a[e];$.value=a,P.value=!0})(a.row)},{default:o((()=>[i("拷贝")])),_:2},1032,["onClick"]),l(_,{icon:"edit",type:"primary",link:"",onClick:e=>(a=>{ua(),T.value="编辑角色",F.value="edit";for(const e in G.value)G.value[e]=a[e];ua(),P.value=!0})(a.row)},{default:o((()=>[i("编辑")])),_:2},1032,["onClick"]),l(_,{icon:"delete",type:"primary",link:"",onClick:e=>{return t=a.row,void s.confirm("此操作将永久删除该角色, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{0===(await k({authorityId:t.authorityId})).code&&(n({type:"success",message:"删除成功!"}),1===M.value.length&&J.value>1&&J.value--,X())})).catch((()=>{n({type:"info",message:"已取消删除"})}));var t}},{default:o((()=>[i("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])]),l(Q,{modelValue:P.value,"onUpdate:modelValue":w[4]||(w[4]=a=>P.value=a),title:T.value},{footer:o((()=>[r("div",S,[l(_,{onClick:oa},{default:o((()=>[i("取 消")])),_:1}),l(_,{type:"primary",onClick:ia},{default:o((()=>[i("确 定")])),_:1})])])),default:o((()=>[l(L,{ref_key:"authorityForm",ref:la,model:G.value,rules:H.value,"label-width":"80px"},{default:o((()=>[l(W,{label:"父级角色",prop:"parentId"},{default:o((()=>[l(z,{modelValue:G.value.parentId,"onUpdate:modelValue":w[1]||(w[1]=a=>G.value.parentId=a),style:{width:"100%"},disabled:"add"===F.value,options:D.value,props:{checkStrictly:!0,label:"authorityName",value:"authorityId",disabled:"disabled",emitPath:!1},"show-all-levels":!1,filterable:""},null,8,["modelValue","disabled","options"])])),_:1}),l(W,{label:"角色ID",prop:"authorityId"},{default:o((()=>[l(K,{modelValue:G.value.authorityId,"onUpdate:modelValue":w[2]||(w[2]=a=>G.value.authorityId=a),disabled:"edit"===F.value,autocomplete:"off"},null,8,["modelValue","disabled"])])),_:1}),l(W,{label:"角色姓名",prop:"authorityName"},{default:o((()=>[l(K,{modelValue:G.value.authorityName,"onUpdate:modelValue":w[3]||(w[3]=a=>G.value.authorityName=a),autocomplete:"off"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"]),R.value?(e(),u(na,{key:0,modelValue:R.value,"onUpdate:modelValue":w[5]||(w[5]=a=>R.value=a),"custom-class":"auth-drawer","with-header":!1,size:"40%",title:"角色配置"},{default:o((()=>[l(da,{"before-leave":ta,type:"border-card"},{default:o((()=>[l(ra,{label:"角色菜单"},{default:o((()=>[l(V,{ref_key:"menus",ref:Z,row:O.value,onChangeRow:Y},null,8,["row"])])),_:1}),l(ra,{label:"角色api"},{default:o((()=>[l(x,{ref_key:"apis",ref:aa,row:O.value,onChangeRow:Y},null,8,["row"])])),_:1}),l(ra,{label:"资源权限"},{default:o((()=>[l(E,{ref_key:"datas",ref:ea,authority:M.value,row:O.value,onChangeRow:Y},null,8,["authority","row"])])),_:1})])),_:1})])),_:1},8,["modelValue"])):d("",!0)])}}});export{D as default};
